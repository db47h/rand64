TARGETS := splitmix64 xoroshiro128plus xoroshiro128starstar xoshiro256plus xoshiro256starstar

.PHONY: all

all: $(TARGETS)
	
splitmix64: splitmix64.c splitmix64_main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xoroshiro128plus: splitmix64.c xoroshiro128plus.c main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xoroshiro128starstar: splitmix64.c xoroshiro128starstar.c main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xoshiro256plus: splitmix64.c xoshiro256plus.c main.c
	$(CC) -Wall -DSTATE=4 -o $@ $^

xoshiro256starstar: splitmix64.c xoshiro256starstar.c main.c
	$(CC) -Wall -DSTATE=4 -o $@ $^

clean:
	rm -f *.o $(TARGETS)
