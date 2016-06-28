// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package splitmix64_test

import (
	"testing"

	"github.com/db47h/rand64/splitmix64"
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

func TestNew(t *testing.T) {
	rng := splitmix64.New(SEED1)
	for _, v := range values {
		n := rng.Uint64()
		if n != v {
			t.Fatalf("Expected %X, got %X", v, n)
		}
	}
}
