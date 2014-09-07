// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The iorand package implements a ran64.Source64 wrapper around an io.Reader.

This implementation does not require seeding and can be used for example to
read random numbers from crypto.rand.Reader.

Since there may be errors when reading from an io.Reader, special care must be
taken when using this source. If an error occurs, IoRand.Uint64() returns 0 and
IoRand.Err will be set to the latest error.
*/
package iorand

import (
	"encoding/binary"
	"io"
)

const (
	_N = 1024 / 8 // 1kb state buffer
)

// IoRand is a wrapper around an io.Reader that implements the rand64.Source64 interface.
type IoRand struct {
	r         io.Reader
	byteOrder binary.ByteOrder
	Err       error // latest io error
}

// New returns a new ran64.Source64 wrapper around an io.Reader using the provided binary.ByteOrder
// for reads. Since IoRand will not buffer its input, providing a buffered Reader is recommended.
func New(r io.Reader, bo binary.ByteOrder) *IoRand {
	return &IoRand{r: r, byteOrder: bo, Err: nil}
}

// no-op dummy function as IoRand does not support Seeding
func (r *IoRand) Seed64(seed uint64) {
	// nothing to do
}

// no-op dummy function as IoRand does not support Seeding
func (r *IoRand) SeedFromSlice(key []uint64) {
	// nothing to do
}

// Uint64 generates a random number on [0, 2^64-1]-interval
// Returns 0 AND sets r.Err if an error occurs.
func (r *IoRand) Uint64() (n uint64) {
	var b [8]byte
	if _, err := io.ReadFull(r.r, b[:]); err != nil {
		r.Err = err
		return 0
	}
	return r.byteOrder.Uint64(b[:])
}
