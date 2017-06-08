#include <asm/thread_info.h>
#include <linux/errno.h>
#include <linux/module.h>
#include <linux/sched.h>

static int test_param = 10;
module_param(test_param, int, S_IRUGO | S_IWUSR);
MODULE_PARM_DESC(test_param, "a test parameter");

static int print_all_processes_init(void) {
  struct task_struct *p;
  // printk(KERN_INFO "initializing module print_all_processes\n");
  // printk(KERN_INFO "================print_all_processes========\n");
  // printk(KERN_INFO "pid\tstate\n");
  // printk(KERN_INFO "thread_info size : %d\n", (int)sizeof(struct
  // thread_info));
  // struct mm_struct *mm;
  // struct file *flp;

  for_each_process(p) {
    if (p->pid == 1) {
      printk(KERN_INFO "stack : %p\n", p->stack);
    }
    // mm = get_task_mm(p);
    // flp = mm->exe_file;
    // if(flp)
    //  printk(KERN_INFO "%d\t%d\t%s\n", (int)p->pid, (int)p->state,
    //  flp->f_path.dentry.d_name.name);
  };

  // dump_stack();
  // show_state();
  return 0;
}

static void print_all_processes_exit(void) {
  printk(KERN_INFO "unload module print_all_processes\n");
}

module_init(print_all_processes_init);
module_exit(print_all_processes_exit);

MODULE_AUTHOR("FEILENGCUI");
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("A MODULE PRINT ALL PROCESSES");
