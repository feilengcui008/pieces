#ifndef _TAN_NET_H_
#define _TAN_NET_H_

#include <fcntl.h>
#include <inttypes.h>
#include <netdb.h>
#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

#include "error.h"
#include "macros.h"

BEGIN_EXTERN_C()

// basic net related operations
int createSocket(int family, int type, int protocol, int nonblocking);
int connectAddress(const char *address, uint16_t port);
int bindAndListenV4(const char *address, uint16_t port, int backlog);
int bindAndListenByAddrinfo(const char *address, const char *port, int backlog);
int setNonblock(int fd);

// epoll related
int createEpoll(int size);
int updateEvent(int epfd, int fd, int events, int op);
int waitPoll(int epfd, struct epoll_event *events, int max_events, int timeout);

END_EXTERN_C()

#endif  // end _TAN_NET_H_