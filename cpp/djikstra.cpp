#include <iostream>
#include <string>
#include <queue>
#include <vector>
#include <math.h>
using namespace std;

struct Edge
{
	int cost;
	Node *node;
};

struct Coordinate
{
	float x;
	float y;
	float z;
	// 3d constructor
	Coordinate(float _x, float _y, float _z)
	{
		x = _x;
		y = _y;
		z = _z;
	};
	// 2d constructor
	Coordinate(float _x, float _y) {
		x = _x;
		y = _y;
		z = 0.0f;
	};
	float getDistanceToCoord(Coordinate coord) {
		// euclidean distance in 3 dimensions
		return sqrt(
			pow((coord.x - x), 2) + 
			pow((coord.y - y), 2) + 
			pow((coord.z - z), 2));
	};
};

struct Node 
{
	Coordinate *coord;
	std::string name; 
	std::vector<Edge*> edges;

	Node(Coordinate *_coord, char *cstr)
	{ 
		coord = _coord;
	};

	void AddEdge(Node *n, float _cost) {
		Edge edge {
			_cost,
			n,
		};
		edges.push_back(&edge);
	};
};

int main() {
	cout << "Graph build";

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