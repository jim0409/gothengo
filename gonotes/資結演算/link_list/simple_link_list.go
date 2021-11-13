package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0], Next: nil}
	// make a copy of head to iterate it ...
	currNode := head

	for i := 1; i < len(arr); i++ {
		// create an instance `tmpNode` to be the next ListNode ...
		tmpNode := &ListNode{Val: arr[i], Next: nil}
		// assign the location of tmpNode as currNode.Next
		currNode.Next = tmpNode
		// redeclare the origin currNode(pointer) as currNode.Next
		currNode = currNode.Next
	}

	return head
}

func (l *ListNode) PrintList() {
	if l == nil {
		return
	}
	fmt.Printf("%d ", l.Val)
	l.Next.PrintList()
}


// Pop would delete the first elements of the ListNode and return it
func (l *ListNode) Pop() *ListNode{
	headNode := l					// 宣告用來做後面的記憶體釋放用的
	// fmt.Println(headNode,l)		// &{1 0xc0000761f0} &{1 0xc0000761f0}
	l = headNode.Next
	// fmt.Println(headNode,l)		// &{1 0xc0000761f0} &{2 0xc000076200}
	headNode = nil	 				// 防止記憶體洩漏
	// fmt.Println(headNode,l)		// <nil> &{2 0xc000076200}
	return l
}

func (l *ListNode) AddNodes(v int) *ListNode{
	tmpList := l // 複製一份 tmpList 出來
	currList := tmpList // 宣告一個記憶體指標來做迭代

	for currList.Next != nil{
		currList = currList.Next	// 當 Next 不為 Nil 的時候，將指標往下指
	}

	currList.Next = &ListNode{Val:v, Next: nil} // 蝶帶到最後一個指標時，修改指標指向的參數 <===> 間接修改掉 tmpList 的指標指向的參數
	l = nil	// 必要的話，可以對 l 做 nil 指向，釋放記憶體 ... 另一種方法，直接對l做一份 currList，之後再回傳 l <優: 節省記憶體; 缺: 失去平行運算的能力>

	return tmpList
}

func (l *ListNode)removeElements(v int) *ListNode {
	tmpList := l
	currNode := tmpList

	for currNode.Next != nil {
		if currNode.Next.Val == v {
			currNode.Next = currNode.Next.Next
		}else {
			currNode = currNode.Next
		}
	}
	l = nil	// 必要的話，可以對 l 做 nil 指向，釋放記憶體 ... 另一種方法，直接對l做一份 currList，之後再回傳 l <優: 節省記憶體; 缺: 失去平行運算的能力>

	return tmpList
}

func (l *ListNode) reverse() *ListNode{
	tmpList := l
	currNode := tmpList

	addressLink := []*ListNode{}
	for currNode.Next != nil{
		addressLink = append(addressLink, currNode)
		currNode = currNode.Next
	}
	addressLink = append(addressLink, currNode)


	fmt.Println(addressLink)

	resultList := &ListNode{}
	tmpNode := resultList
	for i:= len(addressLink)-1; i > 0 ;i--{
		// resultList = addressLink[i]
		fmt.Println(addressLink[i])
		tmpNode = addressLink[i]
		tmpNode = tmpNode.Next
	}

	fmt.Println(resultList)

	return resultList
}


func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	listNodes := GenList(a)
	listNodes.PrintList()
	fmt.Println()

	listNodes = listNodes.Pop()
	listNodes.PrintList()
	fmt.Println()

	listNodes = listNodes.AddNodes(8)
	listNodes.PrintList()
	fmt.Println()

	listNodes = listNodes.removeElements(3)
	listNodes.PrintList()
	fmt.Println()

	listNodes = listNodes.reverse()
	listNodes.PrintList()
	fmt.Println()
}
