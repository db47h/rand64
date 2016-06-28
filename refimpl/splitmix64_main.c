

#include <stdint.h>
#include <stdio.h>

extern uint64_t sm64;
extern uint64_t sm64_next();

#define SEED1 1387366483214

int main() {
	int i;
	sm64 = SEED1;
	for (i = 0; i < 10; i++) {
		uint64_t z = sm64_next();
		printf("\t0x%016lX,\n", z);
	}
	return 0;
}
