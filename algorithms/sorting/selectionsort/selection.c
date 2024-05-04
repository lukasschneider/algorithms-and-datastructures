#include <stdio.h>

void swap(int *a, int *b) {
    int t = *a;
    *a = *b;
    *b = t;
}

void selection_sor(int *arr, int *n) {
    for (int i = 0; i < *n; i++) {
        int min = i;
        for (int j = i + 1; j < *n; j++) {
            if (arr[j] < arr[min]) {
                min = j;
            }
        }
        swap(&arr[i], &arr[min]);
    }
}

void printArray(int *arr, int *n) {
    for (int i = 0; i < *n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int n = 5;
    int a[] = {1,25,6,10,8};
    printArray(a, &n);
    selection_sor(a, &n);
    printArray(a, &n);

}
