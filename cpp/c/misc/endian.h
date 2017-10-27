#ifndef ENDIAN_H_
#define ENDIAN_H_

#include "macros.h"

BEGIN_EXTERN_C()

// 0: little endian, 1: big endian
inline int checkEndian() {
    union {
        int a;
        char b;
    } u;
    u.a = 1;
    return u.b == 1 ? 0 : 1;
}

END_EXTERN_C()

#endif  // end ENDIAN_H_
