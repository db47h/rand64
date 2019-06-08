// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package mt19937 implements a 64-bit version of the Mersenne Twister
pseudo-random number generator (MT19937 PRNG).

The state size is 312 uint64.

This is a pure Go implementation based on the mt19937-64.c C implementation
by Makoto Matsumoto and Takuji Nishimura.

More information on the Mersenne Twister algorithm and other implementations
are available from http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html

See included LICENSE_MT for original C code and license.
*/
package mt19937

const (
	_NN      = 312
	_MM      = 156
	_MatrixA = 0xB5026F5AA96619E9
	_UM      = 0xFFFFFFFF80000000 // Most significant 33 bits
	_LM      = 0x7FFFFFFF         // Least significant 31 bits
)

var mag01 = [...]uint64{
	0,
	_MatrixA,
}

// Rng wraps the state data for the MT19937 pseudo-random number generator.
//
type Rng struct {
	state [_NN]uint64 // State vector
	// index is used in descending order (opposite to the original C version).
	// Thus, an uninitialized Rng has index = 0 instead of NN+1. This allows the use
	// of a Rng{} struct literal as a valid PRNG and exhibit the same behavior as the
	// C version if generating values without seeding.
	index uint64
}

// Seed uses the provided uint64 seed value to initialize the generator to a deterministic state.
//
// If Seed is 0, the generator will be seeded with the same default value as in the original C code (5489).
//
func (rng *Rng) Seed(seed int64) {
	var i uint64
	mt := rng.state[:]

	if seed == 0 {
		seed = 5489
	}

	mt[0] = uint64(seed)
	for i = 1; i < _NN; i++ {
		mt[i] = 6364136223846793005*(mt[i-1]^(mt[i-1]>>62)) + i
	}
	rng.index = 1
}

// SeedFromSlice initializes the state array with data from slice key. This function behaves
// exactly like init_by_array() in the original C code.
//
func (rng *Rng) SeedFromSlice(key []uint64) {
	var (
		i uint64 = 1
		j uint64
		k = uint64(len(key))
	)
	mt := rng.state[:]

	rng.Seed(19650218)

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

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
func (rng *Rng) Uint64() uint64 {
	var i int
	var x uint64
	mt := rng.state[:]
	mti := rng.index

	if mti <= 1 { // generate _NN words at once
		// seed if needed
		if mti == 0 {
			rng.Seed(5489)
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

		mti = _NN + 1
	}

	x = mt[_NN+1-mti]
	rng.index = mti - 1

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (rng *Rng) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}
