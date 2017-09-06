/* 
 * demo for using shmget + shmctl + shmat + shmdt
 *
 */

#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define SHM_KEY 20170802
#define SHM_SIZE 4096

void usage() {
    fprintf(stdout, "./bin create|read|write\n");
}

int create() {
    // create, with RW permission
    int shmid = shmget(SHM_KEY, SHM_SIZE, 0666|IPC_CREAT);
    if (shmid < 0) {
        fprintf(stderr, "%s\n", "create shmget failed");
        return shmid;
    }
    // get info
    struct shmid_ds shmds;
    int ret = shmctl(shmid, IPC_STAT, &shmds);
    if (ret < 0) {
        return ret;
    }
    sleep(20);
    // marked the shared memory to be destroyed
    // actually destroyed when last process detached
    return shmctl(shmid, IPC_RMID, 0);
}

int read() {
    // get id
    int shmid = shmget(SHM_KEY, 0, 0);
    if (shmid < 0) {
        fprintf(stderr, "%s\n", "read shmget failed");
        return shmid;
    }
    // attach
    void *mem = shmat(shmid, (const void *)0, 0);
    if (NULL == mem) {
        fprintf(stderr, "%s\n", "read shmat failed");
        return -1;
    }
    // read
    char buf[100];
    strncpy(buf, (const char *)mem, sizeof(buf));
    fprintf(stdout, "read value %s\n", buf);

    int ret = shmdt((const void *)mem);
    if (ret < 0) {
        fprintf(stderr, "%s\n", "read shmdt failed");
    }
    
    return 0;
}

int write() {
    // get id
    int shmid = shmget(SHM_KEY, 0, 0);
    if (shmid < 0) {
        fprintf(stderr, "%s\n", "read shmget failed");
        return shmid;
    }
    // attach
    void *mem = shmat(shmid, (const void *)0, 0);
    if (NULL == mem) {
        fprintf(stderr, "%s\n", "read shmat failed");
        return -1;
    }
    // write
    char buf[100] = "hello, world!";
    strncpy((char *)mem, buf, sizeof(buf));
    fprintf(stdout, "write value %s\n", buf);

    int ret = shmdt((const void *)mem);
    if (ret < 0) {
        fprintf(stderr, "%s\n", "write shmdt failed");
    }

    return 0;
}

int main(int argc, char *argv[])
{
    if (argc < 2) {
        usage();
        return -1;
    }
    const char *cmd = argv[1];
    if (!strcmp(cmd, "create")) {
        create();
    } else if (!strcmp(cmd, "read")) {
        read();
    } else if (!strcmp(cmd, "write")) {
        write();
    } else {
        usage();
        return -1;
    }
    return 0;
}
