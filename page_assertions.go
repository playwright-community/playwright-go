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
	expectedValues := toExpectedTextValues([]interface{}{titleOrRegExp}, false, true, nil)
	return pa.expect(
		"to.have.title",
		frameExpectOptions{ExpectedText: expectedValues, Timeout: timeout},
		titleOrRegExp,
		"Page title expected to be",
	)
}

func (pa *pageAssertionsImpl) ToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}

	baseURL := pa.actualPage.Context().(*browserContextImpl).options.BaseURL
	if urlPath, ok := urlOrRegExp.(string); ok && baseURL != nil {
		u, _ := url.Parse(*baseURL)
		u.Path = path.Join(u.Path, urlPath)
		urlOrRegExp = u.String()
	}

	expectedValues := toExpectedTextValues([]interface{}{urlOrRegExp}, false, false, nil)
	return pa.expect(
		"to.have.url",
		frameExpectOptions{ExpectedText: expectedValues, Timeout: timeout},
		urlOrRegExp,
		"Page URL expected to be",
	)
}

func (pa *pageAssertionsImpl) NotToHaveTitle(titleOrRegExp interface{}, options ...PageAssertionsToHaveTitleOptions) error {
	return pa.Not().ToHaveTitle(titleOrRegExp, options...)
}

func (pa *pageAssertionsImpl) NotToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error {
	return pa.Not().ToHaveURL(urlOrRegExp, options...)
}

func (pa *pageAssertionsImpl) Not() PageAssertions {
	return newPageAssertions(pa.actualPage, true, pa.defaultTimeout)
}
