#ifndef _PTHREAD_H
#define _PTHREAD_H

typedef int __TODO_PTHREAD_ATTR_T;
typedef int __TODO_PTHREAD_MUTEXATTR_T;
typedef int __TODO_PTHREAD_MUTEX_T;
typedef int __TODO_PTHREAD_T;

typedef __TODO_PTHREAD_ATTR_T pthread_attr_t;
typedef __TODO_PTHREAD_MUTEXATTR_T pthread_mutexattr_t;
typedef __TODO_PTHREAD_MUTEX_T pthread_mutex_t;
typedef __TODO_PTHREAD_T pthread_t;

#define PTHREAD_MUTEX_INITIALIZER 0

enum
{
  PTHREAD_MUTEX_RECURSIVE
};

int
pthread_mutexattr_init (pthread_mutexattr_t * TODO)
{
}

int
pthread_mutexattr_settype (pthread_mutexattr_t * TODO, int TODO2)
{
}

int
pthread_mutexattr_destroy (pthread_mutexattr_t * TODO)
{
}

int
pthread_mutex_trylock (pthread_mutex_t * TODO)
{
}

int
pthread_create (pthread_t * TODO, const pthread_attr_t * TODO2,
		void *(*TODO3) (void *), void *TODO4)
{
}

int
pthread_join (pthread_t TODO, void * *TODO2)
{
}

#endif /* _PTHREAD_H */
