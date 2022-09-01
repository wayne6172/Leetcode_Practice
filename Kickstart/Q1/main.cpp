#include <cstdio>

int main() {

    int testCase;
    int test = 1;
    char I[100001] = {0}, P[100001] = {0};
    scanf("%d", &testCase);
    while(testCase--) {
        scanf("%s", I);
        scanf("%s", P);

        int i = 0, j = 0, ans = 0;
        while(I[i] && P[j]) {
            I[i] == P[j] ? i++ : ans++;
            j++;
        }

        while(P[j] != '\0') {
            ans++;
            j++;
        }

        if (I[i] != '\0') {
            ans = -1;
        }

        if (ans == -1) {
            printf("Case #%d: IMPOSSIBLE\n", test++);
        }
        else {
            printf("Case #%d: %d\n", test++, ans);
        }
    }

    return 0;
}