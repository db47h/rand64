# rand64

[![Build Status][travisImg]][travis]
[![Go Report Card][goreportImg]][goreport]
[![GoDoc][godocImg]][godoc]

Package rand64 provides pseudo random number generators yielding unsigned 64
bits numbers in the range \[0, 2<sup>64</sup>).

Implementations for the following pseudo random number generators are provided
in their own packages:

- splitmix64, a 64 bits SplittableRandom PRNG. Mostly used as a seeder for the other PRNGs.
- xoshiro256**
- xoshiro256+
- xoroshiro128**
- xoroshiro128+
- io.Reader wrapper for PRNG sources.

These generetors implement rand.Source64, so they can be used as source for
rand.Rand (as of Go 1.8).

Note that some algorithms make use of the bits package from Go 1.9.

Creation of a PRNG looks like this:

```
// create a source with the xoshiro256**
source := &xoshiro.Rng256SS{}

// use it as a source in rand.New
rng := rand.New(&source)

// Seed it from a single int64 (negative values are accepted)
rng.Seed(int64Seed)
```

## Algorithms

### splitmix64

This is a fixed-increment version of Java 8's [SplittableRandom](http://docs.oracle.com/javase/8/docs/api/java/util/SplittableRandom.html) generator. See also the page on [Fast splittable pseudorandom number generators](http://dx.doi.org/10.1145/2714064.2660195).

Go implementation based on a C reference implementation by Sebastiano Vigna.

### xoshiro256** and xoshiro256+

Period 2<sup>256</sup>-1

According to the algorithm authors:

> xoshiro256** (XOR/shift/rotate) is our all-purpose, rock-solid generator (not
> a cryptographically secure generator, though). It has excellent (sub-ns)
> speed, a state space (256 bits) that is large enough for any parallel
> application, and it passes all tests we are aware of.
>
> If, however, one has to generate only 64-bit floating-point numbers (by
> extracting the upper 53 bits) xoshiro256+ is a slightly (≈15%) faster
> generator with analogous statistical properties. For general usage, one has to
> consider that its lowest bits have low linear complexity and will fail
> linearity tests; however, low linear complexity can have hardly any impact in
> practice, and certainly has no impact at all if you generate floating-point
> numbers using the upper bits (we computed a precise estimate of the linear
> complexity of the lowest bits).

### xoroshiro128** and xoroshiro128+

Period 2<sup>128</sup>-1

According to the algorithm authors:

> xoroshiro128** (XOR/rotate/shift/rotate) and xoroshiro128+ have the same speed
> than xoshiro256 and use half of the space; the same comments apply. They are
> suitable only for low-scale parallel applications; moreover, xoroshiro128+
> exhibits a mild dependency in Hamming weights that generates a failure after 5
> TB of output in our test. We believe this slight bias cannot affect any
> application.

Go implementation based on a C reference implementation by David Blackman and
Sebastiano Vigna.

For more information, visit the [xoshiro / xoroshiro generators and the PRNG shootout][PRNGSHoutout] page.

### MT19937-64
period 2<sup>19937</sup>-1

This is a pure Go implementation based on the mt19937-64.c C implementation
by Makoto Matsumoto and Takuji Nishimura.

More information on the Mersenne Twister algorithm and other implementations
are available from http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html

If you use this algorithm, please make sure that you comply with the
accompanying license (see the license file mt19937/LICENSE_MT).

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

    "github.com/db47h/rand64/v3/iorand"
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

These benchmarks where done with go 1.12.5

The last result is for the default PRNG provided by the standard library's
rand.NewSource() for comparison:

| Algorithm     | AMD FX-6300 | Core i5 6200U | ARM Cortex-A7 @900MHz |
|---------------|------------:|--------------:|-----------------:|
| xoshiro256**  |  5.50 ns/op |               |      106.0 ns/op |
| xoshiro256+   |  5.48 ns/op |               |       86.1 ns/op |
| xoroshiro128**|  5.15 ns/op |               |       79.2 ns/op |
| xoroshiro128+ |  5.13 ns/op |               |       62.7 ns/op |
| splitmix64    |  4.31 ns/op |               |       77.5 ns/op |
| MT19937       |  8.91 ns/op |               |      136.0 ns/op |
| GoRand        |  6.62 ns/op |               |       68.4 ns/op |

Note that the benchmarks show slower performance compared to earlier releases.
This is due to the fact that we did call Rng.Uint64 directly instead of going
through the rand.Rand64 interface. In order to do a fair comparison with the Go
standard library's rng, all benchmarks now through a rand.Source64 interface.

## Go module support

rand64 supports go modules. Previous versions 1.x and 2.x have been moved to
their respective branches. Since semver tags with no go.mod seemed to upset go
modules, tags for these versions have been reset.

[PRNGShoutout]: http://xoshiro.di.unimi.it/
[travisImg]: https://travis-ci.org/db47h/rand64.svg?branch=master
[travis]: https://travis-ci.org/db47h/rand64
[goreportImg]: https://goreportcard.com/badge/github.com/db47h/rand64
[goreport]: https://goreportcard.com/report/github.com/db47h/rand64
[godocImg]: https://godoc.org/github.com/db47h/rand64?status.svg
[godoc]: http://godoc.org/github.com/db47h/rand64
