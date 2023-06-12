package playwright

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var testIdAttributeName = "data-testid"

type locatorImpl struct {
	frame    *frameImpl
	selector string
	options  *LocatorLocatorOptions
}

func newLocator(frame *frameImpl, selector string, options ...LocatorLocatorOptions) (*locatorImpl, error) {
	var option *LocatorLocatorOptions
	if len(options) == 1 {
		option = &options[0]
		if option.HasText != nil {
			switch hasText := option.HasText.(type) {
			case *regexp.Regexp:
				pattern, flags := convertRegexp(hasText)
				selector += fmt.Sprintf(" >> :scope:text-matches('%s', '%s')", pattern, flags)
			case string:
				selector += fmt.Sprintf(" >> :scope:has-text('%s')", hasText)
			}
		}
		if option.Has != nil {
			has := option.Has.(*locatorImpl)
			if frame != has.frame {
				return nil, errors.New("inner 'has' locator must belong to the same frame")
			}
			marshaledSelector, err := json.Marshal(has.selector)
			if err != nil {
				return nil, fmt.Errorf("could not marshal selector '%s'", has.selector)
			}
			selector += " >> internal:has=" + string(marshaledSelector)
		}
	}

	return &locatorImpl{frame: frame, selector: selector, options: option}, nil
}

func convertRegexp(reg *regexp.Regexp) (pattern, flags string) {
	matches := regexp.MustCompile(`\(\?([imsU]+)\)(.+)`).FindStringSubmatch(reg.String())

	if len(matches) == 3 {
		pattern = matches[2]
		flags = matches[1]
	} else {
		pattern = reg.String()
	}
	return
}

