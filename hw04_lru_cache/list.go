package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	front *listItem
	back  *listItem
	len   int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *listItem {
	if l.front == nil && l.len == 1 {
		l.front = l.back
	}
	return l.front
}

func (l list) Back() *listItem {
	if l.back == nil && l.len == 1 {
		l.back = l.front
	}
	return l.back
}

func (l *list) PushBack(v interface{}) *listItem {
	newItem := listItem{Value: v}
	if l.back != nil {
		l.back.Prev = &newItem
		newItem.Next = l.back
	} else if l.front != nil {
		l.front.Prev = &newItem
		newItem.Next = l.front
	}
	if l.front == nil {
		l.front = l.back
	}
	l.back = &newItem
	l.len++
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	newItem := listItem{Value: v}
	if l.front != nil {
		l.front.Next = &newItem
		newItem.Prev = l.front
	} else if l.back != nil {
		l.back.Next = &newItem
		newItem.Prev = l.back
	}
	if l.back == nil {
		l.back = l.front
	}
	l.front = &newItem
	l.len++
	return l.front
}

func (l *list) Remove(i *listItem) {
	if l.back == i {
		l.back = i.Next
	}
	if l.front == i {
		l.front = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Prev = nil
	i.Next = nil
	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	if l.front != nil {
		i.Prev = l.front
		l.front.Next = i
	}
	l.front = i
	l.len++
}

func NewList() List {
	return &list{}
}
