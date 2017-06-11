#include <gtest/gtest.h>
#include <iostream>
#include "signature.h"

TEST(signature, uuid) { std::cout << Tan::uuid() << std::endl; }

int main(int argc, char *argv[]) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
