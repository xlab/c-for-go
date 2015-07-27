#ifndef _FCNTL_H_
#define _FCNTL_H_

#include <sys/types.h>

extern int fcntl (int, int, ...);

struct flock
{
  short l_type;
  short l_whence;
  off_t l_start;
  off_t l_len;
  pid_t l_pid;
};

//#define F_DUPFD               0
//#define F_GETFD               1
//#define F_SETFD               2
//#define F_GETFL               3
//#define F_SETFL               4
#define F_GETLK		5
#define F_SETLK		6
#define F_SETLKW      7
//#define F_SETOWN      8
//#define F_GETOWN      9
//#define F_SETSIG      10
//#define F_GETSIG      11

#define F_RDLCK	0
#define F_WRLCK	1
#define F_UNLCK	2

// #define O_ACCMODE       00000003
#define O_RDONLY        00000000
// #define O_WRONLY        00000001
#define O_RDWR          00000002
#define O_CREAT         00000100
#define O_EXCL          00000200
// #define O_NOCTTY        00000400
// #define O_TRUNC         00001000
// #define O_APPEND        00002000
// #define O_NONBLOCK      00004000
// #define O_DSYNC         00010000
// #define FASYNC          00020000
// #define O_DIRECT        00040000
// #define O_LARGEFILE     00100000
// #define O_DIRECTORY     00200000
// #define O_NOFOLLOW      00400000
// #define O_NOATIME       01000000
// #define O_CLOEXEC       02000000

#endif /* _FCNTL_H_ */
