// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build windows

package qemu

import (
	"os"

	"golang.org/x/sys/windows"
)

func terminateProcess(p *os.Process) error {
	return windows.GenerateConsoleCtrlEvent(windows.CTRL_C_EVENT, uint32(p.Pid))
}
