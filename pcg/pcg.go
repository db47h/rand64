// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package pcg provides an implementation of the PCG XSL RR 128/64 pseudo random
number generator.

Full details can be found at the PCG-Random website
(http://www.pcg-random.org/). This version of the code provides a single
member of the PCG family, PCG XSL RR 128/64 (LCG).

Use of this algorithm is governed by a MIT-style license that can be found
in the LICENSE-pcg file.
*/
package pcg

import (
	"math/bits"

	"github.com/db47h/rand64/v3/splitmix64"
)

const (
	mulHi = 2549297995355413924
	mulLo = 4865540595714422341
	incHi = 6364136223846793005
	incLo = 1442695040888963407
)

// Rng encapsulates a PCG XSL RR 128/64 (LCG) PRNG.
//
// This is a permuted congruential generator as defined in
//
// 	PCG: A Family of Simple Fast Space-Efficient Statistically Good Algorithms for Random Number Generation
// 	Melissa E. O'Neill, Harvey Mudd College
// 	https://www.cs.hmc.edu/tr/hmc-cs-2014-0905.pdf
//
type Rng struct {
	LO uint64 // low 64 bits of 128 bits state
	HI uint64 // high 64 bits of 128 bits state
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
//
func (rng *Rng) Seed(seed int64) {
	src := splitmix64.Rng{}
	src.Seed(seed)
	rng.LO = src.Uint64()
	rng.HI = src.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng) Uint64() uint64 {
	hi, lo := bits.Mul64(rng.LO, mulLo)
	hi += rng.HI * mulLo
	hi += rng.LO * mulHi

	var c uint64
	rng.LO, c = bits.Add64(lo, incLo, 0)
	rng.HI, _ = bits.Add64(hi, incHi, c)

	// return hi^lo rotated right by high 6 bits of 128 bits state
	return bits.RotateLeft64(rng.HI^rng.LO, -int(rng.HI>>58))
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
