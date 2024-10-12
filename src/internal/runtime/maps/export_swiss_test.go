// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.swissmap

package maps

import (
	"internal/abi"
	"unsafe"
)

func NewTestTable[K comparable, V any](length uint64) *table {
	var m map[K]V
	mTyp := abi.TypeOf(m)
	mt := (*abi.SwissMapType)(unsafe.Pointer(mTyp))
	return newTable(mt, length)
}