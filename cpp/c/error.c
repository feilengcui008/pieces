#include "error.h"
#include <stdio.h>

BEGIN_EXTERN_C()

void error_exit(const char *msg)
{
  //perror(msg);
  fprintf(stderr, "FILE:%s, LINENO:%d, REASON:%s\n", __FILE__, __LINE__, msg);
  exit(EXIT_FAILURE);
}

END_EXTERN_C()
