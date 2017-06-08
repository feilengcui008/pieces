#ifndef _TAN_MACROS_H_
#define _TAN_MACROS_H_

/* some common useful macros */

// extern "C"
#ifdef __cplusplus
#define BEGIN_EXTERN_C() extern "C" {
#define END_EXTERN_C() }
#else
#define BEGIN_EXTERN_C()
#define END_EXTERN_C()
#endif

// dynamic library api export
#if defined(__GNU__) && __GNUC__ >= 4
#define EXPORT_API __attribute__ ((visibility("default")))
#define EXPORT_DLEXPORT __attribute__ ((visibility("default")))
#else
#define EXPORT_API
#define EXPORT_DLEXPORT
#endif

// alignment
#define MM_ALIGNMENT 8
#define MM_ALIGNMENT_MASK ~(MM_ALIGNMENT - 1)
#define MM_ALIGNMENT_SIZE(size) (((size) + MM_ALIGNMENT - 1) & MM_ALIGNMENT_MASK)

#endif  // end _TAN_MACROS_H_
