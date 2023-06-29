package playwright

import (
	"regexp"
)

type locatorAssertionsImpl struct {
	assertionsBase
}

func newLocatorAssertions(locator Locator, isNot bool, defaultTimeout *float64) *locatorAssertionsImpl {
	return &locatorAssertionsImpl{
		assertionsBase: assertionsBase{
			actualLocator:  locator,
			isNot:          isNot,
			defaultTimeout: defaultTimeout,
		},
	}
}

func (la *locatorAssertionsImpl) ToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error {
	var expression = "to.be.attached"
	var timeout *float64
	if len(options) == 1 {
		if options[0].Attached != nil && !*options[0].Attached {
			expression = "to.be.detached"
		}
		timeout = options[0].Timeout
	}
	return la.expect(
		expression,
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be attached",
	)
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
	return la.expect(
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
	return la.expect(
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
	return la.expect(
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
	return la.expect(
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
	return la.expect(
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
	return la.expect(
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
	return la.expect(
		"to.be.hidden",
		frameExpectOptions{Timeout: timeout},
		nil,
		"Locator expected to be hidden",
	)
}

func (la *locatorAssertionsImpl) ToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error {
	var (
		ratio   *float64
		timeout *float64
	)
	if len(options) == 1 {
		ratio = options[0].Ratio
		timeout = options[0].Timeout
	}
	return la.expect(
		"to.be.in.viewport",
		frameExpectOptions{
			ExpectedNumber: ratio,
			Timeout:        timeout},
		nil,
		"Locator expected to be in viewport",
	)
}

func (la *locatorAssertionsImpl) ToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	return la.expect(
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
		ignoreCase   *bool
	)
	if len(options) == 1 {
		timeout = options[0].Timeout
		useInnerText = options[0].UseInnerText
		ignoreCase = options[0].IgnoreCase
	}

	switch expected.(type) {
	case []string, []*regexp.Regexp:
		expectedText := toExpectedTextValues(convertToInterfaceList(expected), true, true, ignoreCase)
		return la.expect(
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
		expectedText := toExpectedTextValues([]interface{}{expected}, true, true, ignoreCase)
		return la.expect(
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
	expectedText := toExpectedTextValues([]interface{}{value}, false, false, nil)
	return la.expect(
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
		expectedText := toExpectedTextValues(convertToInterfaceList(expected), false, false, nil)
		return la.expect(
			"to.have.class.array",
			frameExpectOptions{
				ExpectedText: expectedText,
				Timeout:      timeout,
			},
			expected,
			"Locator expected to have class",
		)
	default:
		expectedText := toExpectedTextValues([]interface{}{expected}, false, false, nil)
		return la.expect(
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
	return la.expect(
		"to.have.count",
		frameExpectOptions{ExpectedNumber: Float(float64(count)), Timeout: timeout},
		count,
		"Locator expected to have count",
	)
}

func (la *locatorAssertionsImpl) ToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := toExpectedTextValues([]interface{}{value}, false, false, nil)
	return la.expect(
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
	expectedText := toExpectedTextValues([]interface{}{id}, false, false, nil)
	return la.expect(
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
	return la.expect(
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
		ignoreCase   *bool
	)
	if len(options) == 1 {
		timeout = options[0].Timeout
		useInnerText = options[0].UseInnerText
		ignoreCase = options[0].IgnoreCase
	}

	switch expected.(type) {
	case []string, []*regexp.Regexp:
		expectedText := toExpectedTextValues(convertToInterfaceList(expected), false, true, ignoreCase)
		return la.expect(
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
		expectedText := toExpectedTextValues([]interface{}{expected}, false, true, ignoreCase)
		return la.expect(
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
	expectedText := toExpectedTextValues([]interface{}{value}, false, false, nil)
	return la.expect(
		"to.have.value",
		frameExpectOptions{ExpectedText: expectedText, Timeout: timeout},
		value,
		"Locator expected to have Value",
	)
}

func (la *locatorAssertionsImpl) ToHaveValues(values []interface{}, options ...LocatorAssertionsToHaveValuesOptions) error {
	var timeout *float64
	if len(options) == 1 {
		timeout = options[0].Timeout
	}
	expectedText := toExpectedTextValues(values, false, false, nil)
	return la.expect(
		"to.have.values",
		frameExpectOptions{ExpectedText: expectedText, Timeout: timeout},
		values,
		"Locator expected to have Values",
	)
}

func (la *locatorAssertionsImpl) NotToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error {
	return la.Not().ToBeAttached(options...)
}

func (la *locatorAssertionsImpl) NotToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error {
	return la.Not().ToBeChecked(options...)
}

func (la *locatorAssertionsImpl) NotToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error {
	return la.Not().ToBeDisabled(options...)
}

func (la *locatorAssertionsImpl) NotToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error {
	return la.Not().ToBeEditable(options...)
}

func (la *locatorAssertionsImpl) NotToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error {
	return la.Not().ToBeEmpty(options...)
}

func (la *locatorAssertionsImpl) NotToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error {
	return la.Not().ToBeEnabled(options...)
}

func (la *locatorAssertionsImpl) NotToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error {
	return la.Not().ToBeFocused(options...)
}

func (la *locatorAssertionsImpl) NotToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error {
	return la.Not().ToBeHidden(options...)
}

func (la *locatorAssertionsImpl) NotToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error {
	return la.Not().ToBeInViewport(options...)
}

func (la *locatorAssertionsImpl) NotToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error {
	return la.Not().ToBeVisible(options...)
}

func (la *locatorAssertionsImpl) NotToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error {
	return la.Not().ToContainText(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error {
	return la.Not().ToHaveAttribute(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error {
	return la.Not().ToHaveClass(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error {
	return la.Not().ToHaveCount(count, options...)
}

func (la *locatorAssertionsImpl) NotToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error {
	return la.Not().ToHaveCSS(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error {
	return la.Not().ToHaveId(id, options...)
}

func (la *locatorAssertionsImpl) NotToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error {
	return la.Not().ToHaveJSProperty(name, value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error {
	return la.Not().ToHaveText(expected, options...)
}

func (la *locatorAssertionsImpl) NotToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error {
	return la.Not().ToHaveValue(value, options...)
}

func (la *locatorAssertionsImpl) NotToHaveValues(values []interface{}, options ...LocatorAssertionsToHaveValuesOptions) error {
	return la.Not().ToHaveValues(values, options...)
}

func (la *locatorAssertionsImpl) Not() LocatorAssertions {
	return newLocatorAssertions(la.actualLocator, true, la.defaultTimeout)
}
