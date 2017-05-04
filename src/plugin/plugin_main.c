#include <stdio.h>
#include <dlfcn.h>

typedef int (*printf_type)(const char *msg);

int main(int argc, char *argv[])
{
  //const char *path = "/lib/x86_64-linux-gnu/libc.so.6";
  const char *path = "plugin/cplugin.so";
  void *handle = dlopen(path, RTLD_LAZY);
  if (handle == NULL) {
    const char *errmsg = dlerror();
    fprintf(stderr, "dlopen failed\n");
    if (errmsg != NULL) {
      fprintf(stderr, "dlopen errmsg is %s\n", errmsg);
    }
    return 0;
  }
  //void *fn = dlsym(handle, "printf");
  void *fn = dlsym(handle, "myprintf");
  if (fn == NULL) {
    fprintf(stderr, "snprintf is not in libc\n");
  } else {
    ((printf_type)fn)("called by dl symble\n");
  }
  int ret = dlclose(handle);
  if (ret < 0) {
    fprintf(stderr, "close dlopen handle failed with ret %d\n", ret);
  }
  return 0;
}
