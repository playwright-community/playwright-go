package playwright

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

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
				selector += fmt.Sprintf(" >> :scope:text-matches(%s)", convertRegexp(hasText))
			case string:
				selector += fmt.Sprintf(" >> :scope:has-text('%s')", hasText)
			}
		}
		if option.Has != nil {
			has := option.Has.(*locatorImpl)
			if frame != has.frame {
				return nil, errors.New("inner 'has' locator must belong to the same frame")
			}
			selector += " >> has=" + has.selector
		}
	}

	return &locatorImpl{frame: frame, selector: selector, options: option}, nil
}

func convertRegexp(reg *regexp.Regexp) string {
	matches := regexp.MustCompile(`\(\?([imsU]+)\)(.+)`).FindStringSubmatch(reg.String())

	var pattern, flags string
	if len(matches) == 3 {
		pattern = matches[2]
		flags = matches[1]
	} else {
		pattern = reg.String()
	}
	return fmt.Sprintf("'%s', '%s'", pattern, flags)
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

func (l *locatorImpl) First() (Locator, error) {
	return newLocator(l.frame, l.selector+" >> nth=0")
}

func (l *locatorImpl) Focus(options ...FrameFocusOptions) error {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.Focus(l.selector, options...)
}

func (l *locatorImpl) GetAttribute(name string, options ...PageGetAttributeOptions) (string, error) {
	if len(options) == 1 {
		options[0].Strict = Bool(true)
	}
	return l.frame.GetAttribute(l.selector, name, options...)
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

func (l *locatorImpl) Locator(selector string) (Locator, error) {
	return newLocator(l.frame, l.selector+" >> "+selector)
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
