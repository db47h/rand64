// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorshift1024_test

import (
	"fmt"
	"math/rand"

	"github.com/db47h/rand64/v3/xorshift1024"
)

const (
	SEED1 = 1387366483214
)

/* Tests */
func Example() {
	// simple testing function
	// takes a rand64.Source64 and gets a bunch of numbers using
	// math.rand
	testfunc := func(name string, s rand.Source64) {
		fmt.Println(name)
		r64 := rand.New(s)
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r64.Uint32())
		}
		fmt.Println("")
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r64.Uint64())
		}
		fmt.Println("")
		// Play craps
		for i := 0; i < 10; i++ {
			fmt.Printf(" %d%d", r64.Intn(6)+1, r64.Intn(6)+1)
		}
		fmt.Println("")
	}

	// test xorshift1024*
	var s = &xorshift1024.Rng{}
	s.Seed(SEED1)
	testfunc("xorshift1024*", s)

	// Output:
	// xorshift1024*
	//  3332849200 1164738618 456220800 3523432244
	//  10311270752396438174 3766918502733849924 15396074446274990069 15679784721060022461
	//  26 55 15 62 52 26 16 66 34 52
}
