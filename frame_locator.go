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
	return newFrameLocator(fl.frame, fl.frameSelector+" >> control=enter-frame >> "+selector)
}

func (fl *frameLocatorImpl) Last() FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth=-1")
}

func (fl *frameLocatorImpl) Locator(selector string, options ...LocatorLocatorOptions) (Locator, error) {
	return newLocator(fl.frame, fl.frameSelector+" >> control=enter-frame >> "+selector, options...)
}

func (fl *frameLocatorImpl) Nth(index int) FrameLocator {
	return newFrameLocator(fl.frame, fl.frameSelector+" >> nth="+strconv.Itoa(index))
}
