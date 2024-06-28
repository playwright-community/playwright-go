package playwright_test

import (
	"testing"

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

func beforePageClock(t *testing.T, installTime, pauseAtTime int) {
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
		require.Equal(t, 1, calls.Len())
	})

	t.Run("does not trigger without sufficient delay", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(10))
		require.Equal(t, 0, calls.Len())
	})

	t.Run("triggers after sufficient delay", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
		require.Equal(t, 1, calls.Len())
	})

	t.Run("triggers simultaneous timers", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 100); setTimeout(window.stub, 100)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(100))
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
		require.Equal(t, 4, calls.Len())
	})

	t.Run("waits after setTimeout was called", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(window.stub, 150)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor(50))
		require.Equal(t, 0, calls.Len())
		require.NoError(t, page.Clock().RunFor(100))
		require.Equal(t, 1, calls.Len())
	})

	t.Run("triggers event when some throw", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { throw new Error(); }, 100); setTimeout(window.stub, 120)")
		require.NoError(t, err)
		require.Error(t, page.Clock().RunFor(120))
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
		require.Equal(t, 2, calls.Len())
	})

	t.Run("passes 1 minute", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 6000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor("01:00"))
		require.Equal(t, 10, calls.Len())
	})

	t.Run("passes 2 hours 34 minutes and 10 seconds", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 10000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().RunFor("02:34:10"))
		require.Equal(t, 925, calls.Len())
	})

	t.Run("throws for invalid format", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setInterval(window.stub, 10000)")
		require.NoError(t, err)
		require.Error(t, page.Clock().RunFor("12:02:34:10"))
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
		require.Equal(t, 0, calls.Len())
	})

	t.Run("pushes back exeution time for skipped timers", func(t *testing.T) {
		BeforeEach(t)

		calls := pageClockFixture(t)
		beforePageClock(t, 0, 1000)

		_, err := page.Evaluate("setTimeout(() => { window.stub(Date.now()); }, 1000)")
		require.NoError(t, err)
		require.NoError(t, page.Clock().FastForward(2000))
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
		require.Equal(t, [][]any{
			{1000 + 110000},
		}, calls.Get())
	})
}
