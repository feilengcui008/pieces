//#include <asm/thread_info.h>
#include <linux/errno.h>
#include <linux/module.h>
#include <linux/sched.h>

static int test_param = 10;
module_param(test_param, int, S_IRUGO | S_IWUSR);
MODULE_PARM_DESC(test_param, "a test parameter");

static int print_all_processes_init(void) {
    struct task_struct *p;
    struct mm_struct *mm;

    for_each_process(p) {
        if (p->pid == 1) {
            printk(KERN_INFO "pid: %d, state: %d\n", (int)p->pid, (int)p->state);

            mm = get_task_mm(p);
            // print virtual memory address layout
            if (mm != NULL) {
                printk(KERN_INFO 
                        "start_code: %lx, end_code: %lx\n"
                        "start_data: %lx, end_data: %lx\n"
                        "start_brk: %lx, brk: %lx\n"
                        "start_stack: %lx, stack: %p\n"
                        "arg_start: %lx, arg_end: %lx\n"
                        "env_start: %lx, env_end: %lx\n"
                        "mmap_base: %lx\n"
                        "task_size: %lx\n",
                        mm->start_code, mm->end_code,
                        mm->start_data, mm->end_data,
                        mm->start_brk, mm->brk,
                        mm->start_stack, p->stack,
                        mm->arg_start, mm->arg_end,
                        mm->env_start, mm->env_end,
                        mm->mmap_base,
                        mm->task_size);

            }
        }
        // dump_stack();
        // show_state();
        return 0;
    }

}

static void print_all_processes_exit(void) {
    printk(KERN_INFO "unload module print_all_processes\n");
}

module_init(print_all_processes_init);
module_exit(print_all_processes_exit);

MODULE_AUTHOR("FEILENGCUI");
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("A MODULE PRINT ALL PROCESSES");
