// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package xoshiro provides an implementation for a pseudo-random number
generator (PRNG) using the xoshiro256** and xoshiro256+ algorithms.

Period: 2^256-1. State size: 256 bits.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna. For further information: http://xoshiro.di.unimi.it/
*/
package xoshiro

import (
	"math/bits"

	"github.com/db47h/rand64/v3/splitmix64"
)

// Rng256SS encapsulates a xoshiro256** PRNG.
//
// xoshiro256** 1.0 is Black,an & Vigna's all-purpose, rock-solid generator. It
// has excellent (sub-ns) speed, a state (256 bits) that is large enough for any
// parallel application, and it passes all tests the authors are aware of.
//
// For generating just floating-point numbers, xoshiro256+ is even faster.
//
type Rng256SS [4]uint64

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng256SS) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng[0] = src.Uint64()
	rng[1] = src.Uint64()
	rng[2] = src.Uint64()
	rng[3] = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng256SS) Uint64() uint64 {
	result := bits.RotateLeft64(rng[1]*5, 7) * 9

	t := rng[1] << 17

	rng[2] ^= rng[0]
	rng[3] ^= rng[1]
	rng[1] ^= rng[2]
	rng[0] ^= rng[3]

	rng[2] ^= t

	rng[3] = bits.RotateLeft64(rng[3], 45)

	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng256SS) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Rng256P encapsulates a xoshiro256+ PRNG.
//
// xoshiro256+ 1.0 is Blackman & Vigna's best and fastest generator for
// floating-point numbers. The authors suggest to use its upper bits for
// floating-point generation, as it is slightly faster than xoshiro256**. It
// passes all tests the authors are aware of except for the lowest three bits,
// which might fail linearity tests (and just those), so if low linear
// complexity is not considered an issue (as it is usually the case) it can be
// used to generate 64-bit outputs, too.
//
// The authors suggest to use a sign test to extract a random Boolean value, and
// right shifts to extract subsets of bits.
//
// Note that the Go implementation of Rand.Float64 uses the upper bits as suggested.
//
type Rng256P [4]uint64

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng256P) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng[0] = src.Uint64()
	rng[1] = src.Uint64()
	rng[2] = src.Uint64()
	rng[3] = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng256P) Uint64() uint64 {
	result := rng[0] + rng[3]
	t := rng[1] << 17

	rng[2] ^= rng[0]
	rng[3] ^= rng[1]
	rng[1] ^= rng[2]
	rng[0] ^= rng[3]

	rng[2] ^= t

	rng[3] = bits.RotateLeft64(rng[3], 45)

	return result
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng256P) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
