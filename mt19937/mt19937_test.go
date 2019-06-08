package mt19937_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/db47h/rand64/v3/mt19937"
)

func Example() {
	init := []uint64{
		0x12345, 0x23456, 0x34567, 0x45678,
	}
	mt := new(mt19937.Rng)
	mt.SeedFromSlice(init)

	fmt.Println("10 outputs of Rng.Uint64()")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %20d", mt.Uint64())
		if i%5 == 4 {
			fmt.Println()
		}
	}
	for i := 0; i < 1000; i++ {
		mt.Uint64()
	}
	fmt.Println("10 more")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %20d", mt.Uint64())
		if i%5 == 4 {
			fmt.Println()
		}
	}

	// Output:
	// 10 outputs of Rng.Uint64()
	//   7266447313870364031  4946485549665804864 16945909448695747420 16394063075524226720  4873882236456199058
	//  14877448043947020171  6740343660852211943 13857871200353263164  5249110015610582907 10205081126064480383
	// 10 more
	//  14907209235746902445 15452338815569321965 17045090235069538607 15507333859934612093   157175897107904252
	//   2578005313950236321  6502648805754593060 13133523174961431106  2698278206396822833  3278969850082110371
}

func TestRng_Seed(t *testing.T) {
	// we'll just check the default seed values
	r0 := mt19937.Rng{}
	r5489 := mt19937.Rng{}
	r := mt19937.Rng{}
	rc := mt19937.Rng{}

	r0.Seed(0)
	r5489.Seed(5489)
	// keep r unseeded. Should auto-seed to 5489

	v0 := r0.Uint64()
	v5489 := r5489.Uint64()
	if v0 != v5489 {
		t.Fatal("v0 != v5489")
	}
	v := r.Uint64()
	if v != v0 {
		t.Fatal("v != v0")
	}
	var s int64
	for s == 5489 || s == 0 {
		s = time.Now().UnixNano()
	}
	rc.Seed(s)
	vc := rc.Uint64()
	if vc == v {
		t.Fatal("vc == v")
	}
}

func TestRng_Int63(t *testing.T) {
	r := mt19937.Rng{}

	s := time.Now().UnixNano()
	r.Seed(s)
	u := r.Uint64()
	r.Seed(s)
	i := r.Int63()

	if i != int64(u>>1) {
		t.Fatalf("%x != %x", i, int64(u>>1))
	}
}
