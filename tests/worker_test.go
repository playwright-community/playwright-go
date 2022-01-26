package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestWorkerShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	worker, err := page.ExpectWorker(func() error {
		_, err := page.Goto(server.PREFIX + "/worker/worker.html")
		return err
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(page.Workers()))
	require.Equal(t, worker, page.Workers()[0])
	worker = page.Workers()[0]
	require.Contains(t, worker.URL(), "worker.js")
	res, err := worker.Evaluate(`() => self["workerFunction"]()`)
	require.NoError(t, err)
	require.Equal(t, "worker function result", res)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, 0, len(page.Workers()))
}

func TestWorkerShouldEmitCreatedAndDestroyedEvents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	var workerObj playwright.JSHandle
	worker, err := page.ExpectWorker(func() error {
		workerObjInterface, err := page.EvaluateHandle("() => new Worker(URL.createObjectURL(new Blob(['1'], {type: 'application/javascript'})))")
		workerObj = workerObjInterface
		return err
	})
	require.NoError(t, err)
	_, err = worker.ExpectEvent("close", func() error {
		_, err := page.Evaluate("workerObj => workerObj.terminate()", workerObj)
		return err
	})
	require.NoError(t, err)
}

func TestWorkerShouldReportConsoleLogs(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	message, err := page.ExpectEvent("console", func() error {
		_, err := page.Evaluate("() => new Worker(URL.createObjectURL(new Blob(['console.log(1)'], {type: 'application/javascript'})))")
		return err
	})
	require.NoError(t, err)
	require.Equal(t, message.(playwright.ConsoleMessage).Text(), "1")
}

func TestWorkerShouldHaveJSHandlesForConsoleLogs(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	message, err := page.ExpectEvent("console", func() error {
		_, err := page.Evaluate("() => new Worker(URL.createObjectURL(new Blob(['console.log(1,2,3,this)'], {type: 'application/javascript'})))")
		return err
	})
	require.NoError(t, err)
	log := message.(playwright.ConsoleMessage)
	if !isFirefox {
		require.Equal(t, "1 2 3 DedicatedWorkerGlobalScope", log.Text())
	} else {
		require.Equal(t, "1 2 3 JSHandle@object", log.Text())
	}
	require.Equal(t, 4, len(log.Args()))
	origin, err := log.Args()[3].GetProperty("origin")
	require.NoError(t, err)
	val, err := origin.JSONValue()
	require.NoError(t, err)
	require.Equal(t, "null", val)
}

func TestWorkerShouldEvaluate(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	worker, err := page.ExpectWorker(func() error {
		_, err := page.Evaluate("() => new Worker(URL.createObjectURL(new Blob(['console.log(1)'], {type: 'application/javascript'})))")
		return err
	})
	require.NoError(t, err)
	result, err := worker.Evaluate("1+1")
	require.NoError(t, err)
	require.Equal(t, 2, result)
}

func TestWorkershouldReportErrors(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	pageError, err := page.ExpectEvent("pageerror", func() error {
		_, err := page.Evaluate("() => new Worker(URL.createObjectURL(new Blob([`\n" +
			"  setTimeout(() => {\n" +
			"    // Do a console.log just to check that we do not confuse it with an error.\n" +
			"    console.log('hey');\n" +
			"    throw new Error('this is my error');\n" +
			"  })\n" +
			"`], {type: 'application/javascript'})))")
		return err
	})
	require.NoError(t, err)
	require.Contains(t, pageError.(*playwright.Error).Error(), "this is my error")
}

func TestWorkerShouldClearUponCrossProcessNavigation(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	worker, err := page.ExpectWorker(func() error {
		_, err := page.Evaluate("() => new Worker(URL.createObjectURL(new Blob(['console.log(1)'], {type: 'application/javascript'})))")
		return err
	})
	require.NoError(t, err)
	require.Equal(t, len(page.Workers()), 1)
	destroyed := false
	worker.On("close", func() {
		destroyed = true
	})
	_, err = page.Goto(server.CROSS_PROCESS_PREFIX + "/empty.html")
	require.NoError(t, err)
	require.True(t, destroyed)
	require.Equal(t, 0, len(page.Workers()))
}
