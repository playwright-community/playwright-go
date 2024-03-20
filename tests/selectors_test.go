//nolint:staticcheck
package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestSelectorsRegisterShouldWork(t *testing.T) {
	BeforeEach(t)

	tagSelector := `
	{
		create(root, target) {
			return target.nodeName;
		},
		query(root, selector) {
			return root.querySelector(selector);
		},
		queryAll(root, selector) {
			return Array.from(root.querySelectorAll(selector));
		}
	}
	`
	selectorName := "tag_" + browserName
	selector2Name := "tag2_" + browserName

	err := pw.Selectors.Register(selectorName, playwright.Script{})
	require.ErrorContains(t, err, `Either source or path should be specified`)
	// Register one engine before creating context.
	err = pw.Selectors.Register(selectorName, playwright.Script{
		Content: &tagSelector,
	})
	require.NoError(t, err)

	// Register another engine after creating context.
	err = pw.Selectors.Register(selector2Name, playwright.Script{
		Content: &tagSelector,
	})
	require.NoError(t, err)

	require.NoError(t, page.SetContent(`<div><span></span></div><div></div>`))

	ret, err := page.EvalOnSelector(selectorName+"=DIV", `e => e.nodeName`, nil)
	require.NoError(t, err)
	require.Equal(t, "DIV", ret)
	ret, err = page.EvalOnSelector(selectorName+"=SPAN", `e => e.nodeName`, nil)
	require.NoError(t, err)
	require.Equal(t, "SPAN", ret)
	ret, err = page.EvalOnSelectorAll(selectorName+"=DIV", `es => es.length`, nil)
	require.NoError(t, err)
	require.Equal(t, 2, ret)

	ret, err = page.EvalOnSelector(selector2Name+"=DIV", `e => e.nodeName`, nil)
	require.NoError(t, err)
	require.Equal(t, "DIV", ret)
	ret, err = page.EvalOnSelector(selector2Name+"=SPAN", `e => e.nodeName`, nil)
	require.NoError(t, err)
	require.Equal(t, "SPAN", ret)
	ret, err = page.EvalOnSelectorAll(selector2Name+"=DIV", `es => es.length`)
	require.NoError(t, err)
	require.Equal(t, 2, ret)

	// Selector names are case-sensitive.
	_, err = page.QuerySelector("tAG=DIV")
	require.ErrorContains(t, err, `Unknown engine "tAG" while parsing selector tAG=DIV`)

	require.NoError(t, context.Close())
}

func TestSelectorsShouldUseDataTestIdInStrictErrors(t *testing.T) {
	BeforeEach(t)

	pw.Selectors.SetTestIdAttribute("data-custom-id")
	require.NoError(t, page.SetContent(`
	<div>
		<div></div>
		<div>
			<div></div>
			<div></div>
		</div>
	</div>
	<div>
		<div class='foo bar:0' data-custom-id='One'>
		</div>
		<div class='foo bar:1' data-custom-id='Two'>
		</div>
	</div>`))

	err := page.Locator(".foo").Hover(playwright.LocatorHoverOptions{
		Timeout: playwright.Float(500),
	})
	require.ErrorContains(t, err, "strict mode violation")
	require.ErrorContains(t, err, `<div class="foo bar:0`)
	require.ErrorContains(t, err, `<div class="foo bar:1`)
	require.ErrorContains(t, err, `aka getByTestId('One')`)
}

func TestSelectorsShouldWorkWithPath(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, pw.Selectors.Register("foo", playwright.Script{
		Path: playwright.String(Asset("sectionselectorengine.js")),
	}))
	require.NoError(t, page.SetContent(`<section></section>`))

	ret, err := page.EvalOnSelector("foo=whatever", `e => e.nodeName`, nil)
	require.NoError(t, err)
	require.Equal(t, "SECTION", ret)
}
