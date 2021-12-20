#include <stdio.h>

int main(int *argc, char *argv)
{
    void *sh;
    char *s;
    int hdrlen = 8;
    sh = (void *)0x1000;

    // 1
    s = (char *)sh + hdrlen;
    printf("s: %p\n", s); // 0000000000001008

    // 2 先转再加
    s = (char *)sh;
    s += hdrlen;
    printf("s: %p\n", s); // 0000000000001008

    // 3 先加再转
    void *t = sh + hdrlen;
    s = (char *)t;

    printf("s: %p\n", s); // 0000000000001008
}
