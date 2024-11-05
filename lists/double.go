package lists

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/***************************************************
 * Node definition and operations
 ***************************************************/
type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

func (me *Node[T]) AddAfter(node *Node[T]) {
	if me.Next != nil {
		me.Next.Prev = node
	}
	node.Next = me.Next
	node.Prev = me
	me.Next = node
}

func (me *Node[T]) DeleteAfter() *Node[T] {
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
type DoublyLinkedList[T any] struct {
	top_sentinel    *Node[T]
	bottom_sentinel *Node[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{&Node[T]{}, &Node[T]{}}
	list.top_sentinel.Next = list.bottom_sentinel
	list.bottom_sentinel.Prev = list.top_sentinel
	return list
}

func (l *DoublyLinkedList[T]) Add(value T) {
	l.top_sentinel.AddAfter(&Node[T]{value, nil, nil})
}

func (l *DoublyLinkedList[T]) AddList(list DoublyLinkedList[T]) {
	if list.Length() == 0 {
		return
	}

	lastNode := l.LastNode()

	for node := list.top_sentinel.Next; node != list.bottom_sentinel; node = node.Next {
		lastNode.AddAfter(&Node[T]{node.Data, nil, nil})
		lastNode = lastNode.Next
	}
}

func (l *DoublyLinkedList[T]) AddRange(values []T) {
	lastNode := l.LastNode()

	for _, value := range values {
		lastNode.AddAfter(&Node[T]{value, nil, nil})
		lastNode = lastNode.Next
	}
}

func (l *DoublyLinkedList[T]) Append(value T) {
	lastNode := l.LastNode()
	lastNode.AddAfter(&Node[T]{value, nil, nil})
}

func (l *DoublyLinkedList[T]) Clear() {
	l.top_sentinel.Next = l.bottom_sentinel
	l.top_sentinel.Prev = nil
	l.bottom_sentinel.Prev = l.top_sentinel
	l.bottom_sentinel.Next = nil
}

func (l *DoublyLinkedList[T]) Clone() *DoublyLinkedList[T] {
	clone := NewDoublyLinkedList[T]()
	clone.AddList(*l)
	return clone
}

func (l *DoublyLinkedList[T]) Contains(value T) bool {
	node := l.Find(value)
	return node != nil
}

func (l *DoublyLinkedList[T]) Find(value T) *Node[T] {
	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		if reflect.DeepEqual(node.Data, value) {
			return node
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	if l.top_sentinel.Next != l.bottom_sentinel {
		return false
	}
	return true
}

func (l *DoublyLinkedList[T]) LastNode() *Node[T] {
	return l.bottom_sentinel.Prev
}

func (l *DoublyLinkedList[T]) Length() int {
	//likely better to keep a running count
	//TODO: implement a counter in the DoublyLinkedList struct
	var length int
	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		length++
	}
	return length
}

func (l *DoublyLinkedList[T]) Remove(value T) *Node[T] {
	var prev *Node[T]
	for node := l.top_sentinel; node.Next != l.bottom_sentinel; node = node.Next {
		if reflect.DeepEqual(node.Next.Data, value) {
			prev = node
			break
		}
	}
	if prev != nil {
		return prev.DeleteAfter()
	}
	return nil
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) *Node[T] {
	var prev *Node[T]
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

func (l *DoublyLinkedList[T]) ToSlice() []Node[T] {
	var slice []Node[T]

	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		slice = append(slice, *node)
	}
	return slice
}

func (l *DoublyLinkedList[T]) ToString(separator string) string {
	var builder strings.Builder

	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		if node != l.top_sentinel.Next {
			builder.WriteString(separator)
		}
		builder.WriteString(fmt.Sprintf("%v", node.Data))
	}
	return builder.String()
}

func (l *DoublyLinkedList[T]) ToStringMax(separator string, max int) string {
	var builder strings.Builder

	i := 0
	for node := l.top_sentinel.Next; node != l.bottom_sentinel && i < max; node = node.Next {
		if i > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(fmt.Sprintf("%v", node.Data))
		i++
	}
	return builder.String()
}

func (l *DoublyLinkedList[T]) Values() []T {
	var values []T
	for node := l.top_sentinel.Next; node != l.bottom_sentinel; node = node.Next {
		values = append(values, node.Data)
	}
	return values
}

/***************************************************
 * Queue operations on the DoublyLinkedList
 ***************************************************/
func (queue *DoublyLinkedList[T]) Dequeue() T {
	if queue.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	dequeued := queue.LastNode()
	return dequeued.Prev.DeleteAfter().Data
}

func (queue *DoublyLinkedList[T]) Enqueue(value T) {
	queue.Add(value)
}

/***************************************************
 * Deque operations on the DoublyLinkedList
 ***************************************************/
func (deque *DoublyLinkedList[T]) PushBottom(value T) {
	deque.bottom_sentinel.Prev.AddAfter(&Node[T]{value, nil, nil})
}

func (deque *DoublyLinkedList[T]) PushTop(value T) {
	deque.top_sentinel.AddAfter(&Node[T]{value, nil, nil})
}

func (deque *DoublyLinkedList[T]) PopBottom() T {
	if deque.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	popped := deque.bottom_sentinel.Prev
	return popped.Prev.DeleteAfter().Data
}

func (deque *DoublyLinkedList[T]) PopTop() T {
	if deque.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	return deque.top_sentinel.DeleteAfter().Data
}

/***************************************************
 * Stack operations on DoublyLinkedList
 ***************************************************/
func (l *DoublyLinkedList[T]) Push(value T) {
	l.top_sentinel.AddAfter(&Node[T]{value, nil, nil})
}

func (l *DoublyLinkedList[T]) Pop() (T, error) {
	var result T
	if l.IsEmpty() {
		return result, errors.New("Empty Stack")
	}
	cell := l.top_sentinel.DeleteAfter()
	if cell == nil {
		return result, errors.New("Empty Stack")
	}
	return cell.Data, nil
}
