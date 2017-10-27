#define _GNU_SOURCE
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <sys/utsname.h>
#include <unistd.h>
#include <sys/types.h>
#include <sched.h>
#include <fcntl.h>
#include <wait.h>

// mainly setns and unshare system calls

/* 
int setns(int fd, int nstype);

// 不同版本内核/proc/pid/ns下namespace文件情况
CLONE_NEWCGROUP (since Linux 4.6)
fd must refer to a cgroup namespace.

CLONE_NEWIPC (since Linux 3.0)
fd must refer to an IPC namespace.

CLONE_NEWNET (since Linux 3.0)
fd must refer to a network namespace.

CLONE_NEWNS (since Linux 3.8)
fd must refer to a mount namespace.

CLONE_NEWPID (since Linux 3.8)
fd must refer to a descendant PID namespace.

CLONE_NEWUSER (since Linux 3.8)
fd must refer to a user namespace.

CLONE_NEWUTS (since Linux 3.0)
fd must refer to a UTS namespace.

// 特殊的pid namespace 
CLONE_NEWPID behaves somewhat differently from the other nstype
values: reassociating the calling thread with a PID namespace changes
only the PID namespace that child processes of the caller will be
created in; it does not change the PID namespace of the caller
itself.  Reassociating with a PID namespace is allowed only if the
PID namespace specified by fd is a descendant (child, grandchild,
etc.)  of the PID namespace of the caller.  For further details on
PID namespaces, see pid_namespaces(7).

int unshare(int flags);
CLONE_FILES | CLONE_FS | CLONE_NEWCGROUP | CLONE_NEWIPC | CLONE_NEWNET 
| CLONE_NEWNS | CLONE_NEWPID | CLONE_NEWUSER | CLONE_NEWUTS | CLONE_SYSVSEM
*/



#define MAX_PROCPATH_LEN 1024

#define err_exit(msg) \
    do { fprintf(stderr, "%s in file %s in line %d\n", msg, __FILE__, __LINE__);\
        exit(EXIT_FAILURE); } while (0)

void printInfo();
int openAndSetns(const char *path);

int main(int argc, char *argv[]) {
    if (argc < 2) {
        fprintf(stdout, "usage : execname pid(find namespaces of this process)\n");
        return 0;
    }
    printInfo();

    fprintf(stdout, "---- setns for uts ----\n");
    char uts[MAX_PROCPATH_LEN];
    snprintf(uts, MAX_PROCPATH_LEN, "/proc/%s/ns/uts", argv[1]);
    openAndSetns(uts);
    printInfo();

    fprintf(stdout, "---- setns for user ----\n");
    char user[MAX_PROCPATH_LEN];
    snprintf(user, MAX_PROCPATH_LEN, "/proc/%s/ns/user", argv[1]);
    openAndSetns(user);
    printInfo();

    // 注意pid namespace的不同行为，只有后续创建的子进程进入setns设置
    // 的新的pid namespace，本进程不会改变
    fprintf(stdout, "---- setns for pid ----\n");
    char pidpath[MAX_PROCPATH_LEN];
    snprintf(pidpath, MAX_PROCPATH_LEN, "/proc/%s/ns/pid", argv[1]);
    openAndSetns(pidpath);
    printInfo();

    fprintf(stdout, "---- setns for ipc ----\n");
    char ipc[MAX_PROCPATH_LEN];
    snprintf(ipc, MAX_PROCPATH_LEN, "/proc/%s/ns/ipc", argv[1]);
    openAndSetns(ipc);
    printInfo();

    fprintf(stdout, "---- setns for net ----\n");
    char net[MAX_PROCPATH_LEN];
    snprintf(net, MAX_PROCPATH_LEN, "/proc/%s/ns/net", argv[1]);
    openAndSetns(net);
    printInfo();

    // 注意mnt namespace需要放在其他后面，避免mnt namespace改变后
    // 找不到/proc/pid/ns下的文件
    fprintf(stdout, "---- setns for mount ----\n");
    char mount[MAX_PROCPATH_LEN];
    snprintf(mount, MAX_PROCPATH_LEN, "/proc/%s/ns/mnt", argv[1]);
    openAndSetns(mount);
    printInfo();

    // 测试子进程的pid namespace
    int ret = fork();
    if (-1 == ret) {
        err_exit("failed to fork");
    } else if (ret == 0) {
        fprintf(stdout, "********\n");
        fprintf(stdout, "in child process\n");
        printInfo();
        fprintf(stdout, "********\n");
        for (;;) {
            sleep(5);
        }
    } else {
        fprintf(stdout, "child pid : %d\n", ret);
    }
    for (;;) {
        sleep(5);
    }
    waitpid(ret, NULL, 0);
    return 0;
}

void printInfo() {
    pid_t pid;
    struct utsname uts;
    uid_t uid;
    gid_t gid;
    // pid namespace 
    pid = getpid();
    // user namespace 
    uid = getuid();
    gid = getgid();
    // uts namespace 
    uname(&uts);
    fprintf(stdout, "pid : %d\n", pid);
    fprintf(stdout, "uid : %d\n", uid);
    fprintf(stdout, "gid : %d\n", gid);
    fprintf(stdout, "hostname : %s\n", uts.nodename);
}

int openAndSetns(const char *path) {
    int ret = open(path, O_RDONLY, 0);
    if (-1 == ret) {
        fprintf(stderr, "%s\n", strerror(errno));
        err_exit("failed to open fd");
    }
    if (-1 == (ret = setns(ret, 0))) {
        fprintf(stderr, "%s\n", strerror(errno));
        err_exit("failed to setns");
    }
    return ret;
}
