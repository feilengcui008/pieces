#include <pthread.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <malloc.h>

int main(int argc, char **argv) {
    fprintf(stdout, "=========\nmultiple times malloc size smaller than 128K, " \
            "should alloc with moving up brk, and if the exceed 128K, then double\n");
    //int size = 130000; // 小于128K，使用brk向高地址移动，分配堆空间
    int size = 100 * 1024; // 小于128K，使用brk向高地址移动，分配堆空间
    int size1 = 2000;
    char *init_break = (char *)sbrk(0);
    fprintf(stdout, "init break address :%p\n", (char *)sbrk(0));

    fprintf(stdout, "------\n");
    char *p = (char *)malloc(size);
    fprintf(stdout, "malloc size :%d, address :%p\n", size, p);
    char *b = (char *)sbrk(0);
    fprintf(stdout, "break address :%p\n", b);
    fprintf(stdout, "delta : %ld\n", b - p);
    fprintf(stdout, "delta between pre break: %ld\n", b - init_break);

    fprintf(stdout, "------\n");
    char *p1 = (char *)malloc(size1);
    fprintf(stdout, "malloc size :%d, address :%p\n", size1, p1);
    char *b1 = (char *)sbrk(0);
    fprintf(stdout, "break address :%p\n", b1);
    fprintf(stdout, "delta : %ld\n", (char *)sbrk(0) - p1);
    fprintf(stdout, "delta between pre break: %ld\n", b1 - b);

    fprintf(stdout, "------\n");
    char *p2 = (char *)malloc(size);
    fprintf(stdout, "malloc size :%d, address :%p\n", size, p2);
    char *b2 = (char *)sbrk(0);
    fprintf(stdout, "break address :%p\n", b2);
    fprintf(stdout, "delta : %ld\n", (char *)sbrk(0) - p2);
    fprintf(stdout, "delta between pre break: %ld\n", b2 - b1);

    fprintf(stdout, "------\n");
    char *p3 = (char *)malloc(size);
    fprintf(stdout, "malloc size :%d, address :%p\n", size, p3);
    char *b3 = (char *)sbrk(0);
    fprintf(stdout, "break address :%p\n", b3);
    fprintf(stdout, "delta : %ld\n", (char *)sbrk(0) - p3);
    fprintf(stdout, "delta between pre break: %ld\n", b3 - b2);

    fprintf(stdout, "total size for brk: %d\n", size * 3 + size1);

    fprintf(stdout, "=========\nmalloc a block size bigger than 128K, " \
            "should use mmap directly, but if the brk has enough memory to " \
            "grow up, then it will still use the brk\n");
    // 128K，超过128K会使用mmap，但是如果之前brk按2的倍数扩张后，如果剩余空间足够大，则任然使用brk
    int size2 = 131072; // total malloc size < 128K * 4
    // total malloc size > 128K * 4, so brk not enough and use mmap directly 
    //int size2 = 131072 + 2000; 
    char *p4 = (char *)malloc(size2);
    fprintf(stdout, "malloc size :%d, address :%p\n", size2, p4);

    fprintf(stderr, "=========\ntest for process(main thread) stack memory alloc(kind of weird...)\n");
    //int size_arr1 = 1000000; // 大于139K，主线程栈往低地址扩张
    int size_arr1 = 10;
    char arr1[size_arr1];
    arr1[0] = 'a';
    fprintf(stdout, "temp var arr1, size : %d, address in main thread stack: %p\n", size_arr1, arr1);

    //int size_arr2 = 1000000;
    int size_arr2 = 1000;
    char arr2[size_arr2];
    arr2[0] = 'a';
    fprintf(stdout, "temp var arr2, size : %d, address in main thread stack: %p\n", size_arr2, arr2);
    fprintf(stdout, "delta : %ld\n", arr2 - arr1);

    int size_arr3 = 1000000;
    char arr3[size_arr3];
    fprintf(stdout, "temp var arr3 address in main thread lower than 139 K : %p\n", arr3);
    fprintf(stdout, "delta : %ld\n", arr3 - arr2);
    for (;;) {
        sleep(5);
    }
    //free(p);
    return 0;
}
