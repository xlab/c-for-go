#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>

#pragma once

int test_pass_int(int i1, int i2);
const char* test_pass_string(const unsigned char *s1, const char *s2);
uint8_t* test_pass_bytes(unsigned char *b1, size_t n1, uint8_t *b2, size_t n2);
ssize_t test_find_char(char *b, char c);

void test_a4_byte(uint8_t b[4]);
void test_a4_string(const char *s[4]);
void test_a4_s_byte(uint8_t *b[4], size_t n);
void test_a4_s_string(const char **s[4], size_t n);
void test_a2_a2_byte(uint8_t b[2][2]);
void test_a2_a2_string(const char *s[2][2]);
void test_a2_a2_s_byte(uint8_t *b[2][2], size_t n);
void test_a2_a2_s_string(const char **s[2][2], size_t n);
void test_s_s_byte(uint8_t **b, size_t n1, size_t n2);
void test_s_s_string(const char ***s, size_t n1, size_t n2);
void test_a4_s_s_byte(uint8_t **b[4], size_t n1, size_t n2);
void test_a4_s_s_string(const char ***s[4], size_t n1, size_t n2);
void test_a2_a2_s_s_byte(uint8_t **b[2][2], size_t n1, size_t n2);
void test_a2_a2_s_s_string(const uint8_t ***s[2][2], size_t n1, size_t n2);

#define MESSAGE_HEADER "msg"
#define MESSAGE_HEADER_LEN 3
#define MAX_NAME_LEN 255
#define MAX_MESSAGE_LEN 512

typedef struct test_message {
	uint8_t from[MAX_NAME_LEN];
	uint8_t to[MAX_NAME_LEN];
	uint8_t message[MAX_MESSAGE_LEN];
	const char *signature;
	bool sent;
} test_message_t;

size_t test_send_message(test_message_t *m, uint8_t *buf);
