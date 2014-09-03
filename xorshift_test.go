// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand64_test

import (
	"math/rand"
	"github.com/wildservices/rand64"
	"testing"
)

/* Benchmarks */

func BenchmarkXorShift64star(b *testing.B) {
	s := rand64.NewXorShift64star(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkXorShift128plus(b *testing.B) {
	s := rand64.NewXorShift128plus(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}

func BenchmarkXorShift1024star(b *testing.B) {
	s := rand64.NewXorShift1024star(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Uint64()
	}
}


func BenchmarkRandSource(b *testing.B) {
	s := rand.NewSource(SEED1)
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

/* Tests */

