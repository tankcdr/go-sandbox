package lists

import (
	"testing"
)

func TestNode_AddAfter(t *testing.T) {
	t.Parallel()
	aNode := Node[string]{"Apple", nil, nil}
	bNode := Node[string]{"Banana", nil, nil}

	aNode.AddAfter(&bNode)

	if aNode.Next != &bNode {
		t.Errorf("aNode.Next = %v, want %v", aNode.Next, &bNode)
	}

	if bNode.Prev != &aNode {
		t.Errorf("bNode.Prev = %v, want %v", bNode.Prev, &aNode)
	}
}

func TestNode_DeleteAfter(t *testing.T) {
	t.Parallel()
	aNode := Node[string]{"Apple", nil, nil}
	bNode := Node[string]{"Banana", nil, nil}

	aNode.AddAfter(&bNode)
	deleted := aNode.DeleteAfter()

	if deleted != &bNode {
		t.Errorf("deleted = %v, want %v", deleted, &bNode)
	}
	if aNode.Next != nil {
		t.Errorf("aNode.Next = %v, want nil", aNode.Next)
	}
}

func TestNode_DeleteAfter_Panic(t *testing.T) {
	t.Parallel()
	aNode := Node[string]{"Apple", nil, nil}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DeleteAfter did not panic")
		}
	}()

	aNode.DeleteAfter()
}

func TestDoublyLinkedList_NewDoublyLinkedList(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	if list.top_sentinel.Next != list.bottom_sentinel {
		t.Errorf("list.top_sentinel.Next = %v, want %v", list.top_sentinel.Next, list.bottom_sentinel)
	}

	if list.top_sentinel.Prev != nil {
		t.Errorf("list.top_sentinel.Prev = %v, want nil", list.top_sentinel.Prev)
	}

	if list.bottom_sentinel.Prev != list.top_sentinel {
		t.Errorf("list.bottom_sentinel.Prev = %v, want %v", list.bottom_sentinel.Prev, list.top_sentinel)
	}

	if list.bottom_sentinel.Next != nil {
		t.Errorf("list.bottom_sentinel.Next = %v, want nil", list.bottom_sentinel.Next)
	}
}

func TestDoublyLinkedList_AddList(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	other := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)
	other.AddRange(values)

	list.AddList(*other)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestDoublyLinkedList_AddList_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	other := NewDoublyLinkedList[string]()

	list.AddList(*other)

	if list.top_sentinel.Next != list.bottom_sentinel {
		t.Errorf("list.sentinel.Next = %v, want  bottom_sentinal: %v", list.top_sentinel.Next, list.bottom_sentinel)
	}
}

func TestDoublyLinkedList_AddList_WithValues(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	other := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.Add("Date")
	other.AddRange(values)

	list.AddList(*other)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}

	if list.top_sentinel.Next.Data != "Date" {
		t.Errorf("list.sentinel.Next.Data = %q, want %q", list.top_sentinel.Next.Data, "Date")
	}

	if list.Length() != 4 {
		t.Errorf("list.Length() = %d, want 4", list.Length())
	}
}

func TestDoublyLinkedList_Add(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	list.Add("Apple")

	if list.top_sentinel.Next.Data != "Apple" {
		t.Errorf("list.top_sentinel.Next.Data = %s, want Apple", list.top_sentinel.Next.Data)
	}

	if list.bottom_sentinel.Prev.Data != "Apple" {
		t.Errorf("list.bottom_sentinel.Prev.Data = %s, want Apple", list.bottom_sentinel.Prev.Data)
	}

	if list.top_sentinel.Next.Next != list.bottom_sentinel {
		t.Errorf("list.top_sentinel.Next.Next = %v, want %v", list.top_sentinel.Next.Next, list.bottom_sentinel)
	}

	if list.bottom_sentinel.Prev.Prev != list.top_sentinel {
		t.Errorf("list.bottom_sentinel.Prev.Prev = %v, want %v", list.bottom_sentinel.Prev.Prev, list.top_sentinel)
	}
}

func TestDoublyLinkedList_AddRange(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	node := list.top_sentinel.Next
	for i, value := range values {
		if node == nil {
			t.Errorf("node = nil, want &Node[string]{%q, nil}", value)
		} else if node.Data != value {
			t.Errorf("node.Data = %q, want %q", node.Data, value)
		}
		if i < len(values)-1 {
			node = node.Next
		}
	}
}

