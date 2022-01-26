package playwright_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

var pw *playwright.Playwright
var browser playwright.Browser
var context playwright.BrowserContext
var page playwright.Page
var isChromium bool
var isFirefox bool
var isWebKit bool
var browserName = getBrowserName()
var server *testServer
var browserType playwright.BrowserType
var utils *testUtils

func init() {
	if err := mime.AddExtensionType(".js", "application/javascript"); err != nil {
		log.Fatalf("could not add mime extension type: %v", err)
	}
}

func TestMain(m *testing.M) {
	BeforeAll()
	code := m.Run()
	AfterAll()
	os.Exit(code)
}

func BeforeAll() {
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	if browserName == "chromium" || browserName == "" {
		browserType = pw.Chromium
	} else if browserName == "firefox" {
		browserType = pw.Firefox
	} else if browserName == "webkit" {
		browserType = pw.WebKit
	}
	browser, err = browserType.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(os.Getenv("HEADFUL") == ""),
	})
	if err != nil {
		log.Fatalf("could not launch: %v", err)
	}
	isChromium = browserName == "chromium" || browserName == ""
	isFirefox = browserName == "firefox"
	isWebKit = browserName == "webkit"
	server = newTestServer()
	utils = &testUtils{}
}

func AfterAll() {
	server.testServer.Close()
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
}

var DEFAULT_CONTEXT_OPTIONS = playwright.BrowserNewContextOptions{
	AcceptDownloads: playwright.Bool(true),
	HasTouch:        playwright.Bool(true),
}

func BeforeEach(t *testing.T) {
	newContextWithOptions(t, DEFAULT_CONTEXT_OPTIONS)
}

func newContextWithOptions(t *testing.T, contextOptions playwright.BrowserNewContextOptions) {
	var err error
	context, err = browser.NewContext(contextOptions)
	require.NoError(t, err)
	page, err = context.NewPage()
	require.NoError(t, err)
}

func Asset(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %v", err)
	}
	return filepath.Join(cwd, "assets", path)
}

func AfterEach(t *testing.T, closeContext ...bool) {
	if len(closeContext) == 0 {
		if err := context.Close(); err != nil {
			t.Errorf("could not close context: %v", err)
		}
	}
	server.AfterEach()
}

func newTestServer() *testServer {
	ts := &testServer{
		routes:              make(map[string]http.HandlerFunc),
		requestSubscriberes: make(map[string][]chan *http.Request),
	}
	ts.testServer = httptest.NewServer(http.HandlerFunc(ts.serveHTTP))
	ts.PREFIX = ts.testServer.URL
	ts.EMPTY_PAGE = ts.testServer.URL + "/empty.html"
	ts.CROSS_PROCESS_PREFIX = strings.Replace(ts.testServer.URL, "127.0.0.1", "localhost", 1)
	return ts
}

type testServer struct {
	sync.Mutex
	testServer           *httptest.Server
	routes               map[string]http.HandlerFunc
	requestSubscriberes  map[string][]chan *http.Request
	PREFIX               string
	EMPTY_PAGE           string
	CROSS_PROCESS_PREFIX string
}

func (t *testServer) AfterEach() {
	t.Lock()
	defer t.Unlock()
	t.routes = make(map[string]http.HandlerFunc)
	t.requestSubscriberes = make(map[string][]chan *http.Request)
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
	w.Header().Add("Cache-Control", "no-cache, no-store")
	http.FileServer(http.Dir("assets")).ServeHTTP(w, r)
}

func (s *testServer) SetRoute(path string, f http.HandlerFunc) {
	s.Lock()
	defer s.Unlock()
	s.routes[path] = f
}

func (s *testServer) SetRedirect(from, to string) {
	s.SetRoute(from, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, http.StatusFound)
	})
}

func (s *testServer) WaitForRequestChan(path string) <-chan *http.Request {
	s.Lock()
	defer s.Unlock()
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

func (t *testUtils) AttachFrame(page playwright.Page, frameId string, url string) (playwright.Frame, error) {
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

func (tu *testUtils) VerifyViewport(t *testing.T, page playwright.Page, width, height int) {
	require.Equal(t, page.ViewportSize().Width, width)
	require.Equal(t, page.ViewportSize().Height, height)
	innerWidth, err := page.Evaluate("window.innerWidth")
	require.NoError(t, err)
	require.Equal(t, innerWidth, width)
	innerHeight, err := page.Evaluate("window.innerHeight")
	require.NoError(t, err)
	require.Equal(t, innerHeight, height)
}

func (tu *testUtils) AssertEval(t *testing.T, page playwright.Page, script string, expected interface{}) {
	result, err := page.Evaluate(script)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func getBrowserName() string {
	browserName, hasEnv := os.LookupEnv("BROWSER")
	if hasEnv {
		return browserName
	}
	return "chromium"
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
