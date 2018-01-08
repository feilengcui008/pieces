#include "lock_ring_buffer.h"

#include <sys/syscall.h>
#include <unistd.h>
#include <thread>

int main(int argc, char *argv[]) {
    Tan::LockRingBuffer<int, 5> rb;

    auto producer = [&rb]() {
        rb.put(new int(syscall(SYS_gettid)));
        rb.printData();
    };

    const int N = 10;

    std::thread thds[N + 1 + 1];

    for (int i = 0; i < N; i++) {
        thds[i] = std::thread(producer);
    }

    auto consumer = [&rb]() {
        while (1) {
            auto ele = rb.get();
            std::cout << syscall(SYS_gettid) << " " << *ele << std::endl;
            delete ele;
        }
    };

    thds[N] = std::thread(consumer);

    auto prod = [&rb]() {
        while (1) {
            rb.put(new int(syscall(SYS_gettid)));
            rb.printData();
        }
    };

    thds[N + 1] = std::thread(prod);

    for (auto &ele : thds) {
        ele.join();
    }
    return 0;
}
