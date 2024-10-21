package lists

import (
	"strings"
)

/***************************************************
 * Cell definition and operations
 ***************************************************/
type Cell struct {
	Data string
	Next *Cell
}

func (me *Cell) AddAfter(cell *Cell) {
	cell.Next = me.Next
	me.Next = cell
}

func (me *Cell) DeleteAfter() *Cell {
	if me.Next == nil {
		panic("Attempt to delete nonexistent cell")
	}
	after := me.Next
	me.Next = me.Next.Next
	return after
}

/***************************************************
 * LinkedList definition and operations
 ***************************************************/
type LinkedList struct {
	sentinel *Cell
}

func NewLinkedList() *LinkedList {
	return &LinkedList{&Cell{}}
}

func (l *LinkedList) AddRange(values []string) {
	var lastCell *Cell

	for lastCell = l.sentinel; lastCell.Next != nil; lastCell = lastCell.Next {
	}

	for _, value := range values {
		lastCell.AddAfter(&Cell{value, nil})
		lastCell = lastCell.Next
	}
}

func (l *LinkedList) Values() []string {
	var values []string
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		values = append(values, cell.Data)
	}
	return values
}

func (l *LinkedList) toString(separator string) string {
	var builder strings.Builder

	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		builder.WriteString(cell.Data)
		if cell.Next != nil {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

func (l *LinkedList) Length() int {
	var length int
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		length++
	}
	return length
}

func (l *LinkedList) isEmpty() bool {
	if l.sentinel.Next != nil {
		return false
	}
	return true
}

func (l *LinkedList) Find(value string) *Cell {
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		if cell.Data == value {
			return cell
		}
	}
	return nil
}

func (l *LinkedList) Contains(value string) bool {
	cell := l.Find(value)
	return cell != nil
}

func (l *LinkedList) Clear() {
	l.sentinel.Next = nil
}
