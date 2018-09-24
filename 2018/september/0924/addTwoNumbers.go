package main

import "fmt"

/**
https://leetcode-cn.com/problems/add-two-numbers/description/
2. 两数相加
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
func main() {

	// 初始化 struct 的两种方式
	// b和c的初始化形式相同

	a := ListNode{Val: 1}
	b := &ListNode{Val: 1}
	c := new(ListNode)

	// cannot use _ as value
	// var _ = ListNode{0, nil}

	fmt.Println(a.Val)
	fmt.Println(b.Val)
	fmt.Println(c.Val)
	fmt.Println(a)
	fmt.Println(b)

}

func test() {
	node1 := new(ListNode)
	node12 := new(ListNode)
	node13 := new(ListNode)
	node1.Val = 8
	node12.Val = 9
	node13.Val = 9
	node1.Next = node12
	node12.Next = node13

	node2 := new(ListNode)
	node2.Val = 2

	numbers := addTwoNumbers(node1, node2)

	for numbers != nil {
		i := numbers.Val
		fmt.Println(i)
		numbers = numbers.Next
	}
	fmt.Println("end")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 初始化一个  listNode
	node := new(ListNode)
	carry := 0
	var currentNode = node
	var initial = false

	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		if !initial {
			node.Val = sum
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
			initial = true
			continue
		}

		nextNode := new(ListNode)
		nextNode.Val = sum
		nextNode.Next = nil
		currentNode.Next = nextNode
		currentNode = currentNode.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	// copy from leetcode
	dummy := &ListNode{}
	current, carry := dummy, 0

	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		carry, val = val/10, val%10
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
