// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rand64 provides support for pseudo-random number generators (PRNG)
yielding unsinged 64 bits numbers in the range [0, 1<<64).

Implementations for the following pseudo random number generators are provided
in sub-packages: scrambled xorshift (xorshift), 64bits Mersene Twister
(mt19937).

Note that rand64.Rand64 implements rand.Source, so it can be used to proxy
rand64.Source64 sources for rand.Rand and integrate them transparently into
existing code; at the cost of a slight degradation in the statistical quality
of their output.

PRNGs are not seeded at creation time. This is to prevent duplication of
constructors for each seeding method (from single value or from slice).
*/
package rand64

// A Source64 represents a source of uniformly-distributed pseudo-random uint64 values in the range [0, 1<<64).
//
// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
//
// SeedFromSlice seeds the generator's state buffer with values from the given source slice.
//
// Uint64 returns a pseudo-random 64-bit integer in the range [0, 1<<64).
type Source64 interface {
	Seed64(uint64)
	SeedFromSlice([]uint64)
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
func (r *Rand64) Seed(seed int64) { r.src.Seed64(uint64(seed)) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand64) Int63() int64 { return int64(r.src.Uint64() >> 1) }

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (r *Rand64) Seed64(seed uint64) { r.src.Seed64(seed) }

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (r *Rand64) SeedFromSlice(seed []uint64) {
	r.src.SeedFromSlice(seed)
}

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

// Float64Full returns, as a float64, a pseudo-random number in [0.0,1.0]
func (r *Rand64) Float64Full() float64 {
	return float64(r.Uint64n(1<<53)) / (1<<53 - 1)
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

// BulkUint64 returns a slice of n uint64 filled with pseudo-random numbers.
// The resulting slice can be used for example as a seed for "lesser" PRNGs via Source64.SeedFromSlice()
func (r *Rand64) BulkUint64(n uint) []uint64 {
	rng := r.src
	a := make([]uint64, n)
	for i := range a {
		a[i] = rng.Uint64()
	}
	return a
}
