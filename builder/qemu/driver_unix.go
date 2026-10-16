// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build unix

package qemu

import "os"

func terminateProcess(p *os.Process) error {
	return p.Signal(os.Interrupt)
}
