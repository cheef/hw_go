package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	MoveToBack(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	head   *ListItem
	tail   *ListItem
	length int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newHead := ListItem{Value: v}

	l.setHead(&newHead)
	l.length++

	return &newHead
}

func (l *list) PushBack(v interface{}) *ListItem {
	newTail := ListItem{Value: v}

	l.setTail(&newTail)
	l.length++

	return &newTail
}

func (l *list) Remove(i *ListItem) {
	l.unlink(i)
	l.length--
}

func (l *list) unlink(i *ListItem) {
	next := i.Next
	prev := i.Prev

	if next != nil {
		next.Prev = prev
	}

	if prev != nil {
		prev.Next = next
	}

	if l.head == i {
		l.head = next
	}

	if l.tail == i {
		l.tail = prev
	}

	i.Next = nil
	i.Prev = nil
}

func (l *list) setHead(i *ListItem) {
	i.Next = l.head
	i.Prev = nil

	if l.head != nil {
		l.head.Prev = i
	}

	l.head = i

	if l.tail == nil {
		l.tail = i
	}
}

func (l *list) setTail(i *ListItem) {
	i.Next = nil
	i.Prev = l.tail

	if l.tail != nil {
		l.tail.Next = i
	}

	l.tail = i

	if l.head == nil {
		l.head = i
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}

	l.unlink(i)
	l.setHead(i)
}

func (l *list) MoveToBack(i *ListItem) {
	if l.tail == i {
		return
	}

	l.unlink(i)
	l.setTail(i)
}

func NewList() List {
	return new(list)
}
