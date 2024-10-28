// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha512

import (
	"bytes"
	"crypto/internal/fips"
	"errors"
)

func init() {
	fips.CAST("SHA2-512", func() error {
		input := []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
			0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
		}
		want := []byte{
			0xb4, 0xc4, 0xe0, 0x46, 0x82, 0x6b, 0xd2, 0x61,
			0x90, 0xd0, 0x97, 0x15, 0xfc, 0x31, 0xf4, 0xe6,
			0xa7, 0x28, 0x20, 0x4e, 0xad, 0xd1, 0x12, 0x90,
			0x5b, 0x08, 0xb1, 0x4b, 0x7f, 0x15, 0xc4, 0xf3,
			0x8e, 0x29, 0xb2, 0xfc, 0x54, 0x26, 0x5a, 0x12,
			0x63, 0x26, 0xc5, 0xbd, 0xea, 0x66, 0xc1, 0xb0,
			0x8e, 0x9e, 0x47, 0x72, 0x3b, 0x2d, 0x70, 0x06,
			0x5a, 0xc1, 0x26, 0x2e, 0xcc, 0x37, 0xbf, 0xb1,
		}
		h := New()
		h.Write(input)
		if got := h.Sum(nil); !bytes.Equal(got, want) {
			return errors.New("unexpected result")
		}
		return nil
	})
}