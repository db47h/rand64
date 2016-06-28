TARGETS := splitmix64 xoroshiro128plus xorshift128plus xorshift1024star

.PHONY: all

all: $(TARGETS)
	
splitmix64: splitmix64.c splitmix64_main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xoroshiro128plus: splitmix64.c xoroshiro128plus.c main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xorshift128plus: splitmix64.c xorshift128plus.c main.c
	$(CC) -Wall -DSTATE=2 -o $@ $^

xorshift1024star: splitmix64.c xorshift1024star.c main.c
	$(CC) -Wall -DSTATE=16 -o $@ $^

clean:
	rm -f *.o $(TARGETS)