func (l *locatorImpl) All() ([]Locator, error) {
	result := make([]Locator, 0)
	count, err := l.Count()
	if err != nil {
		return nil, err
	}
	for i := 0; i < count; i++ {
		item, err := l.Nth(i)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func (l *locatorImpl) AllInnerTexts() ([]string, error) {
	innerTexts, err := l.frame.EvalOnSelectorAll(l.selector, "ee => ee.map(e => e.innerText)")
	if err != nil {
		return nil, err
	}
	texts := innerTexts.([]interface{})
	result := make([]string, len(texts))
	for i := range texts {
		result[i] = texts[i].(string)
	}
	return result, nil
}

func (l *locatorImpl) AllTextContents() ([]string, error) {
	textContents, err := l.frame.EvalOnSelectorAll(l.selector, "ee => ee.map(e => e.textContent || '')")
	if err != nil {
		return nil, err
	}
	texts := textContents.([]interface{})
	result := make([]string, len(texts))
	for i := range texts {
		result[i] = texts[i].(string)
	}
	return result, nil
}

func (l *locatorImpl) Blur(options ...LocatorBlurOptions) error {
	params := map[string]interface{}{
		"selector": l.selector,
		"strict":   true,
	}
	if len(options) == 1 {
		if options[0].Timeout != nil {
			params["timeout"] = options[0].Timeout
		}
	}
	_, err := l.frame.channel.Send("blur", params)
	return err
}

func (l *locatorImpl) BoundingBox(options ...LocatorBoundingBoxOptions) (*Rect, error) {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	result, err := l.withElement(func(handle ElementHandle) (interface{}, error) {
		return handle.BoundingBox()
	}, option)

	if err != nil {
		return nil, err
	}

	return result.(*Rect), nil
}

func (l *locatorImpl) Check(options ...FrameCheckOptions) error {
	return l.frame.Check(l.selector, options...)
}

func (l *locatorImpl) Clear(options ...LocatorClearOptions) error {
	if len(options) > 0 {
		return l.Fill("", FrameFillOptions{
			Force:       options[0].Force,
			NoWaitAfter: options[0].NoWaitAfter,
			Timeout:     options[0].Timeout,
		})
	} else {
		return l.Fill("")
	}
}

func (l *locatorImpl) Click(options ...PageClickOptions) error {
	return l.frame.Click(l.selector, options...)
}

func (l *locatorImpl) Count() (int, error) {
	return l.frame.queryCount(l.selector)
}

func (l *locatorImpl) Dblclick(options ...FrameDblclickOptions) error {
	return l.frame.Dblclick(l.selector, options...)
}

func (l *locatorImpl) DispatchEvent(typ string, eventInit interface{}, options ...PageDispatchEventOptions) error {
	return l.frame.DispatchEvent(l.selector, typ, eventInit, options...)
}

func (l *locatorImpl) DragTo(target Locator, options ...FrameDragAndDropOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.DragAndDrop(l.selector, target.(*locatorImpl).selector, options...)
}

func (l *locatorImpl) ElementHandle(options ...LocatorElementHandleOptions) (ElementHandle, error) {
	option := PageWaitForSelectorOptions{
		State:  WaitForSelectorStateAttached,
		Strict: Bool(true),
	}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}
	return l.frame.WaitForSelector(l.selector, option)
}

func (l *locatorImpl) ElementHandles() ([]ElementHandle, error) {
	return l.frame.QuerySelectorAll(l.selector)
}

func (l *locatorImpl) Evaluate(expression string, arg interface{}, options ...LocatorEvaluateOptions) (interface{}, error) {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	return l.withElement(func(handle ElementHandle) (interface{}, error) {
		return handle.Evaluate(expression, arg)
	}, option)
}

func (l *locatorImpl) EvaluateAll(expression string, options ...interface{}) (interface{}, error) {
	return l.frame.EvalOnSelectorAll(l.selector, expression, options...)
}

func (l *locatorImpl) EvaluateHandle(expression string, arg interface{}, options ...LocatorEvaluateHandleOptions) (interface{}, error) {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	return l.withElement(func(handle ElementHandle) (interface{}, error) {
		return handle.EvaluateHandle(expression, arg)
	}, option)
}

func (l *locatorImpl) Fill(value string, options ...FrameFillOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Fill(l.selector, value, options...)
}

func (l *locatorImpl) Filter(options ...LocatorLocatorOptions) (Locator, error) {
	return newLocator(l.frame, l.selector, options...)
}

func (l *locatorImpl) First() (Locator, error) {
	return newLocator(l.frame, l.selector+" >> nth=0")
}

func (l *locatorImpl) Focus(options ...FrameFocusOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Focus(l.selector, options...)
}

func (l *locatorImpl) FrameLocator(selector string) FrameLocator {
	return newFrameLocator(l.frame, l.selector+" >> "+selector)
}

func (l *locatorImpl) GetAttribute(name string, options ...PageGetAttributeOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.GetAttribute(l.selector, name, options...)
}

func (l *locatorImpl) GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) (Locator, error) {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByAltTextSelector(text, exact))
}

func (l *locatorImpl) GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) (Locator, error) {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByLabelSelector(text, exact))
}

func (l *locatorImpl) GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) (Locator, error) {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByPlaceholderSelector(text, exact))
}

func (l *locatorImpl) GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) (Locator, error) {
	return l.Locator(getByRoleSelector(role, options...))
}

func (l *locatorImpl) GetByTestId(testId interface{}) (Locator, error) {
	return l.Locator(getByTestIdSelector(getTestIdAttributeName(), testId))
}

func (l *locatorImpl) GetByText(text interface{}, options ...LocatorGetByTextOptions) (Locator, error) {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByTextSelector(text, exact))
}

func (l *locatorImpl) GetByTitle(text interface{}, options ...LocatorGetByTitleOptions) (Locator, error) {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByTitleSelector(text, exact))
}

func (l *locatorImpl) Highlight() error {
	return l.frame.highlight(l.selector)
}

func (l *locatorImpl) Hover(options ...PageHoverOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Hover(l.selector, options...)
}

func (l *locatorImpl) InnerHTML(options ...PageInnerHTMLOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InnerHTML(l.selector, options...)
}

func (l *locatorImpl) InnerText(options ...PageInnerTextOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InnerText(l.selector, options...)
}

func (l *locatorImpl) InputValue(options ...FrameInputValueOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InputValue(l.selector, options...)
}

