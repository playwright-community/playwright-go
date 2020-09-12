package playwright

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var pw *Playwright
var globalTestHelper *TestHelperData

func TestMain(m *testing.M) {
	BeforeAll()
	code := m.Run()
	AfterAll()
	os.Exit(code)
}

func BeforeAll() {
	var err error
	pw, err = Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	browserName := os.Getenv("BROWSER")
	var browserType *BrowserType
	if browserName == "chromium" || browserName == "" {
		browserType = pw.Chromium
	} else if browserName == "firefox" {
		browserType = pw.Firefox
	} else if browserName == "webkit" {
		browserType = pw.WebKit
	}

	browser, err := browserType.Launch(BrowserTypeLaunchOptions{
		Headless: Bool(os.Getenv("HEADFUL") == ""),
	})
	if err != nil {
		log.Fatalf("could not launch: %v", err)
	}
	globalTestHelper = &TestHelperData{
		Playwright: pw,
		Browser:    browser,
		IsChromium: browserName == "chromium" || browserName == "",
		IsFirefox:  browserName == "firefox",
		IsWebKit:   browserName == "webkit",
		server:     newTestServer(),
		assetDir:   "tests/assets/",
		utils:      &testUtils{},
	}
}

func AfterAll() {
	globalTestHelper.server.testServer.Close()
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
}

type TestHelperData struct {
	t          *testing.T
	Playwright *Playwright
	Browser    *Browser
	Context    *BrowserContext
	Page       *Page
	IsChromium bool
	IsFirefox  bool
	IsWebKit   bool
	server     *testServer
	assetDir   string
	utils      *testUtils
}

var CONTEXT_OPTIONS = BrowserNewContextOptions{
	AcceptDownloads: Bool(true),
}

func BeforeEach(t *testing.T) *TestHelperData {
	context, err := globalTestHelper.Browser.NewContext(CONTEXT_OPTIONS)
	require.NoError(t, err)
	globalTestHelper.Context = context
	page, err := context.NewPage()
	require.NoError(t, err)
	globalTestHelper.Page = page
	return globalTestHelper
}

func (t *TestHelperData) Asset(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %v", err)
	}
	return filepath.Join(cwd, "tests", "assets", path)
}

func (t *TestHelperData) AfterEach(closeContext ...bool) {
	if len(closeContext) == 0 {
		if err := t.Context.Close(); err != nil {
			t.t.Errorf("could not close context: %v", err)
		}
	}
	t.server.AfterEach()
}

func newTestServer() *testServer {
	ts := &testServer{
		routes:              make(map[string]http.HandlerFunc),
		requestSubscriberes: make(map[string][]chan *http.Request),
	}
	ts.testServer = httptest.NewServer(http.HandlerFunc(ts.serveHTTP))
	ts.PREFIX = ts.testServer.URL
	ts.EMPTY_PAGE = ts.testServer.URL + "/empty.html"
	return ts
}

type testServer struct {
	sync.Mutex
	testServer          *httptest.Server
	routes              map[string]http.HandlerFunc
	requestSubscriberes map[string][]chan *http.Request
	PREFIX              string
	EMPTY_PAGE          string
}

func (t *testServer) AfterEach() {
	t.Lock()
	t.routes = make(map[string]http.HandlerFunc)
	t.requestSubscriberes = make(map[string][]chan *http.Request)
	t.Unlock()
}

func (t *testServer) serveHTTP(w http.ResponseWriter, r *http.Request) {
	t.Lock()
	defer t.Unlock()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if handlers, ok := t.requestSubscriberes[r.URL.Path]; ok {
		for _, handler := range handlers {
			handler <- r
		}
	}
	if route, ok := t.routes[r.URL.Path]; ok {
		route(w, r)
		return
	}
	http.FileServer(http.Dir("./tests/assets")).ServeHTTP(w, r)
}

func (s *testServer) SetRoute(path string, f http.HandlerFunc) {
	s.routes[path] = f
}

func (s *testServer) SetRedirect(from, to string) {
	s.SetRoute(from, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, http.StatusFound)
	})
}

func (s *testServer) WaitForRequestChan(path string) <-chan *http.Request {
	if _, ok := s.requestSubscriberes[path]; !ok {
		s.requestSubscriberes[path] = make([]chan *http.Request, 0)
	}
	channel := make(chan *http.Request, 1)
	s.requestSubscriberes[path] = append(s.requestSubscriberes[path], channel)
	return channel
}

func Map(vs interface{}, f func(interface{}) interface{}) []interface{} {
	v := reflect.ValueOf(vs)
	vsm := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		vsm[i] = f(v.Index(i).Interface())
	}
	return vsm
}

// ChanToSlice reads all data from ch (which must be a chan), returning a
// slice of the data. If ch is a 'T chan' then the return value is of type
// []T inside the returned interface.
// A typical call would be sl := ChanToSlice(ch).([]int)
func ChanToSlice(ch interface{}, amount int) interface{} {
	chv := reflect.ValueOf(ch)
	slv := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(ch).Elem()), 0, 0)
	for i := 0; i < amount; i++ {
		v, ok := chv.Recv()
		if !ok {
			return slv.Interface()
		}
		slv = reflect.Append(slv, v)
	}
	return slv.Interface()
}

type syncSlice struct {
	sync.Mutex
	slice []interface{}
}

func (s *syncSlice) Append(v interface{}) {
	s.Lock()
	s.slice = append(s.slice, v)
	s.Unlock()
}

func (s *syncSlice) Get() interface{} {
	s.Lock()
	defer s.Unlock()
	return s.slice
}

func newSyncSlice() *syncSlice {
	return &syncSlice{
		slice: make([]interface{}, 0),
	}
}

type testUtils struct {
}

func (t *testUtils) AttachFrame(page *Page, frameId string, url string) (*Frame, error) {
	_, err := page.EvaluateHandle(`async ({ frame_id, url }) => {
		const frame = document.createElement('iframe');
		frame.src = url;
		frame.id = frame_id;
		document.body.appendChild(frame);
		await new Promise(x => frame.onload = x);
		return frame;
	}`, map[string]interface{}{
		"frame_id": frameId,
		"url":      url,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (tu *testUtils) VerifyViewport(t *testing.T, page *Page, width, height int) {
	require.Equal(t, page.ViewportSize().Width, width)
	require.Equal(t, page.ViewportSize().Height, height)
	innerWidth, err := page.Evaluate("window.innerWidth")
	require.NoError(t, err)
	require.Equal(t, innerWidth, width)
	innerHeight, err := page.Evaluate("window.innerHeight")
	require.NoError(t, err)
	require.Equal(t, innerHeight, height)
}

func (tu *testUtils) AssertEval(t *testing.T, page *Page, script string, expected interface{}) {
	result, err := page.Evaluate(script)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}
