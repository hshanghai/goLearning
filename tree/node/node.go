package node

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node TreeNode) Print(){
	fmt.Print(node.Value,"  ")
}

func (node *TreeNode) SetValue(value int)  {
	//nil 可以调用所以添加判断
	if node == nil {
		fmt.Println(" Setting value to nil" + "node. Ignored")
		return
	}
	node.Value = value
}

//遍历树形 先遍历左支数
func (node *TreeNode) Traverse()  {
	if node == nil{
		return
	}
	//例如如果是java 是必须先判断是否为nil
	//if node.left != nil {
	//	node.left.traverse()
	//}
	node.Left.Traverse() //go 语音nil是可以被调用，所以可以直接调用，在程序开头判断一下，为nil就return出去就行
	node.Print()
	node.Right.Traverse()
}

//工厂函数
func CreateNode(values int) *TreeNode  {
	return &TreeNode{Value:values}
}