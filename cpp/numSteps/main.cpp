#include <iostream>
#include <map>

using namespace std;

class NumStepsTopDown {
    

    public:
        static int numWays(int n, map<int,int> &memo) {
        map<int,int>::iterator it;
        it = memo.find(n);
        if (it == memo.end()) {
            memo[n] = numWays(n-3, memo) + numWays(n - 2, memo) + numWays(n-1, memo);
        }
        return memo[n];
        }
        static int getNumWays(int n) {
            map<int, int> memo;
            memo[1] = 1;
            memo[2] = 2;
            memo[3] = 4;
            return numWays(n, memo);
        }

};

class NumStepsBottomUp {

public:
    static int getNumWays(int num) {
        // base cases 3 steps -> 4 ways, 2 steps -> 2 ways, 1 step -> 1 way
        int a = 4, b = 2, c = 1;
        if (num == 1) {
            return c;
        }
        if (num == 2) {
            return b;
        }
        if (num == 3) {
            return a;
        }
        int result;
        for (int i = 4; i <= num; i++) {
            // slide the assignment to the right and assign a the calculated value
            result = a + b + c;
            c = b;
            b = a;
            a = result;
        }
        return result;
    }

};

int main() {
    int n = 12;
    std::cout << NumStepsTopDown::getNumWays(n) << "\n";
    std::cout << NumStepsBottomUp::getNumWays(n) << "\n";
}