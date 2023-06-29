package playwright

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/playwright-community/playwright-go/internal/multierror"
)

var (
	testIdAttributeName    = "data-testid"
	ErrLocatorNotSameFrame = errors.New("inner 'has' or 'hasNot' locator must belong to the same frame")
)

type locatorImpl struct {
	frame    *frameImpl
	selector string
	options  *LocatorLocatorOptions
	err      error
}

func newLocator(frame *frameImpl, selector string, options ...LocatorLocatorOptions) *locatorImpl {
	option := &LocatorLocatorOptions{}
	if len(options) == 1 {
		option = &options[0]
	}
	locator := &locatorImpl{frame: frame, selector: selector, options: option, err: nil}
	if option.HasText != nil {
		selector += fmt.Sprintf(` >> internal:has-text=%s`, escapeForTextSelector(option.HasText, false))
	}
	if option.HasNotText != nil {
		selector += fmt.Sprintf(` >> internal:has-not-text=%s`, escapeForTextSelector(option.HasText, false))
	}
	if option.Has != nil {
		has := option.Has.(*locatorImpl)
		if frame != has.frame {
			locator.err = multierror.Join(locator.err, ErrLocatorNotSameFrame)
		} else {
			selector += fmt.Sprintf(` >> internal:has=%s`, escapeText(has.selector))
		}
	}
	if option.HasNot != nil {
		hasNot := option.HasNot.(*locatorImpl)
		if frame != hasNot.frame {
			locator.err = multierror.Join(locator.err, ErrLocatorNotSameFrame)
		} else {
			selector += fmt.Sprintf(` >> internal:has-not=%s`, escapeText(hasNot.selector))
		}
	}
	locator.selector = selector

	return locator
}

func (l *locatorImpl) Err() error {
	return l.err
}

func (l *locatorImpl) All() ([]Locator, error) {
	result := make([]Locator, 0)
	count, err := l.Count()
	if err != nil {
		return nil, err
	}
	for i := 0; i < count; i++ {
		result = append(result, l.Nth(i))
	}
	return result, nil
}

func (l *locatorImpl) AllInnerTexts() ([]string, error) {
	if l.err != nil {
		return nil, l.err
	}
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
	if l.err != nil {
		return nil, l.err
	}
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

func (l *locatorImpl) And(locator Locator) Locator {
	return newLocator(l.frame, l.selector+` >> internal:and=`+escapeText(locator.(*locatorImpl).selector))
}

func (l *locatorImpl) Or(locator Locator) Locator {
	return newLocator(l.frame, l.selector+` >> internal:or=`+escapeText(locator.(*locatorImpl).selector))
}

func (l *locatorImpl) Blur(options ...LocatorBlurOptions) error {
	if l.err != nil {
		return l.err
	}
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
	if l.err != nil {
		return nil, l.err
	}
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
	if l.err != nil {
		return l.err
	}
	return l.frame.Check(l.selector, options...)
}

func (l *locatorImpl) Clear(options ...LocatorClearOptions) error {
	if l.err != nil {
		return l.err
	}
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
	if l.err != nil {
		return l.err
	}
	return l.frame.Click(l.selector, options...)
}

func (l *locatorImpl) Count() (int, error) {
	if l.err != nil {
		return 0, l.err
	}
	return l.frame.queryCount(l.selector)
}

func (l *locatorImpl) Dblclick(options ...FrameDblclickOptions) error {
	if l.err != nil {
		return l.err
	}
	return l.frame.Dblclick(l.selector, options...)
}

func (l *locatorImpl) DispatchEvent(typ string, eventInit interface{}, options ...PageDispatchEventOptions) error {
	if l.err != nil {
		return l.err
	}
	return l.frame.DispatchEvent(l.selector, typ, eventInit, options...)
}

func (l *locatorImpl) DragTo(target Locator, options ...FrameDragAndDropOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.DragAndDrop(l.selector, target.(*locatorImpl).selector, options...)
}

func (l *locatorImpl) ElementHandle(options ...LocatorElementHandleOptions) (ElementHandle, error) {
	if l.err != nil {
		return nil, l.err
	}
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
	if l.err != nil {
		return nil, l.err
	}
	return l.frame.QuerySelectorAll(l.selector)
}

func (l *locatorImpl) Evaluate(expression string, arg interface{}, options ...LocatorEvaluateOptions) (interface{}, error) {
	if l.err != nil {
		return nil, l.err
	}
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	return l.withElement(func(handle ElementHandle) (interface{}, error) {
		return handle.Evaluate(expression, arg)
	}, option)
}

func (l *locatorImpl) EvaluateAll(expression string, options ...interface{}) (interface{}, error) {
	if l.err != nil {
		return nil, l.err
	}
	return l.frame.EvalOnSelectorAll(l.selector, expression, options...)
}

func (l *locatorImpl) EvaluateHandle(expression string, arg interface{}, options ...LocatorEvaluateHandleOptions) (interface{}, error) {
	if l.err != nil {
		return nil, l.err
	}
	var option PageWaitForSelectorOptions
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
	}

	return l.withElement(func(handle ElementHandle) (interface{}, error) {
		return handle.EvaluateHandle(expression, arg)
	}, option)
}

func (l *locatorImpl) Fill(value string, options ...FrameFillOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Fill(l.selector, value, options...)
}

func (l *locatorImpl) Filter(options ...LocatorLocatorOptions) Locator {
	return newLocator(l.frame, l.selector, options...)
}

func (l *locatorImpl) First() Locator {
	return newLocator(l.frame, l.selector+" >> nth=0")
}

func (l *locatorImpl) Focus(options ...FrameFocusOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Focus(l.selector, options...)
}

func (l *locatorImpl) FrameLocator(selector string) FrameLocator {
	return newFrameLocator(l.frame, l.selector+" >> "+selector)
}

