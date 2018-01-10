#include <iostream>
#include <memory>

using namespace std;

class Test {
   public:
    Test() : a_(0) {}
    Test(int a) : a_(a) {}
    ~Test() { cout << "in dtor" << endl; }
    void func() {
        shared_ptr<Test> p(this);
        cout << "in func " << p.use_count() << endl;
    }

   private:
    int a_;
};

class TestClass : public enable_shared_from_this<TestClass> {
   public:
    TestClass() : a_(0) {}
    TestClass(int a) : a_(a) {}
    ~TestClass() { cout << "in TestClass dtor" << endl; }
    void func() {
        shared_ptr<TestClass> p = shared_from_this();
        cout << "in func " << p.use_count() << endl;
    }

   private:
    int a_;
};

class TestObject {
   public:
    TestObject() = default;
    ~TestObject() { cout << "in TestObject dtor" << endl; }
};

int main(int argc, char *argv[]) {
    {
        cout << "===== shared_from_this =====" << endl;
        shared_ptr<TestClass> p(new TestClass);
        cout << p.use_count() << endl;
        p->func();
        cout << p.use_count() << endl;
    }

    {
        // cout << "===== from raw pointer ======" << endl;
        // TestObject *optr = new TestObject();
        // shared_ptr<TestObject> o1(optr);
        // cout << o1.use_count() << endl;
        // shared_ptr<TestObject> o2(optr);
        // cout << o2.use_count() << endl;
    }

    {
        cout << "===== no shared_from_this =====" << endl;
        shared_ptr<Test> p1(new Test);
        cout << p1.use_count() << endl;
        p1->func();
        cout << p1.use_count() << endl;
    }
    return 0;
}
