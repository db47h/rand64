// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xorshift128 provides an implementation for a pseudo-random number
generator (PRNG) using the xorshift128+ algorithm.

Period: 2^128-1. State size: 128 bits.

Go implementation based on a C reference implementation by Sebastiano Vigna.
For further information: http://xorshift.di.unimi.it/

This generator has been superseeded by xoroshiro128+, which is significantly
faster and has better statistical properties.

The lowest bit is a LFSR (linear-feeedback shift register), and thus it is
slightly less random than the other bits. It is therefore recommended to use a
sign test in order to extract boolean values (i.e. check the high order bit).
*/
package xorshift128

import (
	"github.com/db47h/rand64/splitmix64"
)

// Rng encapsulates a xsorshift128+ PRNG.
//
type Rng struct {
	state [2]uint64
}

// Seed uses the provided uint64 seed value to initialize the generator to a deterministic state.
//
func (rng *Rng) Seed(seed int64) {
	s := splitmix64.Rng{State: uint64(seed)}
	rng.state[0] = s.Uint64()
	rng.state[1] = s.Uint64()
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
//
func (rng *Rng) Uint64() uint64 {
	s1 := rng.state[0]
	s0 := rng.state[1]
	result := s0 + s1
	rng.state[0] = s0
	s1 ^= s1 << 23                                  // a
	rng.state[1] = s1 ^ s0 ^ (s1 >> 18) ^ (s0 >> 5) // b, c
	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
