// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package xorshift1024 provides an implementation for a pseudo-random number
generator (PRNG) using the xorshift1024* algorithm.

Period: 2^1024-1. State size: 1024 bits (16 uint64).

Go implementation based on a C reference implementation by Sebastiano Vigna.
For further information: http://xorshift.di.unimi.it/

The three lowest bits are LFSRs (linear-feeedback shift registers), and thus
they are slightly less random than the other bits. It is therefore recommended
to use a sign test in order to extract boolean values (i.e. check the high order
bit).
*/
package xorshift1024

import (
	"github.com/db47h/rand64/v3/splitmix64"
)

// Rng encapsulates an xorshift 1024* PRNG.
// This is a fast, top-quality generator. If 1024 bits of state are too
// much, try a xoroshiro generator instead.
//
type Rng struct {
	state [16]uint64
	p     int
}

// Seed uses the provided uint64 seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng) Seed(seed int64) {
	sm := splitmix64.Rng{}
	sm.Seed(seed)
	for i := range rng.state {
		rng.state[i] = sm.Uint64()
	}
}

// Uint64 returns an unsigned pseudo-random 64-bit integer.
//
func (rng *Rng) Uint64() uint64 {
	s := rng.state[:]
	s0 := s[rng.p]
	rng.p = (rng.p + 1) & 15
	s1 := s[rng.p]
	s1 ^= s1 << 31                               // a
	s[rng.p] = s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30) // b,c
	return s[rng.p] * 1181783497276652981
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
