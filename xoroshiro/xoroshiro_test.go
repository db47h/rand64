package xoroshiro_test

import (
	"fmt"
	"math/rand"

	"github.com/db47h/rand64/v3/xoroshiro"
)

const (
	SEED1 = 1387366483214
)

func Example() {
	src := xoroshiro.Rng128P{}
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
	// 3672052799 776653214 1122818236 1139848352
	//  14850484681238877506 7018105211938886447 5908230704518956940 2042158984393296588
	//  65 53 21 56 44 16 23 42 55 41
}
