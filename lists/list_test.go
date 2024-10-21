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

	got := list.toString(", ")
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

	if !list.isEmpty() {
		t.Errorf("list.isEmpty() = false, want true")
	}
}

func TestLinkedList_isEmpty_NotEmpty(t *testing.T) {
	t.Parallel()
	list := NewLinkedList()
	values := []string{"Apple", "Banana", "Cherry"}

	list.AddRange(values)

	if list.isEmpty() {
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
