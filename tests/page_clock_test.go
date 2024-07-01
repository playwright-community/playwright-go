package playwright_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func pageClockFixture(t *testing.T) *syncSlice[[]interface{}] {
	t.Helper()
	calls := newSyncSlice[[]interface{}]()
	err := page.ExposeFunction("stub", func(args ...interface{}) interface{} {
		calls.Append(args)
		return nil
	})
	require.NoError(t, err)
	return calls
}

func beforePageClock(t *testing.T, installTime, pauseAtTime interface{}) {
	require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{
		Time: installTime,
	}))
	require.NoError(t, page.Clock().PauseAt(pauseAtTime))
}

func TestPageClockRunFor(t *testing.T) {
	t.Run("tiggers immediately without specified delay", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(0))
		time.Sleep(100 * time.Millisecond) // wait for binding call to resolve
		require.Equal(t, 1, calls.Len())
	})

	t.Run("does not trigger without sufficient delay", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(10))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
	})

	t.Run("triggers after sufficient delay", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 1, calls.Len())
	})

	t.Run("triggers simultaneous timers", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100); setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 2, calls.Len())
	})

	t.Run("triggers multiple simultaneous timers", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate(
			"setTimeout(window.stub, 100); setTimeout(window.stub, 100); setTimeout(window.stub, 99); setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 4, calls.Len())
	})

	t.Run("waits after setTimeout was called", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 150)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(50))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
		require.NoError(t, page.Clock().RunFor(100))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 1, calls.Len())
	})

	t.Run("triggers event when some throw", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { throw new Error(); }, 100); setTimeout(window.stub, 120)")
		require.NoError(t, err)
		require.Error(t, page.Clock().RunFor(120))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 1, calls.Len())
	})

	t.Run("creates updated Date while ticking", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		require.NoError(t, page.Clock().SetSystemTime(0))
		_, err := page.Evaluate("setInterval(() => { window.stub(new Date().getTime()); }, 10)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
		require.Equal(t, [][]interface{}{
			{10},
			{20},
			{30},
			{40},
			{50},
			{60},
			{70},
			{80},
			{90},
			{100},
		}, calls.Get())
	})

	t.Run("passes 8 seconds", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 4000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor("08"))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 2, calls.Len())
	})

	t.Run("passes 1 minute", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 6000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor("01:00"))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 10, calls.Len())
	})

	t.Run("passes 2 hours 34 minutes and 10 seconds", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 10000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor("02:34:10"))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 925, calls.Len())
	})

	t.Run("throws for invalid format", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 10000)")
		require.NoError(t, err)
		require.Error(t, page.Clock().RunFor("12:02:34:10"))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
	})

	t.Run("returns the current now value", func(t *testing.T) {
		BeforeEach(t)

		beforePageClock(t, 0, 1000)

		require.NoError(t, page.Clock().SetSystemTime(0))
		value := 200
		require.NoError(t, page.Clock().RunFor(value))
		ret, err := page.Evaluate("Date.now()")
		require.NoError(t, err)
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, value, ret)
	})
}

func TestPageClockFastForward(t *testing.T) {
	t.Run("ignores timers which wouldnt be run", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { window.stub('should not be logged'); }, 1000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward(500))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
	})

	t.Run("pushes back exeution time for skipped timers", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { window.stub(Date.now()); }, 1000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward(2000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{
			{1000 + 2000},
		}, calls.Get())
	})

	t.Run("supports string time arguments", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { window.stub(Date.now()); }, 100000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward("01:50"))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{
			{1000 + 110000},
		}, calls.Get())
	})
}

