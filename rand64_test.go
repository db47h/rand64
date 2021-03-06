package rand64_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/db47h/rand64/v3/mt19937"
	"github.com/db47h/rand64/v3/pcg"
	"github.com/db47h/rand64/v3/splitmix64"
	"github.com/db47h/rand64/v3/xoroshiro"
	"github.com/db47h/rand64/v3/xoshiro"
)

const (
	SEED1 = 1387366483214
)

// Short example with single value seeding. Since the PRNG's state is larger
// than 64 bits, it will be seeded using values generated by a splitmix64 PRNG.
func Example() {
	// Use a xoshiro256** PRNG as a rand.Source for rand.New
	rng := rand.New(&xoshiro.Rng256SS{})
	// seed the PRNG.
	rng.Seed(time.Now().UnixNano())
	// pull values
	_ = rng.Uint64()
}

/* Benchmarks */

func BenchmarkXoroshiro128starstar(b *testing.B) {
	s := rand.Source64(&xoroshiro.Rng128SS{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkXoroshiro128plus(b *testing.B) {
	s := rand.Source64(&xoroshiro.Rng128P{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkXoshiro256starstar(b *testing.B) {
	s := rand.Source64(&xoshiro.Rng256SS{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkXoshiro256plus(b *testing.B) {
	s := rand.Source64(&xoshiro.Rng256P{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkPCG(b *testing.B) {
	s := rand.Source64(&pcg.Rng{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkSplitmix64(b *testing.B) {
	s := rand.Source64(&splitmix64.Rng{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkMt19937(b *testing.B) {
	s := rand.Source64(&mt19937.Rng{})
	s.Seed(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkGoRand(b *testing.B) {
	s := rand.NewSource(SEED1).(rand.Source64)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}
