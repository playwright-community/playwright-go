package playwright

import (
	"testing"
)

func TestNewPage(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}