func TestPageClockStubTimers(t *testing.T) {
	t.Run("sets initial timestamp", func(t *testing.T) {
		BeforeEach(t)

		beforePageClock(t, 0, 1000)

		require.NoError(t, page.Clock().SetSystemTime(1400))
		ret, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 1400, ret)
	})

	t.Run("replaces global setTimeout", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate(`setTimeout(window.stub, 1000)`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 1, calls.Len())
	})

	t.Run("global fake setTimeout should return id", func(t *testing.T) {
		BeforeEach(t)

		_ = pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		ret, err := page.Evaluate(`setTimeout(window.stub, 1000)`)
		require.NoError(t, err)
		require.IsType(t, int(1), ret)
	})

	t.Run("replaces global clearTimeout", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate(`const to = setTimeout(window.stub, 1000);
      clearTimeout(to);`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
	})

	t.Run("replaces global setInterval", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate(`setInterval(window.stub, 500)`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 2, calls.Len())
	})

	t.Run("replaces global clearInterval", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate(`const to = setInterval(window.stub, 500);
      clearInterval(to);`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, 0, calls.Len())
	})

	t.Run("replaces global performance now", func(t *testing.T) {
		BeforeEach(t)

		beforePageClock(t, 0, 1000)

		chanRet := make(chan interface{}, 1)
		go func() {
			ret, err := page.Evaluate(`
			async () => {
				const prev = performance.now();
				await new Promise(f => setTimeout(f, 1000));
				const next = performance.now();
				return { prev, next };
			}`)
			if err == nil {
				chanRet <- ret
			} else {
				close(chanRet)
			}
		}()
		require.NoError(t, page.Clock().RunFor(1000))
		ret := <-chanRet
		require.Equal(t, map[string]interface{}{
			"prev": 1000,
			"next": 2000,
		}, ret)
	})

	t.Run("fakes Date constructor", func(t *testing.T) {
		BeforeEach(t)

		beforePageClock(t, 0, 1000)

		ret, err := page.Evaluate(`new Date().getTime()`)
		require.NoError(t, err)
		require.Equal(t, 1000, ret)
	})
}

func TestPageClockStubTimersPerformance(t *testing.T) {
	t.Run("replaces global performance time origin", func(t *testing.T) {
		BeforeEach(t)

		beforePageClock(t, 1000, 2000)

		chanRet := make(chan interface{}, 1)
		go func() {
			ret, err := page.Evaluate(`
			async () => {
				const prev = performance.now();
				await new Promise(f => setTimeout(f, 1000));
				const next = performance.now();
				return { prev, next };
			}`)
			if err == nil {
				chanRet <- ret
			} else {
				close(chanRet)
			}
		}()
		require.NoError(t, page.Clock().RunFor(1000))
		origin, err := page.Evaluate(`performance.timeOrigin`)
		require.NoError(t, err)
		require.Equal(t, 1000, origin)
		ret := <-chanRet
		require.Equal(t, map[string]interface{}{
			"prev": 1000,
			"next": 2000,
		}, ret)
	})
}

func TestPageClockPopup(t *testing.T) {
	t.Run("should tick after popup", func(t *testing.T) {
		BeforeEach(t)

		now := time.Date(2015, 9, 25, 0, 0, 0, 0, time.UTC)
		beforePageClock(t, 0, now)

		var popup playwright.Page
		page.OnPopup(func(d playwright.Page) {
			popup = d
		})
		_, _ = page.Evaluate(`window.open('about:blank')`)
		popupTime, _ := popup.Evaluate(`Date.now()`)
		require.Equal(t, int(now.UnixMilli()), popupTime)
		require.NoError(t, page.Clock().RunFor(1000))
		popupTimeAfter, _ := popup.Evaluate(`Date.now()`)
		require.Equal(t, int(now.UnixMilli())+1000, popupTimeAfter)
	})

	t.Run("should tick before popup", func(t *testing.T) {
		BeforeEach(t)

		now := time.Date(2015, 9, 25, 0, 0, 0, 0, time.UTC)
		beforePageClock(t, 0, now)
		require.NoError(t, page.Clock().RunFor(1000))

		var popup playwright.Page
		page.OnPopup(func(d playwright.Page) {
			popup = d
		})
		_, _ = page.Evaluate(`window.open('about:blank')`)
		popupTime, _ := popup.Evaluate(`Date.now()`)
		require.Equal(t, int(now.UnixMilli())+1000, popupTime)
	})

	t.Run("should run time before popup", func(t *testing.T) {
		BeforeEach(t)

		server.SetRoute("/popup.html", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			_, err := w.Write([]byte(`<script>window.time = Date.now()</script>`))
			if err != nil {
				log.Printf("could not write: %v", err)
			}
		})

		_, err := page.Goto(server.EMPTY_PAGE)
		require.NoError(t, err)

		popupChan := make(chan playwright.Page, 1)
		page.OnPopup(func(d playwright.Page) {
			popupChan <- d
		})

		_, err = page.Evaluate(fmt.Sprintf(`window.open('%s/popup.html')`, server.PREFIX))
		require.NoError(t, err)
		popup := <-popupChan
		popupTime, _ := popup.Evaluate(`window.time`)
		require.GreaterOrEqual(t, popupTime, 2000)
	})

	t.Run("should not run time before popup on pause", func(t *testing.T) {
		BeforeEach(t)

		server.SetRoute("/popup.html", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			_, err := w.Write([]byte(`<script>window.time = Date.now()</script>`))
			if err != nil {
				log.Printf("could not write: %v", err)
			}
		})

		beforePageClock(t, 0, 1000)

		_, err := page.Goto(server.EMPTY_PAGE)
		require.NoError(t, err)
		popupChan := make(chan playwright.Page, 1)
		page.OnPopup(func(d playwright.Page) {
			popupChan <- d
		})

		_, err = page.Evaluate(fmt.Sprintf(`window.open('%s/popup.html')`, server.PREFIX))
		require.NoError(t, err)
		popup := <-popupChan
		popupTime, _ := popup.Evaluate(`window.time`)
		require.Equal(t, 1000, popupTime)
	})
}

