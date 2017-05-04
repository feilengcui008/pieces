#include <stdio.h>
#include <dlfcn.h>

typedef int (*printf_type)(const char *msg);
typedef void (*func)();

void test_dlopen_from_c_so() 
{
  const char *path = "plugin/cplugin.so";
  void *handle = dlopen(path, RTLD_LAZY);
  if (handle == NULL) {
    const char *errmsg = dlerror();
    fprintf(stderr, "dlopen failed\n");
    if (errmsg != NULL) {
      fprintf(stderr, "dlopen errmsg is %s\n", errmsg);
    }
    return;
  }
  void *fn = dlsym(handle, "myprintf");
  if (fn == NULL) {
    fprintf(stderr, "myprintf is not in libc\n");
  } else {
    ((printf_type)fn)("called by dl symble\n");
  }
  int ret = dlclose(handle);
  if (ret < 0) {
    fprintf(stderr, "close dlopen handle failed with ret %d\n", ret);
  }
}

void test_dlopen_from_go_so()
{
  const char *path = "plugin/goplugin.so";
  void *handle = dlopen(path, RTLD_LAZY);
  if (handle == NULL) {
    const char *errmsg = dlerror();
    fprintf(stderr, "dlopen failed\n");
    if (errmsg != NULL) {
      fprintf(stderr, "dlopen errmsg is %s\n", errmsg);
    }
    return;
  }
  // not work for c dlopen of go plugin .so file now
  void *fn = dlsym(handle, "ExportedForPlugin");
  //void *fn = dlsym(handle, "runtime.printfloat");
  //void *fn = dlsym(handle, "plugin/unnamed-e1f58c25fd8588f5627fd829684f7dcd66e42b57.ExportedForPlugin");
  if (fn == NULL) {
    fprintf(stderr, "ExportedForPlugin is not in libc\n");
  } else {
    ((func)fn)();
  }
  int ret = dlclose(handle);
  if (ret < 0) {
    fprintf(stderr, "close dlopen handle failed with ret %d\n", ret);
  }
}

int main(int argc, char *argv[])
{
  test_dlopen_from_c_so();
  fprintf(stdout, "====================\n");
  test_dlopen_from_go_so();
  return 0;
}
