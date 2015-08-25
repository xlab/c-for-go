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

char a1;			// type of a1 is C.char
char *a2;			// type of a2 is *C.char
char a3[1];			// type of a3 is [1]C.char
char a4[1][2];		// type of a4 is [1][2]C.char
char a5[1][2][3];	// type of a5 is [1][2][3]C.char
char *a6[1][2][3];	// type of a6 is [1][2][3]*C.char
char **a7[1][2][3];	// type of a7 is [1][2][3]**C.char
char ***a8[1][2];	// type of a8 is [1][2]***C.char
char *a9[1][2];		// type of a9 is [1][2]*C.char
char **a10[1];		// type of a0 is [1]**C.char
char *a11[1];		// type of a1 is [1]*C.char

void b1(char a1) {} 			// as type C.char in _Cfunc_b1
void b2(char *a2) {} 			// as type *C.char in _Cfunc_b2
void b3(char a3[1]) {}  		// as type *C.char in _Cfunc_b3
void b4(char a4[1][2]) {}		// as type *[2]C.char in _Cfunc_b4
void b5(char a5[1][2][3]) {}	// as type *[2][3]C.char in _Cfunc_b5
void b6(char *a6[1][2][3]) {}	// as type *[2][3]*C.char in _Cfunc_b6
void b7(char **a7[1][2][3]) {}	// as type *[2][3]**C.char in _Cfunc_b7
void b8(char ***a8[1][2]) {}	// as type *[2]***C.char in _Cfunc_b8
void b9(char *a9[1][2]) {}		// as type *[2]*C.char in _Cfunc_b9
void b10(char **a10[1]) {}		// as type ***C.char in _Cfunc_b10
void b11(char *a11[1]) {}		// as type **C.char in _Cfunc_b11

