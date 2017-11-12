import timeit
import random
import string

class Naive(object):
    @staticmethod
    def is_palindrome(_string, i, j):
    # runs roughly n times
        while i < j:
            if _string[i] == _string[j]:
                i += 1
                j -= 1
            else:
                return False
        return True

    @staticmethod
    def get_longest_palindrome(_string):
        l = len(_string)
        largest__string = ""
        greatest_length = 0
        start = 0
        # runs n times
        for i in range(l):
            # runs n - 1 times (n)
            for j in range(i + 1, l):
                is_p = Naive.is_palindrome(_string, i, j)
                length = j - i + 1
                # print(is_p, _string[i:j+1])
                if is_p and length > greatest_length:
                    greatest_length = length
                    start = i
        return _string[i:i+greatest_length], greatest_length

class Dyn(object):
    @staticmethod
    def get_longest_palindrome(_string):
        """
        O(n) + O(n) + O(n^2) -> O(n^2)
        """
        longest_start = 0
        n = len(_string)
        max_length = 1
        table = [ [False for x in range(n)] for y in range(n)]

        # singles
        for i in range(n):
            # single length substr are palin
            table[i][i] = True

        # doubles
        for i in range(n - 1):
            if _string[i] == _string[i + 1]:
                table[i][i + 1] = True
                # we havent set a 2 character long max_length
                if max_length < 2:
                    longest_start = i
                max_length = 2
        
        # triples and beyond
        # calc 3s, calc 4s, calc 5s, etc
        # 3s depend on 1s, 4s depend on 2s, 5s depend on 3s, etc
        for k in range(3, n):

            # check all sub_strings for range 0 -> i < string_length - sub_string_length (excluding last)
            # aka if k = 3, and string is 9 long, we want to check for every starting index 
            # from 0 -> 5 inclusive, for each 2 -> 8 inclusive ending index

            for i in range(n - k):
                j = i + k - 1

                if table[i + 1][j - 1] and _string[i] == _string[j]:
                    table[i][j] = True
                    if k > max_length:
                        max_length = k
                        longest_start = i

        return _string[longest_start: longest_start + max_length], max_length


def randomword(length):
    """
    Util function to generate large strings
    """
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(length))


if __name__ == "__main__":
    _input = "babycakessekacbaby"
    # expected output = "cakessekac"

    # million char string
    large_input = randomword(int(10e3))

    print("\n***SMALL INPUT***\n")

    start_time1 = timeit.default_timer()
    result1 = Naive.get_longest_palindrome(_input)
    elapsed1 = timeit.default_timer() - start_time1

    start_time2 = timeit.default_timer()
    result2 = Dyn.get_longest_palindrome(_input)
    elapsed2 = timeit.default_timer() - start_time2

    print("NAIVE: R: {}, T: {}".format(result1, elapsed1))
    print("DYN: R: {}, T: {}".format(result2, elapsed2))

    print("\n***LARGE INPUT***\n")

    start_time3 = timeit.default_timer()
    result3 = Naive.get_longest_palindrome(large_input)
    elapsed3 = timeit.default_timer() - start_time3

    start_time4 = timeit.default_timer()
    result4 = Dyn.get_longest_palindrome(large_input)
    elapsed4 = timeit.default_timer() - start_time4

    print("NAIVE: R: {}, T: {}".format(result3, elapsed3))
    print("DYN: R: {}, T: {}".format(result4, elapsed4))

    