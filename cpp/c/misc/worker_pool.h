#ifndef WORKER_POOL_H_
#define WORKER_POOL_H_

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <unistd.h>

#include "macros.h"

BEGIN_EXTERN_C()

// callback type
typedef void *(*callback)(void *args);
// task queue
typedef struct _task task;
struct _task {
    callback cb;
    void *args;
    struct _task *next;
};

/* thread pool */
typedef struct _pool pool;
struct _pool {
    // thread number
    int thread_number;
    // current task size of queue
    int task_queue_size;
    // max queue size allowed
    int max_queue_size;
    // stop flag
    int running;
    // pt array
    pthread_t *pt;
    // task queue
    task *task_queue_head;
    // mutex for task queue
    pthread_mutex_t queue_mutex;
    // cond for producer and consumer
    pthread_cond_t queue_cond;
    // mutex for running flag
    pthread_mutex_t running_mutex;
    // for initilize worker threads
    pthread_barrier_t countdown_latch;
    // worker thread sleep interval
    int interval;
};

// init pool and start threads
void pool_init(pool *p, int thread_number, int max_queue_size);
// stop pool
void pool_stop(pool *p);
// destroy pool resource
void pool_destroy(pool *p);
// add task into pool
int pool_add_task(pool *p, callback cb, void *data);

END_EXTERN_C()

#endif  // end WORKER_POOL_H_
