#include "worker_pool.h"
#include "macros.h"

BEGIN_EXTERN_C()

// worker thread routine
static void *routine(void *args) {
    pool *p = (pool *)args;
    task *t;
    //fprintf(stdout, "thread_id: %ld\n", syscall(SYS_gettid));
    pthread_barrier_wait(&p->countdown_latch);
    while (p->running) {
        pthread_mutex_lock(&p->queue_mutex);
        while (p->task_queue_size == 0) {
            // unlock when enter wait, lock when wakeup
            pthread_cond_wait(&p->queue_cond, &p->queue_mutex);
        }
        // wake up to do a task
        t = p->task_queue_head;
        p->task_queue_head = p->task_queue_head->next;
        p->task_queue_size--;
        // unlock after wakeup lock
        pthread_mutex_unlock(&p->queue_mutex);
        t->cb(t->args);
        // sleep for a while
        sleep(p->interval);
    }
}

// init pool and start threads
void pool_init(pool *p, int thread_number, int max_queue_size) {
    p->thread_number = thread_number;
    p->max_queue_size = max_queue_size;
    p->task_queue_size = 0;
    p->task_queue_head = NULL;
    p->interval = 1;
    p->pt = (pthread_t *)malloc(sizeof(pthread_t) * thread_number);
    if (!p->pt) {
        perror("malloc pthread_t array failed");
        exit(EXIT_FAILURE);
    }
    pthread_mutex_init(&p->queue_mutex, NULL);
    pthread_cond_init(&p->queue_cond, NULL);
    pthread_mutex_init(&p->running_mutex, NULL);
    pthread_barrier_init(&p->countdown_latch, NULL, thread_number + 1);
    for (int i = 0; i < p->thread_number; i++) {
        pthread_create(&(p->pt[i]), NULL, routine, (void *)p);
    }
    // wait when all worker threads are started
    pthread_barrier_wait(&p->countdown_latch);
    p->running = 1;
}

void pool_stop(pool *p) {
    pthread_mutex_lock(&p->running_mutex);
    p->running = 0;
    pthread_mutex_unlock(&p->running_mutex);
}

void pool_destroy(pool *p) {
    if (!p->running) return;
    p->running = 0;
    // tell all threads we are exiting
    // pthread_cond_broadcast(&p->queue_cond);
    // wait and join all threads
    for (int i = 0; i < p->thread_number; ++i) {
        pthread_join(p->pt[i], NULL);
    }
    free(p->pt);
    // free tasks
    task *temp;
    while ((temp = p->task_queue_head) != NULL) {
        p->task_queue_head = p->task_queue_head->next;
        free(temp);
    }
    pthread_mutex_destroy(&p->queue_mutex);
    pthread_cond_destroy(&p->queue_cond);
    pthread_mutex_destroy(&p->running_mutex);
    pthread_barrier_destroy(&p->countdown_latch);
    free(p);
    p = NULL;
}

static int _pool_add_task(pool *p, task *t) {
    pthread_mutex_lock(&p->queue_mutex);
    if (p->task_queue_size >= p->max_queue_size) {
        pthread_mutex_unlock(&p->queue_mutex);
        fprintf(stderr, "exceeded task queue size\n");
        return 0;
    }
    task *temp = p->task_queue_head;
    if (temp != NULL) {
        while (temp->next != NULL) {
            temp = temp->next;
        }
        temp->next = t;
    } else {
        p->task_queue_head = t;
    }
    p->task_queue_size++;
    pthread_cond_signal(&p->queue_cond);
    pthread_mutex_unlock(&p->queue_mutex);
    return 0;
}

int pool_add_task(pool *p, callback cb, void *data) {
    int ret = 0;
    task *t = (task *)malloc(sizeof(task));
    t->cb = cb;
    t->args = data;
    t->next = NULL;
    if ((ret = _pool_add_task(p, t)) != 0) {
        perror("add wroker failed, reaching max size of task queue\n");
    }
    return ret;
}

END_EXTERN_C()
