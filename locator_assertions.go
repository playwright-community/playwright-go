package playwright

import (
	"reflect"
	"regexp"
)

type locatorAssertionsImpl struct {
	assertionsBase
}

func newLocatorAssertions(locator Locator, isNot bool) LocatorAssertions {
	return &locatorAssertionsImpl{
		assertionsBase: assertionsBase{
			actualLocator: locator,
			isNot:         isNot,
		},
	}
}

func (la *locatorAssertionsImpl) ToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error {
	var expression = "to.be.checked"
	var timeout *float64
	if len(options) == 1 {
		if options[0].Checked != nil && !*options[0].Checked {
			expression = "to.be.unchecked"
		}
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		expression,
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be checked",
	)
}

func (la *locatorAssertionsImpl) ToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.disabled",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be disabled",
	)
}

func (la *locatorAssertionsImpl) ToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.editable",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be editable",
	)
}

func (la *locatorAssertionsImpl) ToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.empty",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be empty",
	)
}

func (la *locatorAssertionsImpl) ToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.enabled",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be enabled",
	)
}

func (la *locatorAssertionsImpl) ToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.focused",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be focused",
	)
}

func (la *locatorAssertionsImpl) ToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.hidden",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be hidden",
	)
}

func (la *locatorAssertionsImpl) ToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.be.visible",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be visible",
	)
}

func (la *locatorAssertionsImpl) ToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error {
	var (
		timeout      *float64
		useInnerText *bool
	)
	if len(options) == 1 {
		timeout = options[0].Timeout
		useInnerText = options[0].UseInnerText
	}

	switch expected.(type) {
	case []string, []*regexp.Regexp:
		expectedText := la.toExpectedTextValues(la.convertToInterfaceList(expected), true, true)
		return la.expectImpl(
			"to.contain.text.array",
			frameExpectOptions{
				ExpectedText: expectedText,
				UseInnerText: useInnerText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to contain text",
		)
	default:
		expectedText := la.toExpectedTextValues([]interface{}{expected}, true, true)
		return la.expectImpl(
			"to.have.text",
			frameExpectOptions{
				ExpectedText: expectedText,
				UseInnerText: useInnerText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to contain text",
		)
	}
}

func (la *locatorAssertionsImpl) ToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := la.toExpectedTextValues([]interface{}{value}, false, false)
	return la.expectImpl(
		"to.have.attribute",
		frameExpectOptions{
			ExpressionArg: name,
			ExpectedText:  expectedText,
			Timeout:       timeout,
		},
		value,
		"Locator expected to have attribute",
	)
}

func (la *locatorAssertionsImpl) ToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	switch expected.(type) {
	case []string, []*regexp.Regexp:
		expectedText := la.toExpectedTextValues(la.convertToInterfaceList(expected), false, false)
		return la.expectImpl(
			"to.have.class.array",
			frameExpectOptions{
				ExpectedText: expectedText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to have class",
		)
	default:
		expectedText := la.toExpectedTextValues([]interface{}{expected}, false, false)
		return la.expectImpl(
			"to.have.class",
			frameExpectOptions{
				ExpectedText: expectedText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to have class",
		)
	}
}

func (la *locatorAssertionsImpl) ToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.have.count",
		frameExpectOptions{ExpectedNumber: &count, Timeout: timeout},
		count,
		"Locator expected to have count",
	)
}

func (la *locatorAssertionsImpl) ToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := la.toExpectedTextValues([]interface{}{value}, false, false)
	return la.expectImpl(
		"to.have.css",
		frameExpectOptions{
			ExpressionArg: name,
			ExpectedText:  expectedText,
			Timeout:       timeout,
		},
		value,
		"Locator expected to have CSS",
	)
}

func (la *locatorAssertionsImpl) ToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := la.toExpectedTextValues([]interface{}{id}, false, false)
	return la.expectImpl(
		"to.have.id",
		frameExpectOptions{ExpectedText: expectedText, Timeout: timeout},
		id,
		"Locator expected to have ID",
	)
}

func (la *locatorAssertionsImpl) ToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expectImpl(
		"to.have.property",
		frameExpectOptions{
			ExpressionArg: name,
			ExpectedValue: value,
			Timeout:       timeout,
		},
		value,
		"Locator expected to have JS Property",
	)
}

func (la *locatorAssertionsImpl) ToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error {
	var (
		timeout      *float64
		useInnerText *bool
	)
	if len(options) == 1 {
		timeout = options[0].Timeout
		useInnerText = options[0].UseInnerText
	}

	switch expected.(type) {
	case []string, []*regexp.Regexp:
		expectedText := la.toExpectedTextValues(la.convertToInterfaceList(expected), false, true)
		return la.expectImpl(
			"to.have.text.array",
			frameExpectOptions{
				ExpectedText: expectedText,
				UseInnerText: useInnerText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to have text",
		)
	default:
		expectedText := la.toExpectedTextValues([]interface{}{expected}, false, true)
		return la.expectImpl(
			"to.have.text",
			frameExpectOptions{
				ExpectedText: expectedText,
				UseInnerText: useInnerText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to have text",
		)
	}
}

func (la *locatorAssertionsImpl) ToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := la.toExpectedTextValues([]interface{}{value}, false, false)
	return la.expectImpl(
		"to.have.value",
		frameExpectOptions{ExpectedText: expectedText, Timeout: timeout},
		value,
		"Locator expected to have Value",
	)
}

func (la *locatorAssertionsImpl) NotToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error {
	return la.not().ToBeChecked(options...)
}

func (la *locatorAssertionsImpl) NotToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error {
	return la.not().ToBeDisabled(options...)
}

func (la *locatorAssertionsImpl) NotToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error {
	return la.not().ToBeEditable(options...)
}

func (la *locatorAssertionsImpl) NotToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error {
	return la.not().ToBeEmpty(options...)
}

func (la *locatorAssertionsImpl) NotToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error {
	return la.not().ToBeEnabled(options...)
}

func (la *locatorAssertionsImpl) NotToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error {
	return la.not().ToBeFocused(options...)
}

func (la *locatorAssertionsImpl) NotToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error {
	return la.not().ToBeHidden(options...)
}

func (la *locatorAssertionsImpl) NotToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error {
	return la.not().ToBeVisible(options...)
}

func (la *locatorAssertionsImpl) NotToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error {
	return la.not().ToContainText(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error {
	return la.not().ToHaveAttribute(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error {
	return la.not().ToHaveClass(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error {
	return la.not().ToHaveCount(count, options...)
}

func (la *locatorAssertionsImpl) NotToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error {
	return la.not().ToHaveCSS(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error {
	return la.not().ToHaveId(id, options...)
}

func (la *locatorAssertionsImpl) NotToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error {
	return la.not().ToHaveJSProperty(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error {
	return la.not().ToHaveText(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error {
	return la.not().ToHaveValue(value, options...)
}

func (la *locatorAssertionsImpl) not() LocatorAssertions {
	return newLocatorAssertions(la.actualLocator, true)
}

func (la *locatorAssertionsImpl) convertToInterfaceList(v interface{}) []interface{} {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		return []interface{}{v}
	}

	list := make([]interface{}, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		list[i] = rv.Index(i).Interface()
	}
	return list
}
