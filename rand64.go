// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package rand64 provides support for pseudo-random number generators (PRNG)
yielding unsinged 64 bits numbers in the range [0, 1<<64).

Implementations for the following pseudo random number generators are provided
in sub-packages: 64 bit SplittableRandom (splitmix64), XOR/rotate/shift/rotate
(xoroshiro), scrambled xorshift (xorshift), 64bits Mersene Twister (mt19937).

Note that rand64.Rand implements rand.Source, so it can be used to proxy
rand64.Source sources for rand.Rand and integrate them transparently into
existing code; at the cost of a slight degradation in the statistical quality
of their output.

Apart from splitmix64 which has a 64 bit state, PRNGs are not seeded at creation
time. This is to prevent duplication of constructors for each seeding method
(from single value or from slice).
*/
package rand64

// A Source represents a source of uniformly-distributed pseudo-random uint64 values in the range [0, 1<<64).
//
// Seed uses the provided uint64 seed value to initialize the generator to a deterministic state.
//
// SeedFromSlice seeds the generator's state buffer with values from the given source slice.
//
// Uint64 returns a pseudo-random 64-bit integer in the range [0, 1<<64).
type Source interface {
	Seed(uint64)
	SeedFromSlice([]uint64)
	Uint64() uint64
}

// A Rand is a source of unsigned 64 bit pseudo random numbers in the range [0,1<<64).
// Most of the methods are the unsigned counterparts of math.Rand.
type Rand struct {
	src Source
}

// New returns a new Rand that uses random values from src
// to generate other random values.
func New(src Source) *Rand { return &Rand{src} }

// Seed uses the provided unsinged seed value to initialize the generator to a deterministic state.
func (r *Rand) Seed(seed int64) { r.src.Seed(uint64(seed)) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 { return int64(r.src.Uint64() >> 1) }

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state.
func (r *Rand) Seed64(seed uint64) { r.src.Seed(seed) }

// SeedFromSlice seeds the generator's state buffer with values from the array argument.
func (r *Rand) SeedFromSlice(seed []uint64) {
	r.src.SeedFromSlice(seed)
}

// Uint64 returns a pseudo-random 64-bit integer in the range [0, 1<<64).
func (r *Rand) Uint64() uint64 { return r.src.Uint64() }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 { return uint32(r.Uint64() >> 32) }

// Uint64n returns, as an uint64, a pseudo-random number in [0,n).
//
// caveat: maximum range is [0, 1<<64-2].
func (r *Rand) Uint64n(n uint64) uint64 {
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
func (r *Rand) Uint32n(n uint32) uint32 {
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
func (r *Rand) Uintn(n uint) uint {
	if n <= 1<<32-1 {
		return uint(r.Uint32n(uint32(n)))
	}
	return uint(r.Uint64n(uint64(n)))
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 {
	// See Go's math/rand source code.
	// 1<<53 is the highest power of two for which float64(1<<n-1)/(1<<n) is != 1
	return float64(r.Uint64n(1<<53)) / (1 << 53)
}

// Float64Closed returns, as a float64, a pseudo-random number in [0.0,1.0]
func (r *Rand) Float64Closed() float64 {
	return float64(r.Uint64n(1<<53)) / (1<<53 - 1)
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 {
	// Same rationale as in Float64
	return float32(r.Uint64n(1<<24)) / (1 << 24)
}

// UPerm returns, as a slice of n uints, a pseudo-random permutation of the unsigned integers [0,n).
func (r *Rand) UPerm(n uint) []uint {
	m := make([]uint, n)
	for i := uint(0); i < n; i++ {
		j := r.Uintn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
func (r *Rand) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i += 8 {
		val := r.src.Uint64()
		for j := 0; i+j < len(p) && j < 8; j++ {
			p[i+j] = byte(val)
			val >>= 8
		}
	}
	return len(p), nil
}
