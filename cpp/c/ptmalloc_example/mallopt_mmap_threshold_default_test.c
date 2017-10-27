#include <stdlib.h>
#include <stdio.h>
#include <malloc.h>
#include <unistd.h>

void test_mmap_threshold() {
    fprintf(stdout, "---------------------------\n");
    int default_mmap_shreshold = 128 * 1024;
    fprintf(stdout, "default mmap threshold size : %d\n", default_mmap_shreshold);
    int alloc_size = 127 * 1024;
    fprintf(stdout, "=======\n");
    fprintf(stdout, "before change mmap threshold alloc size : %d\n", alloc_size);
    char *p = (char *)malloc(alloc_size);
    fprintf(stdout, "address is : %p\n", p);

    fprintf(stdout, "=======\n");
    malloc_stats();
}

int main(int argc, char *argv[]) {
    test_mmap_threshold();
    return 0;
}
