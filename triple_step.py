
# naive
class TripleStep:

	@staticmethod
	def solve(target):
		memo = {}
		return TripleStep.get_ways(target, memo)

	@staticmethod
	def get_ways(target, memo):
		if target in memo:
			return memo[target]
		if target > 0:
			ways = TripleStep.get_ways(target - 1, memo) + TripleStep.get_ways(target - 2, memo) + TripleStep.get_ways(target - 3, memo)
			memo[target] = ways
			return ways
		if target == 0:
			return 1
		return 0

print(TripleStep.solve(7))