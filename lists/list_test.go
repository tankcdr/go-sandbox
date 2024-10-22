package lists

import (
	"testing"
)

func TestCell_AddAfter(t *testing.T) {
	t.Parallel()
	aCell := Cell{"Apple", nil}
	bCell := Cell{"Banana", nil}

	aCell.AddAfter(&bCell)

	if aCell.Next != &bCell {
		t.Errorf("aCell.Next = %v, want %v", aCell.Next, &bCell)
	}
}

func TestCell_DeleteAfter(t *testing.T) {
	t.Parallel()
	aCell := Cell{"Apple", nil}
	bCell := Cell{"Banana", nil}

	aCell.AddAfter(&bCell)
	deleted := aCell.DeleteAfter()

	if deleted != &bCell {
		t.Errorf("deleted = %v, want %v", deleted, &bCell)
	}
	if aCell.Next != nil {
		t.Errorf("aCell.Next = %v, want nil", aCell.Next)
	}
}

func TestCell_DeleteAfter_Panic(t *testing.T) {
	t.Parallel()
	aCell := Cell{"Apple", nil}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DeleteAfter did not panic")
		}
	}()

	aCell.DeleteAfter()
}

func TestLinkedList_NewLinkedList(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	if list.sentinel.Next != nil {
		t.Errorf("list.sentinel.Next = %v, want nil", list.sentinel.Next)
	}
}

func TestLinkedList_AddRange(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	cell := list.sentinel.Next
	for i, value := range values {
		if cell == nil {
			t.Errorf("cell = nil, want &Cell{%q, nil}", value)
		} else if cell.Data != value {
			t.Errorf("cell.Data = %q, want %q", cell.Data, value)
		}
		if i < len(values)-1 {
			cell = cell.Next
		}
	}
}

func TestLinkedList_Values(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
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

func TestLinkedList_toString(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.ToString(", ")
	want := "Apple, Banana, Cherry"

	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestLinkedList_Length_NotEmpty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	got := list.Length()
	want := len(values)

	if got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}

func TestLinkedList_Length_Empty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	got := list.Length()
	want := 0

	if got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}

func TestLinkedList_isEmpty_Empty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	if !list.IsEmpty() {
		t.Errorf("list.isEmpty() = false, want true")
	}
}

func TestLinkedList_isEmpty_NotEmpty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	if list.IsEmpty() {
		t.Errorf("list.isEmpty() = true, want false")
	}
}

func TestLinkedList_Find(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		cell := list.Find(value)
		if cell == nil {
			t.Errorf("cell = nil, want &Cell{%q, nil}", value)
		} else if cell.Data != value {
			t.Errorf("cell.Data = %q, want %q", cell.Data, value)
		}
	}
}

func TestLinkedList_Find_NotFound(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	cell := list.Find("Date")

	if cell != nil {
		t.Errorf("cell = %v, want nil", cell)
	}
}

func TestLinkedList_Contains(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		if !list.Contains(value) {
			t.Errorf("list.Contains(%q) = false, want true", value)
		}
	}
}

func TestLinkedList_Contains_NotFound(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	if list.Contains("Date") {
		t.Errorf("list.Contains(%q) = true, want false", "Date")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for _, value := range values {
		removed := list.Remove(value)
		if removed == nil {
			t.Errorf("list.Remove(%q) = false, want true", value)
		}
	}
}

func TestLinkedList_Remove_NotFound(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	removed := list.Remove("Date")

	if removed != nil {
		t.Errorf("list.Remove(%q) = true, want false", "Date")
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	for i := 0; i < len(values); i++ {
		removed := list.RemoveAt(0)
		if removed == nil {
			t.Errorf("list.RemoveAt(%d) = false, want true", i)
		}
	}

	if list.sentinel.Next != nil {
		t.Errorf("list.sentinel.Next = %v, want nil", list.sentinel.Next)
	}

	if list.Length() != 0 {
		t.Errorf("list.Length() = %d, want 0", list.Length())
	}

}

func TestLinkedList_RemoveAt_OutOfRange(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	removed := list.RemoveAt(8)

	if removed != nil {
		t.Errorf("list.RemoveAt(8) = true, want false")
	}
}

func TestLinkedList_LastNode(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestLinkedList_Append(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)
	list.Append("Date")

	lastNode := list.LastNode()

	if lastNode.Data != "Date" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestLinkedList_Add(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	list.Add("Apple")

	if list.sentinel.Next == nil {
		t.Errorf("list.sentinel.Next = nil, want &Cell{Apple, nil}")
	} else if list.sentinel.Next.Data != "Apple" {
		t.Errorf("list.sentinel.Next.Data = %q, want %q", list.sentinel.Next.Data, "Apple")
	}
}

func TestLinkedList_AddList(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	other := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)
	other.AddRange(values)

	list.AddList(*other)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}
}

func TestLinkedList_AddList_Empty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	other := NewLinkedList()

	list.AddList(*other)

	if list.sentinel.Next != nil {
		t.Errorf("list.sentinel.Next = %v, want nil", list.sentinel.Next)
	}
}

func TestLinkedList_AddList_WithValues(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	other := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.Add("Date")
	other.AddRange(values)

	list.AddList(*other)

	lastNode := list.LastNode()

	if lastNode.Data != "Cherry" {
		t.Errorf("lastNode.Data = %q, want %q", lastNode.Data, "Date")
	}

	if list.sentinel.Next.Data != "Date" {
		t.Errorf("list.sentinel.Next.Data = %q, want %q", list.sentinel.Next.Data, "Date")
	}

	if list.Length() != 4 {
		t.Errorf("list.Length() = %d, want 4", list.Length())
	}
}

func TestLinkedList_Clear(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	list.Clear()

	if list.sentinel.Next != nil {
		t.Errorf("list.sentinel.Next = %v, want nil", list.sentinel.Next)
	}
}

func TestLinkedList_Clear_Empty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	list.Clear()

	if list.sentinel.Next != nil {
		t.Errorf("list.sentinel.Next = %v, want nil", list.sentinel.Next)
	}
}

func TestLinkedList_ToSlice(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
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

func TestLinkedList_ToSlice_Empty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()

	got := list.ToSlice()

	if len(got) != 0 {
		t.Errorf("len(got) = %d, want 0", len(got))
	}
}

func TestLinkedList_ToSlice_WithValues(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
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

func TestLinkedList_Clone(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
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

func TestLinkedList_Push(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	list.Push("Date")

	if list.sentinel.Next.Data != "Date" {
		t.Errorf("list.sentinel.Next.Data = %q, want %q", list.sentinel.Next.Data, "Date")
	}
}

func TestLinkedList_Pop(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	popped := list.Pop()

	if popped == "" {
		t.Errorf("popped = nil, want &Cell{Cherry, nil}")
	} else if popped != "Apple" {
		t.Errorf("popped.Data = %q, want %q", popped, "Apple")
	}
}
