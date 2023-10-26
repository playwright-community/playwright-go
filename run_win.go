//go:build windows

package playwright

import "syscall"

var defaultSysProcAttr = &syscall.SysProcAttr{HideWindow: true}
