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
)

// Wrap crypto/rand in an IoRand
func ExampleNew() {
	// first, wrap crypto/rand.Reader in a buffered bufio.Reader
	bufferedReader := bufio.NewReader(rand.Reader)
	// Create the new IoRand
	ior := iorand.New(bufferedReader, binary.LittleEndian)
	// use it as rand.Source for rand64.New
	rng := rand64.New(ior)
	// get random numbers...
	for i := 0; i < 4; i++ {
		_ = rng.Uint64()
	}

	// the randutil package provides a utility function just to do that in a one
	// liner and quickly get a "good" seed for other PRNGs
}
