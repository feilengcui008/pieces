#ifndef ERROR_H_
#define ERROR_H_

#include <stdio.h>
#include <stdlib.h>
#include "macros.h"

BEGIN_EXTERN_C()

void error_exit(const char *msg) {
    fprintf(stderr, "FILE:%s, LINENO:%d, REASON:%s\n", __FILE__, __LINE__, msg);
    exit(EXIT_FAILURE);
}

END_EXTERN_C()

#endif  // end ERROR_H_
