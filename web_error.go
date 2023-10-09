package playwright

type webErrorImpl struct {
	err  *Error
	page Page
}

func (e *webErrorImpl) Page() Page {
	return e.page
}

func (e *webErrorImpl) Error() error {
	return e.err
}

func newWebError(page Page, err *Error) WebError {
	return &webErrorImpl{
		err:  err,
		page: page,
	}
}