func TestDoublyLinkedList_Append(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)
	list.Append("Date")

	lastNode := list.LastNode()

	if lastNode.Data != "Date" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestDoublyLinkedList_Clone(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	clone := list.Clone()

	if clone == list {
		t.Errorf("clone = list, want different instances")
	}

	if clone.ToString(", ") != list.ToString(", ") {
		t.Errorf("clone.ToString() = %q, want %q", clone.ToString(", "), list.ToString(", "))
	}
}

func TestDoublyLinkedList_ToStringMax(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.ToStringMax(", ", 2)
	want := "Apple, Banana"

	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestDoublyLinkedList_Values(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.Values()
	for i, value := range values {
		if i < len(got) {
			if got[i] != value {
				t.Errorf("got[%d] = %q, want %q", i, got[i], value)
			}
		} else {
			t.Errorf("got[%d] = _, want %q", i, value)
		}
	}
}

func TestDoublyLinkedList_toString(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.ToString(", ")
	want := "Apple, Banana, Cherry"

	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestDoublyLinkedList_Length_NotEmpty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.Length()
	want := len(values)

	if got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}

func TestDoublyLinkedList_Length_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	got := list.Length()
	want := 0

	if got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}

func TestDoublyLinkedList_isEmpty_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	if !list.IsEmpty() {
		t.Errorf("list.isEmpty() = false, want true")
	}
}

func TestDoublyLinkedList_isEmpty_NotEmpty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	if list.IsEmpty() {
		t.Errorf("list.isEmpty() = true, want false")
	}
}

func TestDoublyLinkedList_Find(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		node := list.Find(value)
		if node == nil {
			t.Errorf("node = nil, want &Node[string]{%q, nil}", value)
		} else if node.Data != value {
			t.Errorf("node.Data = %q, want %q", node.Data, value)
		}
	}
}

func TestDoublyLinkedList_Find_NotFound(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	node := list.Find("Date")

	if node != nil {
		t.Errorf("node = %v, want nil", node)
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		if !list.Contains(value) {
			t.Errorf("list.Contains(%q) = false, want true", value)
		}
	}
}

func TestDoublyLinkedList_Contains_NotFound(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	if list.Contains("Date") {
		t.Errorf("list.Contains(%q) = true, want false", "Date")
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		removed := list.Remove(value)
		if removed == nil {
			t.Errorf("list.Remove(%q) = false, want true", value)
		}
	}
}

func TestDoublyLinkedList_Remove_NotFound(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	removed := list.Remove("Date")

	if removed != nil {
		t.Errorf("list.Remove(%q) = true, want false", "Date")
	}
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for i := 0; i < len(values); i++ {
		removed := list.RemoveAt(0)
		if removed == nil {
			t.Errorf("list.RemoveAt(%d) = false, want true", i)
		}
	}

	if list.top_sentinel.Next != list.bottom_sentinel {
		t.Errorf("list.sentinel.Next = %v, want bottom_sentinel", list.top_sentinel.Next)
	}

	if list.Length() != 0 {
		t.Errorf("list.Length() = %d, want 0", list.Length())
	}

}

func TestDoublyLinkedList_RemoveAt_OutOfRange(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	removed := list.RemoveAt(8)

	if removed != nil {
		t.Errorf("list.RemoveAt(8) = true, want false")
	}
}

func TestDoublyLinkedList_LastNode(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	list.Clear()

	if list.top_sentinel.Next != list.bottom_sentinel {
		t.Errorf("list.sentinel.Next = %v, want bottom_sentinel", list.top_sentinel.Next)
	}
}

func TestDoublyLinkedList_Clear_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	list.Clear()

	if list.top_sentinel.Next != list.bottom_sentinel {
		t.Errorf("list.sentinel.Next = %v, want bottom_sentinel", list.top_sentinel.Next)
	}
}

func TestDoublyLinkedList_ToSlice(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.ToSlice()
	for i, value := range values {
		if i < len(got) {
			if got[i].Data != value {
				t.Errorf("got[%d] = %q, want %q", i, got[i].Data, value)
			}
		} else {
			t.Errorf("got[%d] = _, want %q", i, value)
		}
	}
}

