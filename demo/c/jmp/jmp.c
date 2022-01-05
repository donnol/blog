#include <stdio.h>
#include <setjmp.h>

// jmp_buf定义在setjmp.h:
// /* Calling environment, plus possibly a saved signal mask.  */
// struct __jmp_buf_tag
//   {
//     /* NOTE: The machine-dependent definitions of `__sigsetjmp'
//        assume that a `jmp_buf' begins with a `__jmp_buf' and that
//        `__mask_was_saved' follows it.  Do not move these members
//        or add others before it.  */
//     __jmp_buf __jmpbuf;		/* Calling environment.  */
//     int __mask_was_saved;	/* Saved the signal mask?  */
//     __sigset_t __saved_mask;	/* Saved signal mask.  */
//   };
// 
// typedef struct __jmp_buf_tag jmp_buf[1];
// 
// That means that jmp_buf is not a pointer, but an array with a single structure in it. So you use it like a normal array of structures:
// 
// jmp_buf _env;
// 
// _env[0].__jmpbuf[x] = y;

// [c __attribute__](https://sites.google.com/site/emmoblin/gcc-tech/gun-attribute)

int main()
{
    // 一个缓冲区，用来暂存环境变量
    jmp_buf buf;
    printf("line1 \n");

    // 保存此刻的上下文信息
    int ret = setjmp(buf);
    printf("ret = %d, buf = %p \n", ret, buf);

    // 检查返回值类型
    if (0 == ret)
    {
        // 返回值0：说明是正常的函数调用返回
        printf("line2 \n");

        // 主动跳转到 setjmp 那条语句处
        longjmp(buf, 1);
    }
    else
    {
        // 返回值非0：说明是从远程跳转过来的
        printf("line3 \n");
    }

    printf("line4 \n");

    return 0;
}
