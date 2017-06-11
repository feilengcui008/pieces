// use AddressSanitizer to check memory addressing related errors during runtime
// clang++-3.5 -fsanitize=address xx.cpp && ./a.out
// or g++-4.8 -fsanitize=address xx.cpp && ./a.out  (some features like leak
// sanitizer not reported by g++)

int global_arr[9];

void use_after_free() {
    int *n = new int(1);
    delete n;
    *n = 1;
}

void global_out_of_bound() { global_arr[10] = 1; }

void heap_out_of_bound() {
    int *h = new int[9];
    h[10] = 11;
    delete[] h;
}

void stack_out_of_bound() {
    int arr[4] = {1, 2, 3, 4};
    arr[5] = 12;
}

void memory_leak() { int *n = new int[10]; }

void double_free() {
    int *n = new int(1);
    delete n;
    delete n;
}

void invalid_free() {
    int *n = new int[9];
    delete n;
}

int main(int argc, char *argv[]) {
    // use_after_free();
    // global_out_of_bound();
    // heap_out_of_bound();
    // stack_out_of_bound();
    // memory_leak();
    // double_free();
    invalid_free();
    return 0;
}
