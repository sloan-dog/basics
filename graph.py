
"""
{name0: {data:data,adj:[name1,name2,name3]}}
"""
class AdjacencyMap(dict):

	def __init__(self, pattern_string=None):
		super(AdjacencyMap, self).__init__()
		if pattern_string:
			self.create_from_pattern(pattern_string)

	def create_from_pattern(self, pattern_string):
		nodes = pattern_string.strip().split("/")
		for node in nodes:
			n, d, a = node.split(":")
			d = int(d)
			adj = []
			for _as in a.split(","):
				if _as != "":
					adj.append(_as)
			self[n] = {"data":d,"adj":adj}

class Node(object):
	def __init__(self, name, data):
		# earlier I tried having a default arg of adjacent=[]
		# however default args are created as a pointer
		# so all instances of class were appending to the same list pointer
		self.adjacent_nodes = []
		self.data = data
		self.name = name

	def add_adjacent(self, child_node):
		self.adjacent_nodes.append(child_node)

	def get_adj(self):
		return [node.name for node in self.adjacent_nodes]


class Graph(object):
	def __init__(self, nodes={}):
		self.nodes = nodes

	def create_from_adjacency_map(self, _map):
		for name in _map:
			self.create_node(name, _map)

	def create_node(self, name, _map):
		if name in self.nodes:
			# node has already been created
			return self.nodes[name]
		item = _map[name]
		data = item["data"]
		adj = item["adj"]
		n = Node(name, data)
		self.nodes[name] = n
		for a in adj:
			# recursively create the adjecent nodes as well
			n.add_adjacent(self.create_node(a, _map))
			# get the newly added adj nodes
			a_adj = n.get_adj()
		return n

if __name__ == "__main__":
	g = Graph()
	p = "a:7:b/b:5:c,a/c:8:b/d:12:a,c"
	am = AdjacencyMap(p)
	g.create_from_adjacency_map(am)
	for name in g.nodes:
		print("name {} : adj {}".format(name, g.nodes[name].get_adj()))

"""
/name0:7:name1/name1:5:name0,name2/name2:8:name1
"""
		

		


