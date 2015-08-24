#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

const uint8_t* test_proxy_string(const unsigned char *s1, const char *s2);
uint8_t* test_proxy_bytes(unsigned char *b1, size_t len1, char *b2, size_t len2);
ssize_t test_find_byte(char *b, size_t len, char c);

void test_a4_byte(char *b[4], size_t len);
void test_a4_string(const char *s[4]);
void test_a4s_byte(char *b[4], size_t len);
void test_a4s_string(const char **s[4], size_t len);
void test_a2a2_byte(char s[2][2]);
void test_a2a2_string(const char *s[2][2]);
void test_ss_byte(char **s, size_t len1, size_t len2);
void test_ss_string(const char ***s, size_t len1, size_t len2);
void test_a4ss_byte(char **s[4], size_t len1, size_t len2);
void test_a4ss_string(const char ***s[4], size_t len1, size_t len2);

#define MAX_NAME_LEN 255
const ssize_t MAX_MESSAGE_LEN 512

typedef message struct {
	uint8_t *from[MAX_NAME_LEN];
	uint8_t *to[MAX_NAME_LEN];
	uint8_t *message[MAX_MESSAGE_LEN];
	const uint8_t *signature;
	bool sent;
} message_t;

size_t test_send_message(message_t *m, uint8_t *buf);
