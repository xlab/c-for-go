#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>

#pragma once

int foo_pass_int(int i1, int i2);
const char* foo_pass_string(const unsigned char *s1, const char *s2);
uint8_t* foo_pass_bytes(unsigned char *b1, size_t n1, uint8_t *b2, size_t n2);
ssize_t foo_find_char(char *s, char c);

void foo_a4_byte(uint8_t b[4]);
void foo_a4_string(const char *s[4]);
void foo_a4_s_byte(uint8_t *b[4], size_t n);
void foo_a4_s_string(const char **s[4], size_t n);
void foo_a2_a2_byte(uint8_t b[2][2]);
void foo_a2_a2_string(const char *s[2][2]);
void foo_a2_a2_s_byte(uint8_t *b[2][2], size_t n);
void foo_a2_a2_s_string(const char **s[2][2], size_t n);
void foo_s_s_byte(uint8_t **b, size_t n1, size_t n2);
void foo_s_s_string(const char ***s, size_t n1, size_t n2);
void foo_a4_s_s_byte(uint8_t **b[4], size_t n1, size_t n2);
void foo_a4_s_s_string(const char ***s[4], size_t n1, size_t n2);
void foo_a2_a2_s_s_byte(uint8_t **b[2][2], size_t n1, size_t n2);
void foo_a2_a2_s_s_string(const uint8_t ***s[2][2], size_t n1, size_t n2);

#define FOO_MESSAGE_HEADER "msg"
#define FOO_MESSAGE_HEADER_LEN 3
#define FOO_NAME_LEN_MAX 50
#define FOO_MESSAGE_LEN_MAX 140

typedef struct foo_message {
	uint8_t from[FOO_NAME_LEN_MAX];
	uint8_t to[FOO_NAME_LEN_MAX];
	uint8_t message[FOO_MESSAGE_LEN_MAX];
	const char *signature;
	bool sent;
} foo_message_t;

size_t foo_send_message(foo_message_t *m, uint8_t *buf);
