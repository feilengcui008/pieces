#include <pthread.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void *routine(void *args) {
    int size1 = 10;
    //int size2 = 10000;
    int size2 = 1000;
    int size3 = 10;
    fprintf(stdout, "========\n");
    char arr[size1];
    memset(arr, 0, sizeof(arr)); 
    char arr1[size2];
    memset(arr1, 0, sizeof(arr1)); 
    char arr2[size3];
    memset(arr2, 0, sizeof(arr2)); 
    fprintf(stdout, "temp var arr size : %d,  address in child thread : %p\n", size1, arr);
    fprintf(stdout, "temp var arr1 size : %d,  address in child thread : %p\n", size2, arr1);
    fprintf(stdout, "temp var arr2 size : %d,  address in child thread : %p\n", size3, arr2);
    fprintf(stdout, "delta arr1 - arr: %ld\n", arr1 - arr);
    fprintf(stdout, "delta arr2 - arr1: %ld\n", arr2 - arr1);

    for(;;) {
        sleep(5);
    }
}

int main(int argc, char *argv[]) {
    pthread_t pt; // 4
    pthread_t pt1; // 4
    int ret;  // 4
    // pthread max stack size(can be changed): 0x800000 = 8M
    // char bigArr[0x800000 - 10000]; // SEGMENT FAULT
    //char arr1[144000];
    char arr1[144];
    //arr1[0] = 'a';
    memset(arr1, 0, sizeof(arr1)); 
    fprintf(stdout, "temp var arr1 address in main thread lower than 139 K : %p\n", arr1);
    //char arr2[100];
    char arr2[1];
    memset(arr2, 0, sizeof(arr2)); 
    fprintf(stdout, "temp var arr2 address in main thread lower than 139 K : %p\n", arr2);
    fprintf(stdout, "delta : %ld\n", arr2 - arr1);
    //char arr3[100];
    char arr3[10];
    memset(arr3, 0, sizeof(arr3)); 
    fprintf(stdout, "temp var arr3 address in main thread lower than 139 K : %p\n", arr3);
    fprintf(stdout, "delta : %ld\n", arr3 - arr2);
    ret = pthread_create(&pt, NULL, routine, NULL);
    ret = pthread_create(&pt1, NULL, routine, NULL);
    pthread_join(pt, NULL); 
    pthread_join(pt1, NULL); 
    return 0;
}
