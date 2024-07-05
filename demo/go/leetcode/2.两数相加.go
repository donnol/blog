package leetcode

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToSlice() []int {
	r := make([]int, 0)
	for l := l; l != nil; l = l.Next {
		r = append(r, l.Val)
	}
	return r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}

	var v int
	var p int
	var l = r
	for {
		if l1 != nil && l2 != nil {
			v = l1.Val + l2.Val
		} else if l1 != nil {
			v = l1.Val
		} else if l2 != nil {
			v = l2.Val
		} else {
			break
		}

		if p == 1 {
			v += p
			p = 0
		}
		if v >= 10 {
			p = 1
			v -= 10
		}
		l.Val = v

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

		if l1 != nil || l2 != nil {
			n := &ListNode{}
			l.Next = n
			l = n
		}
	}
	if p == 1 {
		l.Next = &ListNode{Val: p}
	}
	return r
}
