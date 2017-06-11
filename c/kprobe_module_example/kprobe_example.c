#include <linux/kernel.h>
#include <linux/kprobes.h>
#include <linux/module.h>
#include <linux/sched.h>
//#include <asm/thread_info.h>

/*
 * kprobe 可以使用两种方式提供注入点
 * 1. 某个函数名或者模块名:函数名，register_kprobe会自动解析到地址
 * 2. 地址
 */

/* For each probe you need to allocate a kprobe structure */
static struct kprobe kp = {
    //.symbol_name	= "do_fork",
    //.symbol_name	= "dev_queue_xmit",
    .symbol_name = "ip_local_out",
    //.symbol_name	= "x86_64_start_kernel",
};

/* kprobe pre_handler: called just before the probed instruction is executed */
static int handler_pre(struct kprobe *p, struct pt_regs *regs) {
    /*
    printk(KERN_INFO "-----------------------------------------------------\n");
    dump_stack();
    printk(KERN_INFO "-----------------------------------------------------\n");
  #ifdef CONFIG_X86
          printk(KERN_INFO "pre_handler: p->addr = 0x%p, ip = %lx,"
                          " flags = 0x%lx\n",
                  p->addr, regs->ip, regs->flags);
  #endif
  #ifdef CONFIG_PPC
          printk(KERN_INFO "pre_handler: p->addr = 0x%p, nip = 0x%lx,"
                          " msr = 0x%lx\n",
                  p->addr, regs->nip, regs->msr);
  #endif
  #ifdef CONFIG_MIPS
          printk(KERN_INFO "pre_handler: p->addr = 0x%p, epc = 0x%lx,"
                          " status = 0x%lx\n",
                  p->addr, regs->cp0_epc, regs->cp0_status);
  #endif
  #ifdef CONFIG_TILEGX
          printk(KERN_INFO "pre_handler: p->addr = 0x%p, pc = 0x%lx,"
                          " ex1 = 0x%lx\n",
                  p->addr, regs->pc, regs->ex1);
  #endif

          // A dump_stack() here will give a stack backtrace
    if (current->pid != 0) {
      printk(KERN_INFO "============pid:%d, ppid:%d\n", current->pid,
  current->parent != NULL ? current->parent->pid : -1);
    }
    */
    dump_stack();
    return 0;
}

/* kprobe post_handler: called after the probed instruction is executed */
static void handler_post(struct kprobe *p, struct pt_regs *regs,
                         unsigned long flags) {
    /*
    printk(KERN_INFO
  "-------------------------iiii----------------------------\n");
    dump_stack();
    printk(KERN_INFO
  "-------------------------iiii----------------------------\n");
  #ifdef CONFIG_X86
          printk(KERN_INFO "post_handler: p->addr = 0x%p, flags = 0x%lx\n",
                  p->addr, regs->flags);
  #endif
  #ifdef CONFIG_PPC
          printk(KERN_INFO "post_handler: p->addr = 0x%p, msr = 0x%lx\n",
                  p->addr, regs->msr);
  #endif
  #ifdef CONFIG_MIPS
          printk(KERN_INFO "post_handler: p->addr = 0x%p, status = 0x%lx\n",
                  p->addr, regs->cp0_status);
  #endif
  #ifdef CONFIG_TILEGX
          printk(KERN_INFO "post_handler: p->addr = 0x%p, ex1 = 0x%lx\n",
                  p->addr, regs->ex1);
  #endif
    if (current->pid != 0) {
      printk(KERN_INFO "pid:%d, ppid:%d\n", current->pid, current->parent !=
  NULL
  ? current->parent->pid : -1);
    }
    */
}

/*
 * fault_handler: this is called if an exception is generated for any
 * instruction within the pre- or post-handler, or when Kprobes
 * single-steps the probed instruction.
 */
static int handler_fault(struct kprobe *p, struct pt_regs *regs, int trapnr) {
    printk(KERN_INFO "fault_handler: p->addr = 0x%p, trap #%dn", p->addr,
           trapnr);
    /* Return 0 because we don't handle the fault. */
    return 0;
}

static int __init kprobe_init(void) {
    int ret;
    kp.pre_handler = handler_pre;
    kp.post_handler = handler_post;
    kp.fault_handler = handler_fault;

    ret = register_kprobe(&kp);
    if (ret < 0) {
        printk(KERN_INFO "register_kprobe failed, returned %d\n", ret);
        return ret;
    }
    printk(KERN_INFO "Planted kprobe at %p\n", kp.addr);
    return 0;
}

static void __exit kprobe_exit(void) {
    unregister_kprobe(&kp);
    printk(KERN_INFO "kprobe at %p unregistered\n", kp.addr);
}

module_init(kprobe_init) module_exit(kprobe_exit) MODULE_LICENSE("GPL");
