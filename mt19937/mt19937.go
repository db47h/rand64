// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The mt19937 package implements a 64-bit version of the Mersenne Twister
pseudo-random number generator.

This is a pure Go implementation based on the mt19937-64.c C implementation
by Makoto Matsumoto and Takuji Nishimura.

More information on the Mersenne Twister algorithm and other implementations
are available from
http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html

See included LICENSE_MT for original C code and license (BSD).
*/
package mt19937

import "github.com/wildservices/rand64"

const (
	_NN       = 312
	_MM       = 156
	_MATRIX_A = 0xB5026F5AA96619E9
	_UM       = 0xFFFFFFFF80000000 // Most significant 33 bits
	_LM       = 0x7FFFFFFF         // Least significant 31 bits
)

var mag01 [2]uint64 = [...]uint64{
	0,
	_MATRIX_A,
}

type mt19937 struct {
	mt  [_NN]uint64 // State vector
	mti uint64      // mti==NN+1 means mt[NN] is not initialized
}

// New returns a new pseudo-random Source64 using the MT19937 algorithm seeded with the given value.
//
// If Seed is 0, a default seed will be used (5489).
func New(seed uint64) rand64.Source64 {
	var mt mt19937
	mt.Seed64(seed)
	return &mt
}

// Seed64 uses the provided uint64 seed value to initialize the generator to a deterministic state
func (rng *mt19937) Seed64(seed uint64) {
	var mti uint64
	mt := rng.mt[:]

	if seed == 0 {
		seed = 5489 // same default seed as original C code
	}

	mt[0] = seed
	for mti = 1; mti < _NN; mti++ {
		mt[mti] = 6364136223846793005*(mt[mti-1]^(mt[mti-1]>>62)) + mti
	}
	rng.mti = mti
}

// Seed uses the provided int64 seed value to initialize the generator to a deterministic state.
// Seeds < 0 are accepted.
func (rng *mt19937) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// SeedBySlice initializes the state array with data from slice key
func (rng *mt19937) SeedBySlice(key []uint64) {
	var i uint64 = 1
	var j uint64
	var k uint64 = uint64(len(key))
	mt := rng.mt[:]

	rng.Seed64(19650218)

	if _NN > k {
		k = _NN
	}
	for ; k != 0; k-- {
		mt[i] = (mt[i] ^ ((mt[i-1] ^ (mt[i-1] >> 62)) * 3935559000370003845)) + key[j] + j // non linear
		i++
		j++
		if i >= _NN {
			mt[0] = mt[_NN-1]
			i = 1
		}
		if j >= uint64(len(key)) {
			j = 0
		}
	}
	for k = _NN - 1; k != 0; k-- {
		mt[i] = (mt[i] ^ ((mt[i-1] ^ (mt[i-1] >> 62)) * 2862933555777941757)) - i // non linear
		i++
		if i >= _NN {
			mt[0] = mt[_NN-1]
			i = 1
		}
	}
	mt[0] = 1 << 63
}

// Uint64 generates a random number on [0, 2^64-1]-interval
func (rng *mt19937) Uint64() uint64 {
	var i int
	var x uint64
	mt := rng.mt[:]
	mti := rng.mti

	if mti >= _NN { // generate _NN words at once
		// if Seed64 has not been called, panic. This can't happen
		if mti == _NN+1 {
			panic("MT19937 unseeded")
		}

		for i = 0; i < _NN-_MM; i++ {
			x = (mt[i] & _UM) | (mt[i+1] & _LM)
			mt[i] = mt[i+_MM] ^ (x >> 1) ^ mag01[x&1]
		}
		for ; i < _NN-1; i++ {
			x = (mt[i] & _UM) | (mt[i+1] & _LM)
			mt[i] = mt[i+(_MM-_NN)] ^ (x >> 1) ^ mag01[x&1]
		}
		x = (mt[_NN-1] & _UM) | (mt[0] & _LM)
		mt[_NN-1] = mt[_MM-1] ^ (x >> 1) ^ mag01[x&1]

		mti = 0
	}

	x = mt[mti]

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	rng.mti = mti + 1

	return x
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *mt19937) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
