#include "proxy_test.h"

const char *test_proxy_string(const char *s1, const char *s2, const char *s3) {
	if (strcmp(s1, "a") == 0 && 
		strcmp(s2, "b") == 0 && 
		strcmp(s3, "c") == 0) {
		 	return "HELLO";
		}
	return "BYE";
}

unsigned char* test_proxy_bytes(unsigned char *buf, char *buf_two) {
	unsigned char* result = (unsigned char*)malloc(4);
	memcpy(result, buf, 2);
	memcpy(&result[2], buf_two, 2);
	return result;
}

const char *****test_proxy_ptr3_string(const char ****str[4]) {
	return str;
}
