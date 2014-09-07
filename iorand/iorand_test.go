// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iorand_test

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/wildservices/rand64"
	"github.com/wildservices/rand64/iorand"
	"github.com/wildservices/rand64/randutil"
)

// Wrap crypto/rand in an IoRand
func ExampleNew() {
	// first, wrap rand.Reader in a buffered bufio.Reader
	bufferedReader := bufio.NewReader(rand.Reader)
	// Create the new IoRand
	ior := iorand.New(bufferedReader, binary.LittleEndian)
	// wrap it in a rand64.Rand64 for access to utility functions
	rng := rand64.New(ior)
	for i := 0; i < 4; i++ {
		v := rng.Uint64()
	}

	// A one liner to quickly get a "good" seed for other PRNGs
	// Here we build a slice of 512 uint64 to be used with Source64.SeedFromSlice
	seedSlice := rand64.New(iorand.New(rand.Reader, binary.LittleEndian)).BulkUint64(512)
	// the randutil.GenerateSeed() does the same:
	seedSlice := randutil.GenerateSeed(512)
}
