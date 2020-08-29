package playwright

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var pw *Playwright

func TestMain(m *testing.M) {
	var err error
	pw, err = Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	code := m.Run()
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	os.Exit(code)
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
}

func NewTestHelper(t *testing.T) *TestHelperData {
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
	require.NoError(t, err)
	context, err := browser.NewContext(BrowserNewContextOptions{
		AcceptDownloads: Bool(true),
	})
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	th := &TestHelperData{
		t:          t,
		Playwright: pw,
		Browser:    browser,
		Context:    context,
		Page:       page,
		IsChromium: browserName == "chromium" || browserName == "",
		IsFirefox:  browserName == "firefox",
		IsWebKit:   browserName == "webkit",
		server:     newTestServer(),
		assetDir:   "tests/assets/",
	}
	return th
}

func (t *TestHelperData) AfterEach() {
	if err := t.Browser.Close(); err != nil {
		t.t.Errorf("could not close browser: %v", err)
	}
	t.server.Stop()
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
	testServer          *httptest.Server
	routes              map[string]http.HandlerFunc
	requestSubscriberes map[string][]chan *http.Request
	PREFIX              string
	EMPTY_PAGE          string
}

func (t *testServer) Stop() {
	t.testServer.Close()
	t.routes = make(map[string]http.HandlerFunc)
	t.requestSubscriberes = make(map[string][]chan *http.Request)
}

func (t *testServer) serveHTTP(w http.ResponseWriter, r *http.Request) {
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
