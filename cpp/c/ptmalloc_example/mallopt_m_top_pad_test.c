#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <unistd.h>

const int default_pad_size = 128 * 1024;
const int default_trim_threshold = 128 * 1024;
int pad_size;

void alloc_size_stat(int size) {
    int alloc_size = size;
    fprintf(stdout, "allocated size : %d\n", alloc_size);
    char *address_before = (char *)sbrk(0);
    fprintf(stdout, "the break address before malloc : %p\n", address_before);
    char *p = (char *)malloc(alloc_size);
    char *address_after = (char *)sbrk(0);
    fprintf(stdout, "the break address after malloc : %p\n", address_after);
    fprintf(stdout, "the delta of address is : %ld\n", address_after - address_before);
    fprintf(stdout, "the real padding because of page alignment : %ld\n", 
            address_after - address_before - pad_size);
    free(p);
    char *address_after_free = (char *)sbrk(0);
    fprintf(stdout, "the break address after free : %p\n", address_after_free);
    fprintf(stdout, "======= malloc stats ======\n");
    malloc_stats();
}

/*
 * M_TOP_PAD : 使用brk调整break时需要加上的大小，
 * 并且padding的实际大小要与内存页边界对齐，默认为128 * 1024 
 */
void test_m_top_pad() {
    fprintf(stdout, "---------------------------\n");
    int changed_to = 256 * 1024;
    //int changed_to = 100 * 1024;
    //int changed_to = 128 * 1024 + 1;
    //int changed_to = 128 * 1024 + 100;
    //int changed_to = 128 * 1024 + 4095;
    fprintf(stdout, "after change default M_TOP_PAD : %d to %d\n", default_pad_size, changed_to);
    mallopt(M_TOP_PAD, changed_to);
    //mallopt(M_TRIM_THRESHOLD, 10 * 1024);
    pad_size = changed_to;
    int malloc_size = 4097;
    alloc_size_stat(malloc_size);
}

int main(int argc, char **argv) {
    test_m_top_pad();  
    return 0;
}
