#include "ConcurrentSingleton.h"
#include <stdio.h>
#include <sys/syscall.h>
#include <unistd.h>
#include <array>
#include <thread>

int main(int argc, char *argv[]) {
    auto func = []() {
        auto s = ConcurrentSingleton<int>::instance();
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
