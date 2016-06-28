# rand64

[![Build Status](https://travis-ci.org/db47h/rand64.svg?branch=master)](https://travis-ci.org/db47h/rand64)
[![Go Report Card](https://goreportcard.com/badge/github.com/db47h/rand64)](https://goreportcard.com/report/github.com/db47h/rand64)  [![GoDoc](https://godoc.org/github.com/db47h/rand64?status.svg)](https://godoc.org/github.com/db47h/rand64)

Package rand64 provides support for pseudo random number generators
yielding unsinged 64 bits numbers in the range \[0, 2<sup>64</sup>).

Go's built in PRNG returns 63 bits positive integers in \[0, 2<sup>63</sup>)
and the algorithm used is an additive lagged Fibonacci generator:
LFib(2<sup>63</sup>, 607, 273, +).

Although the built in rand.Rand is meant to be extensible, the limitation to 63
bits makes it unsuitable in some cases, like when interoperating with other
systems or software where full blown 64 bit integers are expected.

Implementations for the following pseudo random number generators are provided
in sub-packages:

 - scrambled xorshift (xorshift), with two variants passing the BigCrunch test
   without systematic errors.
 - 64bits Mersene Twister (mt19937).
 - io.Reader wrapper.

Note that rand64.Rand64 implements rand.Source, so it can be used to proxy
rand64.Source64 sources for rand.Rand and integrate them transparently into
existing code; at the cost of a slight degradation in the statistical quality
of their output.

PRNGs are not seeded at creation time. This is to prevent duplication of
constructors for each seeding method (from single value of from slice).

# Algorithms

Scrambled xorshift algorithms by George Marsaglia, Sebastiano Vigna. Go
implementation based on a C reference implementation by S. Vigna. For further
information: http://xorshift.di.unimi.it/

Tests through the [ent program][ent], with a deliberately fairly low sample
count, are included. Note that these results are informational only and may
vary between runs (especially the chi square distribution).

[ent]: http://www.fourmilab.ch/random/

## xorshift128+
period 2<sup>128</sup>-1

This is the fastest generator of the series, passing BigCrush without
systematic errors, but due to the relatively short period it is
acceptable only for applications with a very mild amount of parallelism;
otherwise, use a xorshift1024\* generator.

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

## xorshift64\*
period 2<sup>64</sup>-1

This is a decent generator (failing BigCrunch only on the MatrixRank
test). It is used internally to seed the state buffers for the other
algorithms from a single value. Using it as a general purpose PRNG is however
not recommended since xorshift128+ is noticeably faster with better statistical
quality and a much longer period.

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
crypto/rand in a Source64 and use its output to seed the faster PRNGs.

# Benchmarks

The last result is for the default PRNG provided by the standrd library's
rand.NewSource() for comparison:

    BenchmarkXorShift64star-6    200000000     8.92 ns/op
    BenchmarkXorShift128plus-6   300000000     5.28 ns/op
    BenchmarkXorShift1024star-6  200000000     9.38 ns/op
    
    BenchmarkMt19937-6           100000000    10.50 ns/op
    
    BenchmarkRandSource-6        200000000     7.89 ns/op

# TODO

 - xorshift4096\* implementation. Passes BigCrunch, same speed as xorshift1204\*,
but much longer period and bigger state buffer for applications that might
need it.
 - Look into maximally equidistributed Mersenne Twister MEMT19937-II - http://www3.ocn.ne.jp/~harase/

# Documentation

You can browse the package documentation online at http://godoc.org/github.com/db47h/rand64
