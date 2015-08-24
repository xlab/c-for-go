#include <stdio.h>

typedef struct {
	const char** names;
	size_t size;
	unsigned char flag;
} data_t;

typedef struct {
	unsigned char flag[4];
} lol_t;

void print_names(data_t* d) {
	printf("# printing %lu names:\n", d->size);
	for (int i = 0; i < d->size; i++) {
		printf("name: %s\n", d->names[i]);
	}
}

void lol(lol_t *l) {
	printf("Flag0: %d\n", l->flag[0]);
	printf("Flag1: %d\n", l->flag[1]);
	printf("Flag2: %d\n", l->flag[2]);
	printf("Flag3: %d\n", l->flag[3]);
}

void stringCube(const char**** str, int x, int y, int z) {
	// printf("Cube dimensions: x=%d, y=%d, z=%d\n", x, y, z);
	for (int i = 0; i < x; i++) {
		for (int j = 0; j < y; j++) {
		 	for (int k = 0; k < z; k++) {
		 		// printf("text at [%d][%d][%d] = %s\n", i, j, k, str[i][j][k]);
		 		str[i][j][k] = "HELLO";
			}
		}
	}
}
