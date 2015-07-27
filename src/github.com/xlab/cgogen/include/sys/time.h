#ifndef _SYS_TIME_H_
#define _SYS_TIME_H_

#include <sys/types.h>

struct timeval
{
  time_t tv_sec;
  long tv_usec;
};

#endif /* _SYS_TIME_H_ */
