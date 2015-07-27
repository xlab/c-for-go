#ifndef _ASSERT_H_
#define _ASSERT_H_

void
__TODO_ASSERTION_FAILED (char *file, int line, char *msg)
{
}

#endif /* _ASSERT_H_ */

#undef assert

#ifdef NDEBUG
#define assert(ignore) ((void)0)
#else
#define assert(x) ((void)((x) ? 0 : (__TODO_ASSERTION_FAILED(__FILE__, __LINE__, #x), 0)))
#endif /* NDEBUG */
