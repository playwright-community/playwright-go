package playwright

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Map(vs interface{}, f func(interface{}) interface{}) []interface{} {
	v := reflect.ValueOf(vs)
	vsm := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		vsm[i] = f(v.Index(i).Interface())
	}
	return vsm
}

func TestConsoleShouldWork(t *testing.T) {
	helper := NewTestHelper(t)
	messages := make(chan *ConsoleMessage)
	helper.Page.Once("console", func(args ...interface{}) {
		messages <- args[0].(*ConsoleMessage)
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
	helper.Browser.Close()
}

func TestConsoleShouldEmitSameLogTwice(t *testing.T) {
	helper := NewTestHelper(t)
	messages := make(chan string, 2)
	helper.Page.On("console", func(args ...interface{}) {
		messages <- args[0].(*ConsoleMessage).Text()
	})
	_, err := helper.Page.Evaluate(`() => { for (let i = 0; i < 2; ++i ) console.log("hello"); } `)
	require.NoError(t, err)
	m1 := <-messages
	m2 := <-messages
	require.Equal(t, []string{"hello", "hello"}, []string{m1, m2})
	helper.Browser.Close()
}

func TestConsoleShouldUseTextForStr(t *testing.T) {
	helper := NewTestHelper(t)
	messages := []*ConsoleMessage{}
	helper.Page.On("console", func(args ...interface{}) {
		messages = append(messages, args[0].(*ConsoleMessage))
	})
	_, err := helper.Page.Evaluate(`() => console.log("Hello world")`)
	require.NoError(t, err)
	require.Len(t, messages, 1)
	require.Equal(t, "Hello world", messages[0].String())
	helper.Browser.Close()
}

func TestConsoleShouldWorkForDifferentConsoleAPICalls(t *testing.T) {
	helper := NewTestHelper(t)
	messages := []*ConsoleMessage{}
	helper.Page.On("console", func(args ...interface{}) {
		messages = append(messages, args[0].(*ConsoleMessage))
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
	helper.Browser.Close()
}

func TestConsoleShouldNotFailForWindowObjects(t *testing.T) {
	helper := NewTestHelper(t)
	messages := make(chan *ConsoleMessage)
	helper.Page.Once("console", func(args ...interface{}) {
		messages <- args[0].(*ConsoleMessage)
	})
	_, err := helper.Page.Evaluate("() => console.error(window)")
	require.NoError(t, err)
	message := <-messages
	require.Equal(t, "JSHandle@object", message.Text())
	helper.Browser.Close()
}

func TestConsoleShouldTriggerCorrectLog(t *testing.T) {
	helper := NewTestHelper(t)
	messages := make(chan *ConsoleMessage)
	helper.Page.Once("console", func(args ...interface{}) {
		messages <- args[0].(*ConsoleMessage)
	})
	require.NoError(t, helper.Page.Goto("about:blank"))
	// TODO: use server
	_, err := helper.Page.Evaluate("url => fetch(url).catch(e => {})", "https://www.test-cors.org/")
	require.NoError(t, err)
	message := <-messages
	require.Contains(t, message.Text(), "Access-Control-Allow-Origin")
	require.Equal(t, "error", message.Type())
	helper.Browser.Close()
}
