package list

import "strconv"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
}

func (list *List) Push(x int) {
	newNode := &Node{next: list.head, data: x}
	list.head = newNode
}

func (list *List) Pop() (x int) {
	if list.head == nil {
		panic("Popping empty List")
	}

	x = list.head.data
	list.head = list.head.next
	return
}

func (list *List) String() string {
	str := "{"
	for cur := list.head; cur != nil; cur = cur.next {
		str += strconv.Itoa(cur.data)
		if cur.next != nil {
			str += " "
		}
	}
	str += "}"
	return str
}
