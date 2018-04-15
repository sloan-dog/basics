from collections import deque

class Node(object):
	def __init__(self, col_idx, row_idx):
		self.col_idx = col_idx
		self.row_idx = row_idx

class RobotGrid:

	@staticmethod
	def solveBFS(grid):
		q = deque()
		n = Node(0,0)
		q.append(n)
		height = len(grid)
		width = len(grid[0])
		while len(q) > 0:
			n = q.popleft()
			# we are at rightmost south most square
			if n.col_idx == width - 1 and n.row_idx == height - 1:
				return True
			if RobotGrid.is_good_square(grid, n.col_idx + 1, n.row_idx):
				nright = Node(n.col_idx + 1, n.row_idx)
				q.append(nright)
			if RobotGrid.is_good_square(grid, n.col_idx, n.row_idx + 1):
				ndown = Node(n.col_idx, n.row_idx + 1)
				q.append(ndown)
		return False

	@staticmethod
	def solve_recursive(grid):
		visited = set()
		l = len(grid)
		w = len(grid[0])
		return RobotGrid.get_path(grid, w - 1, l - 1, visited)
	
	@staticmethod
	def get_path(grid, col_idx, row_idx, visited):
		# print("trying ", col_idx, row_idx)
		if (row_idx, col_idx) in visited:
			return True

		if col_idx < 0 or row_idx < 0 or grid[row_idx][col_idx] < 1:
			# print(col_idx, row_idx, " turns up false early")
			return False

		visited.add((row_idx, col_idx))

		is_at_origin = row_idx == 0 and col_idx == 0

		if (is_at_origin or \
			RobotGrid.get_path(grid, col_idx - 1, row_idx, visited) or\
			RobotGrid.get_path(grid, col_idx, row_idx - 1, visited)):
			# print(col_idx, row_idx, " turns up true")
			return True
		return False


	@staticmethod
	def is_good_square(grid, col_idx, row_idx):
		if row_idx < len(grid) and \
		col_idx < len(grid[0]) and \
		grid[row_idx][col_idx] != 0:
			return True
		return False
		

"""
grid format
1 = good
0 = bad
[
	[1 1 1 1 1 1 1 1 1]
	[1 1 1 1 1 1 0 1 1]
	[1 1 1 1 1 1 1 1 0]
	[1 1 1 0 1 1 1 1 1]
]
"""

grid_possible = [ \
	[1, 1, 1, 1, 1, 1, 1, 1, 1], \
	[1, 1, 1, 1, 1, 0, 0, 1, 1], \
	[1, 1, 1, 1, 1, 1, 1, 1, 0], \
	[1, 1, 1, 0, 1, 1, 1, 1, 1] \
]

grid_impossible = [ \
	[1, 1, 1, 1, 1, 1],
	[0, 0, 0, 0 ,0 ,0],
	[1, 0, 1, 1, 1, 1]
]

grid_big = [ [1] * 400 ] * 400



# BFS
print(RobotGrid.solveBFS(grid_possible))
print(RobotGrid.solveBFS(grid_impossible))

# Recursive
print(RobotGrid.solve_recursive(grid_possible))
print(RobotGrid.solve_recursive(grid_impossible))
print(RobotGrid.solve_recursive(grid_big))

