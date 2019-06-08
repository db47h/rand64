package iorand_test

import (
	"bufio"
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/db47h/rand64/v3/iorand"
)

// Wrap crypto/rand in an IoRand
func ExampleNew() {
	// first, wrap crypto/rand.Reader in a buffered bufio.Reader
	bufferedReader := bufio.NewReader(crand.Reader)
	// Create the new IoRand
	ior := iorand.New(bufferedReader, binary.LittleEndian)
	// use it as rand.Source64
	rng := rand.New(ior)
	// get random numbers...
	for i := 0; i < 4; i++ {
		_ = rng.Uint64()
		_ = rng.Int63()
	}

	// Output:
}

func TestIoRand_Uint64(t *testing.T) {
	b := make([]byte, 8)
	var ctrl uint64
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		v := rand.Int31n(256)
		b[i] = byte(v)
		ctrl = (ctrl << 8) | uint64(v)
	}
	src := iorand.New(bytes.NewReader(b), binary.BigEndian)
	v := src.Uint64()
	if v != ctrl {
		t.Fatalf("%x != %x", v, ctrl)
	}
	v = src.Uint64()
	if v != 0 || src.Err != io.EOF {
		t.Fatalf("v = %x, err = %v", v, src.Err)
	}
}
