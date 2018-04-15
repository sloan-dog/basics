#!/bin/usr/python2.7
class Solution:
    answers = []

    def stepnum(self, N, M):
        length = 0
        temporary = M
        while temporary != 0:
            temporary = int(temporary / 10)
            length += 1

        self.answers = []
        self.dfs(length, N, M)
        return sorted(answers)

    def dfs(length, N, M, num=0):
        if (num >= N and num <= M):
            self.answers.append(num)
        
        if length == 0:
            return;
        if num == 0:
            for i in range(1, 10):
                self.dfs(length - 1, N, M, i)
            return;

        last_digit = num % 10
        if last_digit == 0:
            self.dfs(length - 1, N, M, num * 10 + last_digit + 1)
        elif last_digit == 9:
            self.dfs(length - 1, N, M, num * 10 + last_digit - 1)
        else:
            self.dfs(length - 1, N, M, num * 10 + last_digit - 1)
            self.dfs(length - 1, N, M, num * 10 + last_digit + 1)
    
if __name__ == "__main__":
    s = Solution()

    # Input format
    # N/n
    # M/n

    N = int(input().strip())
    M = int(input().strip())

    print(s.stepnum(N, M))
