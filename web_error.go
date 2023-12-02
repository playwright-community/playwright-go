package playwright

type webErrorImpl struct {
	err  error
	page Page
}

func (e *webErrorImpl) Page() Page {
	return e.page
}

func (e *webErrorImpl) Error() error {
	return e.err
}

func newWebError(page Page, err error) WebError {
	return &webErrorImpl{
		err:  err,
		page: page,
	}
}
