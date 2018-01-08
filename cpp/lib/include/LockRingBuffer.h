#ifndef LOCK_RING_BUFFER_H
#define LOCK_RING_BUFFER_H

#include <array>
#include <condition_variable>
#include <cstdint>
#include <iostream>
#include <mutex>

namespace Tan {

template <class T, uint64_t size>
class LockRingBuffer {
   public:
    LockRingBuffer();
    ~LockRingBuffer();
    void put(T *ele);
    T *get();

   public:
    static inline uint64_t nextPowerOfTwo(int64_t n) {
        if (LockRingBuffer::isPowerOfTwo(n)) return n;
        n |= n >> 1;
        n |= n >> 2;
        n |= n >> 4;
        n |= n >> 8;
        n |= n >> 16;
        n |= n >> 32;
        return n + 1;
    }

    static inline bool isPowerOfTwo(int64_t n) { return !(n & (n - 1)); }

    // for test
    void printData() {
        std::unique_lock<std::mutex> lock(mux_);
        for (int i = 0; i < len_; i++) {
            if (nullptr != data_[i]) {
                std::cout << i << ":" << *data_[i] << "\t";
            }
        }
        std::cout << "== head:" << head_ << " tail:" << tail_;
        std::cout << std::endl;
    }

   private:
    uint64_t mask_, len_;
    uint64_t head_, tail_;
    std::mutex mux_;
    std::condition_variable empty_, full_;
    T **data_;
};

template <class T, uint64_t size>
LockRingBuffer<T, size>::LockRingBuffer() {
    head_ = tail_ = 0;
    len_ = nextPowerOfTwo(size);
    data_ = new T *[len_];
    mask_ = len_ - 1;
}

template <class T, uint64_t size>
LockRingBuffer<T, size>::~LockRingBuffer() {
    if (nullptr != data_) {
        delete[] data_;
    }
}

template <class T, uint64_t size>
T *LockRingBuffer<T, size>::get() {
    std::unique_lock<std::mutex> lock(mux_);
    auto pred = [this]() -> bool {
        // do not wait anymore if this is true, namely not empty
        return this->tail_ < this->head_;
    };
    empty_.wait(lock, pred);
    T *ret = data_[tail_++ & mask_];
    full_.notify_one();
    return ret;
}

template <class T, uint64_t size>
void LockRingBuffer<T, size>::put(T *ele) {
    std::unique_lock<std::mutex> lock(mux_);
    auto pred = [this]() -> bool {
        return this->head_ < this->tail_ + this->len_;
    };
    full_.wait(lock, pred);
    data_[head_++ & mask_] = ele;
    empty_.notify_one();
}

}  // namespace Tan

#endif  // LOCK_RING_BUFFER_H
