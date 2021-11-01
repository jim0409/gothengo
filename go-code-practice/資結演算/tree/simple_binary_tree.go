package main

import "fmt"

// this is a test example to implement an int array to binary tree

// SimpleNode is a Unit of Tree
type SimpleNode struct {
	Val   int
	Left  *SimpleNode
	Right *SimpleNode
}


func ImplementTree(arr []int) *SimpleNode {
	rtNode := SimpleNode{arr[0],nil,nil}

	for i:=1; i < len(arr); i++ {
		rtNode.InsertNode(arr[i])
	}

	return &rtNode
}

func (n *SimpleNode) InsertNode(v int) {
	
	if n == nil {
		n = &SimpleNode{v, nil, nil}

	} else {

		if v < n.Val {
			if n.Left == nil{
				n.Left = &SimpleNode{v,nil,nil}
			}else{
				n.Left.InsertNode(v)
			}
		} else {
			if n.Right == nil{
				n.Right = &SimpleNode{v,nil,nil}
			}else{
				n.Right.InsertNode(v)
			}
		}
	}
}

func (n *SimpleNode) Traversal(s string){
	if n.Right != nil{
		
		n.Right.Traversal("right")
	}
	if n.Left != nil{
		n.Left.Traversal("left")
	}
	fmt.Printf("%s__%d\n",s ,n.Val)
}


func (n *SimpleNode) preOrderTraversal() {
	fmt.Printf("%d ",n.Val)
	if n.Left != nil{
		n.Left.preOrderTraversal()
	}
	if n.Right != nil{
		n.Right.preOrderTraversal()
	}
}

func (n *SimpleNode) inOrderTraversal(){
	if n.Left != nil{
		n.Left.inOrderTraversal()
	}
	fmt.Printf("%d ",n.Val)
	if n.Right != nil{
		n.Right.inOrderTraversal()
	}
}


func (n *SimpleNode) postOrderTraversal(){
	if n.Left != nil{
		n.Left.postOrderTraversal()
	}
	if n.Right != nil{
		n.Right.postOrderTraversal()
	}
	fmt.Printf("%d ",n.Val)

}

func main() {
	a := []int{6,3,5,4,7,8,9,2}
	ta := ImplementTree(a)
	fmt.Println(ta)
	ta.Traversal("root")

	fmt.Println(ta.PreOrderTraversal())
	ta.preOrderTraversal()
	fmt.Println()

	// fmt.Println(ta.InOrderTraversal())
	ta.inOrderTraversal()
	fmt.Println()

	// fmt.Println(ta.PostOrderTraversal())
	ta.postOrderTraversal()
	fmt.Println()

}


func (n *SimpleNode) PreOrderTraversal() []int{
	if n == nil{
		return nil
	}

	traver := []int{n.Val}
	l := n.Left.PreOrderTraversal()
	r := n.Right.PreOrderTraversal()
	traver = append(traver,l...)
	traver = append(traver,r...)

	return traver
}


func (n *SimpleNode) InOrderTraversal() []int{
	if n == nil{
		return nil
	}
	traver := []int{n.Val}
	l := n.Left.InOrderTraversal()
	traver = append(traver,l...)
	r := n.Right.InOrderTraversal()

	traver = append(traver,r...)

	return traver
}

func (n *SimpleNode) PostOrderTraversal() []int{
	if n == nil{
		return nil
	}
	traver := n.Left.PostOrderTraversal()
	traver = append(traver, n.Val)

	r := n.Right.PostOrderTraversal()
	traver = append(traver,r...)

	return traver
}