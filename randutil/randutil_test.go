// Copyright 2014 Denis Bernard. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randutil_test

import (
	"github.com/db47h/rand64"
	"github.com/db47h/rand64/randutil"
	"github.com/db47h/rand64/xorshift"
)

func ExampleGenerateSeed() {
	// A one liner and quickly get a "good" seed for other PRNGs. This will
	// generate a slice of 16 uint64 values
	seedSlice := randutil.GenerateSeed(16)
	// now create the rand64.Rand object with our PRNG source of choice
	rng := rand64.New(xorshift.New1024star())
	// and seed it.
	rng.SeedFromSlice(seedSlice)

	// we could even compact this in two lines:
	rng = rand64.New(xorshift.New1024star())
	// and seed it.
	rng.SeedFromSlice(randutil.GenerateSeed(16))
}
