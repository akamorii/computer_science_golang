package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linked_list struct {
	head *Node
}

func (ll *Linked_list) Append(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll Linked_list) Show_arr() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll Linked_list) Take_element(index int) int {
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data
}
