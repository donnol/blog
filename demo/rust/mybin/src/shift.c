#include <stdio.h>
#include "shift.h"

// warning: implicit declaration of function ‘shift_left’ [-Wimplicit-function-declaration] -- 在使用之前先要声明，或者在头文件声明然后导入它
// int shift_left(int a, int n);
// int shift_right(int a, int n);

int main(int argc, char **argv)
{
    int r = shift_left(10, 3);
    printf("%d\n", r); // 80 = 10 * 2^3

    // error: redefinition of ‘r’ -- 不能重复定义
    // int r = shift_right(10, 3);
    int rr = shift_right(10, 3);
    printf("%d\n", rr); // 1 = 10 / 2^3
}

int shift_left(int a, int n)
{
    return a << n; // a * 2^n
}

int shift_right(int a, int n)
{
    return a >> n; // a / 2^n
}
