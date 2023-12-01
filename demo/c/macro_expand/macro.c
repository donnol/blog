#include <stdio.h>
#include <stdlib.h>

#define SUM(a, b) ((a) + (b))
#define SUB(a, b) ((a) - (b))

int main()
{
    int a = 2;
    int b = 1;
    int c = 0;

    c = SUM(a, b) + SUB(a, b);
    printf("%d\n", c);
}
