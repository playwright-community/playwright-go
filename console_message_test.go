package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConsoleShouldWork(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messages := make(chan *ConsoleMessage, 1)
	helper.Page.Once("console", func(message *ConsoleMessage) {
		messages <- message
	})
	_, err := helper.Page.Evaluate(`() => console.log("hello", 5, {foo: "bar"})`)
	require.NoError(t, err)
	message := <-messages
	require.Equal(t, message.Text(), "hello 5 JSHandle@object")
	require.Equal(t, message.String(), "hello 5 JSHandle@object")
	require.Equal(t, message.Type(), "log")
	jsonValue1, err := message.Args()[0].JSONValue()
	require.NoError(t, err)
	require.Equal(t, "hello", jsonValue1)
	jsonValue2, err := message.Args()[1].JSONValue()
	require.NoError(t, err)
	require.Equal(t, 5, jsonValue2)
	jsonValue3, err := message.Args()[2].JSONValue()
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{
		"foo": "bar",
	}, jsonValue3)
}

func TestConsoleShouldEmitSameLogTwice(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messages := make(chan string, 2)
	helper.Page.On("console", func(message *ConsoleMessage) {
		messages <- message.Text()
	})
	_, err := helper.Page.Evaluate(`() => { for (let i = 0; i < 2; ++i ) console.log("hello"); } `)
	require.NoError(t, err)
	m1 := <-messages
	m2 := <-messages
	require.Equal(t, []string{"hello", "hello"}, []string{m1, m2})
}

func TestConsoleShouldUseTextForStr(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messages := make(chan *ConsoleMessage, 1)
	helper.Page.On("console", func(message *ConsoleMessage) {
		messages <- message
	})
	_, err := helper.Page.Evaluate(`() => console.log("Hello world")`)
	require.NoError(t, err)
	message := <-messages
	require.Equal(t, "Hello world", message.String())
}

func TestConsoleShouldWorkForDifferentConsoleAPICalls(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messagesChan := make(chan *ConsoleMessage, 6)
	helper.Page.On("console", func(message *ConsoleMessage) {
		messagesChan <- message
	})
	// All console events will be reported before 'page.evaluate' is finished.
	_, err := helper.Page.Evaluate(
		`() => {
      // A pair of time/timeEnd generates only one Console API call.
      console.time('calling console.time');
      console.timeEnd('calling console.time');
      console.trace('calling console.trace');
      console.dir('calling console.dir');
      console.warn('calling console.warn');
      console.error('calling console.error');
      console.log(Promise.resolve('should not wait until resolved!'));
	}`)
	messages := ChanToSlice(messagesChan, 6).([]*ConsoleMessage)
	require.NoError(t, err)
	require.Equal(t, []interface{}{
		"timeEnd",
		"trace",
		"dir",
		"warning",
		"error",
		"log",
	}, Map(messages, func(msg interface{}) interface{} {
		return msg.(*ConsoleMessage).Type()
	}))

	require.Contains(t, messages[0].Text(), "calling console.time")
	require.Equal(t, []interface{}{
		"calling console.trace",
		"calling console.dir",
		"calling console.warn",
		"calling console.error",
		"JSHandle@promise",
	}, Map(messages[1:], func(msg interface{}) interface{} {
		return msg.(*ConsoleMessage).Text()
	}))
}

func TestConsoleShouldNotFailForWindowObjects(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messages := make(chan *ConsoleMessage, 1)
	helper.Page.Once("console", func(message *ConsoleMessage) {
		messages <- message
	})
	_, err := helper.Page.Evaluate("() => console.error(window)")
	require.NoError(t, err)
	message := <-messages
	require.Equal(t, "JSHandle@object", message.Text())
}

func TestConsoleShouldTriggerCorrectLog(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messages := make(chan *ConsoleMessage, 1)
	helper.Page.Once("console", func(message *ConsoleMessage) {
		messages <- message
	})
	_, err := helper.Page.Goto("about:blank")
	require.NoError(t, err)
	_, err = helper.Page.Evaluate("url => fetch(url).catch(e => {})", helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	message := <-messages
	require.Contains(t, message.Text(), "Access-Control-Allow-Origin")
	require.Equal(t, "error", message.Type())
}

func TestConsoleShouldHaveLocation(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	messageEvent, err := helper.Page.ExpectEvent("console", func() error {
		_, err := helper.Page.Goto(helper.server.PREFIX + "/consolelog.html")
		return err
	}, func(m *ConsoleMessage) bool {
		return m.Text() == "yellow"
	})
	require.NoError(t, err)
	message := messageEvent.(*ConsoleMessage)
	require.Equal(t, message.Type(), "log")
	require.Equal(t, helper.server.PREFIX+"/consolelog.html", message.Location().URL)
	require.Equal(t, 7, message.Location().LineNumber)
}
