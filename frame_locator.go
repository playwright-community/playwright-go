package playwright

import (
	"fmt"
	"strconv"

	"github.com/playwright-community/playwright-go/internal/multierror"
)

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

func (fl *frameLocatorImpl) GetByAltText(text interface{}, options ...FrameLocatorGetByAltTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByAltTextSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByLabel(text interface{}, options ...FrameLocatorGetByLabelOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByLabelSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByPlaceholder(text interface{}, options ...FrameLocatorGetByPlaceholderOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByPlaceholderSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByRole(role AriaRole, options ...FrameLocatorGetByRoleOptions) Locator {
	if len(options) == 1 {
		return fl.Locator(getByRoleSelector(role, LocatorGetByRoleOptions(options[0])))
	}
	return fl.Locator(getByRoleSelector(role))
}

func (fl *frameLocatorImpl) GetByTestId(testId interface{}) Locator {
	return fl.Locator(getByTestIdSelector(getTestIdAttributeName(), testId))
}

func (fl *frameLocatorImpl) GetByText(text interface{}, options ...FrameLocatorGetByTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return fl.Locator(getByTextSelector(text, exact))
}

func (fl *frameLocatorImpl) GetByTitle(text interface{}, options ...FrameLocatorGetByTitleOptions) Locator {
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

func (fl *frameLocatorImpl) Locator(selectorOrLocator interface{}, options ...FrameLocatorLocatorOptions) Locator {
	var option LocatorLocatorOptions
	if len(options) == 1 {
		option = LocatorLocatorOptions(options[0])
	}

	selector, ok := selectorOrLocator.(string)
	if ok {
		return newLocator(fl.frame, fl.frameSelector+" >> internal:control=enter-frame >> "+selector, option)
	}
	locator, ok := selectorOrLocator.(*locatorImpl)
	if ok {
		if fl.frame != locator.frame {
			locator.err = multierror.Join(locator.err, ErrLocatorNotSameFrame)
			return locator
		}
		return newLocator(locator.frame,
			fmt.Sprintf("%s >> internal:control=enter-frame >> %s", fl.frameSelector, locator.selector),
			option,
		)
	}
	return &locatorImpl{
		frame:    fl.frame,
		selector: fl.frameSelector,
		err:      fmt.Errorf("invalid locator parameter: %v", selectorOrLocator),
	}
}

func (fl *frameLocatorImpl) Nth(index int) FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth="+strconv.Itoa(index))
}
