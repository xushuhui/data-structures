package set

import "data-structures/linked_list"

type LinkedListSet struct {
	LinkedList *linked_list.LinkedList
}

func LinkedListSetConstructor() *LinkedListSet {
	return &LinkedListSet{
		LinkedList: linked_list.LinkedListConstructor(),
	}
}
func (this *LinkedListSet) Add(e interface{}) {
	if !this.Contains(e) {
		this.LinkedList.AddFirst(e)
	}
}
func (this *LinkedListSet) Remove(e interface{}) {
	this.LinkedList.RemoveElement(e)
}
func (this *LinkedListSet) Contains(e interface{}) bool {
	return this.LinkedList.Contains(e)
}
func (this *LinkedListSet) GetSize() int {
	return this.LinkedList.GetSize()
}
func (this *LinkedListSet) IsEmpty() bool {
	return this.LinkedList.IsEmpty()
}
