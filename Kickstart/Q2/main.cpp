#include <cstdio>

int main() {
    int testCase;
    int test, i, j, temp, need;

    scanf("%d", &testCase);
    for(test = 1; test <= testCase; test++) {
        char N[130000] = {0};
        char ans[130000] = {0};
        scanf("%s", N);

        i = temp = 0;
        while(N[i]) {
            temp = temp + (N[i] - '0');
            i++;
        }
        j = i;

        need = 9 - (temp % 9);
        if(need == 9) {
            need = 0;
        }

        i = need == 0 ? 1 : 0;
        while(N[i] && need >= N[i] - '0') {
            i++;
        }

        N[j] = '\0';
        while(j > i) {
            N[j] = N[j - 1];
            j--;
        }

        N[i] = need + '0';

        printf("Case #%d: %s\n", test, N);
    }

    return 0;
}