package bingalie

import (
	"fmt"
	"testing"
)

//singly-linked list
//单链表的结构体，假设每个节点的值为int类型
type ListNode struct {
	     Val int
	     Next *ListNode
}

//单链表的增删改查 todo
//在节点target前面新增一个节点，新增节点的值为x
func AddNodePrev(head, target *ListNode, x int) *ListNode {
	curr := head
	if head == target {
		head = &ListNode{
			Val:  x,
			Next: curr,
		}
		fmt.Println("head is ",head)  //todo 为什么不写返回head，直接改的head的值，调这个函数的外面，这个head不会改变呢？ 传参是个指针呀
		return head
	}
	for curr != nil {
		if curr.Next == target {
			curr.Next = &ListNode{
				Val:  x,
				Next: target,
			}
			return head
		}
		curr = curr.Next
	}
	return head
}

//在节点target后面新增一个节点，新增节点的值为x
func AddNodeNext(target *ListNode, x int)  {
	if target.Next == nil {
		target.Next = &ListNode{
			Val:  x,
			Next: nil,
		}
	}else {
		target.Next = &ListNode{
			Val:  x,
			Next: target.Next,
		}
	}
}

//在节点target前面删除一个节点
func DelNodePre(head, target *ListNode) *ListNode {
	if target == head {
		return head
	}
	if head.Next == target {
		return target
	}
	curr := head
	for curr != nil {
		if curr.Next.Next == target {
			curr.Next = target
			return head
		}
		curr = curr.Next
	}
	return head
}

//在节点target后面删除一个节点
func DelNodeNext(target *ListNode)  {
	if target.Next == nil {
		return
	}
	target.Next = target.Next.Next
	return
}

//查询所有节点值
func GetAllNodesVal(head *ListNode) []int {
	allVals := make([]int, 0)
	curr := head
	for curr != nil {
		allVals = append(allVals, curr.Val)
		curr = curr.Next
	}
	return allVals
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	//初始值，当前节点为头结点，前一个节点为空
	var prev *ListNode
	curr := head
	//迭代一个链表
	for curr != nil {
		next := curr.Next //保存下一个节点
		curr.Next = prev  //将当前节点指向前一个节点
		prev = curr       //当前节点变为前一个节点
		curr = next       //下一个节点变为当前节点
	}
	return prev
}

//-------------------------------------------------------------

func TestLinked(t *testing.T)  {
	node1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	head := DelNodePre(node1, node2)
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, ",")
		curr = curr.Next
	}

}