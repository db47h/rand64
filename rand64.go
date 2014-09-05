// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rand64 provides support for pseudo random number generators yielding
unsinged 64 bits numbers in the range [0, 1<<64).

Implementations for various pseudo random number generators are provided in
sub-packages.

Note that the various Source64 implementations can be integrated
transparently into existing code (ie. code using rand.Rand) since they also
provide a rand.Source interface; at the cost of a slight degradation in their
overall "randomness" quality.

*/
package rand64

import (
	"math/rand"
)

// A Source64 represents a source of uniformly-distributed pseudo-random uint64 values in the range [0, 1<<64).
//
// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
//
// Uint64 returns a pseudo-random 64-bit integer in the range [0, 1<<64).
type Source64 interface {
	rand.Source
	Seed64(uint64)
	Uint64() uint64
}

// A Rand64 is a source of unsigned 64 bit pseudo random numbers in the range [0,1<<64).
type Rand64 struct {
	src Source64
}

// New returns a new Rand64 that uses random values from src
// to generate other random values.
func New(src Source64) *Rand64 { return &Rand64{src} }

// Seed uses the provided unsinged seed value to initialize the generator to a deterministic state.
func (r *Rand64) Seed(seed int64) { r.src.Seed(seed) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand64) Int63() int64 { return r.src.Int63() }

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (r *Rand64) Seed64(seed uint64) { r.src.Seed64(seed) }

// Uint64 returns a pseudo-random 64-bit integer in the range [0, 1<<64).
func (r *Rand64) Uint64() uint64 { return r.src.Uint64() }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand64) Uint32() uint32 { return uint32(r.Uint64() >> 32) }

// Uint64n returns, as an uint64, a pseudo-random number in [0,n).
//
// caveat: maximum range is [0, 1<<64-2].
func (r *Rand64) Uint64n(n uint64) uint64 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Uint64() & (n - 1)
	}
	max := (1 << 64) - 1 - ((1<<64-1)%n+1)%n // == (1<<64)-1 - (1<<64)%n
	v := r.Uint64()
	for v > max {
		v = r.Uint64()
	}
	return v % n
}

// Uint32n returns, as a uint32, a pseudo-random number in [0,n).
//
// caveat: maximum range is [0, 1<<32-2].
func (r *Rand64) Uint32n(n uint32) uint32 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Uint32() & (n - 1)
	}
	max := uint32((1 << 32) - 1 - (1<<32)%uint64(n))
	v := r.Uint32()
	for v > max {
		v = r.Uint32()
	}
	return v % n
}

// Uintn returns, as a uint, a pseudo-random number in [0,n).
func (r *Rand64) Uintn(n uint) uint {
	if n <= 1<<32-1 {
		return uint(r.Uint32n(uint32(n)))
	}
	return uint(r.Uint64n(uint64(n)))
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand64) Float64() float64 {
	// See Go's math/rand source code.
	// 1<<53 is the highest power of two for which float64(1<<n-1)/(1<<n) is != 1
	return float64(r.Uint64n(1<<53)) / (1 << 53)
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand64) Float32() float32 {
	// Same rationale as in Float64
	return float32(r.Uint64n(1<<24)) / (1 << 24)
}

// UPerm returns, as a slice of n uints, a pseudo-random permutation of the unsigned integers [0,n).
func (r *Rand64) UPerm(n uint) []uint {
	m := make([]uint, n)
	for i := uint(0); i < n; i++ {
		j := r.Uintn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}
