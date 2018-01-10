#ifndef CONCURRENT_SINGLETON_H
#define CONCURRENT_SINGLETON_H

#include <memory>
#include <mutex>

template <class T>
class ConcurrentSingleton {
   public:
    ConcurrentSingleton() = delete;
    ~ConcurrentSingleton() {}

    ConcurrentSingleton &operator=(ConcurrentSingleton &r) = delete;
    ConcurrentSingleton &operator=(ConcurrentSingleton &&r) = delete;

    ConcurrentSingleton(ConcurrentSingleton &t) = delete;
    ConcurrentSingleton(ConcurrentSingleton &&t) = delete;

    static T &instance() {
        if (nullptr == ConcurrentSingleton::data_) {
            std::unique_lock<std::mutex> lock(ConcurrentSingleton::mux_);
            if (nullptr == ConcurrentSingleton::data_) {
                ConcurrentSingleton::data_ = std::make_shared<T>(T());
            }
        }
        return *ConcurrentSingleton::data_.get();
    }

   private:
    static std::mutex mux_;
    static std::shared_ptr<T> data_;
};

template <class T>
std::mutex ConcurrentSingleton<T>::mux_;

template <class T>
std::shared_ptr<T> ConcurrentSingleton<T>::data_;

#endif  // end CONCURRENT_SINGLETON_H
