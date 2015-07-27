#ifndef _STDARG_H_
#define _STDARG_H_

typedef void *va_list;

#define va_start(ap, parmN) __TODO_VA_START(ap, parmN)
#define va_arg(ap, type) *(type*)__TODO_VA_ARG(ap)
#define va_end(ap) __TODO_VA_END(ap)
//#define va_copy(dst, src) __TODO_VA_COPY(dst, src)

void
__TODO_VA_START ()
{
}

void *
__TODO_VA_ARG (va_list ap)
{
}

void
__TODO_VA_END (va_list ap)
{
}

//void __TODO_VA_COPY(void *dst, void *src) {}

#endif /* _STDARG_H_ */
