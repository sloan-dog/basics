#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <deque>
#include <map>
using namespace std;

struct Node
{
    int value;
    int level;
    
    Node(int x, int y)
    {
        value = x;
        level = y;
    }
};

std::vector<int> get_max_factors(int num, std::map<int, std::vector<int>> *fac_memo)
{
    std::map<int, std::vector<int>>::iterator it;
    std::vector<int> int_vec = std::vector<int>();
    it = (*fac_memo).find(num);
    if (it != (*fac_memo).end()) {
        return (*fac_memo)[num];
    }
    int s = std::sqrt(num);
    for (int i = 2; i <= s; i ++) {
        int divides = num % i;
        if (divides == 0) {
            // this is causing bad allocation?
            int v = num / i;
            int_vec.push_back(v);
        }
    }
    (*fac_memo)[num] = int_vec;
    // cout << " vec size " << to_string(int_vec.size()) << "\n";
    return int_vec;
}

int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    std::map<int, int> memo = std::map<int, int>();
    std::map<int, int>::iterator it;
    
    std::map<int, std::vector<int>> fac_memo = std::map<int, std::vector<int>>();
    std::deque<Node*> q = std::deque<Node*>();
    int q_count;
    cin >> q_count;
    int target;
    for (int i = 0; i < q_count; i++) {
        cin >> target;
        it = memo.find(target);
        if (it != memo.end()) {
            cout << to_string(memo[target]) << "\n";
        } else {
            // track for this target number the smallest cost of a number
            std::map<int, int> least_cost = std::map<int, int>();
            std::map<int, int>::iterator o_it;

            Node *_n = new Node(target, 0);
            q.push_back(_n);
            while (!q.empty()) {
                // cout << "val " << to_string(_n -> value) << " lvl " << to_string(_n -> level) << "\n";
                Node *n_ptr = q.front();
                q.pop_front();
                int lvl = n_ptr -> level;
                int val = n_ptr -> value;
                delete n_ptr;

                o_it = least_cost.find(val);
                // lets see if we already have a cost for this value
                if ( o_it != least_cost.end() ) {
                    if ( least_cost[val] < lvl ) {
                        // we have a new lowest cost
                        least_cost[val] = lvl;
                    } else {
                        // we're looking at a more expensive cost
                        // lets get the next node
                        continue;
                    }
                } else {
                    least_cost[val] = lvl;
                }
                if (val == 0) {
                    memo[target] = lvl;
                    cout << to_string(lvl) << "\n";
                    break;
                }
                std::vector<int> int_vec = get_max_factors(val, &fac_memo);
                for (int i = 0; i < int_vec.size(); i++) {
                    // cout << "fac " << to_string(int_vec[i]) << "\n";
                    Node *__n = new Node(int_vec[i], lvl + 1);
                    q.push_back(__n);
                }
                Node *__n = new Node(val - 1, lvl + 1);
                q.push_back(__n);
            }
            q.clear();
        }
    }
    return 0;
}
