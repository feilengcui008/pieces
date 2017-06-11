#include "net.h"
#include <sys/epoll.h>

BEGIN_EXTERN_C();

int createSocket(int family, int type, int protocol, int nonblocking) {
    int socketfd = socket(family, type, protocol);
    if (socketfd < 0) {
        error_exit("error create socket");
    }
    if (nonblocking) {
        setNonblock(socketfd);
    }
    return socketfd;
}

int connectAddress(const char *address, uint16_t port) { return -1; }

int bindAndListenV4(const char *address, uint16_t port, int backlog) {
    struct sockaddr_in si;
    si.sin_family = AF_INET;
    si.sin_port = htons(port);
    // si.sin_add
    return -1;
}

int bindAndListenByAddrinfo(const char *address, const char *port,
                            int backlog) {
    // int ret = bind
    struct addrinfo hints;
    struct addrinfo *result, *rp;
    int sfd, s;
    socklen_t socklen;
    int ret;

    memset(&hints, 0, sizeof(struct addrinfo));
    hints.ai_family = AF_UNSPEC;    /* Allow IPv4 or IPv6 */
    hints.ai_socktype = SOCK_DGRAM; /* Datagram socket */
    hints.ai_flags = AI_PASSIVE;    /* For wildcard IP address */
    hints.ai_protocol = 0;          /* Any protocol */
    hints.ai_canonname = NULL;
    hints.ai_addr = NULL;
    hints.ai_next = NULL;

    s = getaddrinfo(address, port, &hints, &result);
    if (s < 0) {
        error_exit("getaddrinfo error");
    }
    /* getaddrinfo() returns a list of address structures.
       Try each address until we successfully bind(2).
       If socket(2) (or bind(2)) fails, we (close the socket
       and) try the next address. */
    for (rp = result; rp != NULL; rp = rp->ai_next) {
        sfd = createSocket(rp->ai_family, rp->ai_socktype, rp->ai_protocol, 0);
        if (sfd < 0) {
            continue;
        }
        if (bind(sfd, rp->ai_addr, rp->ai_addrlen) == 0) break; /* Success */
        close(sfd);
    }
    if (rp == NULL) { /* No address succeeded */
        error_exit("Could not bind\n");
    }
    ret = bind(sfd, rp->ai_addr, socklen);
    if (ret < 0) {
        close(sfd);
        error_exit("error bind");
    }
    freeaddrinfo(result); /* No longer needed */
    ret = listen(sfd, backlog);
    if (ret < 0) {
        close(sfd);
        error_exit("error listen");
    }
    return ret;
}

int setNonblock(int fd) {
    int ret = fcntl(fd, F_GETFL);
    if (ret < 0) {
        error_exit("get fd flag error\n");
    }
    ret = fcntl(fd, F_SETFL, ret | O_NONBLOCK);
    if (ret < 0) {
        error_exit("set fd nonblocking error\n");
    }
    return ret;
}

int createEpoll(int size) {
    int ret = epoll_create(size);
    if (ret < 0) {
        error_exit("create epoll error");
    }
    return ret;
}

int updateEvent(int epfd, int fd, int events, int op) {
    struct epoll_event ee;
    ee.data.fd = fd;
    ee.events = events;
    int ret = epoll_ctl(epfd, op, fd, &ee);
    if (ret < 0) {
        error_exit("error when update events");
    }
    return ret;
}

int waitPoll(int epfd, struct epoll_event *events, int max_events,
             int timeout) {
    int ret = epoll_wait(epfd, events, max_events, timeout);
    if (ret < 0) {
        error_exit("epoll_wait error");
    }
    return ret;
}

END_EXTERN_C()
