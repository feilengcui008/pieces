#include <iostream>
#include <memory>

// class template
template <class T, class P>
class TestClass {
 public:
  TestClass() { std::cout << "in class TestClass<T, P>" << std::endl; }
  ~TestClass() = default;
};
// class template partial specilization
template <class T, class P>
class TestClass<T, P *> {
 public:
  TestClass() { std::cout << "in class TestClass<T, P*>" << std::endl; }
  ~TestClass() = default;
};
// class template partial specialization
template <class T>
class TestClass<T, int *> {
 public:
  TestClass() { std::cout << "in class TestClass<T, int *>" << std::endl; }
  ~TestClass() = default;
};
// class tempalte specialization
template <>
class TestClass<bool, int> {
 public:
  TestClass() { std::cout << "in class TestClass<bool, int>" << std::endl; }
  ~TestClass() = default;
};

void class_template_test() {
  std::shared_ptr<TestClass<bool, int>> p1 =
      std::make_shared<TestClass<bool, int>>(TestClass<bool, int>());
  std::shared_ptr<TestClass<float, int *>> p2 =
      std::make_shared<TestClass<float, int *>>(TestClass<float, int *>());
  std::shared_ptr<TestClass<float, double *>> p3 =
      std::make_shared<TestClass<float, double *>>(
          TestClass<float, double *>());
  std::shared_ptr<TestClass<float, double>> p4 =
      std::make_shared<TestClass<float, double>>(TestClass<float, double>());
}

// function template
template <class T, class P>
void testFunc(T t, P p) {
  std::cout << "in testFunc<T, P>" << std::endl;
}
// function template specialization
template <>
void testFunc<bool, int>(bool a, int b) {
  std::cout << "in testFunc<bool, int>" << std::endl;
}
// function tempalte partial specilization is not allowed
// template <class T>
// void testFunc<T, bool>(T t, bool p) {}

// function template overload
template <class T, class P>
void testFunc(T t, P *p) {
  std::cout << "in testFunc<T, P*>" << std::endl;
}
// function template overload
template <class T, class P>
void testFunc(T t, int *a) {
  std::cout << "in testFunc<T, int*>" << std::endl;
}

void function_template_test() {
  int a = 1;
  double b = 2;
  testFunc(true, 1);
  testFunc<float, double>(1, &a);
  testFunc<float, double>(1, &b);
  testFunc<float, double>(1, 2);
}

//////////////////////////////// main //////////////////////////

int main(int argc, char *argv[]) {
  class_template_test();
  function_template_test();
  return 0;
}
