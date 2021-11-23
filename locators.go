package playwright

type locatorImpl struct {
	frame    Frame
	selector string
}

func (l *locatorImpl) Check(options ...FrameCheckOptions) error {
	option := FrameCheckOptions{}
	if len(options) == 0 {
		option = FrameCheckOptions{
			Strict: Bool(true),
		}
	} else {
		option = options[0]
		option.Strict = Bool(true)
	}
	return l.frame.Check(l.selector, option)
}

func (l *locatorImpl) Click(options ...PageClickOptions) error {
	option := PageClickOptions{}
	if len(options) == 0 {
		option = PageClickOptions{
			Strict: Bool(true),
		}
	} else {
		option = options[0]
		option.Strict = Bool(true)
	}
	return l.frame.Click(l.selector, option)
}

func (l *locatorImpl) DispatchEvent(typ string, eventInit interface{}, options ...PageDispatchEventOptions) {
	l.frame.DispatchEvent(l.selector, typ, eventInit, options...)
}

func (l *locatorImpl) Dblclick(options ...FrameDblclickOptions) error {
	option := FrameDblclickOptions{}
	if len(options) == 0 {
		option = FrameDblclickOptions{
			Strict: Bool(true),
		}
	} else {
		option = options[0]
	}
	return l.frame.Dblclick(l.selector, option)
}

func (l *locatorImpl) elementHandle(timeout float64) (ElementHandle, error) {
	return l.frame.WaitForSelector(l.selector, PageWaitForSelectorOptions{
		Timeout: Float(timeout),
		Strict:  Bool(true),
		State:   WaitForSelectorStateAttached,
	})
}

func (l *locatorImpl) elementHandles(timeout float64) ([]ElementHandle, error) {
	return l.frame.QuerySelectorAll(l.selector)
}

func newLocator(frame Frame, selector string) *locatorImpl {
	return &locatorImpl{
		frame:    frame,
		selector: selector,
	}
}
