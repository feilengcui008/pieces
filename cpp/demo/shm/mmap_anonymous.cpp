#include <sys/mman.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>

#define SHM_SIZE 1024

int main(int argc, char *argv[]) {
    
    void *mem = mmap(NULL, SHM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED | MAP_ANONYMOUS, 0, 0);
    if (MAP_FAILED == mem) {
        fprintf(stderr, "mmap failed\n");
        return -1;
    }

    int ret = fork();
    if (ret < 0) {
        fprintf(stderr, "fork failed\n");
        return ret;
    } else if (ret > 0) {
        // write msg in parent process
        char buf[100] = "hello, world!";
        strncpy((char *)mem, (const char *)buf, sizeof(buf));
        fprintf(stdout, "in parent write value %s\n", buf);
        munmap(mem, SHM_SIZE);
        sleep(2); 
    } else {
        // read msg in child process
        sleep(1);
        char buf[100];
        strncpy(buf, (const char *)mem, sizeof(buf));
        fprintf(stdout, "in child read value %s\n", buf);
        munmap(mem, SHM_SIZE);
    }
    return 0;
}
