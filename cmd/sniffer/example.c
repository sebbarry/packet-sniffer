#include <stdio.h>
#include <stdlib.h>


int main() {
    int i, n; 
    int *a; 
    printf("Number of elements to be entered: ");
    scanf("%d", &n);
    a = (int*) calloc(n, sizeof(int));
    printf("Enter %d numbers: \n", n);
    for(i = 0; i<n;i++) { //enter the numbers sequentually that follow in a loop
        scanf("%d", &a[i]); //allocating value to the dereferenced pionter value
    }
    printf("The numbers entered are: "); 
    for(i = 0; i < n; i++) {
        printf("%d", a[i]);
    }
    free(a);
    return(0);
}
