/*
 * demo for using shm_open + ftruncate + mmap + munmap + shm_unlink
 */
#include <fcntl.h>
#include <sys/mman.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>


const char *dev_shm_file = "20170802";
#define SHM_SIZE 4096

void usage() {
    fprintf(stdout, "./bin read|write|delete\n");
}

int read() {
    // open file in /dev/shm/...
    int fd = shm_open(dev_shm_file, O_CREAT | O_RDWR, 0666);
    if (fd < 0) {
        fprintf(stderr, "%s\n", "read shm_open failed");
        return fd;
    }
    // ftruncate
    ftruncate(fd, SHM_SIZE);

    // mmap
    void *mem = mmap(NULL, SHM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (MAP_FAILED == mem) {
        fprintf(stderr, "%s\n", "read mmap failed");
        return -1;
    }

    // read shared memory
    char buf[100];
    strncpy(buf, (const char *)mem, sizeof(buf));
    fprintf(stdout, "read value %s\n", buf);

    // unmap
    int ret = munmap((void *)mem, SHM_SIZE);
    if (ret < 0) {
        fprintf(stderr, "munmap failed\n");
        return ret;
    }
    
    return 0;
}

int write() {
    // open file in /dev/shm/...
    int fd = shm_open(dev_shm_file, O_CREAT | O_RDWR, 0666);
    if (fd < 0) {
        fprintf(stderr, "%s\n", "write shm_open failed");
        return fd;
    }
    // ftruncate
    ftruncate(fd, SHM_SIZE);

    // mmap
    void *mem = mmap(NULL, SHM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (MAP_FAILED == mem) {
        fprintf(stderr, "%s\n", "write mmap failed");
        return -1;
    }

    // write shared memory
    char buf[100] = "hello, world!";
    strncpy((char *)mem, (const char *)buf, sizeof(buf));
    fprintf(stdout, "write value %s\n", buf);

    // unmap
    int ret = munmap((void *)mem, SHM_SIZE);
    if (ret < 0) {
        fprintf(stderr, "munmap failed\n");
        return ret;
    }
    
    return 0;
}

int del() {
    int ret = shm_unlink(dev_shm_file);
    if (ret < 0) {
        fprintf(stderr, "%s\n", "shm_unlink failed");
        return -1;
    }
    return ret;
}

int main(int argc, char *argv[])
{
    
    if (argc < 2) {
        usage();
        return -1;
    }
    const char *cmd = argv[1];
    if (!strcmp(cmd, "read")) {
        read();
    } else if (!strcmp(cmd, "write")) {
        write();
    } else if (!strcmp(cmd, "delete")) {
        del();
    } else {
        usage();
        return -1;
    }
    return 0;
}
