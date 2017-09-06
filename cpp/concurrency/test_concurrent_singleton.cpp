#include "concurrent_singleton.h"
#include <stdio.h>
#include <thread>
#include <unistd.h>
#include <array>
#include <sys/syscall.h>

int main(int argc, char *argv[]) {

  auto func = []() {
    auto s = Tan::ConcurrentSingleton<int>::instance();
    fprintf(stdout, "%ld, %p\n", syscall(SYS_gettid), &s);
  };

  static const int N = 100;

  std::array<std::thread, N> arr;
  for (int i = 0; i < N; i++) {
    arr[i] = std::thread(func);
  }

  for (auto &ele : arr) {
    ele.join();
  }

  return 0;
}
