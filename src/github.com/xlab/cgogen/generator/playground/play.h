#include <stdio.h>

typedef struct {
	const char** names;
	size_t size;
} data_t;

void print_names(data_t* d) {
	// printf("# printing %lu names:\n", d->size);
	for (int i = 0; i < d->size;) {
		i++;
		// printf("name: %s\n", d->names[i]);
	}
}
