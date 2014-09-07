// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The xorshift package provides implementations for pseudo-random number
generators (PRNG) using scrambled xorshift algorithms, with two of them passing
the BigCrunch battery test of the TestU01 test suite without systematic errors.

Algorithms by George Marsaglia, Sebastiano Vigna. Go implementation based on a
C reference implementation by S. Vigna.
For further information: http://xorshift.di.unimi.it/

xorshift128+

Period: 2^128-1. State size: 2 uint64.

This is the fastest generator of the series, passing BigCrush without systematic
errors, but due to the relatively short period it is acceptable only
for applications with a very mild amount of parallelism; otherwise, use
a xorshift1024* generator.

xorshift1024*

Period: 2^1024-1. State size: 16 uint64.

This is a fast, top-quality generator, also passing BigCrunch without systematic
errors. If 1024 bits of state are too much, try a xorshift128+ or xorshift64*
generator.

xorshift64*

Period: 2^64-1. State size: 1 uint64.

This is a decent generator (failing BigCrunch only on the MatrixRank test). It
is used internally to seed the state buffers for the other algorithms. Using
it as a general purpose PRNG is however not recommended since xorshift128+ is
noticably faster with better output quality and a much longer period.

PRNGs are created by calling New<algorithm name>() and must be seeded before use.

Benchmarks

The last result is for the default PRNG provided by rand.NewSource() for comparison:

	BenchmarkXorShift64star     100000000       11.7 ns/op
	BenchmarkXorShift128plus    500000000        7.3 ns/op
	BenchmarkXorShift1024star   100000000       12.5 ns/op

	BenchmarkRandSource         100000000       11.4 ns/op

TODO

xorshift4096* implementation.
*/
package xorshift

import (
	"github.com/wildservices/rand64"
)

// helper function to seed a state array dst given a slice of src uint64
// if the source slice is too short, we complete with values generated
// by an xorshift64* PRNG
func seedArray(dst, src []uint64) {
	i := copy(dst, src)
	// fill in the missing bits
	if i < len(dst) {
		var seed uint64
		if i > 0 {
			seed = dst[i-1]
		}
		s64 := New64star()
		s64.Seed64(seed)
		for ; i < len(dst); i++ {
			dst[i] = s64.Uint64()
		}
	}
}

// xorshift64*
// This is a good generator if you're short on memory, but otherwise we
// rather suggest to use a xorshift128+ (for maximum speed) or
// xorshift1024* (for speed and very long period) generator.
type xs64star uint64

// New64star returns a new pseudo-random number Source64 using the xorshift64* algorithm.
func New64star() rand64.Source64 {
	var rng xs64star
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (s *xs64star) Seed64(seed uint64) {
	if seed == 0 {
		seed = 89482311
	}
	*s = xs64star(seed)
}

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (s *xs64star) SeedFromSlice(seed []uint64) {
	if len(seed) == 0 {
		s.Seed64(0)
	} else {
		s.Seed64(seed[0])
	}
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (s *xs64star) Uint64() uint64 {
	*s ^= *s >> 12 // a
	*s ^= *s << 25 // b
	*s ^= *s >> 27 // c
	return uint64(*s * 2685821657736338717)
}

// xorshift128+
// this is the fastest generator passing BigCrush without systematic
// errors, but due to the relatively short period it is acceptable only
// for applications with a very mild amount of parallelism; otherwise, use
// a xorshift1024* generator
type xs128plus [2]uint64

// New128plus returns a new pseudo-random number Source64 using the xorshift128+ algorithm.
func New128plus() rand64.Source64 {
	var rng xs128plus
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs128plus) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	rng.SeedFromSlice([]uint64{seed}) // TODO: a better way to setup s[]?
}

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (rng *xs128plus) SeedFromSlice(seed []uint64) {
	seedArray(rng[:], seed)
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs128plus) Uint64() uint64 {
	s1 := rng[0]
	s0 := rng[1]
	rng[0] = s0
	s1 ^= s1 << 23                               // a
	rng[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26)) // b, c
	return rng[1] + s0
}

// xorshift1024*
// This is a fast, top-quality generator. If 1024 bits of state are too
// much, try a xorshift128+ or a xorshift64* generator.
type xs1024star struct {
	state [16]uint64
	p     int
}

// New1024star returns a new pseudo-random number Source64 using the xorshift1024* algorithm.
func New1024star() rand64.Source64 {
	return &xs1024star{}
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs1024star) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	// We use xorshit64* to seed the state
	xs64 := New64star()
	xs64.Seed64(seed)
	for p := range rng.state {
		rng.state[p] = xs64.Uint64()
	}
}

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (rng *xs1024star) SeedFromSlice(seed []uint64) {
	seedArray(rng.state[:], seed)
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs1024star) Uint64() uint64 {
	s := rng.state[:]
	s0 := s[rng.p]
	rng.p = (rng.p + 1) & 15
	s1 := s[rng.p]
	s1 ^= s1 << 31 // a
	s1 ^= s1 >> 11 // b
	s0 ^= s0 >> 30 // c
	s[rng.p] = s0 ^ s1
	return s[rng.p] * 1181783497276652981
}
