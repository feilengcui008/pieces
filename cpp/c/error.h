#ifndef _TAN_ERROR_H_
#define _TAN_ERROR_H_

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "macros.h"

BEGIN_EXTERN_C()

void error_exit(const char *msg);

END_EXTERN_C()

#endif  // end _TAN_ERROR_H_
