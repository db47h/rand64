// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package randutil provides various utility functions around the rand64 package.
*/
package randutil

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"

	"github.com/db47h/rand64"
	"github.com/db47h/rand64/iorand"
)

// GenerateSeed creates a slice of n uint64 filled with random numbers generated
// by Go's default cryptograhically secure PRNG.
//
// The returned slice can be used as an argument to Source.SeedFromSlice()
func GenerateSeed(n uint) []uint64 {
	r := bufio.NewReaderSize(rand.Reader, int(n&(1<<31-1)))
	rng := rand64.New(iorand.New(r, binary.LittleEndian))
	b := make([]uint64, n)
	for i := range b {
		b[i] = rng.Uint64()
	}
	return b
}
