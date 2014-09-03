// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Algorithms by George Marsaglia, Sebastiano Vigna. Go implementation based on a
C reference implementation by S. Vigna.
For further information: http://xorshift.di.unimi.it/

TODO: xorshift4096*
*/

package rand64

// xorshift64*
// This is a good generator if you're short on memory, but otherwise we
// rather suggest to use a xorshift128+ (for maximum speed) or
// xorshift1024* (for speed and very long period) generator.
type xs64starSource uint64

// NewXorShift64star returns a new pseudo-random Source64 using the xorshift64* algorithm
// seeded with the given value.
func NewXorShift64star(seed uint64) Source64 {
	var xs64 xs64starSource
	xs64.Seed64(seed)
	return &xs64
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (s *xs64starSource) Seed64(seed uint64) {
	if seed == 0 {
		seed = 89482311
	}
	*s = xs64starSource(seed)
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (s *xs64starSource) Seed(seed int64) {
	s.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (s *xs64starSource) Uint64() uint64 {
	*s ^= *s >> 12 // a
	*s ^= *s << 25 // b
	*s ^= *s >> 27 // c
	return uint64(*s * 2685821657736338717)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (s *xs64starSource) Int63() int64 {
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
type xs128plusSource struct {
	s [2]uint64
}

// NewXorShift128plus returns a new pseudo-random Source64 using the xorshift128+ algorithm
// seeded with the given value.
func NewXorShift128plus(seed uint64) Source64 {
	var rng xs128plusSource
	rng.Seed64(seed)
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs128plusSource) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	if seed == 0 {
		seed = 89482311
	}
	xs64 := NewXorShift64star(seed) // TODO: a better way to setup s[]?
	rng.s[0] = uint64(seed)
	rng.s[1] = xs64.Uint64()
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (rng *xs128plusSource) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs128plusSource) Uint64() uint64 {
	s1 := rng.s[0]
	s0 := rng.s[1]
	rng.s[0] = s0
	s1 ^= s1 << 23                                 // a
	rng.s[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26)) // b, c
	return rng.s[1] + s0
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *xs128plusSource) Int63() int64 {
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
type xs1024starSource struct {
	s [16]uint64
	p int
}

// NewXorShift1024star returns a new pseudo-random rand.Source using the xorshift1024* algorithm
// seeded with the given value.
func NewXorShift1024star(seed uint64) Source64 {
	var rng xs1024starSource
	rng.Seed64(seed)
	return &rng
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (rng *xs1024starSource) Seed64(seed uint64) {
	// The state must be seeded so that it is not everywhere zero.
	// We use xorshit64* to seed the state
	xs64 := NewXorShift64star(seed)
	for p := range rng.s {
		rng.s[p] = xs64.Uint64()
	}
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (rng *xs1024starSource) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
func (rng *xs1024starSource) Uint64() uint64 {
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
func (rng *xs1024starSource) Int63() int64 {
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
