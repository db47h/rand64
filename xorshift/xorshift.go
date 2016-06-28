// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xorshift provides implementations for pseudo-random number
generators (PRNG) using scrambled xorshift algorithms, with two of them passing
the BigCrunch battery test of the TestU01 test suite without systematic errors.

Go implementation based on a C reference implementation by Sebastiano Vigna.
For further information: http://xorshift.di.unimi.it/

xorshift128+

Period: 2^128-1. State size: 128 bits.

This generator has been replaced by xoroshiro128+, which is significantly
faster and has better statistical properties.

This is the fastest generator of the xorshift series, passing BigCrush without
systematic errors, but due to the relatively short period it is acceptable only
for applications with a very mild amount of parallelism; otherwise, use
a xorshift1024* generator.

xorshift1024*

Period: 2^1024-1. State size: 1024 bits (16 uint64).

This is a fast, top-quality generator, also passing BigCrunch without systematic
errors. If 1024 bits of state are too much, try a xorshift128+ or xorshift64*
generator.
*/
package xorshift

import (
	"github.com/db47h/rand64"
	"github.com/db47h/rand64/internal/util"
)

// xorshift128+
// this is the fastest generator passing BigCrush without systematic
// errors, but due to the relatively short period it is acceptable only
// for applications with a very mild amount of parallelism; otherwise, use
// a xorshift1024* generator
type xs128plus [2]uint64

// New128plus returns a new pseudo-random number Source using the xorshift128+ algorithm.
func New128plus() rand64.Source {
	return new(xs128plus)
}

// Seed uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs128plus) Seed(seed uint64) {
	util.SeedSlice((*rng)[:], seed)
}

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (rng *xs128plus) SeedFromSlice(seed []uint64) {
	util.SeedFromSlice(rng[:], seed)
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs128plus) Uint64() uint64 {
	s1 := rng[0]
	s0 := rng[1]
	result := s0 + s1
	rng[0] = s0
	s1 ^= s1 << 23                            // a
	rng[1] = s1 ^ s0 ^ (s1 >> 18) ^ (s0 >> 5) // b, c
	return result
}

// xorshift1024*
// This is a fast, top-quality generator. If 1024 bits of state are too
// much, try a xorshift128+ or a xorshift64* generator.
type xs1024star struct {
	state [16]uint64
	p     int
}

// New1024star returns a new pseudo-random number Source using the xorshift1024* algorithm.
func New1024star() rand64.Source {
	return new(xs1024star)
}

// Seed uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs1024star) Seed(seed uint64) {
	util.SeedSlice(rng.state[:], seed)
}

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (rng *xs1024star) SeedFromSlice(seed []uint64) {
	util.SeedFromSlice(rng.state[:], seed)
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs1024star) Uint64() uint64 {
	s := rng.state[:]
	s0 := s[rng.p]
	rng.p = (rng.p + 1) & 15
	s1 := s[rng.p]
	s1 ^= s1 << 31                               // a
	s[rng.p] = s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30) // b,c
	return s[rng.p] * 1181783497276652981
}
