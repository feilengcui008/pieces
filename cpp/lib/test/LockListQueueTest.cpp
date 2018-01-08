#include "lock_list_queue.h"

#include <stdio.h>
#include <sys/syscall.h>
#include <unistd.h>
#include <memory>
#include <thread>

int main(int argc, char *argv[]) {
    Tan::LockListQueue<int> lq;

    auto producer = [&lq]() {
        auto begin = time(NULL);
        while (time(NULL) < begin + 3) {
            lq.put(new Tan::Node<int>(new int(syscall(SYS_gettid))));
        }
    };

    auto consumer = [&lq]() {
        auto begin = time(NULL);
        while (time(NULL) < begin + 5) {
            auto nd = lq.get();
            fprintf(stdout, "get ele %d\n", *(nd->get()));
            delete nd;
        }
    };

    const int N = 3;
    std::thread thds[N + 1];
    for (int i = 0; i < N; ++i) {
        thds[i] = std::thread(producer);
    }
    thds[N] = std::thread(consumer);

    for (auto &ele : thds) {
        ele.join();
    }

    return 0;
}
