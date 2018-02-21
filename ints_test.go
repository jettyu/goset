package goset_test

import (
	"testing"

	"github.com/jettyu/goset"
)

func TestInts(t *testing.T) {
	s := goset.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
	if !s.Equal(goset.IntSlice([]int{0, 1, 2, 3, 4, 5, 6})) {
		t.Fatal(s.Items())
	}
	if !s.Has(goset.IntElement(0), 0) {
		t.Fatal(s)
	}
	if s.Has(goset.IntElement(0), 1) {
		t.Fatal(s)
	}
	if !s.Has(goset.IntElement(3), 2) {
		t.Fatal(s)
	}
	if s.Has(goset.IntElement(10), 0) {
		t.Fatal(s)
	}
	if s.Insert(goset.IntSlice([]int{1, 5, 7, 8})) != 2 {
		t.Fatal(s)
	}
	if s.Erase(goset.IntSlice([]int{7, 9})) != 1 {
		t.Fatal(s)
	}
	if s.Erase(goset.IntSlice([]int{6, 8})) != 2 {
		t.Fatal(s)
	}
	if s.Erase(goset.IntSlice([]int{0, 1})) != 2 {
		t.Fatal(s)
	}
	if !s.Equal(goset.IntSlice([]int{2, 3, 4, 5})) {
		t.Fatal(s.Items())
	}
}

func TestUnion(t *testing.T) {
	arr1 := []int{0, 2, 4, 5}
	arr2 := []int{1, 2, 3, 5, 6}

	it3 := goset.Union(goset.IntSlice(arr1), goset.IntSlice(arr2))
	except := []int{0, 1, 2, 3, 4, 5, 6}
	if it3.Len() != len(except) {
		t.Fatal(it3)
	}

	for i, v := range except {
		if int(it3.Elem(i).(goset.IntElement)) != v {
			t.Fatal(except, it3)
		}
	}
}

func TestIntersection(t *testing.T) {
	arr1 := []int{0, 2, 4, 5}
	arr2 := []int{1, 2, 3, 5, 6}
	it3 := goset.Intersection(goset.IntSlice(arr1), goset.IntSlice(arr2))
	except := []int{2, 5}
	if it3.Len() != len(except) {
		t.Fatal(it3)
	}

	for i, v := range except {
		if int(it3.Elem(i).(goset.IntElement)) != v {
			t.Fatal(except, it3)
		}
	}
}

func TestDifference(t *testing.T) {
	arr1 := []int{0, 2, 4, 5}
	arr2 := []int{1, 2, 3, 5, 6}
	it3 := goset.Difference(goset.IntSlice(arr1), goset.IntSlice(arr2))
	except := []int{0, 1, 3, 4, 6}
	if it3.Len() != len(except) {
		t.Fatal(it3)
	}

	for i, v := range except {
		if int(it3.Elem(i).(goset.IntElement)) != v {
			t.Fatal(except, it3)
		}
	}
}
