#include <unistd.h>
#include "timer.h"
#include <gtest/gtest.h>

TEST(Timer, basic_test)
{
  {
    Tan::Timer t;
    ::sleep(2);
  }
}

int main(int argc, char *argv[])
{
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
