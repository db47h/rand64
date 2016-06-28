// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iorand_test

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"

	"github.com/db47h/rand64"
	"github.com/db47h/rand64/iorand"
	"github.com/db47h/rand64/randutil"
)

// Wrap crypto/rand in an IoRand
func ExampleNew() {
	// first, wrap rand.Reader in a buffered bufio.Reader
	bufferedReader := bufio.NewReader(rand.Reader)
	// Create the new IoRand
	ior := iorand.New(bufferedReader, binary.LittleEndian)
	// use it as rand.Source for rand64.New
	rng := rand64.New(ior)
	// get random numbers...
	for i := 0; i < 4; i++ {
		_ = rng.Uint64()
	}

	// A one liner to quickly get a "good" seed for other PRNGs
	// Here we build a slice of 512 uint64 to be used with Source64.SeedFromSlice
	seedSlice := rand64.New(iorand.New(rand.Reader, binary.LittleEndian)).BulkUint64(512)
	// the function randutil.GenerateSeed() does the same:
	seedSlice = randutil.GenerateSeed(512)

	// the following is just to make the example compile.
	_ = seedSlice
}
