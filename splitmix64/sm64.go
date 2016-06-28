// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package splitmix64 implements a 64 bit SplittableRandom PRNG.

This is a fixed-increment version of Java 8's SplittableRandom generator.

Period: 2^64. State size: 64 bits.

See http://dx.doi.org/10.1145/2714064.2660195 and
http://docs.oracle.com/javase/8/docs/api/java/util/SplittableRandom.html

It is a very fast generator passing BigCrush, and it can be useful if for some
reason you absolutely want only 64 bits of state; otherwise, we rather suggest
to use a xoroshiro128+ (for moderately parallel computations) or xorshift1024*
(for massively parallel computations) generator.
*/
package splitmix64

import "github.com/db47h/rand64"

// splitmix64 encapsulates the splitmix64 PRNG.
type splitmix64 uint64

// New returns a new pseudo-random number Source using the splitmix64
// algorithm seeded with the provided seed.
func New(seed uint64) rand64.Source {
	var sm64 = splitmix64(seed)
	return &sm64
}

func (rng *splitmix64) Seed(seed uint64) {
	*rng = splitmix64(seed)
}

func (rng *splitmix64) SeedFromSlice(src []uint64) {
	if len(src) != 0 {
		*rng = splitmix64(src[0])
	} else {
		*rng = 0
	}
	return
}

func (rng *splitmix64) Uint64() uint64 {
	*rng += 0x9E3779B97F4A7C15
	z := uint64(*rng)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}
