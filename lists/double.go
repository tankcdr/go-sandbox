package lists

import (
	"strings"
)

/***************************************************
 * Node definition and operations
 ***************************************************/
type Node struct {
	Data string
	Next *Node
	Prev *Node
}

func (me *Node) AddAfter(node *Node) {
	if me.Next != nil {
		me.Next.Prev = node
	}
	node.Next = me.Next
	node.Prev = me
	me.Next = node
}

func (me *Node) DeleteAfter() *Node {
	if me.Next == nil {
		panic("Attempt to delete nonexistent node")
	}

	after := me.Next

	if after.Next != nil {
		after.Next.Prev = me
	}

	me.Next = after.Next

	after.Next, after.Prev = nil, nil
	return after
}

/***************************************************
 * DoublyLinkedList definition and operations
 ***************************************************/
type DoublyLinkedList struct {
	top_sentinel    *Node
	bottom_sentinel *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	list := &DoublyLinkedList{&Node{}, &Node{}}
	list.top_sentinel.Next = list.bottom_sentinel
	list.bottom_sentinel.Prev = list.top_sentinel
	return list
}

func (l *DoublyLinkedList) Add(value string) {
	l.top_sentinel.AddAfter(&Node{value, nil, nil})
}

func (l *DoublyLinkedList) AddList(list DoublyLinkedList) {
	if list.Length() == 0 {
		return
	}

	lastNode := l.LastNode()

	for node := list.top_sentinel.Next; node != list.bottom_sentinel; node = node.Next {
		lastNode.AddAfter(&Node{node.Data, nil, nil})
		lastNode = lastNode.Next
	}
}

func (l *DoublyLinkedList) AddRange(values []string) {
	lastNode := l.LastNode()

	for _, value := range values {
		lastNode.AddAfter(&Node{value, nil, nil})
		lastNode = lastNode.Next
	}
}

func (l *DoublyLinkedList) Append(value string) {
	lastNode := l.LastNode()
	lastNode.AddAfter(&Node{value, nil, nil})
}

func (l *DoublyLinkedList) Clear() {
	l.top_sentinel.Next = l.bottom_sentinel
	l.top_sentinel.Prev = nil
	l.bottom_sentinel.Prev = l.top_sentinel
	l.bottom_sentinel.Next = nil
}

func (l *DoublyLinkedList) Clone() *DoublyLinkedList {
	clone := NewDoublyLinkedList()
	clone.AddList(*l)
	return clone
}

func (l *DoublyLinkedList) Contains(value string) bool {
	node := l.Find(value)
	return node != nil
}

func (l *DoublyLinkedList) Find(value string) *Node {
	for node := l.top_sentinel.Next; node != nil; node = node.Next {
		if node.Data == value {
			return node
		}
	}
	return nil
}

func (l *DoublyLinkedList) IsEmpty() bool {
	if l.top_sentinel.Next != l.bottom_sentinel {
		return false
	}
	return true
}

func (l *DoublyLinkedList) LastNode() *Node {
	return l.bottom_sentinel.Prev
}

func (l *DoublyLinkedList) Length() int {
	//likely better to keep a running count
	//TODO: implement a counter in the DoublyLinkedList struct
	var length int
	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		length++
	}
	return length
}

func (l *DoublyLinkedList) Remove(value string) *Node {
	var prev *Node
	for node := l.top_sentinel; node.Next != l.bottom_sentinel; node = node.Next {
		if node.Next.Data == value {
			prev = node
			break
		}
	}
	if prev != nil {
		return prev.DeleteAfter()
	}
	return nil
}

func (l *DoublyLinkedList) RemoveAt(index int) *Node {
	var prev *Node
	for i, node := 0, l.top_sentinel; node.Next != l.bottom_sentinel; i, node = i+1, node.Next {
		if i == index {
			prev = node
			break
		}
	}
	if prev != nil {
		return prev.DeleteAfter()
	}
	return nil
}

func (l *DoublyLinkedList) ToSlice() []Node {
	var slice []Node

	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		slice = append(slice, *node)
	}
	return slice
}

func (l *DoublyLinkedList) ToString(separator string) string {
	var builder strings.Builder

	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		if node != l.top_sentinel.Next {
			builder.WriteString(separator)
		}
		builder.WriteString(node.Data)
	}
	return builder.String()
}

func (l *DoublyLinkedList) ToStringMax(separator string, max int) string {
	var builder strings.Builder

	i := 0
	for node := l.top_sentinel.Next; node != l.bottom_sentinel && i < max; node = node.Next {
		if i > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(node.Data)
		i++
	}
	return builder.String()
}

func (l *DoublyLinkedList) Values() []string {
	var values []string
	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		values = append(values, node.Data)
	}
	return values
}

/***************************************************
 * Queue operations on the DoublyLinkedList
 ***************************************************/
func (l *DoublyLinkedList) Dequeue() string {
	if l.IsEmpty() {
		return ""
	}

	dequeued := l.LastNode()
	return dequeued.Prev.DeleteAfter().Data
}

func (l *DoublyLinkedList) Enqueue(value string) {
	l.Add(value)
}
