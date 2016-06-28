// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorshift_test

import (
	"fmt"
	"math/rand"

	"github.com/db47h/rand64"
	"github.com/db47h/rand64/xorshift"
)

const (
	SEED1 = 1387366483214
)

/* Tests */
func Example() {
	// simple testing function
	// takes a rand64.Source and gets a bunch of numbers using
	// math.rand and rand64
	testfunc := func(name string, s rand64.Source) {
		fmt.Println(name)
		// Using Rand
		r64 := rand64.New(s)
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
			fmt.Printf(" %d%d", r64.Uintn(6)+1, r64.Uintn(6)+1)
		}
		fmt.Println("")

		// since rand64.Rand implements rand.Source, it can
		// be used to proxy a Source to rand.Rand
		r := rand.New(r64)
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r.Int63())
		}
		fmt.Println("")
	}

	// create a new xorshift128+ source
	var s = xorshift.New128plus()
	// Seed it before use
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	s.Seed(SEED1)

	// test it
	testfunc("xorshift128+", s)

	// test xorshift1024*
	s = xorshift.New1024star()
	s.Seed(SEED1)
	testfunc("xorshift1024*", s)

	// Output:
	// xorshift128+
	//  3672052799 2300942069 2356831912 2316732845
	//  5560898047753517047 9806550241747869425 16344204150069124721 7133254478284829050
	//  64 11 26 14 23 64 23 23 13 52
	//  802413841702959598 2251227033134975276 4988225046549461352 2188638676389822986
	// xorshift1024*
	//  3332849200 1164738618 456220800 3523432244
	//  10311270752396438174 3766918502733849924 15396074446274990069 15679784721060022461
	//  36 33 14 53 43 46 16 66 61 33
	//  8307002403806671045 2041967637359636555 2088934487125395476 7936776852298221278
}
