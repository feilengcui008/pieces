#ifndef _TAN_WORKER_POOL_H_
#define _TAN_WORKER_POOL_H_

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <unistd.h>

/*
 * a simple task worker pool using thread
 * maybe some optimizations:
 * 1. add queue for each worker thread
 * 2. add running signal pipe for main thread,
 * so it can be used in multi-process programming
 * 3. add pre-allocated task object pool
 * 4. add thread struct for storing thread specific
 * data like thread id
 * 5. etc
 *
 */

/* task related */
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

#endif  // end _TAN_WORKER_POOL_H_
