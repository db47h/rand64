// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The xorshift package provides implementations for pseudo random number
generators using scrambled xorshift algorithms, with two of them passing the
BigCrunch battery test of the TestU01 test suite without systematic errors.

Algorithms by George Marsaglia, Sebastiano Vigna. Go implementation based on a
C reference implementation by S. Vigna.
For further information: http://xorshift.di.unimi.it/

xorshift128+ (period 2^128-1)

This is the fastest generator of the series, passing BigCrush without systematic
errors, but due to the relatively short period it is acceptable only
for applications with a very mild amount of parallelism; otherwise, use
a xorshift1024* generator.

xorshift1024* (period 2^1024-1)

This is a fast, top-quality generator, also passing BigCrunch without systematic
errors. If 1024 bits of state are too much, try a xorshift128+ or xorshift64*
generator.

xorshift64* (period 2^64-1)

This is a decent generator (failing BigCrunch only on the MatrixRank test). It
is used internally to seed the state buffers for the other algorithms. Using
it as a general purpose PRNG is however not recommended since xorshift128+ is
noticably faster with better output quality and a much longer period.

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

import "github.com/wildservices/rand64"

// xorshift64*
// This is a good generator if you're short on memory, but otherwise we
// rather suggest to use a xorshift128+ (for maximum speed) or
// xorshift1024* (for speed and very long period) generator.
type xs64star uint64

// New64star returns a new pseudo-random Source64 using the xorshift64* algorithm
// seeded with the given value.
func New64star(seed uint64) rand64.Source64 {
	var xs64 xs64star
	xs64.Seed64(seed)
	return &xs64
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (s *xs64star) Seed64(seed uint64) {
	if seed == 0 {
		seed = 89482311
	}
	*s = xs64star(seed)
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (s *xs64star) Seed(seed int64) {
	s.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (s *xs64star) Uint64() uint64 {
	*s ^= *s >> 12 // a
	*s ^= *s << 25 // b
	*s ^= *s >> 27 // c
	return uint64(*s * 2685821657736338717)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (s *xs64star) Int63() int64 {
	*s ^= *s >> 12 // a
	*s ^= *s << 25 // b
	*s ^= *s >> 27 // c
	return int64((*s * 2685821657736338717) >> 1)
}

// xorshift128+
// this is the fastest generator passing BigCrush without systematic
// errors, but due to the relatively short period it is acceptable only
// for applications with a very mild amount of parallelism; otherwise, use
// a xorshift1024* generator
type xs128plus struct {
	s [2]uint64
}

// New128plus returns a new pseudo-random Source64 using the xorshift128+ algorithm
// seeded with the given value.
func New128plus(seed uint64) rand64.Source64 {
	var rng xs128plus
	rng.Seed64(seed)
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs128plus) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	if seed == 0 {
		seed = 89482311
	}
	xs64 := New64star(seed) // TODO: a better way to setup s[]?
	rng.s[0] = uint64(seed)
	rng.s[1] = xs64.Uint64()
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (rng *xs128plus) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs128plus) Uint64() uint64 {
	s1 := rng.s[0]
	s0 := rng.s[1]
	rng.s[0] = s0
	s1 ^= s1 << 23                                 // a
	rng.s[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26)) // b, c
	return rng.s[1] + s0
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *xs128plus) Int63() int64 {
	s1 := rng.s[0]
	s0 := rng.s[1]
	rng.s[0] = s0
	s1 ^= s1 << 23                                 // a
	rng.s[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26)) // b, c
	return int64((rng.s[1] + s0) >> 1)
}

// xorshift1024*
// This is a fast, top-quality generator. If 1024 bits of state are too
// much, try a xorshift128+ or a xorshift64* generator.
type xs1024star struct {
	s [16]uint64
	p int
}

// New1024star returns a new pseudo-random Source64 using the xorshift1024* algorithm
// seeded with the given value.
func New1024star(seed uint64) rand64.Source64 {
	var rng xs1024star
	rng.Seed64(seed)
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs1024star) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	// We use xorshit64* to seed the state
	xs64 := New64star(seed)
	for p := range rng.s {
		rng.s[p] = xs64.Uint64()
	}
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (rng *xs1024star) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs1024star) Uint64() uint64 {
	s := rng.s[:]
	s0 := s[rng.p]
	rng.p = (rng.p + 1) & 15
	s1 := s[rng.p]
	s1 ^= s1 << 31 // a
	s1 ^= s1 >> 11 // b
	s0 ^= s0 >> 30 // c
	s[rng.p] = s0 ^ s1
	return s[rng.p] * 1181783497276652981
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *xs1024star) Int63() int64 {
	s := rng.s[:]
	s0 := s[rng.p]
	rng.p = (rng.p + 1) & 15
	s1 := s[rng.p]
	s1 ^= s1 << 31 // a
	s1 ^= s1 >> 11 // b
	s0 ^= s0 >> 30 // c
	s[rng.p] = s0 ^ s1
	return int64((s[rng.p] * 1181783497276652981) >> 1)
}
