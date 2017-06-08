#ifndef _TAN_TIMER_H_
#define _TAN_TIMER_H_

#include <time.h>
#include <iostream>
#include <string>

namespace Tan {
class Timer {
 public:
  Timer();
  ~Timer();

 private:
  time_t start_;
  time_t stop_;
  time_t delta_;
};
}  // end namespace Tan

#endif  // end _TAN_TIMER_H_
