#include <stdio.h>

// gcc ./demo/c/compare.c -o ./demo/c/out
// ./demo/c/out

int main(int *argc, char *argv)
{
    int a = 11;

    // 当符号写对了，在左在右皆可
    if(a == 11) {
        printf("%d\n", a); // 11
    }
    if(11 == a) {
        printf("%d\n", a); // 11
    }

    // 当不小心把==写成了=，会变为赋值操作，从而导致a的值变为了10；而这明显不是我的本意
    if(a = 10) {
        printf("%d\n", a); // 10
    }
    // error: lvalue required as left operand of assignment
    // if(10 = a) {
    //     printf("%d\n", a);
    // }

    // 从而衍生出一种习惯，将常量写在if条件的左边，将变量写在右边，避免将`==`写成了`=`导致逻辑出错。
}
