#include <iostream>
#include <stdio.h>

using namespace std;

typedef struct {
    unsigned char a;
    unsigned char b;
} TestStruct;

int main() {
    TestStruct t;
    //t.a = 1;
    fprintf(stdout, a: %pn, &t.a);
    fprintf(stdout, b: %pn, &t.b);
    t.a = 0;
    t.b = 1;
    std::cout << *((unsigned short *)&t) << std::endl;

    t.b = 0;
    t.a = 1;
    std::cout << *((unsigned short *)&t) << std::endl;
    return 0;
}

