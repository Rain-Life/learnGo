package tree

import "fmt"

type Node struct {
	Value int
	Left, Right * Node
}

//这是给结构体定义的方法，它是放在结构体外边的(括号中的是接收者，可以理解成其它语言中的this)
func (node Node) Print() {
	fmt.Print(node.Value)
}//你也可能写成这样func Print(node treeNode)，只是调用的时候，方式变了而已

func (node *Node) SetValue(value int)  {
	node.Value = value
}

/**
 * 在C/C++中，局部变量是分配在栈中，而全局变量分配在堆中。在堆上分配需要手动释放，分配在栈中，在使用完之后会销毁栈
 * 在go中，我们不需要知道变量被分配在堆中还是栈中，无所谓。根据场景，一个变量可能分配到堆中，也可能分配到栈中
 * 比如下边的构造函数中，创建的这个treeNode虽然是局部变量，但是它要返回出去给别的代码使用，所以它虽然是局部变量，但是
 * 它被分配在了堆中。如果它不需要给别的程序使用，它就会被分配到栈中
*/

//构造函数
func CreateNode(value int) *Node {
	return &Node{Value:value}
}


