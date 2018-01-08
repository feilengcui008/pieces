#include <malloc.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>


/*
 * 每次malloc用brk系统调用移动堆的break向kernel请求内存时，
 * 需要移动的大小包括三部分，然后这三部分总大小与页边界对齐
 * 为break最终需要移动的位置:
 * size = data_size + top_padding + minsize
 * data_size : 实际用户请求的大小
 * top_padding : 即mallopt参数M_TOP_PAD，默认为128 * 1024 
 * minsize : 每个内部内存chunck最小占用内存大小
 *
 * 如果请求的大小在fastbin,smallbin等之中，不需要调用brk从
 * kernel获取内存，所以不会移动break
 *
 */

void brk_test()
{
    char *origin_brk = (char *)sbrk(0);
    fprintf(stdout, "original brk address : %p\n", origin_brk);
    fprintf(stdout, "--------\n");
    int size = 1024 * 100;
    char *p1 = (char *)malloc(size);
    //char *p1 = (char *)malloc(1024 * 4+ 1);
    char *b1 = (char *)sbrk(0);
    fprintf(stdout, "p1 : %p\n", p1);
    fprintf(stdout, "brk1 : %p\n", b1);
    fprintf(stdout, "brk1 - p1 : %ld\n", b1 - p1);
    fprintf(stdout, "brk1 - p1 - size(%d) : %ld\n", size, b1 - p1 - size);
    fprintf(stdout, "because of page alignment : brk1 - origin_brk - 131072 - size(%d): %ld\n",\
            size, b1 - origin_brk - 131072 - size);

    fprintf(stdout, "--------\n");
    // 4096 - 16(第一个chunck的minsize) - 16(第二个chunck的minsize) - (24) (这个24怎么回事?)
    int size2 = 128 * 1024 + 4096 - 56;
    char *p2 = (char *)malloc(size2);
    char *b2 = (char *)sbrk(0);
    fprintf(stdout, "p2 : %p\n", p2);
    fprintf(stdout, "brk2 : %p\n", b2);
    fprintf(stdout, "brk2 - p2 : %ld\n", b2 - p2);
    fprintf(stdout, "brk2 - (p2 + size2(%d)) : %ld\n", size2, b2 - (p2 + size2));

    fprintf(stdout, "--------\n");
    int size3 = 1;
    char *p3 = (char *)malloc(1);
    char *b3 = (char *)sbrk(0);
    fprintf(stdout, "p3 : %p\n", p3);
    fprintf(stdout, "brk3 : %p\n", b3);
    fprintf(stdout, "brk3 - brk2 - size3(%d) - 131072 : %ld\n", size3, b3 - b2 - size3 - 131072);
    fprintf(stdout, "%ld\n", b3 - (p3 + size3));

    free(p1);
    free(p2);
    free(p3);
}

int main(int argc, char **argv) {
    brk_test();  
    return 0;
}
