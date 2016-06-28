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
// The returned slice can be used as an argument to Source64.SeedFromSlice()
func GenerateSeed(n uint) []uint64 {
	r := bufio.NewReaderSize(rand.Reader, int(n&(1<<31-1)))
	s := iorand.New(r, binary.LittleEndian)
	return rand64.New(s).BulkUint64(n)
	// TODO: IoRand.Uint64 should be checkd for errors here
}
