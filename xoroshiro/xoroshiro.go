// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xoroshiro provides an implementation for a pseudo-random number
generator (PRNG) using the xoshiro128plus algorithm.

Period: 2^128-1. State size: 128 bits.

This is the successor to xorshift128+. It is the fastest full-period generator
passing BigCrush without systematic failures, but due to the relatively short
period it is acceptable only for applications with a mild amount of parallelism;
otherwise, use a xorshift1024* generator.

Beside passing BigCrush, this generator passes the PractRand test suite up to
(and included) 16TB, with the exception of binary rank tests, which fail due to
the lowest bit being an LFSR; all other bits pass all tests. We suggest to use a
sign test to extract a random Boolean value.

The state must be seeded so that it is not everywhere zero. If you have only a
64-bit seed, the Seed() function will seed an internal splitmix64 generator and
use its output to fill the initial state of the xoroshiro generator.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna. For further information: http://xorshift.di.unimi.it/
*/
package xoroshiro

import (
	"github.com/db47h/rand64"
	"github.com/db47h/rand64/internal/util"
)

type xrsr128plus [2]uint64

// New128plus returns a new pseudo-random number Source using the xoroshiro128+ algorithm.
func New128plus() rand64.Source {
	return new(xrsr128plus)
}

func (rng *xrsr128plus) Seed(seed uint64) {
	util.SeedSlice((*rng)[:], seed)
}

func (rng *xrsr128plus) SeedFromSlice(seed []uint64) {
	util.SeedFromSlice(rng[:], seed)
}

func (rng *xrsr128plus) Uint64() uint64 {
	s0 := (*rng)[0]
	s1 := (*rng)[1]
	result := s0 + s1

	s1 ^= s0
	// go 1.6.2 and 1.7 properly assemble the pattern (x << n) | (x >> (64 - n))
	// to a ROLQ n, x on x86_64. Earlier versions not tested.
	(*rng)[0] = ((s0 << 55) | (s0 >> (64 - 55))) ^ s1 ^ (s1 << 14)
	(*rng)[1] = ((s1 << 36) | (s1 >> (64 - 36)))

	return result
}
