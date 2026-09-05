package playwright

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// The frame's waits used to read the state they wait on, then subscribe to the
// event that updates it; an event the dispatch goroutine delivered in between
// reached no listener and was never replayed. These tests aim the event at
// that gap: setNavigationWaiter registers the page's "framedetached" rejection
// immediately before the wait subscribes, so a goroutine that emits the
// moment that listener appears lands in the window often enough that the old
// code failed within a few rounds under -race (and most runs without it); the
// fixed code re-checks after subscribing and never loses one.

const raceRounds = 200

// raceTimeout bounds each round's wait. A round that hits the gap returns at
// once, so the budget is only ever paid by a lost wakeup; it is generous so
// that a stalled CI runner cannot turn a slow round into a failure.
var raceTimeout = Float(2000)

// raceRoundsNeedParallelism gives the emitting goroutine its own P: at
// GOMAXPROCS=1 the waiting goroutine does not yield until it blocks in Wait(),
// which is after it has subscribed, and the old code passes trivially.
func raceRoundsNeedParallelism(t *testing.T) {
	t.Helper()
	if runtime.GOMAXPROCS(0) >= 2 {
		return
	}
	prev := runtime.GOMAXPROCS(2)
	t.Cleanup(func() { runtime.GOMAXPROCS(prev) })
}

// emitWhenSubscribing runs emit on its own goroutine once the wait under test
// has registered its page rejections, i.e. is past its first look at the
// frame's state and about to subscribe. The returned channel closes when emit
// has returned, so a round cannot leave an emit in flight for the next one.
func emitWhenSubscribing(page *pageImpl, emit func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for page.ListenerCount("framedetached") == 0 {
			runtime.Gosched()
		}
		emit()
	}()
	return done
}

func TestWaitForLoadStateSeesAStateRecordedBeforeItSubscribed(t *testing.T) {
	raceRoundsNeedParallelism(t)
	_, frame, _ := newTimeoutSemanticsFixture(t, 200)
	for round := range raceRounds {
		frame.loadStates.Clear()
		done := emitWhenSubscribing(frame.page, func() {
			frame.onLoadState(map[string]any{"add": "load"})
		})
		require.NoError(t, frame.waitForLoadStateImpl("load", raceTimeout), "round %d", round)
		<-done
		require.Zero(t, frame.ListenerCount("loadstate"), "round %d: listener left behind", round)
		require.Zero(t, frame.page.ListenerCount(""), "round %d: page listener left behind", round)
	}
}

func TestWaitForURLSeesANavigationCommittedBeforeItSubscribed(t *testing.T) {
	raceRoundsNeedParallelism(t)
	_, frame, _ := newTimeoutSemanticsFixture(t, 200)
	frame.page.browserContext.options = &BrowserNewContextOptions{}
	const target = "https://example.com/next"
	for round := range raceRounds {
		frame.loadStates.Clear()
		frame.Lock()
		frame.url = "about:blank"
		frame.Unlock()
		done := emitWhenSubscribing(frame.page, func() {
			frame.onFrameNavigated(map[string]any{"url": target, "name": ""})
			frame.onLoadState(map[string]any{"add": "load"})
		})
		require.NoError(t, frame.WaitForURL(target, FrameWaitForURLOptions{Timeout: raceTimeout}), "round %d", round)
		<-done
		require.Equal(t, target, frame.URL())
		require.Zero(t, frame.ListenerCount(""), "round %d: frame listener left behind", round)
		require.Zero(t, frame.page.ListenerCount(""), "round %d: page listener left behind", round)
	}
}
