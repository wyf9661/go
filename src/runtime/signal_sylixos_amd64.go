// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sylixos && amd64

package runtime

import (
	"unsafe"
)

type sigctxt struct {
	info *siginfo
	ctxt unsafe.Pointer
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) regs() *sigcontext { return (*sigcontext)(c.ctxt) }

func (c *sigctxt) rax() uint64 { return c.regs().rax }
func (c *sigctxt) rbx() uint64 { return c.regs().rbx }
func (c *sigctxt) rcx() uint64 { return c.regs().rcx }
func (c *sigctxt) rdx() uint64 { return c.regs().rdx }
func (c *sigctxt) rdi() uint64 { return c.regs().rdi }
func (c *sigctxt) rsi() uint64 { return c.regs().rsi }
func (c *sigctxt) rbp() uint64 { return c.regs().rbp }
func (c *sigctxt) rsp() uint64 { return c.regs().rsp }
func (c *sigctxt) r8() uint64  { return c.regs().r8 }
func (c *sigctxt) r9() uint64  { return c.regs().r9 }
func (c *sigctxt) r10() uint64 { return c.regs().r10 }
func (c *sigctxt) r11() uint64 { return c.regs().r11 }
func (c *sigctxt) r12() uint64 { return c.regs().r12 }
func (c *sigctxt) r13() uint64 { return c.regs().r13 }
func (c *sigctxt) r14() uint64 { return c.regs().r14 }
func (c *sigctxt) r15() uint64 { return c.regs().r15 }

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) rip() uint64 { return c.regs().rip }

func (c *sigctxt) rflags() uint64 { return c.regs().rflags }
func (c *sigctxt) cs() uint64     { return uint64(c.regs().cs) }
func (c *sigctxt) ss() uint64     { return uint64(c.regs().ss) }

// ACOINFO TODO
func (c *sigctxt) fs() uint64 { return uint64(0) }
func (c *sigctxt) gs() uint64 { return uint64(0) }

func (c *sigctxt) sigcode() uint64 { return uint64(c.info.si_code) }
func (c *sigctxt) sigaddr() uint64 { return c.info.si_addr }

func (c *sigctxt) set_rip(x uint64)     { c.regs().rip = x }
func (c *sigctxt) set_rsp(x uint64)     { c.regs().rsp = x }
func (c *sigctxt) set_sigcode(x uint64) { c.info.si_code = int32(x) }
func (c *sigctxt) set_sigaddr(x uint64) { c.info.si_addr = x }
