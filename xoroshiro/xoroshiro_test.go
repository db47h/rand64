// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xoroshiro_test

import (
	"fmt"

	"github.com/db47h/rand64"
	"github.com/db47h/rand64/xoroshiro"
)

const (
	SEED1 = 1387366483214
)

func Example() {
	src := xoroshiro.New128plus()
	src.Seed(SEED1)
	rng := rand64.New(src)
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint32())
	}
	fmt.Println("")
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint64())
	}
	fmt.Println("")
	// Play craps
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d%d", rng.Uintn(6)+1, rng.Uintn(6)+1)
	}

	// Output:
	//  3672052799 3619036596 1817626404 4154021231
	//  13508242557925574888 11509836612120350102 17607668528363997996 9787171209907982739
	//  13 64 44 22 66 61 16 46 22 43
}
