

class Node {
    constructor(data, left_node=null, right_node=null) {
        this.data = data
        this.left_node = left_node
        this.right_node = right_node
    }

    setLeft(node) {
        this.left_node = node
    }

    setRight(node) {
        this.right_node = node
    }
}

class Tree {

    constructor(head_node) {
        this.root_node = head_node
    }
    
    __traverseInOrder(root) {
        // left, root, right
        if (root) {
            if (root.left_node) {
                this.__traverseInOrder(root.left_node)
            }
            console.log(root.data)
            if (root.right_node) {
                this.__traverseInOrder(root.right_node)
            }
        }
        
    }

    __traversePostOrder(root) {
        if (root) {
            if (root.left_node) {
                this.__traversePostOrder(root.left_node)
            }
            if (root.right_node) {
                this.__traversePostOrder(root.right_node)
            }
            console.log(root.data)
        }
    }

    __traversePreOrder(root) {
        if (root) {
            console.log(root.data)
            if (root.left_node) {
                this.__traversePreOrder(root.left_node)
            }
            if (root.right_node) {
                this.__traversePreOrder(root.right_node)
            }
        }
    }

    traverse(pattern) {
        switch (pattern) {
            case "inorder":
                console.log("inorder")
                this.__traverseInOrder(this.root_node)
                break
            case "preorder":
                console.log("preorder")
                this.__traversePreOrder(this.root_node)
                break
            case "postorder":
                console.log("postorder")
                this.__traversePostOrder(this.root_node)
                break
            default:
                console.log(new Error(`invalid order pattern -> ${pattern}`))
        }
    }
}


/*
       5
    6     2
   7 3   8 14
           9 20
*/
n1 = new Node(5)
n2 = new Node(6)
n3 = new Node(2)

n1.setLeft(n2)
n1.setRight(n3)

n4 = new Node(7)
n5 = new Node(3)

n2.setLeft(n4)
n2.setRight(n5)

n6 = new Node(8)
n7 = new Node(14)

n3.setLeft(n6)
n3.setRight(n7)

n8 = new Node(9)
n9 = new Node(20)

n7.setLeft(n8)
n7.setRight(n9)

t = new Tree(n1)

t.traverse("inorder")
t.traverse("postorder")
t.traverse("preorder")