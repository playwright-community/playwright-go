package playwright_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPageRequestGC(t *testing.T) {
	BeforeEach(t)

	_, err := page.Evaluate(`() => {
		globalThis.objectToDestroy = { hello: 'world' };
		globalThis.weakRef = new WeakRef(globalThis.objectToDestroy);
	}`)
	require.NoError(t, err)

	require.NoError(t, page.RequestGC())
	ret, err := page.Evaluate(`() => globalThis.weakRef.deref()`)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"hello": "world"}, ret)

	require.NoError(t, page.RequestGC())
	ret, err = page.Evaluate(`() => globalThis.weakRef.deref()`)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"hello": "world"}, ret)

	_, err = page.Evaluate(`() => globalThis.objectToDestroy = null`)
	require.NoError(t, err)

	require.NoError(t, page.RequestGC())
	ret, err = page.Evaluate(`() => globalThis.weakRef.deref()`)
	require.NoError(t, err)
	require.Nil(t, ret)
}
