#include <stdio.h>

static int (*bpf_ktime_get_ns)(void) = (void *) 5;

int xxx() {
    return 3333;
}

int yyy() {
    return 4444;
}

int main(int *argc, char *argv)
{
    int a = 1, b; // 相当于：int a = 1, b = 0; 表示a和b都是int类型，其中定义a的同时为其赋值为1；看起来跟go的多赋值很像，其实有很大不同。
    printf("%d, %d\n", a, b);

    bpf_ktime_get_ns = xxx;

    int x = bpf_ktime_get_ns(); // 如果没有上面给`bpf_ktime_get_ns`赋值一个函数，会报：段错误 (核心已转储)
    printf("%d\n", x);

    bpf_ktime_get_ns = yyy;

    int y = bpf_ktime_get_ns();
    printf("%d\n", y);
}
