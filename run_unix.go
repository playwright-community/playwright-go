//go:build !windows

package playwright

import "syscall"

var defaultSysProcAttr = &syscall.SysProcAttr{}

// for WritableStream.Copy
const defaultCopyBufSize = 1024 * 1024
