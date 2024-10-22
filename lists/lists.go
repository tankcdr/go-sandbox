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

func (l *LinkedList) Add(value string) {
	l.sentinel.AddAfter(&Cell{value, nil})
}

func (l *LinkedList) AddList(list LinkedList) {
	lastCell := l.LastNode()
	//easy way
	//lastCell.Next = list.sentinel.Next
	//but want my own copies
	for cell := list.sentinel.Next; cell != nil; cell = cell.Next {
		lastCell.AddAfter(&Cell{cell.Data, nil})
		lastCell = lastCell.Next
	}
}

func (l *LinkedList) AddRange(values []string) {
	lastCell := l.LastNode()

	for _, value := range values {
		lastCell.AddAfter(&Cell{value, nil})
		lastCell = lastCell.Next
	}
}

func (l *LinkedList) Append(value string) {
	lastCell := l.LastNode()
	lastCell.AddAfter(&Cell{value, nil})
}

func (l *LinkedList) Clear() {
	l.sentinel.Next = nil
}

func (l *LinkedList) Clone() *LinkedList {
	clone := NewLinkedList()
	clone.AddList(*l)
	return clone
}

func (l *LinkedList) Contains(value string) bool {
	cell := l.Find(value)
	return cell != nil
}

func (l *LinkedList) Find(value string) *Cell {
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		if cell.Data == value {
			return cell
		}
	}
	return nil
}

func (l *LinkedList) IsEmpty() bool {
	if l.sentinel.Next != nil {
		return false
	}
	return true
}

func (l *LinkedList) LastNode() *Cell {
	var lastCell *Cell

	for lastCell = l.sentinel; lastCell.Next != nil; lastCell = lastCell.Next {
	}

	return lastCell
}

func (l *LinkedList) Length() int {
	//likely better to keep a running count
	//TODO: implement a counter in the LinkedList struct
	var length int
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		length++
	}
	return length
}

func (l *LinkedList) Remove(value string) *Cell {
	var prev *Cell
	for cell := l.sentinel; cell.Next != nil; cell = cell.Next {
		if cell.Next.Data == value {
			prev = cell
			break
		}
	}
	if prev != nil {
		return prev.DeleteAfter()
	}
	return nil
}

func (l *LinkedList) RemoveAt(index int) *Cell {
	var prev *Cell
	for i, cell := 0, l.sentinel; cell.Next != nil; i, cell = i+1, cell.Next {
		if i == index {
			prev = cell
			break
		}
	}
	if prev != nil {
		return prev.DeleteAfter()
	}
	return nil
}

func (l *LinkedList) ToSlice() []Cell {
	var slice []Cell

	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		slice = append(slice, *cell)
	}
	return slice
}

func (l *LinkedList) ToString(separator string) string {
	var builder strings.Builder

	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		if cell != l.sentinel.Next {
			builder.WriteString(separator)
		}
		builder.WriteString(cell.Data)
	}
	return builder.String()
}

func (l *LinkedList) ToStringMax(separator string, max int) string {
	var builder strings.Builder

	i := 0
	for cell := l.sentinel.Next; cell != nil && i < max; cell = cell.Next {
		if i > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(cell.Data)
		i++
	}
	return builder.String()
}

func (l *LinkedList) Values() []string {
	var values []string
	for cell := l.sentinel.Next; cell != nil; cell = cell.Next {
		values = append(values, cell.Data)
	}
	return values
}

/***************************************************
 * Stack operations on LinkedList
 ***************************************************/
func (l *LinkedList) Push(value string) {
	l.sentinel.AddAfter(&Cell{value, l.sentinel.Next})
}

func (l *LinkedList) Pop() string {
	if l.IsEmpty() {
		return ""
	}
	cell := l.sentinel.DeleteAfter()
	if cell == nil {
		return ""
	}
	return cell.Data
}

/***************************************************
 * Cycle Detection and Handling operations on LinkedList
 ***************************************************/
func (l *LinkedList) HasLoop() bool {
	//implements Floyd's cycle-finding algorithm
	slow := l.sentinel
	fast := l.sentinel

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
