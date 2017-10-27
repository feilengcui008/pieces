obj-m := print_all_processes.o
KERNEL = /lib/modules/$(shell uname -r)/build
PWD = $(shell pwd)
default:
	make -C $(KERNEL) M=$(PWD)
clean:
	rm -f *.o *.ko *.mod.c *.order *.symvers
