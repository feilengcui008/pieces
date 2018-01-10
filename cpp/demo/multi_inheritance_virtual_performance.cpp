#include <stdio.h>
#include <sys/time.h>
#include <time.h>

/*
 * 简单测试：
 * 1. 多态访问成员函数和成员变量
 * 2. 继承顺序造成的影响
 *
 */

class BaseFirst {
   public:
    virtual void Test() {}
    int f_;
};

class BaseSecond {
   public:
    virtual void Test() {}
    int s_;
};

class Derived : public BaseFirst, public BaseSecond {
   public:
    void Test() override {}
};

int main(int argc, char *argv[]) {
    Derived *d = new Derived();
    const long N = 1000000000;
    BaseFirst *f = d;
    BaseSecond *s = d;
    fprintf(stdout, "%p %p %p\n", d, f, s);

    struct timeval begin;
    ::gettimeofday(&begin, NULL);
    for (int i = 0; i < N; ++i) {
        // BaseFirst *f = d;
        f->Test();
    }
    struct timeval end;
    ::gettimeofday(&end, NULL);
    fprintf(stdout, "First: %ld\n", end.tv_sec * 1000 + end.tv_usec / 1000 -
                                        begin.tv_sec * 1000 -
                                        begin.tv_usec / 1000);

    ::gettimeofday(&begin, NULL);
    for (int i = 0; i < N; ++i) {
        // BaseSecond *s = d;
        s->Test();
    }
    ::gettimeofday(&end, NULL);
    fprintf(stdout, "Second: %ld\n", end.tv_sec * 1000 + end.tv_usec / 1000 -
                                         begin.tv_sec * 1000 -
                                         begin.tv_usec / 1000);

    ::gettimeofday(&begin, NULL);
    for (int i = 0; i < N; ++i) {
        // BaseFirst *f = d;
        f->f_;
    }
    ::gettimeofday(&end, NULL);
    fprintf(stdout, "First data member: %ld\n",
            end.tv_sec * 1000 + end.tv_usec / 1000 - begin.tv_sec * 1000 -
                begin.tv_usec / 1000);

    ::gettimeofday(&begin, NULL);
    for (int i = 0; i < N; ++i) {
        // BaseSecond *s = d;
        s->s_;
    }
    ::gettimeofday(&end, NULL);
    fprintf(stdout, "Second data member: %ld\n",
            end.tv_sec * 1000 + end.tv_usec / 1000 - begin.tv_sec * 1000 -
                begin.tv_usec / 1000);
    return 0;
}
