#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <unistd.h>

void test_mmap_threshold() {
  fprintf(stdout, "---------------------------\n");
  int changed_to = 100 * 1024;
  mallopt(M_MMAP_THRESHOLD, changed_to);
  int alloc_size1 = 110 * 1024;
  fprintf(stdout, "after change mmap threshold to %d, alloc size : %d\n", changed_to, alloc_size1);
  char *p1 = (char *)malloc(alloc_size1);
  fprintf(stdout, "address is : %p\n", p1);

  fprintf(stdout, "=======\n");
  malloc_stats();
}

int main(int argc, char **argv) {
  test_mmap_threshold();
  return 0;
}
