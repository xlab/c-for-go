#include <stdio.h>
#include <string.h>
#include <stdlib.h>
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

#define FOO_MESSAGE_HEADER "msg:"
#define FOO_MESSAGE_HEADER_LEN 4
#define FOO_ID_LEN 4

typedef struct foo_attachment_raw {
	uint8_t *data;
	size_t size;
} foo_attachment_t;

typedef struct foo_message {
	uint8_t from_id[FOO_ID_LEN];
	uint8_t to_id[FOO_ID_LEN];
	char *message;
	size_t attachments_len;
	foo_attachment_t* attachments;
	_Bool sent;
} foo_message_t;

foo_message_t* foo_new_message();
size_t foo_send_message(foo_message_t *m, uint8_t *buf);

// doesn't work in CGO yet
// int foo_anon_test(struct {int n;} a, struct {int n;} b);

struct foo_anon_tag {
	int n;
};

int foo_pass_anon_tag(struct foo_anon_tag a, struct foo_anon_tag b);

struct foo_has_anon_tag {
	int a;
	struct foo_inner_anon_tag {
		int n;
	} b;
};

struct foo_has_anon {
	int a;
	struct {
		int anon_n;
	};
};

typedef int foo_fcb(int);
