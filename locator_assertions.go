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
	expression := "to.be.attached"
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
	expression := "to.be.checked"
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
			Timeout:        timeout,
		},
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
		expectedText, err := toExpectedTextValues(convertToInterfaceList(expected), true, true, ignoreCase)
		if err != nil {
			return err
		}
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
		expectedText, err := toExpectedTextValues([]interface{}{expected}, true, true, ignoreCase)
		if err != nil {
			return err
		}
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
	var ignoreCase *bool
	if len(options) == 1 {
		timeout = options[0].Timeout
		ignoreCase = options[0].IgnoreCase
	}
	expectedText, err := toExpectedTextValues([]interface{}{value}, false, false, ignoreCase)
	if err != nil {
		return err
	}
	return la.expect(
		"to.have.attribute.value",
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
		expectedText, err := toExpectedTextValues(convertToInterfaceList(expected), false, false, nil)
		if err != nil {
			return err
		}
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
		expectedText, err := toExpectedTextValues([]interface{}{expected}, false, false, nil)
		if err != nil {
			return err
		}
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
	expectedText, err := toExpectedTextValues([]interface{}{value}, false, false, nil)
	if err != nil {
		return err
	}
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
	expectedText, err := toExpectedTextValues([]interface{}{id}, false, false, nil)
	if err != nil {
		return err
	}
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
		expectedText, err := toExpectedTextValues(convertToInterfaceList(expected), false, true, ignoreCase)
		if err != nil {
			return err
		}
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
		expectedText, err := toExpectedTextValues([]interface{}{expected}, false, true, ignoreCase)
		if err != nil {
			return err
		}
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
	expectedText, err := toExpectedTextValues([]interface{}{value}, false, false, nil)
	if err != nil {
		return err
	}
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
	expectedText, err := toExpectedTextValues(values, false, false, nil)
	if err != nil {
		return err
	}
	return la.expect(
		"to.have.values",
		frameExpectOptions{ExpectedText: expectedText, Timeout: timeout},
		values,
		"Locator expected to have Values",
	)
}

func (la *locatorAssertionsImpl) Not() LocatorAssertions {
	return newLocatorAssertions(la.actualLocator, true, la.defaultTimeout)
}
