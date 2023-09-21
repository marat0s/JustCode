package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	current := temp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next, list1 = list1, list1.Next
		} else {
			current.Next, list2 = list2, list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return temp.Next
}

// посмотрел в гугле тк не понял как выводить
func printList(head *ListNode) { //head это 1 элемент листа
	for head != nil { //если лист не пустой то делай
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ") //для красоты
		}
	}
	fmt.Println()
}

func main() {

	l1 := &ListNode{Val: 10, Next: &ListNode{Val: 1000, Next: &ListNode{Val: 1337}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}

	printList(mergeTwoLists(l1, l2))
}
