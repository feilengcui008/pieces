#!/usr/bin/env python
# -*- encoding:utf-8 -*-

""" a simple worker pool based on Queue and threading """

import threading
import Queue
import time


WORKER_RUNNING = 0
WORKER_STOPPED = 1


class Task(object):
    """ a simple wrapper for task """

    def __init__(self, func, *args, **kargs):
        self.func = func
        self.args = args
        self.kargs = kargs

    def execute(self):
        """ execute task  """
        return self.func(self.args, self.kargs)


class Worker(threading.Thread):
    """ a thread worker """

    def __init__(self, task_queue, name):
        threading.Thread.__init__(self, name=name)
        self.task_queue = task_queue
        self.state = WORKER_STOPPED
        self.started = False
        self.state_lock = threading.Lock()

    def is_running(self):
        """ this worker is running """
        self.state_lock.acquire()
        flag = (self.state == WORKER_RUNNING)
        self.state_lock.release()
        return flag

    def run(self):
        """ start worker thread """
        if self.started:
            return
        self.started = True
        self.state = WORKER_RUNNING
        while self.state == WORKER_RUNNING:
            try:
                task = self.task_queue.get(block=True, timeout=2)
                # no matter task succeeded or failed, we notify queue.join()
                task.task_done()
                task.execute(task.args, task.kargs)
            except Exception:
                pass

    def terminate(self):
        """ terminate this worker """
        # wait until thread started, kind of ugly for now...
        while not self.started:
            time.sleep(2)
        self.state_lock.acquire()
        self.state = WORKER_STOPPED
        self.state_lock.release()


class WorkerPool(object):
    """ a worker pool """
    workers = []

    def __init__(self, worker_num):
        self.worker_num = worker_num
        self.task_queue = Queue.Queue()
        for i in range(worker_num):
            self.workers.append(Worker(self.task_queue, "worker-%d" % (i,)))

    def add_task(self, task):
        """ add by Task object """
        self.task_queue.put(task)

    def add_task_by_func(self, func, args):
        """ add by caller """
        task = Task(func, args)
        self.task_queue.put(task)

    def start(self):
        """ start all workers """
        for worker in self.workers:
            worker.start()

    def terminate(self):
        """ stop worker """
        # until all tasks have been processed
        self.task_queue.join()
        for worker in self.workers:
            worker.terminate()

    def join(self):
        """ join worker """
        for worker in self.workers:
            worker.join()


def test_worker_pool():
    """ test """
    pool = WorkerPool(5)
    pool.start()
    pool.terminate()
    pool.join()