func (l *locatorImpl) GetAttribute(name string, options ...PageGetAttributeOptions) (string, error) {
	if l.err != nil {
		return "", l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.GetAttribute(l.selector, name, options...)
}

func (l *locatorImpl) GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByAltTextSelector(text, exact))
}

func (l *locatorImpl) GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByLabelSelector(text, exact))
}

func (l *locatorImpl) GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByPlaceholderSelector(text, exact))
}

func (l *locatorImpl) GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator {
	return l.Locator(getByRoleSelector(role, options...))
}

func (l *locatorImpl) GetByTestId(testId interface{}) Locator {
	return l.Locator(getByTestIdSelector(getTestIdAttributeName(), testId))
}

func (l *locatorImpl) GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByTextSelector(text, exact))
}

func (l *locatorImpl) GetByTitle(text interface{}, options ...LocatorGetByTitleOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return l.Locator(getByTitleSelector(text, exact))
}

func (l *locatorImpl) Highlight() error {
	if l.err != nil {
		return l.err
	}
	return l.frame.highlight(l.selector)
}

func (l *locatorImpl) Hover(options ...PageHoverOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Hover(l.selector, options...)
}

func (l *locatorImpl) InnerHTML(options ...PageInnerHTMLOptions) (string, error) {
	if l.err != nil {
		return "", l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InnerHTML(l.selector, options...)
}

func (l *locatorImpl) InnerText(options ...PageInnerTextOptions) (string, error) {
	if l.err != nil {
		return "", l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InnerText(l.selector, options...)
}

func (l *locatorImpl) InputValue(options ...FrameInputValueOptions) (string, error) {
	if l.err != nil {
		return "", l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.InputValue(l.selector, options...)
}

func (l *locatorImpl) IsChecked(options ...FrameIsCheckedOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsChecked(l.selector, options...)
}

func (l *locatorImpl) IsDisabled(options ...FrameIsDisabledOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsDisabled(l.selector, options...)
}

func (l *locatorImpl) IsEditable(options ...FrameIsEditableOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsEditable(l.selector, options...)
}

func (l *locatorImpl) IsEnabled(options ...FrameIsEnabledOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsEnabled(l.selector, options...)
}

func (l *locatorImpl) IsHidden(options ...FrameIsHiddenOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsHidden(l.selector, options...)
}

func (l *locatorImpl) IsVisible(options ...FrameIsVisibleOptions) (bool, error) {
	if l.err != nil {
		return false, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.IsVisible(l.selector, options...)
}

func (l *locatorImpl) Last() Locator {
	return newLocator(l.frame, l.selector+" >> nth=-1")
}

func (l *locatorImpl) Locator(selector string, options ...LocatorLocatorOptions) Locator {
	return newLocator(l.frame, l.selector+" >> "+selector, options...)
}

func (l *locatorImpl) Nth(index int) Locator {
	return newLocator(l.frame, l.selector+" >> nth="+strconv.Itoa(index))
}

func (l *locatorImpl) Page() (Page, error) {
	if l.err != nil {
		return nil, l.err
	}
	return l.frame.Page(), nil
}

func (l *locatorImpl) Press(key string, options ...PagePressOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Press(l.selector, key, options...)
}

func (l *locatorImpl) Screenshot(options ...LocatorScreenshotOptions) ([]byte, error) {
	if l.err != nil {
		return nil, l.err
	}
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
	if l.err != nil {
		return l.err
	}
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
	if l.err != nil {
		return nil, l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SelectOption(l.selector, values, options...)
}

func (l *locatorImpl) SelectText(options ...LocatorSelectTextOptions) error {
	if l.err != nil {
		return l.err
	}
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
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SetChecked(l.selector, checked, options...)
}

func (l *locatorImpl) SetInputFiles(files []InputFile, options ...FrameSetInputFilesOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.SetInputFiles(l.selector, files, options...)
}

func (l *locatorImpl) Tap(options ...FrameTapOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Tap(l.selector, options...)
}

func (l *locatorImpl) TextContent(options ...FrameTextContentOptions) (string, error) {
	if l.err != nil {
		return "", l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.TextContent(l.selector, options...)
}

func (l *locatorImpl) Type(text string, options ...PageTypeOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Type(l.selector, text, options...)
}

func (l *locatorImpl) Uncheck(options ...FrameUncheckOptions) error {
	if l.err != nil {
		return l.err
	}
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Uncheck(l.selector, options...)
}

func (l *locatorImpl) WaitFor(options ...PageWaitForSelectorOptions) error {
	if l.err != nil {
		return l.err
	}
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
	if l.err != nil {
		return nil, l.err
	}
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

func (l *locatorImpl) expect(expression string, options frameExpectOptions) (*frameExpectResult, error) {
	if l.err != nil {
		return nil, l.err
	}
	overrides := map[string]interface{}{
		"selector":   l.selector,
		"expression": expression,
	}
	if options.ExpectedValue != nil {
		overrides["expectedValue"] = serializeArgument(options.ExpectedValue)
		options.ExpectedValue = nil
	}
	response, err := l.frame.channel.SendReturnAsDict("expect", options, overrides)
	if err != nil {
		return nil, err
	}
	var (
		received interface{}
		matches  bool
		log      []string
	)
	responseMap := response.(map[string]interface{})

	if v, ok := responseMap["received"]; ok {
		received = parseResult(v)
	}
	if v, ok := responseMap["matches"]; ok {
		matches = v.(bool)
	}
	if v, ok := responseMap["log"]; ok {
		for _, l := range v.([]interface{}) {
			log = append(log, l.(string))
		}
	}
	return &frameExpectResult{Received: received, Matches: matches, Log: log}, nil
}
