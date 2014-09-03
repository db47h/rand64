// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand64_test

import (
	"fmt"
	"math/rand"
	"github.com/wildservices/rand64"
)

const (
	SEED1 = 1387366483214
)

func Example() {
	// simple testing function 
	// takes a rand64.Source64 and gets a bunch of numbers using
	// math.rand and rand64
	testfunc := func(name string, s rand64.Source64) {
		// Source64 can be used transparently as a initial source for rand.New()
		r := rand.New(s)

		fmt.Println(name)
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r.Int31())
		}
		fmt.Println("")
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r.Int63())
		}
		fmt.Println("")

		// Using Rand64
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
		for i:= 0; i < 10; i++ {
			fmt.Printf(" %d%d", r64.Uintn(6)+1, r64.Uintn(6)+1)
		}
		fmt.Println("")

	}

	// create a new xorshift64+ source and seed it
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	s := rand64.NewXorShift64star(SEED1)

	testfunc("xorshift64*", s)

	// test the other two PRNGs
	s = rand64.NewXorShift128plus(SEED1)
	testfunc("xorshift128+", s)
	s = rand64.NewXorShift1024star(SEED1)
	testfunc("xorshift1024*", s)

	// Output: 
	// xorshift64*
	//  2126484320 783965051 551935797 50417112
	//  5851565507445785224 8026083636041850260 8187893579230876416 277956737880193187
	//  2076756930 421197795 2990677840 2457281038
	//  13952525632776470525 7844513538077537461 12960612483867938010 13344456975593097334
	//  55 53 54 55 31 63 54 42 63 62
	// xorshift128+
	//  759068672 259910602 1650331849 1982956468
	//  3417908859450429503 9173671420798619923 5739421417420422918 3071615256721963023
	//  3483664268 1300950988 1543694279 2963043518
	//  5784970938708832231 13898142144918159171 13709112874101900137 12630090110977981649
	//  12 23 21 62 62 16 36 53 33 44
	// xorshift1024*
	//  97577894 1018512248 997437015 1933861394
	//  3772016827592473926 1952124802007388788 2418079848771310875 8945061587646145557
	//  1264365492 1651718160 1955547191 4057194534
	//  17068160306806490531 9601692629104924013 16655558537308429035 14454256620255781544
	//  31 16 11 32 35 31 16 52 36 16
}
