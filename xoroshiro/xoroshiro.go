// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package xoroshiro provides an implementation for a pseudo-random number
generator (PRNG) using the xoroshiro128** and xoroshiro128+ algorithms.

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
// xoroshiro128+ 1.0 is Blackman & Vigna's best and fastest small-state
// generator for floating-point numbers. They suggest to use its upper bits for
// floating-point generation, as it is slightly faster than xoroshiro128**. It
// passes all tests the authors are aware of except for the four lower bits,
// which might fail linearity tests (and just those), so if low linear
// complexity is not considered an issue (as it is usually the case) it can be
// used to generate 64-bit outputs, too; moreover, this generator has a very
// mild Hamming-weight dependency making our test
// (http://prng.di.unimi.it/hwd.php) fail after 5 TB of output; the authors
// believe this slight bias cannot affect any application. If you are concerned,
// use xoroshiro128** or xoshiro256+.
//
// The authors suggest to use a sign test to extract a random Boolean value, and
// right shifts to extract subsets of bits.
//
// Note that the Go implementation of Rand.Float64 uses the upper bits as suggested.
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
// xoroshiro128** 1.0 is Blackman & Vigna's all-purpose, rock-solid, small-state
// generator. It is extremely (sub-ns) fast and it passes all tests the authors
// are aware of, but its state space is large enough only for mild parallelism.
//
// For generating just floating-point numbers, xoroshiro128+ is even faster (but
// it has a very mild bias, see notes in the comments).
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
