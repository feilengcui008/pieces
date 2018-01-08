#!/usr/bin/python 

import random
import time 

def genDict(filename = "dic.txt", n = 5000):
    random.seed(time.time())
    with open(filename, "w") as fp:
        for i in range(n):
            line = "key" + str(i + 1) + " " + str(random.random()) + "\n"
            fp.write(line)

if __name__ == '__main__':
    genDict()
