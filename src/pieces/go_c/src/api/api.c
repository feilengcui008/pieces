#include <stdio.h>
#include <stdlib.h>
#include "api.h"

void printHello() {
  fprintf(stdout, "hello\n"); 
}

void printString(const char *s) {
  fprintf(stdout, "%s\n", s);
}

void printCStruct(struct CStruct s) {
  fprintf(stdout, "%d, %f\n", s.a, s.type);
}

void printTestStruct(TestStruct s) {
  fprintf(stdout, "%d, %f\n", s.a, s.type);
}

TestStruct *allocTestStruct() {
  TestStruct *ts = (TestStruct *)malloc(sizeof(TestStruct));
  if (ts) {
    ts->a = 3333;
    ts->type = 333.33;
  }
  return ts;
}

void testVoidPointer(void *p, int len)
{
  char *ptr = (char *)p;
  int i;
  for (i = 0; i < len; ++i) {
    ptr[i] = 'a';
  }
}

void setStruct(void **p) {
  struct CStruct **temp = (struct CStruct **)p;
  (*temp)->a = 123;
  (*temp)->type = 9.9;
}

void printStruct(void *p) {
  struct CStruct *temp = (struct CStruct *)p;
  fprintf(stdout, "%d, %f\n", temp->a, temp->type);
}

void freeStruct(void *p) {
  free((struct CStruct *)p);
}
