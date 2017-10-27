#define _GNU_SOURCE
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/utsname.h>
#include <sched.h>

#define error_exit(msg)  do { perror(msg); exit(EXIT_FAILURE); } while (0)
#define STACK_SIZE (1024 * 1024) // stack size for clone

static int shell_routine(void *arg) {
    const char *binary = "/bin/bash";
    char *const argv[] = {
        "/bin/bash",
        NULL
    };
    char *const envp[] = { NULL };

    /* wrappers for execve */
    // has const char * as argument list
    // execl 
    // execle  => has envp
    // execlp  => need search PATH 

    // has char *const arr[] as argument list 
    // execv 
    // execvpe => need search PATH and has envp
    // execvp  => need search PATH 

    //int ret = execv(binary, argv);
    int ret = execve(binary, argv, envp);
    if (ret < 0) {
        error_exit("execve error");
    }
    return ret;
}

void create_shell_process() {
    void *stack;
    void *stack_buttom;
    pid_t pid;

    stack = malloc(STACK_SIZE);
    if (stack == NULL) {
        error_exit("malloc");
    }
    stack_buttom = (char *)stack + STACK_SIZE;

    pid = clone(shell_routine, stack_buttom, CLONE_NEWUTS | CLONE_NEWNS | \
            CLONE_NEWPID | CLONE_NEWUSER | CLONE_NEWIPC | CLONE_NEWNET | SIGCHLD, NULL);
    if (pid == -1) {
        error_exit("clone");
    }
    printf("clone() returned %ld\n", (long) pid);

    if (waitpid(pid, NULL, 0) == -1) {
        error_exit("waitpid");
    }
    printf("child has terminated\n");

    exit(EXIT_SUCCESS);
}



static int child_routine(void *arg) {
    struct utsname uts;
    if (sethostname((char *)arg, strlen((char *)arg)) == -1) {
        error_exit("sethostname");
    }
    if (uname(&uts) == -1) {
        error_exit("uname");
    }
    printf("uts.nodename in child:  %s\n", uts.nodename);

    sleep(200);
    return 0;
}

int main(int argc, char *argv[]) {
    void *stack;
    void *stack_buttom;
    pid_t pid;
    struct utsname uts;

    if (argc < 2) {
        fprintf(stderr, "Usage: %s <child-hostname>\n", argv[0]);
        exit(EXIT_SUCCESS);
    }

    stack = malloc(STACK_SIZE);
    if (stack == NULL) {
        error_exit("malloc");
    }
    stack_buttom = (char *)stack + STACK_SIZE;

    pid = clone(child_routine, stack_buttom, CLONE_NEWUTS | SIGCHLD, (void *)argv[1]);
    if (pid == -1) {
        error_exit("clone");
    }
    printf("clone() returned %ld\n", (long)pid);
    sleep(1);

    if (uname(&uts) == -1) {
        error_exit("uname");
    }
    printf("uts.nodename in parent: %s\n", uts.nodename);

    if (waitpid(pid, NULL, 0) == -1) {
        error_exit("waitpid");
    }
    printf("child has terminated\n");
    exit(EXIT_SUCCESS);
}
