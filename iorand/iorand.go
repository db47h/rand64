// Copyright (c) 2014-2019, Denis Bernard <db047h@gmail.com>
// Use of this source code is governed by the ISC license that
// can be found in the LICENSE file.

/*
Package iorand provides a ran64.Source64 wrapper around an io.Reader.

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

// IoRand is a wrapper around an io.Reader that implements the rand64.Source64 interface.
//
type IoRand struct {
	r         io.Reader
	ByteOrder binary.ByteOrder
	Err       error // latest io error
}

// New returns a new IoRand wrapper around an io.Reader using the provided binary.ByteOrder
// for reads. Since IoRand will not buffer its input, providing a buffered Reader is recommended.
//
func New(r io.Reader, bo binary.ByteOrder) *IoRand {
	return &IoRand{r: r, ByteOrder: bo, Err: nil}
}

// Seed is a no-op as IoRand does not support Seeding.
//
func (r *IoRand) Seed(seed int64) {
	// nothing to do
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
// Returns 0 and sets r.Err to a non-nil value if an error occurs.
//
func (r *IoRand) Uint64() (n uint64) {
	var b [8]byte
	if _, err := io.ReadFull(r.r, b[:]); err != nil {
		r.Err = err
		return 0
	}
	return r.ByteOrder.Uint64(b[:])
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
//
func (r *IoRand) Int63() int64 {
	return int64(r.Uint64() >> 1)
}
