package playwright

import (
	"net/url"
	"path"
)

type pageAssertionsImpl struct {
	assertionsBase
	actualPage Page
}

func newPageAssertions(page Page, isNot bool, defaultTimeout *float64) *pageAssertionsImpl {
	return &pageAssertionsImpl{
		assertionsBase: assertionsBase{
			actualLocator:  page.Locator(":root"),
			isNot:          isNot,
			defaultTimeout: defaultTimeout,
		},
		actualPage: page,
	}
}

func (pa *pageAssertionsImpl) ToHaveTitle(titleOrRegExp interface{}, options ...PageAssertionsToHaveTitleOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedValues, err := toExpectedTextValues([]interface{}{titleOrRegExp}, false, true, nil)
	if err != nil {
		return err
	}
	return pa.expect(
		"to.have.title",
		frameExpectOptions{ExpectedText: expectedValues, Timeout: timeout},
		titleOrRegExp,
		"Page title expected to be",
	)
}

func (pa *pageAssertionsImpl) ToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error {
	var timeout *float64
	var ignoreCase *bool
	if len(options) == 1 {
		timeout = options[0].Timeout
		ignoreCase = options[0].IgnoreCase
	}

	baseURL := pa.actualPage.Context().(*browserContextImpl).options.BaseURL
	if urlPath, ok := urlOrRegExp.(string); ok && baseURL != nil {
		u, _ := url.Parse(*baseURL)
		u.Path = path.Join(u.Path, urlPath)
		urlOrRegExp = u.String()
	}

	expectedValues, err := toExpectedTextValues([]interface{}{urlOrRegExp}, false, false, ignoreCase)
	if err != nil {
		return err
	}
	return pa.expect(
		"to.have.url",
		frameExpectOptions{ExpectedText: expectedValues, Timeout: timeout},
		urlOrRegExp,
		"Page URL expected to be",
	)
}

func (pa *pageAssertionsImpl) Not() PageAssertions {
	return newPageAssertions(pa.actualPage, true, pa.defaultTimeout)
}
