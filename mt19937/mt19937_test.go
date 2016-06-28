// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mt19937_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/db47h/rand64/mt19937"
)

const (
	SEED1 = 1387366483214
)

func ExampleMt19937_Uint64() {
	init := []uint64{
		0x12345, 0x23456, 0x34567, 0x45678,
	}
	mt := mt19937.New()
	mt.SeedFromSlice(init)

	fmt.Println("10 outputs of mt19937.Uint64()")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %20d", mt.Uint64())
		if i%5 == 4 {
			fmt.Println()
		}
	}
	for i := 0; i < 1000; i++ {
		mt.Uint64()
	}
	fmt.Println("10 more")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %20d", mt.Uint64())
		if i%5 == 4 {
			fmt.Println()
		}
	}

	// Output:
	// 10 outputs of mt19937.Uint64()
	//   7266447313870364031  4946485549665804864 16945909448695747420 16394063075524226720  4873882236456199058
	//  14877448043947020171  6740343660852211943 13857871200353263164  5249110015610582907 10205081126064480383
	// 10 more
	//  14907209235746902445 15452338815569321965 17045090235069538607 15507333859934612093   157175897107904252
	//   2578005313950236321  6502648805754593060 13133523174961431106  2698278206396822833  3278969850082110371
}

/* Benchmarks */

func BenchmarkMt19937(b *testing.B) {
	s := mt19937.New()
	s.Seed64(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkRandSource(b *testing.B) {
	s := rand.NewSource(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}
