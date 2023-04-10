package GoDoublyLinkedList

import (
	"fmt"
)

type list[T interface{}] struct {
	start *node[T]
	last  *node[T]
}

//type elem interface {
//	int | byte | string
//}
type node[T interface{}] struct {
	value T
	next  *node[T]
}

func (l *list[T]) isEmpty() bool {
	return l.start != nil
}
func (l *list[T]) View() {

	for s := l.start; s != nil; s = s.next {
		fmt.Println(s.value)
	}
}

func (l *list[T]) add(value T) {

	if l.start == nil {
		l.start = &node[T]{value: value}
		l.last = l.start
	} else {
		l.last.next = &node[T]{value: value}
		l.last = l.last.next
	}
}

func ConvertSliceToList[T interface{}](slice []T) *list[T] {
	l := new(list[T])
	for i := 0; i < len(slice); i++ {
		l.add(slice[i])
	}

	return l
}
