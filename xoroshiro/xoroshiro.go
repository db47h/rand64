// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xoroshiro provides an implementation for a pseudo-random number
generator (PRNG) using the xoshiro128plus algorithm.

Period: 2^128-1. State size: 128 bits.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna. For further information: http://xorshift.di.unimi.it/

This is the successor to xorshift128+. It is the fastest full-period generator
passing BigCrush without systematic failures, but due to the relatively short
period it is acceptable only for applications with a mild amount of parallelism;
otherwise, use a xorshift1024* generator.

The lowest bit is a LFSR (linear-feeedback shift register), and thus it is
slightly less random than the other bits. It is therefore recommended to use a
sign test in order to extract boolean values (i.e. check the high order bit).
*/
package xoroshiro

import (
	"math/bits"

	"github.com/db47h/rand64/splitmix64"
)

// Rng encapsulates a splitmix64 PRNG.
//
type Rng struct {
	s0, s1 uint64
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng.s0 = src.Uint64()
	rng.s1 = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng) Uint64() uint64 {
	s0 := rng.s0
	s1 := rng.s1
	result := s0 + s1
	s1 ^= s0
	rng.s1 = bits.RotateLeft64(s1, 36)
	rng.s0 = bits.RotateLeft64(s0, 55) ^ s1 ^ (s1 << 14)

	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
