#include "worker_pool.h"

void *callbacktest(void *args) {
  fprintf(stdout, "from thread: %d --- passed parameter: %d\n",
          (int)syscall(SYS_gettid), (int)(*(int *)(args)));
  return NULL;
}

int main(int argc, char **argv) {
  pool *p = (pool *)malloc(sizeof(pool));
  if (p == NULL) {
    fprintf(stderr, "malloc pool failed\n");
  }
  pool_init(p, 4, 10);
  int args[10];
  for (int i = 0; i < 11; i++) {
    args[i] = i;
  }
  for (int i = 0; i < 11; i++) {
    pool_add_task(p, &callbacktest, &args[i]);
  }
  // wait for task finished
  sleep(3);
  pool_stop(p);
  pool_destroy(p);
  return 0;
}
