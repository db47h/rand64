# rand64

[![Build Status][travisImg]][travis]
[![Go Report Card][goreportImg]][goreport]
[![GoDoc][godocImg]][godoc]

Package rand64 provides pseudo random number generators yielding unsinged 64
bits numbers in the range \[0, 2<sup>64</sup>).

Implementations for the following pseudo random number generators are provided
in their own packages:

- splitmix64, a 64 bits SplittableRandom PRNG. Mostly used as a seeder for the other PRNGs.
- xorshift128+
- xorshift1024*
- xoroshiro128+
- io.Reader wrapper for PRNG sources.

These generetors implement rand.Source64, so they can be used as source for
rand.Rand (as of Go 1.8).

Note that some algorithms make use of the bits package from Go 1.9.

Creation of a PRNG looks like this:

```
source := xoroshiro.Rng{}           // create source
// wrap it in a rand.Rand wrapper that will provide more utility functions
rng := rand.New(&source)
rng.Seed(int64Seed)              // Seed it from a single int64 (negative values are accepted)
```

## Algorithms

### splitmix64

This is a fixed-increment version of Java 8's [SplittableRandom](http://docs.oracle.com/javase/8/docs/api/java/util/SplittableRandom.html) generator. See also the page on [Fast splittable pseudorandom number generators](http://dx.doi.org/10.1145/2714064.2660195).

Go implementation based on a C reference implementation by Sebastiano Vigna.

### xorshift128+

Period 2<sup>128</sup>-1

A 64-bit version of Saito and Matsumoto's XSadd generator. Instead of using a
multiplication, it returns the sum (in Z/264Z) of two consecutive output of a
xorshift generator. This generator is presently used in the JavaScript engines
of Chrome, Firefox, Safari, Microsoft Edge. If you find its period too short
for large-scale parallel simulations, use xorshift1024*.

Go implementation based on a C reference implementation by Sebastiano Vigna.

For more information, visit the [xoroshiro+ / xorshift* / xorshift+ generators and the PRNG shootout][PRNGSHoutout] page.

### xorshift1024*

Period 2<sup>1024</sup>-1

A fast, high-quality PRNG. Numbers are obtained by scrambling the output of a
Marsaglia xorshift generator with a 64-bit invertible multiplier.

Go implementation based on a C reference implementation by Sebastiano Vigna.

For more information, visit the [xoroshiro+ / xorshift* / xorshift+ generators and the PRNG shootout][PRNGSHoutout] page.

### xoroshiro128+

Period 2<sup>128</sup>-1

According to the algorithm authors:

> xoroshiro128+ (XOR/rotate/shift/rotate) is the successor to xorshift128+.
> Instead of perpetuating Marsaglia's tradition of xorshift as a basic
> operation, xoroshiro128+ uses a carefully handcrafted shift/rotate-based
> linear transformation designed in collaboration with David Blackman. The
> result is a significant improvement in speed (well below a nanosecond per
> integer) and a significant improvement in statistical quality, as detected by
> the long-range tests of PractRand. xoroshiro128+ is our current suggestion for
> replacing low-quality generators commonly found in programming languages. It
> is the default generator in Erlang.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna.

For more information, visit the [xoroshiro+ / xorshift* / xorshift+ generators and the PRNG shootout][PRNGSHoutout] page.

### MT19937-64
period 2<sup>19937</sup>-1

This is a pure Go implementation based on the mt19937-64.c C implementation
by Makoto Matsumoto and Takuji Nishimura.

More information on the Mersenne Twister algorithm and other implementations
are available from http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html

### io.Reader wrapper

Not an actual PRNG.

The IoRand package is a simple wrapper for reading byte streams via io.Reader.
It can be used as a wrapper around crypto/rand to build a rand.Source64 for
the math/rand package:

```go
package main

import (
    "bufio"
    crand "crypto/rand"
    "encoding/binary"
    "math/rand"

    "github.com/db47h/rand64/iorand"
)

// Wrap crypto/rand in an IoRand
func main() {
    // first, wrap crypto/rand.Reader in a buffered bufio.Reader
    bufferedReader := bufio.NewReader(crand.Reader)
    // Create the new IoRand Source
    ior := iorand.New(bufferedReader, binary.LittleEndian)
    // use it as rand.Source64
    rng := rand.New(ior)
    // get random numbers...
    for i := 0; i < 4; i++ {
        _ = rng.Uint64()
    }
}
```

## Benchmarks

These benchmarks where done with go 1.9.1.

According to the [PRNG shootout][PRNGSHoutout] page, xoroshiro should be faster
than the xorshift algorithms. This is not the case on any of the tested
hardware (even slower than xorshift1024* on AMD), so there's probably some room
for optimization in the Go implementation.

The last result is for the default PRNG provided by the standard library's
rand.NewSource() for comparison:

| Algorithm     | AMD FX-6300 | Core i5 6200U | Celeron-M 410 | ARM Cortex-A7    |
|---------------|------------:|--------------:|--------------:|-----------------:|
| xoroshiro128+ |  3.96 ns/op |    2.42 ns/op | 13.6 ns/op    |       33.4 ns/op |
| xorshift128+  |  3.53 ns/op |    2.26 ns/op | 11.6 ns/op    |       29.5 ns/op |
| xorshift1024* |  3.75 ns/op |    3.01 ns/op | 18.9 ns/op    |       50.7 ns/op |
| splitmix64    |  2.20 ns/op |    1.90 ns/op |  5.4 ns/op    |       15.3 ns/op |
| MT19937       |  8.82 ns/op |    6.18 ns/op | 53.3 ns/op    |      137.0 ns/op |
| GoRand        |  7.54 ns/op |    4.61 ns/op | 22.8 ns/op    |       80.4 ns/op |

## Documentation

For the xorshift and xoroshiro generators, the lowest bits of the generated
values are LSFRs (linear-feeedback shift registers), and thus they are slightly
less random than the other bits. It is therefore recommended to use a sign test
in order to extract boolean values (i.e. check the high order bit).

[PRNGShoutout]: http://xoroshiro.di.unimi.it/
[travisImg]: https://travis-ci.org/db47h/rand64.svg?branch=master
[travis]: https://travis-ci.org/db47h/rand64
[goreportImg]: https://goreportcard.com/badge/github.com/db47h/rand64
[goreport]: https://goreportcard.com/report/github.com/db47h/rand64
[godocImg]: https://godoc.org/github.com/db47h/rand64?status.svg
[godoc]: http://godoc.org/github.com/db47h/rand64
