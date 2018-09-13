package main

import (
	"testing"
)

func TestDeleteElement(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 0}
	r, _ := deleteElement(a, 5)
	e := []int{1, 2, 3, 4, 5}
	if len(r) != len(e) {
		t.Errorf("incorrect, got: %d, want: %d.", r, e)
	}

	for i := range r {
		if r[i] != e[i] {
			t.Errorf("incorrect, got: %d, want: %d.", r, e)
		}
	}
}
func TestDeleteElementWithInvalidIndex(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 0}
	r, _ := deleteElement(a, 10)
	e := []int{1, 2, 3, 4, 5, 0}
	if len(r) != len(e) {
		t.Errorf("incorrect, got: %d, want: %d.", r, e)
	}

	for i := range r {
		if r[i] != e[i] {
			t.Errorf("incorrect, got: %d, want: %d.", r, e)
		}
	}
}

func TestSorting1(t *testing.T) {
	a := []int{1, 0, -1, 1, 4, 5}
	r := sortAndRemoveDuplicate(a)
	e := []int{0, 1, 4, 5}

	if len(r) != len(e) {
		t.Errorf("incorrect, got: %d, want: %d.", r, e)
	}

	for i := range r {
		if r[i] != e[i] {
			t.Errorf("incorrect, got: %d, want: %d.", r, e)
		}
	}
}
func TestSorting2(t *testing.T) {
	// covers i == 0 && !(s[i] < 0)
	a := []int{1, 0, 1, 4, 5}
	r := sortAndRemoveDuplicate(a)
	e := []int{0, 1, 4, 5}

	if len(r) != len(e) {
		t.Errorf("incorrect, got: %d, want: %d.", r, e)
	}

	for i := range r {
		if r[i] != e[i] {
			t.Errorf("incorrect, got: %d, want: %d.", r, e)
		}
	}
}

func TestDeleteElements(t *testing.T) {
	a := []int{1, 5, 2, 0, 1, 4}
	p := []int{-1, 0, 10, 2}
	e := []int{5, 0, 1, 4}
	r := deleteElements(a, p)
	if len(r) != len(e) {
		t.Errorf("incorrect, got: %d, want: %d.", r, e)
	}

	for i := range r {
		if r[i] != e[i] {
			t.Errorf("incorrect, got: %d, want: %d.", r, e)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
