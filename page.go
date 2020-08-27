package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"reflect"
)

type Page struct {
	ChannelOwner
	frames    []*Frame
	workers   []*Worker
	mainFrame *Frame
}

func (b *Page) Goto(url string) error {
	return b.mainFrame.Goto(url)
}

func (b *Page) URL() string {
	return b.mainFrame.URL()
}

func (b *Page) SetContent(content string, options ...PageSetContentOptions) error {
	return b.mainFrame.SetContent(content, options...)
}

func (b *Page) Content() (string, error) {
	return b.mainFrame.Content()
}

func (b *Page) Title() (string, error) {
	return b.mainFrame.Title()
}

func (b *Page) Workers() []*Worker {
	return b.workers
}

func (b *Page) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return b.mainFrame.Evaluate(expression, options...)
}

func (b *Page) EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
	return b.mainFrame.EvaluateOnSelector(selector, expression, options...)
}

func (b *Page) EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	return b.mainFrame.EvaluateOnSelectorAll(selector, expression, options...)
}

func (b *Page) Screenshot(options ...PageScreenshotOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := b.channel.Send("screenshot", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%v", err)
	}
	image, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%v", err)
	}
	if path != nil {
		if err := ioutil.WriteFile(*path, image, 0644); err != nil {
			return nil, err
		}
	}
	return image, nil
}

func (b *Page) PDF(options ...PagePdfOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := b.channel.Send("pdf", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%v", err)
	}
	pdf, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%v", err)
	}
	if path != nil {
		if err := ioutil.WriteFile(*path, pdf, 0644); err != nil {
			return nil, err
		}
	}
	return pdf, nil
}

func (b *Page) QuerySelector(selector string) (*ElementHandle, error) {
	return b.mainFrame.QuerySelector(selector)
}

func (p *Page) Click(selector string, options ...PageClickOptions) error {
	return p.mainFrame.Click(selector, options...)
}

func (p *Page) WaitForEvent(event string, predicate ...interface{}) interface{} {
	evChan := make(chan interface{}, 1)
	p.Once(event, func(ev ...interface{}) {
		if len(predicate) == 0 {
			evChan <- ev[0]
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
			}
		}
	})
	return <-evChan
}

func (p *Page) WaitForRequest(url interface{}, options ...interface{}) *Request {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *Request) bool {
		if matcher != nil {
			return matcher.Match(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("request", predicate).(*Request)
}

func (p *Page) WaitForResponse(url interface{}, options ...interface{}) *Response {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *Response) bool {
		if matcher != nil {
			return matcher.Match(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("response", predicate).(*Response)
}

func (p *Page) ExpectEvent(event string, cb func() error) (interface{}, error) {
	return newExpectWrapper(p.WaitForEvent, []interface{}{event}, cb)
}

func (p *Page) ExpectConsoleMessage(cb func() error) (*ConsoleMessage, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"console"}, cb)
	return response.(*ConsoleMessage), err
}

func (p *Page) ExpectedDialog(cb func() error) (*Download, error) {
	download, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return download.(*Download), err
}

func (p *Page) ExpectDownload(cb func() error) (*Download, error) {
	download, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return download.(*Download), err
}

func (p *Page) ExpectFileChooser(cb func() error) (*ConsoleMessage, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"console"}, cb)
	return response.(*ConsoleMessage), err
}

// func (p *Page) ExpectLoadState(cb func() error) (*ConsoleMessage, error) {
// 	response, err := newExpectWrapper(p.WaitForLoadState, []interface{}{"console"}, cb)
// 	return response.(*ConsoleMessage), err
// }

// func (p *Page) ExpectNavigation(cb func() error) (*ConsoleMessage, error) {
// 	response, err := newExpectWrapper(p.WaitForNavigation, []interface{}{"console"}, cb)
// 	return response.(*ConsoleMessage), err
// }

func (p *Page) ExpectPopup(cb func() error) (*Page, error) {
	popup, err := newExpectWrapper(p.WaitForEvent, []interface{}{"popup"}, cb)
	return popup.(*Page), err
}

func (p *Page) ExpectResponse(url interface{}, cb func() error, options ...interface{}) (*Response, error) {
	response, err := newExpectWrapper(p.WaitForResponse, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return response.(*Response), err
}

func (p *Page) ExpectRequest(url interface{}, cb func() error, options ...interface{}) (*Request, error) {
	popup, err := newExpectWrapper(p.WaitForRequest, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return popup.(*Request), err
}

func (p *Page) ExpectWorker(cb func() error) (*Worker, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"worker"}, cb)
	return response.(*Worker), err
}

func newPage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Page {
	channelOwner := initializer["mainFrame"].(*Channel).object
	bt := &Page{
		mainFrame: channelOwner.(*Frame),
		workers:   make([]*Worker, 0),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("console", func(payload ...interface{}) {
		bt.Emit("console", payload[0].(map[string]interface{})["message"].(*Channel).object)
	})
	bt.channel.On("download", func(payload ...interface{}) {
		bt.Emit("download", payload[0].(map[string]interface{})["download"].(*Channel).object)
	})
	bt.channel.On("popup", func(payload ...interface{}) {
		bt.Emit("popup", payload[0].(map[string]interface{})["page"].(*Channel).object)
	})
	bt.channel.On("request", func(payload ...interface{}) {
		bt.Emit("request", payload[0].(map[string]interface{})["request"].(*Channel).object)
	})
	bt.channel.On("requestFailed", func(payload ...interface{}) {
		bt.Emit("requestFailed", payload[0].(map[string]interface{})["request"].(*Channel).object, payload[0].(map[string]interface{})["failureText"])
	})
	bt.channel.On("requestFinished", func(payload ...interface{}) {
		bt.Emit("requestFinished", payload[0].(map[string]interface{})["request"].(*Channel).object)
	})
	bt.channel.On("response", func(payload ...interface{}) {
		bt.Emit("response", payload[0].(map[string]interface{})["response"].(*Channel).object)
	})
	bt.channel.On("route", func(payload ...interface{}) {
		bt.Emit("route", payload[0].(map[string]interface{})["request"].(*Channel).object)
	})
	bt.channel.On("worker", func(payload ...interface{}) {
		worker := payload[0].(map[string]interface{})["worker"].(*Channel).object.(*Worker)
		worker.page = bt
		bt.workers = append(bt.workers, worker)
		bt.Emit("worker", worker)
	})
	return bt
}
