from collections import deque

class Node(object):
	def __init__(self, val, left=None, right=None):
		self.val = val
		self.left = left
		self.right = right

def sorted_list_to_bst(arr, start, end):
	if start > end:
		return None

	mid = int( (start + end) / 2 )
	n = Node(arr[mid])
	n.left = sorted_list_to_bst(arr, start, mid - 1)
	n.right = sorted_list_to_bst(arr, mid + 1, end)
	return n 

def in_order_traversal(node):
	if node:
		print(node.val)
		in_order_traversal(node.left)
		in_order_traversal(node.right)

if __name__ == "__main__":
	arr = [1,2,3,4,5,6,6,7,8,9,10,11,11,12]
	head = sorted_list_to_bst(arr, 0, len(arr) - 1)
	in_order_traversal(head)






