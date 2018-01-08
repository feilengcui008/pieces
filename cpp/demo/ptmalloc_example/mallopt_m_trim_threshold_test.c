#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <unistd.h>


const int default_pad_size = 128 * 1024;
const int default_trim_threshold = 128 * 1024;

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
            address_after - address_before - default_pad_size);
    free(p);
    char *address_after_free = (char *)sbrk(0);
    fprintf(stdout, "the break address after free : %p\n", address_after_free);
    fprintf(stdout, "======= malloc stats ======\n");
    malloc_stats();
}

/*
 * M_TRIM_THRESHOLD : 当break顶端的空闲内存超过此选项的阈值时
 * ptmalloc会将内存还给kernel
 */
void test_m_trim_threshold() {
    fprintf(stdout, "---------------------------\n");
    // less than 128 * 1024
    int m_trim_threshold = 100 * 1024;
    fprintf(stdout, "default M_TRIM_THRESHOLD : %d, changed to %d\n", default_trim_threshold, m_trim_threshold);
    mallopt(M_TRIM_THRESHOLD, m_trim_threshold);
    int malloc_size = 4097;
    alloc_size_stat(malloc_size);
}

int main(int argc, char **argv) {
    test_m_trim_threshold();
    return 0;
}
