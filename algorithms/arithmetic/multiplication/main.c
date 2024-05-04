#include <stdio.h>

int main() {
    int x, y, z;
    x = 7;
    y = 10;
    z = 0;
    while (y > 0) {
        if (y % 2 != 0) {
            y--;
            z = z + x;
        } else {
            y = y / 2;
            x = x + x;
        }
    }

    printf("Ergebnis: %d\n", z);
}

