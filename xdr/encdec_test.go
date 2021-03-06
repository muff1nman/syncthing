// Copyright (C) 2014 Jakob Borg and Contributors (see the CONTRIBUTORS file).
// All rights reserved. Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package xdr_test

import (
	"bytes"
	"testing"
	"testing/quick"
)

// Contains all supported types
type TestStruct struct {
	I    int
	I16  int16
	UI16 uint16
	I32  int32
	UI32 uint32
	I64  int64
	UI64 uint64
	BS   []byte
	S    string
}

func TestEncDec(t *testing.T) {
	fn := func(t0 TestStruct) bool {
		bs := t0.MarshalXDR()
		var t1 TestStruct
		err := t1.UnmarshalXDR(bs)
		if err != nil {
			t.Fatal(err)
		}

		// Not comparing with DeepEqual since we'll unmarshal nil slices as empty
		if t0.I != t1.I ||
			t0.I16 != t1.I16 || t0.UI16 != t1.UI16 ||
			t0.I32 != t1.I32 || t0.UI32 != t1.UI32 ||
			t0.I64 != t1.I64 || t0.UI64 != t1.UI64 ||
			bytes.Compare(t0.BS, t1.BS) != 0 ||
			t0.S != t1.S {
			t.Logf("%#v", t0)
			t.Logf("%#v", t1)
			return false
		}
		return true
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}
