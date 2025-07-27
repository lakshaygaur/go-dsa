package ds

import "fmt"

type MyLinkedList struct {
	next *MyLinkedList
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	temp := this
	for i := 0; i < index && temp.next != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		return -1
	} else {
		return temp.val
	}
}

func (this *MyLinkedList) Print() {
	temp := this

	for temp != nil {
		fmt.Print(temp.val, ",")
		temp = temp.next
	}
}

func (this *MyLinkedList) AddAtHead(val int) {

	// treat this as head
	if this.next == nil {
		this.val = val
	} else {
		newHead := &MyLinkedList{}
		newHead.val = val
		newHead.next = this.next // points the next to head's next
		this.next = newHead      // points the head to current new
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	temp := this
	newtail := &MyLinkedList{}
	newtail.val = val

	for temp.next != nil { // reach last
		temp = temp.next
	}
	temp.next = newtail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	temp := this
	newMem := &MyLinkedList{}
	newMem.val = val
	prev := temp

	for i := 0; i <= index && temp.next != nil; i++ {
		prev = temp
		temp = temp.next
	}
	prev.next = newMem
	newMem.next = temp
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	temp := this
	prev := temp
	for i := 0; i < index && temp.next != nil; i++ {
		prev = temp
		temp = temp.next
	}
	prev.next = temp.next
}
