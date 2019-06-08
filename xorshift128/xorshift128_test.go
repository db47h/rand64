// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorshift128_test

import (
	"fmt"
	"math/rand"

	"github.com/db47h/rand64/v2/xorshift128"
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

	s := &xorshift128.Rng{}
	s.Seed(SEED1)
	testfunc("xorshift128+", s)

	// Output:
	// xorshift128+
	//  3672052799 2300942069 2356831912 2316732845
	//  5560898047753517047 9806550241747869425 16344204150069124721 7133254478284829050
	//  35 41 43 12 15 32 45 45 42 64
}
