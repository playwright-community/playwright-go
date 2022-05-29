package playwright

import (
	"fmt"
	"regexp"
	"strings"
)

type playwrightAssertionsImpl struct {
}

func NewPlaywrightAssertions() PlaywrightAssertions {
	return &playwrightAssertionsImpl{}
}

func (pa *playwrightAssertionsImpl) ExpectLocator(locator Locator) LocatorAssertions {
	return newLocatorAssertions(locator, false)
}

func (pa *playwrightAssertionsImpl) ExpectPage(page Page) PageAssertions {
	return newPageAssertions(page, false)
}

type expectedTextValue struct {
	Str                 *string `json:"string"`
	RegexSource         *string `json:"regexSource"`
	RegexFlags          *string `json:"regexFlags"`
	MatchSubstring      *bool   `json:"matchSubstring"`
	NormalizeWhiteSpace *bool   `json:"normalizeWhiteSpace"`
}

type frameExpectOptions struct {
	ExpressionArg  interface{}         `json:"expressionArg"`
	ExpectedText   []expectedTextValue `json:"expectedText"`
	ExpectedNumber *int                `json:"expectedNumber"`
	ExpectedValue  interface{}         `json:"expectedValue"`
	UseInnerText   *bool               `json:"useInnerText"`
	IsNot          bool                `json:"isNot"`
	Timeout        *float64            `json:"timeout"`
}

type frameExpectResult struct {
	Matches  bool        `json:"matches"`
	Received interface{} `json:"received"`
	Log      []string    `json:"log"`
}

type assertionsBase struct {
	actualLocator Locator
	isNot         bool
}

func (b *assertionsBase) expectImpl(
	expression string,
	options frameExpectOptions,
	expected interface{},
	message string,
) error {
	options.IsNot = b.isNot
	if options.Timeout == nil {
		options.Timeout = Float(5000)
	}
	if options.IsNot {
		message = strings.ReplaceAll(message, "expected to", "expected not to")
	}
	result, err := b.actualLocator.(*locatorImpl).expect(expression, options)
	if err != nil {
		return err
	}

	if result.Matches == b.isNot {
		actual := result.Received
		log := strings.Join(result.Log, "\n")
		if log != "" {
			log = "\nCall log:\n" + log
		}
		if expected != nil {
			return fmt.Errorf("%s '%v'\nActual value: %v %s", message, expected, actual, log)
		}
		return fmt.Errorf("%s\nActual value: %v %s", message, actual, log)
	}

	return nil
}

func (b *assertionsBase) toExpectedTextValues(
	items []interface{},
	matchSubstring bool,
	normalizeWhiteSpace bool,
) []expectedTextValue {
	var out []expectedTextValue
	for _, item := range items {
		switch item := item.(type) {
		case string:
			out = append(out, expectedTextValue{
				Str:                 String(item),
				MatchSubstring:      Bool(matchSubstring),
				NormalizeWhiteSpace: Bool(normalizeWhiteSpace),
			})
		case *regexp.Regexp:
			pattern, flags := splitRegexpString(item)
			out = append(out, expectedTextValue{
				RegexSource:         String(pattern),
				RegexFlags:          String(flags),
				MatchSubstring:      Bool(matchSubstring),
				NormalizeWhiteSpace: Bool(normalizeWhiteSpace),
			})
		}
	}
	return out
}
