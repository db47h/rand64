# OVERVIEW

Package rand64 provides support for pseudo random number generators
yielding unsinged 64 bits numbers in the range [0, 2<sup>64</sup>).

It also provides implementations for random number generators using
scrambled xorshift algorithms, with two of them passing the BigCrunch
battery test of the TestU01 test suite.

Note that the various Source64 implementations can also be integrated,
transparently into existing code (ie. code using rand.Rand) since they
also provide a rand.Source interface; at the cost of a slight
degradation in their overall "randomness" quality.

# Algorithms

Three algorithms are provided.

## xorshift128+
period 2<sup>128</sup>-1

This is the fastest generator of the series, passing BigCrush without
systematic errors, but due to the relatively short period it is
acceptable only for applications with a very mild amount of parallelism;
otherwise, use a xorshift1024\* generator.

## xorshift1024\* 
period 2<sup>1024</sup>-1

This is a fast, top-quality generator, also passing BigCrunch without
systematic errors. If 1024 bits of state are too much, try a
xorshift128+ or xorshift64\* generator.

## xorshift64\* 
period 2<sup>64</sup>-1

This is a decent generator (failing BigCrunch only on the MatrixRank
test). It is used internally to seed the state buffers for the other
algorithms. Using it as a general purpose PRNG is however not
recommended since xorshift128+ is noticably faster with better output
quality and a much longer period.

Algorithms by George Marsaglia, Sebastiano Vigna. Go implementation
based on a C reference implementation by S. Vigna. For further
information: http://xorshift.di.unimi.it/

# Benchmarks

The last result is for the default PRNG provided by rand.NewSource() for
comparison:

    BenchmarkXorShift64star     100000000       11.7 ns/op
    BenchmarkXorShift128plus    500000000        7.3 ns/op
    BenchmarkXorShift1024star   100000000       12.5 ns/op
    
    BenchmarkRandSource         100000000       11.4 ns/op

# TODO

xorshift4096\* implementation. Passes BigCrunch, same speed as xorshift1204\*,
but much longer period and bigger state buffer for applications that might
need it.

# Documentation

You can browse the package documentation online at http://godoc.org/github.com/wildservices/rand64
