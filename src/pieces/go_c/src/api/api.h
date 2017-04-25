#ifndef _API_H_
#define _API_H_

struct CStruct {
  int a;
  float type;
};

typedef struct {
  int a;
  float type;
} TestStruct;

void printHello();
void printString(const char *s);
void printCStruct(struct CStruct s);
void printTestStruct(TestStruct s);

TestStruct *allocTestStruct();

void testVoidPointer(void *p, int len);
void setStruct(void **p);
void printStruct(void *p);
void freeStruct(void *p);


#endif
