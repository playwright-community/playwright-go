package playwright

import "strconv"

type frameLocatorImpl struct {
	frame         *frameImpl
	frameSelector string
}

func newFrameLocator(frame *frameImpl, frameSelector string) *frameLocatorImpl {
	return &frameLocatorImpl{frame: frame, frameSelector: frameSelector}
}

func (fl *frameLocatorImpl) First() FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth=0")
}

func (fl *frameLocatorImpl) FrameLocator(selector string) FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> internal:control=enter-frame >> "+selector)
}

func (fl *frameLocatorImpl) GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByAltTextSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByLabelSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByPlaceholderSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator {
	return fl.Locator(getByRoleSelector(role, options...))
}

func (fl *frameLocatorImpl) GetByTestId(testId interface{}) Locator {
	return fl.Locator(getByTestIdSelector(getTestIdAttributeName(), testId))
}

func (fl *frameLocatorImpl) GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByTextSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByTitle(text interface{}, options ...LocatorGetByTitleOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByTitleSelector(text, exact))
}

func (fl *frameLocatorImpl) Last() FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth=-1")
}

func (fl *frameLocatorImpl) Locator(selector string, options ...FrameLocatorOptions) Locator {
	var option LocatorLocatorOptions
	if len(options) == 1 {
		option = LocatorLocatorOptions(options[0])
	}
	return newLocator(fl.frame, fl.frameSelector+" >> internal:control=enter-frame >> "+selector, option)
}

func (fl *frameLocatorImpl) Nth(index int) FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth="+strconv.Itoa(index))
}
