# rand64

[![Build Status](https://travis-ci.org/db47h/rand64.svg?branch=master)](https://travis-ci.org/db47h/rand64)
[![Go Report Card](https://goreportcard.com/badge/github.com/db47h/rand64)](https://goreportcard.com/report/github.com/db47h/rand64)  [![GoDoc](https://godoc.org/github.com/db47h/rand64?status.svg)](https://godoc.org/github.com/db47h/rand64)

Package rand64 provides support for pseudo random number generators
yielding unsinged 64 bits numbers in the range \[0, 2<sup>64</sup>).

Go's built in PRNG returns 63 bits positive integers in \[0, 2<sup>63</sup>)
and the algorithm used is an additive lagged Fibonacci generator:
LFib(2<sup>63</sup>, 607, 273, +).

Although the built in rand.Rand is meant to be extensible, the limitation to 63
bits makes it unsuitable in some cases.

Implementations for the following pseudo random number generators are provided
in sub-packages:

 - splitmix64, a 64 bits SplittableRandom PRNG. Mostly used as a seeder for the
   other PRNGs.
 - scrambled xorshift (xorshift), with two variants passing the BigCrunch test
   without systematic errors.
 - xoroshiro128, the successor to xorshift128+. It is the fastest full-period
   generator passing BigCrush without systematic failures.
 - 64bits Mersene Twister (mt19937).
 - io.Reader wrapper for PRNG sources.

Note that rand64.Rand implements rand.Source, so it can be used to proxy
rand64.Source sources for rand.Rand and integrate them transparently into
existing code; at the cost of a slight degradation in the statistical quality
of their output.

PRNGs are not seeded at creation time (except for splitmix64). This is to
prevent needless duplication of initial state data during initialization.

Creation of a PRNG looks like this:

```
source := xoroshiro.New128plus()    // create source
source.Seed(uint64Seed)             // Seed it from a single Uint64
source.SeedFromSlice(initialState)  // OR from a slice of Uints
// wrap it in a rand64.Rand wrapper that will provide more utility functions
rng := rand64.New(source)
```

Or even shorter:

```
rng := rand64.New(xoroshiro.New128plus()) // one liner to create wrapper and source
rng.Seed64(uint64Seed)                    // or rng.SeedFromSlice(...)
```

# Algorithms

Scrambled xorshift algorithms by George Marsaglia, Sebastiano Vigna. Go
implementation based on a C reference implementations by David Blackman and
Sebastiano Vigna. For further information: http://xorshift.di.unimi.it/

Tests through the [ent program][ent], with a deliberately fairly low sample
count, are included. Note that these results are informational only and may
vary between runs (especially the chi square distribution).

[ent]: http://www.fourmilab.ch/random/

## xoroshiro128+
period 2<sup>128</sup>-1

This is the successor to xorshift128+. It is the fastest[1] full-period generator
passing BigCrush without systematic failures, but due to the relatively short
period it is acceptable only for applications with a mild amount of parallelism;
otherwise, use a xorshift1024* generator.

Beside passing BigCrush, this generator passes the PractRand test suite up to
(and included) 16TB, with the exception of binary rank tests, which fail due to
the lowest bit being an LFSR; all other bits pass all tests. We suggest to use
a sign test to extract a random Boolean value.

The state must be seeded so that it is not everywhere zero. If you have a
64-bit seed, we suggest to seed a splitmix64 generator and use its output to
fill s.

## xorshift128+
period 2<sup>128</sup>-1

This generator has been replaced by xoroshiro128plus, which is significantly[1]
faster and has better statistical properties.

Passes BigCrush without systematic errors, but due to the relatively short
period it is acceptable only for applications with a very mild amount of
parallelism; otherwise, use a xorshift1024\* generator.

Test with ent over 10,000 64 bits samples:

	Entropy = 7.997753 bits per byte.

	Optimum compression would reduce the size
	of this 80000 byte file by 0 percent.

	Chi square distribution for 80000 samples is 250.99, and randomly
	would exceed this value 50.00 percent of the times.

	Arithmetic mean value of data bytes is 127.6051 (127.5 = random).
	Monte Carlo value for Pi is 3.127878197 (error 0.44 percent).
	Serial correlation coefficient is -0.002355 (totally uncorrelated = 0.0).

Tests runs of this algorithm to compute π with the Monte Carlo method yielded
π = 3.1415924162 after 40,002,000,000 iterations, with a stable approximation
at the 6<sup>th</sup> decimal.

## xorshift1024\*
period 2<sup>1024</sup>-1

This is a fast, top-quality generator, also passing BigCrunch without
systematic errors. If 1024 bits of state are too much, try a
xorshift128+ or xorshift64\* generator.

Test with ent:

	Entropy = 7.997792 bits per byte.

	Optimum compression would reduce the size
	of this 80000 byte file by 0 percent.

	Chi square distribution for 80000 samples is 246.32, and randomly
	would exceed this value 50.00 percent of the times.

	Arithmetic mean value of data bytes is 127.8454 (127.5 = random).
	Monte Carlo value for Pi is 3.132678317 (error 0.28 percent).
	Serial correlation coefficient is -0.001606 (totally uncorrelated = 0.0).

## MT19937-64
period 2<sup>19937</sup>-1

This is a pure Go implementation based on the mt19937-64.c C implementation
by Makoto Matsumoto and Takuji Nishimura.

More information on the Mersenne Twister algorithm and other implementations
are available from http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html

Test with ent:

	Entropy = 7.997964 bits per byte.

	Optimum compression would reduce the size
	of this 80000 byte file by 0 percent.

	Chi square distribution for 80000 samples is 225.61, and randomly
	would exceed this value 90.00 percent of the times.

	Arithmetic mean value of data bytes is 127.6890 (127.5 = random).
	Monte Carlo value for Pi is 3.144978624 (error 0.11 percent).
	Serial correlation coefficient is 0.000496 (totally uncorrelated = 0.0).

## io.Reader wrapper
Not an actual PRNG.

The IoRand package contains a wrapper for reading data streams via io.Reader.
There is a helper function in the randutil package that uses it to wrap
crypto/rand in a Source and use its output to seed the faster PRNGs.

# Benchmarks

These benchmarks where done with go-tip (1.7 beta 2 as of writing) where
xorshift1024*, splitmix64, MT19937 and the standard Go PRNG got a significant
performance boost over Go 1.6.2.

The last result is for the default PRNG provided by the standrd library's
rand.NewSource() for comparison:

    Splitmix64           	300000000	         4.35 ns/op
    XorShift128roplus    	300000000	         5.14 ns/op
    XorShift128plus      	300000000	         4.92 ns/op
    XorShift1024star     	200000000	         6.46 ns/op
    Mt19937              	200000000	         9.12 ns/op
    GoRand              	200000000	         6.40 ns/op

[1]: According to the authors of the algorithms, the xoroshiro128+ algorithm should be significantly faster than xorshift128+. As seen in the benchmarks, it is in fact just a little slower eventhough Go 1.6.2 and 1.7 properly *do* translate the simulated rotate into a single rotate instruction on x86_64.

# Documentation

You can browse the package documentation online at http://godoc.org/github.com/db47h/rand64
