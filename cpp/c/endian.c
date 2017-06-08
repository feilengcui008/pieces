#include "endian.h"

// 0 : little endian
// 1 : big endian
int checkEndian()
{
  union {
    int a;
    char b;
  } u;
  u.a = 1;
  return u.b == 1 ? 0 : 1;
}
