#include "foo.h"

int foo_pass_int(int i1, int i2) {
	return i1 + i2;
}

const char* foo_pass_string(const unsigned char *s1, const char *s2) {
	if (strcmp((const char*)s1, "a") == 0 && strcmp(s2, "b") == 0) {
		return "ab";
	}
	return "";
}

uint8_t* foo_pass_bytes(unsigned char *b1, size_t n1, uint8_t *b2, size_t n2) {
	uint8_t* result = (uint8_t*)malloc(n1 + n2);
	memcpy(result, b1, n1);
	memcpy(&result[n1], b2, n2);
	return result;
}

ssize_t foo_find_char(char *s, char c) {
	for (ssize_t i = 0; s[i] != 0; i++) {
		if (s[i] == c) {
			return i;
		}
	}
	return -1;
}

size_t foo_send_message(foo_message_t *m, uint8_t *buf) {
	if (m == 0) {
		return 0;
	}
	size_t size = 0;
	memcpy(buf, FOO_MESSAGE_HEADER, FOO_MESSAGE_HEADER_LEN);
	size += FOO_MESSAGE_HEADER_LEN;
	memcpy(buf + size, m->from, FOO_NAME_LEN_MAX);
	size += FOO_NAME_LEN_MAX;
	memcpy(buf + size, m->to, FOO_NAME_LEN_MAX);
	size += FOO_NAME_LEN_MAX;
	memcpy(buf + size, m->message, FOO_MESSAGE_LEN_MAX);
	size += FOO_MESSAGE_LEN_MAX;
	memcpy(buf + size, m->signature, strlen(m->signature));
	size += strlen(m->signature);
	m->sent = true;
	return size;
}

void foo_a4_byte(uint8_t b[4]) {
	for (int i = 0; i < 4; i++) {
		b[i]++;
	}
}

void foo_a4_string(const char *s[4]) {
	for (int i = 0; i < 4; i++) {
		s[i] = "go";
	}
}

void foo_a4_s_byte(uint8_t *b[4], size_t n) {
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < n; j++) {
			b[i][j]++;
		}
	}
}

void foo_a4_s_string(const char **s[4], size_t n) {
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < n; j++) {
			s[i][j] = "go";
		}
	}
}

void foo_a2_a2_byte(uint8_t b[2][2]) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			b[i][j]++;
		}
	}
}

void foo_a2_a2_string(const char *s[2][2]) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			s[i][j] = "go";
		}
	}
}

void foo_a2_a2_s_byte(uint8_t *b[2][2], size_t n) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			for (int k = 0; k < n; k++) {
				b[i][j][k]++;
			}
		}
	}
}

void foo_a2_a2_s_string(const char **s[2][2], size_t n) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			for (int k = 0; k < n; k++) {
				s[i][j][k] = "go";
			}
		}
	}
}

void foo_s_s_byte(uint8_t **b, size_t n1, size_t n2) {
	for (int i = 0; i < n1; i++) {
		for (int j = 0; j < n2; j++) {
			b[i][j]++;
		}
	}
}

void foo_s_s_string(const char ***s, size_t n1, size_t n2) {
	for (int i = 0; i < n1; i++) {
		for (int j = 0; j < n2; j++) {
			s[i][j] = "go";
		}
	}
}

void foo_a4_s_s_byte(uint8_t **b[4], size_t n1, size_t n2) {
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < n1; j++) {
			for (int k = 0; k < n2; k++) {
				b[i][j][k]++;
			}
		}
	}
}

void foo_a4_s_s_string(const char ***s[4], size_t n1, size_t n2) {
	for (int i = 0; i < 4; i++) {
		for (int j = 0; j < n1; j++) {
			for (int k = 0; k < n2; k++) {
				s[i][j][k] = "go";
			}
		}
	}
}

void foo_a2_a2_s_s_byte(uint8_t **b[2][2], size_t n1, size_t n2) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			for (int k = 0; k < n1; k++) {
				for (int l = 0; l < n2; l++) {
					b[i][j][k][l]++;
				}
			}
		}
	}
}

void foo_a2_a2_s_s_string(const uint8_t ***s[2][2], size_t n1, size_t n2) {
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			for (int k = 0; k < n1; k++) {
				for (int l = 0; l < n2; l++) {
					s[i][j][k][l] = (uint8_t*)"go";
				}
			}
		}
	}
}
