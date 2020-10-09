package main

import (
	"fmt"
	"learnGo/tree"
)

//对别人写好的包进行扩展
type myTreeNode struct {
	node *tree.Node
}

//增加一个后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.Node
	root  = tree.Node{Value:3}
	//fmt.Println(root)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	nodes := []tree.Node{
		{Value:3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}