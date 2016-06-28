#include <stdint.h>
#include <stdio.h>

extern uint64_t sm64;
extern uint64_t sm64_next(void);

extern uint64_t s[STATE];
extern uint64_t next(void);

#define SEED1 1387366483214

int main() {
	int i;
	sm64 = SEED1;
	for (i = 0; i < STATE; i++) {
		s[i] = sm64_next();
	}
	
	for (i = 0; i < 4; i++) {
		uint64_t z = next();
		printf(" %u", (uint32_t)(z >> 32));
	}
	puts("");
	for (i = 0; i < 4; i++) {
		uint64_t z = next();
		printf(" %lu", z);
	}
	puts("");
	for (i = 0; i < 10; i ++) {
		uint32_t v1 = ((uint32_t)(next() >> 32)) % 6 + 1;
		uint32_t v2 = ((uint32_t)(next() >> 32)) % 6 + 1;
		printf(" %u%u", v1, v2);
	}
	puts("");
	for (i = 0; i < 4; i++) {
		uint64_t z = next();
		printf(" %lu", z>>1);
	}
	puts("");
	return 0;
}
