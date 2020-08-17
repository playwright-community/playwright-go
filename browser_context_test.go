package playwright

import (
	"testing"
)

func TestNewPage(t *testing.T) {
	helper := NewTestHelper(t)
	helper.Close(t)
}