func (l *locatorImpl) IsChecked(options ...FrameIsCheckedOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsChecked(l.selector, options...)
}

func (l *locatorImpl) IsDisabled(options ...FrameIsDisabledOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsDisabled(l.selector, options...)
}

func (l *locatorImpl) IsEditable(options ...FrameIsEditableOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsEditable(l.selector, options...)
}

func (l *locatorImpl) IsEnabled(options ...FrameIsEnabledOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsEnabled(l.selector, options...)
}

func (l *locatorImpl) IsHidden(options ...FrameIsHiddenOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsHidden(l.selector, options...)
}

func (l *locatorImpl) IsVisible(options ...FrameIsVisibleOptions) (bool, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsVisible(l.selector, options...)
}

func (l *locatorImpl) Last() (Locator, error) {
	return newLocator(l.frame, l.selector+" >> nth=-1")
}

func (l *locatorImpl) Locator(selector string, options ...LocatorLocatorOptions) (Locator, error) {
	return newLocator(l.frame, l.selector+" >> "+selector, options...)
}

func (l *locatorImpl) Nth(index int) (Locator, error) {
	return newLocator(l.frame, l.selector+" >> nth="+strconv.Itoa(index))
}

func (l *locatorImpl) Page() Page {
	return l.frame.Page()
}

func (l *locatorImpl) Press(key string, options ...PagePressOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Press(l.selector, key, options...)
}

func (l *locatorImpl) Screenshot(options ...LocatorScreenshotOptions) ([]byte, error) {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	result, err := l.withElement(func(handle ElementHandle) (interface{}, error) {
		var screenshotOption ElementHandleScreenshotOptions
		if len(options) == 1 {
			screenshotOption = ElementHandleScreenshotOptions(options[0])
		}
		return handle.Screenshot(screenshotOption)
	}, option)

	if err != nil {
		return nil, err
	}

	return result.([]byte), nil
}

func (l *locatorImpl) ScrollIntoViewIfNeeded(options ...LocatorScrollIntoViewIfNeededOptions) error {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	_, err := l.withElement(func(handle ElementHandle) (interface{}, error) {
		var opt ElementHandleScrollIntoViewIfNeededOptions
		if len(options) == 1 {
			opt.Timeout = options[0].Timeout
		}
		return nil, handle.ScrollIntoViewIfNeeded(opt)
	}, option)

	return err
}

func (l *locatorImpl) SelectOption(values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SelectOption(l.selector, values, options...)
}

func (l *locatorImpl) SelectText(options ...LocatorSelectTextOptions) error {
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	_, err := l.withElement(func(handle ElementHandle) (interface{}, error) {
		var opt ElementHandleSelectTextOptions
		if len(options) == 1 {
			opt = ElementHandleSelectTextOptions(options[0])
		}
		return nil, handle.SelectText(opt)
	}, option)

	return err
}

func (l *locatorImpl) SetChecked(checked bool, options ...FrameSetCheckedOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SetChecked(l.selector, checked, options...)
}

func (l *locatorImpl) SetInputFiles(files []InputFile, options ...FrameSetInputFilesOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SetInputFiles(l.selector, files, options...)
}

func (l *locatorImpl) Tap(options ...FrameTapOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Tap(l.selector, options...)
}

func (l *locatorImpl) TextContent(options ...FrameTextContentOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.TextContent(l.selector, options...)
}

func (l *locatorImpl) Type(text string, options ...PageTypeOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Type(l.selector, text, options...)
}

func (l *locatorImpl) Uncheck(options ...FrameUncheckOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Uncheck(l.selector, options...)
}

func (l *locatorImpl) WaitFor(options ...PageWaitForSelectorOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	_, err := l.frame.WaitForSelector(l.selector, options...)
	return err
}

