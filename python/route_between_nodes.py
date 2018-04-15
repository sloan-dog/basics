from collections import deque
import timeit
class Node(object):
	def __init__(self, val, children=None):
		if children == None:
			self.children = []
			self.val = val

	def add_child(self, child_node):
		self.children.append(child_node)

	def add_children(self, children_node_list):
		self.children.extend(children_node_list)


class Graph(object):
	def __init__(self, nodes=None):
		if nodes == None:
			self.nodes = []
		self.nodes = nodes


class BFSSolver(object):

	@staticmethod
	def solve(start_node, end_node):
		if not start_node or not end_node:
			return False
		q = deque()

		visited = set()

		q.append(start_node)
		while len(q) > 0:
			n = q.popleft()
			# avoid cycles
			visited.add(n)
			children = n.children
			for _n in children:
				if _n not in visited:
					# avoid cycles
					q.append(_n)
				if _n == end_node:
					return True
		return False

class DFSSolver(object):

	@staticmethod
	def __dfs(node, visited):
		for child_node in node.children:
			if child_node not in visited:
				visited.add(child_node)
				DFSSolver.__dfs(child_node, visited)

	@staticmethod
	def solve(start_node, target_node):
		visited = set()

		visited.add(start_node)
		DFSSolver.__dfs(start_node, visited)

		return target_node in visited

if __name__ == "__main__":
	n1 = Node(1)
	n2 = Node(2)
	n3 = Node(3)
	n4 = Node(4)
	n5 = Node(5)
	n6 = Node(6)
	n7 = Node(7)

	n1.add_children([n2, n6])
	n6.add_child(n2)
	n2.add_children([n4, n5])
	n4.add_children([n4])
	n3.add_child(n7)

	"""

	n1: [n2, n6]
	n2: [n4, n5]
	n3: [n7]

	n4, n3 cycle for testing
	proves value of visited set
	n4: [n5]
	n5: [n4]
	n6: [n2]
	n7: []
	"""

	g = Graph([n1])
	s1 = timeit.default_timer()
	print(BFSSolver.solve(n1, n5))
	elapsed1 = timeit.default_timer() - s1
	s2 = timeit.default_timer()
	print(DFSSolver.solve(n1, n5))
	elapsed2 = timeit.default_timer() - s2
	print("BFS: {}\nDFS: {}".format(elapsed1, elapsed2))
