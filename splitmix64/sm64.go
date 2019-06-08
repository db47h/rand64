// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package splitmix64 implements a 64 bit SplittableRandom PRNG.

This is a fixed-increment version of Java 8's SplittableRandom generator.

Period: 2^64. State size: 64 bits.

See http://dx.doi.org/10.1145/2714064.2660195 and
http://docs.oracle.com/javase/8/docs/api/java/util/SplittableRandom.html

It is a very fast generator passing BigCrush. It is used in the xoroshiro128+
and xorshift1024* generators to initialize their state arrays.
*/
package splitmix64

// Rng encapsulates a splitmix64 PRNG. The State value is exported so that
// the generator can be initialized and seeded in a single line of code:
//
//     rng := splitmix64.Rng{seed}
//     // is equivalent to
//     rng := splitmix64.Rng{}
//     rng.Seed(int64(seed))
//
type Rng struct {
	State uint64 // Internal state value
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng) Seed(seed int64) {
	rng.State = uint64(seed)
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng) Uint64() uint64 {
	rng.State += 0x9E3779B97F4A7C15
	z := rng.State
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