func (l *locatorImpl) withElement(
	callback func(handle ElementHandle) (interface{}, error),
	options ...PageWaitForSelectorOptions,
) (interface{}, error) {
	handle, err := l.frame.WaitForSelector(l.selector, options...)
	if err != nil {
		return nil, err
	}

	result, err := callback(handle)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getByRoleSelector(role AriaRole, options ...LocatorGetByRoleOptions) string {
	props := make(map[string]string)
	if len(options) == 1 {
		if options[0].Checked != nil {
			props["checked"] = fmt.Sprintf("%v", options[0].Checked)
		}
		if options[0].Disabled != nil {
			props["disabled"] = fmt.Sprintf("%v", options[0].Disabled)
		}
		if options[0].Selected != nil {
			props["selected"] = fmt.Sprintf("%v", options[0].Selected)
		}
		if options[0].Expanded != nil {
			props["expanded"] = fmt.Sprintf("%v", options[0].Expanded)
		}
		if options[0].IncludeHidden != nil {
			props["includeHidden"] = fmt.Sprintf("%v", options[0].IncludeHidden)
		}
		if options[0].Level != nil {
			props["level"] = fmt.Sprintf("%d", options[0].Level)
		}
		if options[0].Name != nil {
			exact := false
			if options[0].Exact != nil {
				exact = *options[0].Exact
			}
			switch options[0].Name.(type) {
			case string:
				props["name"] = escapeForAttributeSelector(options[0].Name.(string), exact)
			case *regexp.Regexp:
				pattern, flag := convertRegexp(options[0].Name.(*regexp.Regexp))
				props["name"] = fmt.Sprintf(`/%s/%s`, pattern, flag)
			}
		}
		if options[0].Pressed != nil {
			props["pressed"] = fmt.Sprintf("%v", options[0].Pressed)
		}
	}
	propsStr := ""
	for k, v := range props {
		propsStr += "[" + k + "=" + v + "]"
	}
	return fmt.Sprintf("internal:role=%s%s", role, propsStr)
}

func escapeForAttributeSelector(value string, exact bool) string {
	suffix := "i"
	if exact {
		suffix = ""
	}
	return fmt.Sprintf(`"%s"%s`, strings.Replace(value, `"`, `\"`, -1), suffix)
}

func escapeForTextSelector(text interface{}, exact bool) string {
	switch text := text.(type) {
	case *regexp.Regexp:
		pattern, flag := convertRegexp(text)
		return fmt.Sprintf(`/%s/%s`, pattern, flag)
	default:
		if exact {
			return fmt.Sprintf(`"%s"s`, text.(string))
		}
		return fmt.Sprintf(`"%s"i`, text.(string))
	}
}

func getByAttributeTextSelector(attrName string, text interface{}, exact bool) string {
	switch text := text.(type) {
	case *regexp.Regexp:
		pattern, flag := convertRegexp(text)
		return fmt.Sprintf(`internal:attr=[%s=/%s/%s]`, attrName, pattern, flag)
	default:
		return fmt.Sprintf(`internal:attr=[%s=%s]`, attrName, escapeForAttributeSelector(text.(string), exact))
	}
}

func getByTextSelector(text interface{}, exact bool) string {
	return fmt.Sprintf(`internal:text=%s`, escapeForTextSelector(text, exact))
}

func getByPlaceholderSelector(text interface{}, exact bool) string {
	return getByAttributeTextSelector("placeholder", text, exact)
}

func getByTitleSelector(text interface{}, exact bool) string {
	return getByAttributeTextSelector("title", text, exact)
}

func getByAltTextSelector(text interface{}, exact bool) string {
	return getByAttributeTextSelector("alt", text, exact)
}

func getByLabelSelector(text interface{}, exact bool) string {
	return fmt.Sprintf(`internal:label=%s`, escapeForTextSelector(text, exact))
}

func getTestIdAttributeName() string {
	return testIdAttributeName
}

func setTestIdAttributeName(name string) {
	testIdAttributeName = name
}

func getByTestIdSelector(testIdAttributeName string, testId interface{}) string {
	switch testId := testId.(type) {
	case *regexp.Regexp:
		pattern, flag := convertRegexp(testId)
		return fmt.Sprintf(`internal:testid=[%s=/%s/%s]`, testIdAttributeName, pattern, flag)
	default:
		return fmt.Sprintf(`internal:testid=[%s=%s]`, testIdAttributeName, escapeForAttributeSelector(testId.(string), true))
	}
}