func TestPageClockFixedTime(t *testing.T) {
	t.Run("does not fake methods", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().SetFixedTime(0))
		// Should not stall.
		_, err := page.Evaluate(`new Promise(f => setTimeout(f, 1))`)
		require.NoError(t, err)
	})

	t.Run("allows setting time multiple times", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().SetFixedTime(100))
		ret, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 100, ret)

		require.NoError(t, page.Clock().SetFixedTime(200))
		ret, err = page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 200, ret)
	})

	t.Run("fixed times is not affected by clock manipulation", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().SetFixedTime(100))
		ret, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 100, ret)

		require.NoError(t, page.Clock().FastForward(20))
		ret, err = page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 100, ret)
	})

	t.Run("allows installing fake timers after setting time", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)

		require.NoError(t, page.Clock().SetFixedTime(100))
		ret, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 100, ret)

		require.NoError(t, page.Clock().SetFixedTime(200))
		_, err = page.Evaluate(`setTimeout(() => window.stub(Date.now()))`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(0))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]interface{}{
			{200},
		}, calls.Get())
	})
}

func TestPageClockWhileRunning(t *testing.T) {
	t.Run("should progress time", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		//nolint:staticcheck
		page.WaitForTimeout(1000)
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.LessOrEqual(t, 1000, now)
		require.LessOrEqual(t, now, 2000)
	})

	t.Run("should run for", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(10000))
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.LessOrEqual(t, 10000, now)
		require.LessOrEqual(t, now, 11000)
	})

	t.Run("should fast forward", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward(10000))
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.LessOrEqual(t, 10000, now)
		require.LessOrEqual(t, now, 11000)
	})

	t.Run("should pause", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		//nolint:staticcheck
		page.WaitForTimeout(1000)
		require.NoError(t, page.Clock().Resume())
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.LessOrEqual(t, 0, now)
		require.LessOrEqual(t, now, 1000)
	})

	t.Run("should pause and fast forward", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		require.NoError(t, page.Clock().FastForward(1000))
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 2000, now)
	})

	t.Run("should set system time on pause", func(t *testing.T) {
		BeforeEach(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		now, err := page.Evaluate(`Date.now()`)
		require.NoError(t, err)
		require.Equal(t, 1000, now)
	})
}

func TestPageClockWhileOnPause(t *testing.T) {
	t.Run("fast forward should not run nested immediate", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		_, err = page.Evaluate(`
			setTimeout(() => {
					window.stub('outer');
					setTimeout(() => window.stub('inner'), 0);
			}, 1000);`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}}, calls.Get())
		require.NoError(t, page.Clock().FastForward(1))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}, {"inner"}}, calls.Get())
	})

	t.Run("run for should not run nested immediate", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		_, err = page.Evaluate(`
			setTimeout(() => {
					window.stub('outer');
					setTimeout(() => window.stub('inner'), 0);
			}, 1000);`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}}, calls.Get())
		require.NoError(t, page.Clock().RunFor(1))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}, {"inner"}}, calls.Get())
	})

	t.Run("run for should not run nested immediate from microtask", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)

		require.NoError(t, page.Clock().Install(playwright.ClockInstallOptions{Time: 0}))
		_, err := page.Goto("data:text/html,")
		require.NoError(t, err)
		require.NoError(t, page.Clock().PauseAt(1000))
		_, err = page.Evaluate(`
			setTimeout(() => {
					window.stub('outer');
					void Promise.resolve().then(() => setTimeout(() => window.stub('inner'), 0));
			}, 1000);`)
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(1000))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}}, calls.Get())
		require.NoError(t, page.Clock().RunFor(1))
		time.Sleep(100 * time.Millisecond)
		require.Equal(t, [][]any{{"outer"}, {"inner"}}, calls.Get())
	})
}
