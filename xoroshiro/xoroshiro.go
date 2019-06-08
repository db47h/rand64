// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xoroshiro provides an implementation for a pseudo-random number
generator (PRNG) using the xoroshiro128** and xoroshiro128+ algorithm.

Period: 2^128-1. State size: 128 bits.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna. For further information: http://xoshiro.di.unimi.it/
*/
package xoroshiro

import (
	"math/bits"

	"github.com/db47h/rand64/v3/splitmix64"
)

// Rng128P encapsulates a xoroshiro128+ PRNG.
//
type Rng128P struct {
	s0, s1 uint64
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng128P) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng.s0 = src.Uint64()
	rng.s1 = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng128P) Uint64() uint64 {
	s0 := rng.s0
	s1 := rng.s1
	result := s0 + s1
	s1 ^= s0
	rng.s0 = bits.RotateLeft64(s0, 24) ^ s1 ^ (s1 << 16)
	rng.s1 = bits.RotateLeft64(s1, 37)

	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng128P) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Rng128SS encapsulates a xoroshiro128** PRNG.
//
type Rng128SS struct {
	s0, s1 uint64
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng128SS) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng.s0 = src.Uint64()
	rng.s1 = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng128SS) Uint64() uint64 {
	s0 := rng.s0
	s1 := rng.s1
	result := bits.RotateLeft64(s0*5, 7) * 9

	s1 ^= s0
	rng.s0 = bits.RotateLeft64(s0, 24) ^ s1 ^ (s1 << 16)
	rng.s1 = bits.RotateLeft64(s1, 37)

	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng128SS) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
