#include <functional>
#include <iostream>
#include <memory>

class TestClass {
   public:
    ~TestClass() { std::cout << "dtor" << std::endl; }
    void echo() { std::cout << "echo" << std::endl; }
    void echoInt(int a) { std::cout << "echo " << a << std::endl; }
    static void staticEcho() { std::cout << "static echo" << std::endl; }
};

void globalEcho() { std::cout << "global echo" << std::endl; }

int main(int argc, char *argv[]) {
    std::function<void()> func;
    std::function<void(int)> func_int;
    std::function<void()> func1 = &globalEcho;
    std::function<void()> func2 = &TestClass::staticEcho;
    {
        // bind会增加智能指针的引用计数
        std::shared_ptr<TestClass> ptr =
            std::make_shared<TestClass>(TestClass());
        std::cout << "in inner scope " << ptr.use_count() << std::endl;  // 1
        func = std::bind(&TestClass::echo, ptr);
        std::cout << "in inner scope " << ptr.use_count() << std::endl;  // 2
        func_int = std::bind(&TestClass::echoInt, ptr, std::placeholders::_1);
        std::cout << "in inner scope " << ptr.use_count() << std::endl;  // 3
    }
    // will not segfault
    func();
    func_int(-1);
    func1();
    func2();
    return 0;
}
