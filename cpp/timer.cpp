#include "timer.h"
#include <time.h>
#include <iostream>
#include <string>

namespace Tan {
Timer::Timer() {
    time(&start_);
    char *str = ctime(&start_);
    std::cout << "start time:" << start_ << " ---- " << std::string(str)
              << std::endl;
}
Timer::~Timer() {
    time(&stop_);
    char *str = ctime(&stop_);
    std::cout << "stop time:" << stop_ << " ---- " << std::string(str)
              << std::endl;
    delta_ = stop_ - start_;
    std::cout << "consume time:" << delta_ << std::endl;
}
}  // end namespace Tan
