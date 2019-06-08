package xoshiro_test

import (
	"fmt"
	"math/rand"

	"github.com/db47h/rand64/v3/xoshiro"
)

const (
	SEED1 = 1387366483214
)

func ExampleRng256P() {
	src := xoshiro.Rng256P{}
	src.Seed(SEED1)
	rng := rand.New(&src)
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint32())
	}
	fmt.Println("")
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint64())
	}
	fmt.Println("")
	// Play craps
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d%d", rng.Intn(6)+1, rng.Intn(6)+1)
	}

	// Output:
	// 2171228231 173189436 470990771 2412873522
	//  1816180620218953111 7257068590675658289 8111314002208617320 6106779797696663770
	//  46 33 61 55 34 41 65 44 22 33
}

func ExampleRng256SS() {
	src := xoshiro.Rng256SS{}
	src.Seed(SEED1)
	rng := rand.New(&src)
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint32())
	}
	fmt.Println("")
	for i := 0; i < 4; i++ {
		fmt.Printf(" %d", rng.Uint64())
	}
	fmt.Println("")
	// Play craps
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d%d", rng.Intn(6)+1, rng.Intn(6)+1)
	}

	// Output:
	// 1703513406 2124925634 3601057375 1523263934
	//  14206081294295289219 1400819388980187612 655760235528857176 11230280953057933127
	//  11 13 64 51 53 15 16 55 12 61
}
