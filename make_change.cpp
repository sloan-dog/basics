#include <iostream>
#include <string>
#include <map>
#include <utility>
#include <vector>
using namespace std;

typedef std::pair<int, int> memo_key;
typedef std::map<memo_key, int> memo_map;
typedef std::vector<int> int_vec;
typedef std::pair<memo_key, int> memo_pair;

class ChangeMaker
{
private:
	int target;
	int_vec * denoms;
	memo_map memo;

	long make_change(const int amt, const int_vec * d, const int idx)
	{
		memo_key key = memo_key(amt, idx);
		memo_map::iterator memo_itr = memo.find(key);
		if (memo_itr != memo.end()) {
			return memo[key];
		}
		int denom = (*d)[idx];
		int divisble = amt % denom;
		if (idx == (*d).size() - 1) {
			return ( (divisble == 0) ? 1 : 0);
		}

		long ways = 0;
		for(int i = 0; denom * i <= amt; i ++)
		{
			int remaining = amt - denom * i;
			ways += make_change(remaining, d, idx + 1);
		}
		memo.insert(memo_pair(key, ways));
		return ways;

	}

public:
	ChangeMaker(int amount, int_vec * denominations)
	{
		target = amount;
		denoms = denominations;
	}
	long get_change()
	{
		int i = 0;
		return make_change(target, denoms, i);
	}
};

int main()
{
	/* Input looks like this
	250 26
	8 47 13 24 25 31 32 35 3 19 40 48 1 4 17 38 22 30 33 15 44 46 36 9 20 49
	250 = amount
	26 = num denoms
	2nd line is denoms
	expected result for this input -> 3542323427
	*/
	int amt;
	int num_denoms;
	cin >> amt >> num_denoms;
	int_vec d(num_denoms);
	for (int c_i = 0; c_i < num_denoms; c_i++) {
		cin >> d[c_i];
	}
	ChangeMaker * cm = new ChangeMaker(amt, &d);
	long result = cm -> get_change();
	cout << to_string(result) << "\n";
}