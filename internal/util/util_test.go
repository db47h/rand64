// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util_test

import (
	"testing"

	"github.com/db47h/rand64/internal/util"
)

const (
	SEED1 = 1387366483214
)

var values = [...]uint64{
	0xDDE04155BF79DF63,
	0xFCFED2E9D540B529,
	0x4C5AA74B9BE7FF3E,
	0xA38A0EF197E488D9,
	0xEDA0BA12AA8B5343,
	0x94AC0EE844BA7CB6,
	0x644375EBE6F55AAF,
	0xBD7DF1EF1C84093D,
	0xDBDB00E0A41BE9AB,
	0xC7A8EB53EB467566,
}

func TestSeedSlice(t *testing.T) {
	var s [10]uint64
	util.SeedSlice(s[:], SEED1)
	for i := range s {
		if s[i] != values[i] {
			t.Fatalf("At index %d, got 0x%X, expected 0x%X", i, s[i], values[i])
		}
	}
}

func TestSeedFromSlice(t *testing.T) {
	var dst [10]uint64
	var src = []uint64{1, SEED1}

	util.SeedFromSlice(dst[:], src)
	if dst[0] != 1 || dst[1] != SEED1 {
		t.Fatal()
	}
	for i, v := range dst[2:] {
		if v != values[i] {
			t.Fatalf("At index %d, got 0x%X, expected 0x%X", i+2, v, values[i])
		}
	}
}
