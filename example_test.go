// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand64_test

import (
	"fmt"
	"github.com/wildservices/rand64"
	"github.com/wildservices/rand64/xorshift"
	"math/rand"
)

const (
	SEED1 = 1387366483214
)

func Example() {
	// simple testing function
	// takes a rand64.Source64 and gets a bunch of numbers using
	// math.rand and rand64
	testfunc := func(name string, s rand64.Source64) {
		fmt.Println(name)
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
		for i := 0; i < 10; i++ {
			fmt.Printf(" %d%d", r64.Uintn(6)+1, r64.Uintn(6)+1)
		}
		fmt.Println("")

		// Source64 can also be used transparently as a initial source for rand.New()
		r := rand.New(s)
		for i := 0; i < 4; i++ {
			fmt.Printf(" %d", r.Int63())
		}
		fmt.Println("")
	}

	// create a new xorshift64+ source
	s := xorshift.New64star()
	// Seed it before use
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	s.Seed64(SEED1)

	testfunc("xorshift64*", s)

	// test the other two PRNGs
	s = xorshift.New128plus()
	s.Seed64(SEED1)
	testfunc("xorshift128+", s)
	s = xorshift.New1024star()
	s.Seed64(SEED1)
	testfunc("xorshift1024*", s)

	// Output:
	// xorshift64*
	//  4252968640 1567930103 1103871594 100834224
	//  11703131014891570448 16052167272083700520 16375787158461752832 555913475760386374
	//  14 55 63 53 55 53 54 55 31 63
	//  2491964529228605837 2984422860736700565 8255282211660800230 6933471558157501315
	// xorshift128+
	//  1518137344 519821205 3300663698 3965912936
	//  6835817718900859006 18347342841597239847 11478842834840845837 6143230513443926046
	//  35 63 66 54 12 23 21 62 62 16
	//  7186098100741600992 419266794253563720 4425471115904323875 4503902914122463897
	// xorshift1024*
	//  195155788 2037024496 1994874030 3867722788
	//  7544033655184947852 3904249604014777577 4836159697542621750 17890123175292291115
	//  11 61 54 55 31 16 11 32 35 31
	//  2545656951200074019 4059416866932157570 8849131518188499907 379980398101807463
}
