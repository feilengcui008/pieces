#ifndef TAN_LOCK_LIST_QUEUE_H_
#define TAN_LOCK_LIST_QUEUE_H_

#include <memory>
#include <cstdint>
#include <mutex>

namespace Tan {

template <class T>
class Node {
 public:
  Node() : Node(nullptr) {}
  Node(T *val) : next_(nullptr) { data_ = std::shared_ptr<T>(val); }
  std::shared_ptr<T> get() { return data_; }
  virtual ~Node() {}

 public:
  Node<T> *next_;
  std::shared_ptr<T> data_;
};

template <class T>
class LockListQueue {
 public:
  LockListQueue() {}
  ~LockListQueue() {}
  inline uint64_t getLen() const {
    std::unique_lock<std::mutex> lock(mux_);
    return len_;
  }

  inline Node<T> *get() {
    std::unique_lock<std::mutex> lock(mux_);
    if (len_ == 0) return nullptr;
    auto ret = head_;
    head_ = head_->next_;
    len_--;
    return ret;
  }

  inline void put(Node<T> *ele) {
    std::unique_lock<std::mutex> lock(mux_);
    if (tail_ != nullptr) {
      tail_->next_ = ele;
      tail_ = tail_->next_;
    }
    len_++;
  }

 private:
  uint64_t len_;
  Node<T> *head_, *tail_;
  std::mutex mux_;
};

}  // namespace Tan

#endif  // end TAN_LOCK_LRU_H_
