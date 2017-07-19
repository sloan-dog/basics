#include <iostream>
#include <string>
#include <queue>
using namespace std;

class Node 
{
	public:
		int value;
		Node *left;
		Node *right;

	Node(int x)
	{ 
		value = x;
		left = 0;
		right = 0;
	}

	void add_node(Node *node_ptr, int dir)
	{
		if (dir == 1)
		{
			right = node_ptr;
		} else {
			left = node_ptr;
		}
	}
};

class BreadthFirst
{
private:
	queue<Node*> q;
	static void touch_node(Node *n)
	{
		cout << "value " << std::to_string(n -> value) << "\n";
	}

public:
	BreadthFirst()
	{
		q = queue<Node*>();
	}

	void traverse(Node *n)
	{
		q.push(n);
		while (!q.empty())
		{
			Node *n_ptr = q.front();
			q.pop();
			touch_node(n_ptr);
			if (n_ptr -> left) 
			{
				q.push(n_ptr -> left);
			}
			if (n_ptr -> right)
			{
				q.push(n_ptr -> right);
			}
		}
	}

};

class TreeTraverse
{
public:

	static void in_order(Node *n)
	{
		if (!n) 
		{
			return;
		}

		in_order( n -> left);
		touch_node(n);
		in_order( n -> right);
	}
	static void pre_order(Node *n)
	{
		if (!n)
		{
			return;
		}

		touch_node(n);
		pre_order(n -> left);
		pre_order( n -> right);

	}
	static void post_order(Node *n)
	{
		if (!n)
		{
			return;
		}

		post_order( n -> left);
		post_order( n -> right);
		touch_node(n);
	}
private:
	static void touch_node(Node *n)
	{
		cout << "value " << std::to_string(n -> value) << "\n";
	}
}; 



int main() {
	cout << "Binary tree build";

	Node n1 = Node(1);
	Node n2 = Node(2);
	Node n3 = Node(3);
	Node n4 = Node(4);
	Node n5 = Node(5);
	Node n6 = Node(6);
	Node n7 = Node(7);
	Node n8 = Node(8);

	/*

		    n1
		  /	   \
	    n3	    n2
	  /   \    /  \
 	 n7	  n6  n5  n4
			       \
			        n8	 

	*/
	n1.add_node(&n2, 1);
	n1.add_node(&n3, 0);

	n2.add_node(&n4, 1);
	n2.add_node(&n5, 0);
	n4.add_node(&n8, 1);
	n3.add_node(&n6, 1);
	n3.add_node(&n7, 0);

	// Seperate build messages from our output
	cout << "\n";

	cout << "--IN ORDER--" << "\n";
	TreeTraverse::in_order(&n1);

	cout << "--POST ORDER--" << "\n";
	TreeTraverse::post_order(&n1);

	cout << "--PRE ORDER--" << "\n";
	TreeTraverse::pre_order(&n1);

	cout << "--BREADTH FIRST--" << "\n";
	BreadthFirst b = BreadthFirst();
	b.traverse(&n1);

	return 0;
};