func TestDoublyLinkedList_ToSlice_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	got := list.ToSlice()

	if len(got) != 0 {
		t.Errorf("len(got) = %d, want 0", len(got))
	}
}

func TestDoublyLinkedList_ToSlice_WithValues(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)
	list.Append("Date")

	values = append(values, "Date")

	got := list.ToSlice()
	for i, value := range values {
		if i < len(got) {
			if got[i].Data != value {
				t.Errorf("got[%d] = %q, want %q", i, got[i].Data, value)
			}
		} else {
			t.Errorf("got[%d] = _, want %q", i, value)
		}
	}
}

func TestDoublyLinkedList_Enqueue(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.Enqueue(value)
	}

	for i, value := range values {
		node := list.bottom_sentinel.Prev
		for j := 0; j < i; j++ {
			node = node.Prev
		}
		if node.Data != value {
			t.Errorf("node.Data = %q, want %q", node.Data, value)
		}
	}
}

func TestDoublyLinkedList_Dequeue(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.Enqueue(value)
	}

	for _, value := range values {
		dequeued := list.Dequeue()
		if dequeued == "" {
			t.Errorf("list.Dequeue() = false, want true")
		}
		if dequeued != value {
			t.Errorf("dequeued = %q, want %q", dequeued, value)
		}
	}
}

func TestDoublyLinkedList_Dequeue_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	dequeued := list.Dequeue()

	if dequeued != "" {
		t.Errorf("dequeued = %q, want \"\"", dequeued)
	}
}

func TestDoublyLinkedList_PushBottom(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.PushBottom(value)
	}

	for i, value := range values {
		node := list.top_sentinel.Next
		for j := 0; j < i; j++ {
			node = node.Next
		}
		if node.Data != value {
			t.Errorf("node.Data = %q, want %q", node.Data, value)
		}
	}
}

func TestDoublyLinkedList_PopBottom(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.PushBottom(value)
	}

	for i := len(values) - 1; i >= 0; i-- {
		popped := list.PopBottom()
		if popped == "" {
			t.Errorf("list.PopBottom() = false, want true")
		}
		if popped != values[i] {
			t.Errorf("popped = %q, want %q", popped, values[i])
		}
	}
}

func TestDoublyLinkedList_PopBottom_Empty(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	popped := list.PopBottom()

	if popped != "" {
		t.Errorf("popped = %q, want \"\"", popped)
	}
}

func TestDoublyLinkedList_PushTop(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.PushTop(value)
	}

	if list.ToString(", ") != "Cherry, Banana, Apple" {
		t.Errorf("list: %v\n", list.ToString(", "))
	}
}

func TestDoublyLinkedList_PopTop(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	for _, value := range values {
		list.PushTop(value)
	}

	for i := len(values) - 1; i >= 0; i-- {
		popped := list.PopTop()
		if popped == "" {
			t.Errorf("list.PopTop() = false, want true")
		}
		if popped != values[i] {
			t.Errorf("popped = %q, want %q", popped, values[i])
		}
	}
}

func TestDoublyLinkedList_Pop_String(t *testing.T) {
	t.Parallel()

	list := NewDoublyLinkedList[string]()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	popped, error := list.Pop()

	if error != nil {
		t.Errorf("error = %v, want nil", error)
	}

	if popped == "" {
		t.Errorf("popped = nil, want &Node{Cherry, nil}")
	} else if popped != "Apple" {
		t.Errorf("popped.Data = %q, want %q", popped, "Apple")
	}
}

func TestDoublyLinkedList_Pop_Int(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[int]()
	values := []int{1, 2, 3}

	list.AddRange(values)

	popped, error := list.Pop()

	if error != nil {
		t.Errorf("error = %v, want nil", error)
	}

	if popped != 1 {
		t.Errorf("popped.Data = %q, want %q", popped, int(1))
	}
}

func TestDoublyLinkedList_Pop_Empty_String(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[string]()

	_, error := list.Pop()

	if error == nil {
		t.Errorf("error = nil, want error")
	}
}

func TestDoublyLinkedList_Pop_Empty_Int(t *testing.T) {
	t.Parallel()
	list := NewDoublyLinkedList[int]()

	_, error := list.Pop()

	if error == nil {
		t.Errorf("error = nil, want error")
	}
}
