package goset_test

import (
	"testing"

	"github.com/jettyu/goset"
)

func TestInts(t *testing.T) {
	s := goset.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
	if !s.Equal([]int{0, 1, 2, 3, 4, 5, 6}) {
		t.Fatal(s.Items())
	}
	if !s.Has(0, 0) {
		t.Fatal(s)
	}
	if s.Has(0, 1) {
		t.Fatal(s)
	}
	if !s.Has(3, 2) {
		t.Fatal(s)
	}
	if s.Has(10, 0) {
		t.Fatal(s)
	}
	if s.Insert([]int{1, 5, 7, 8}) != 2 {
		t.Fatal(s)
	}

	if s.Erase([]int{7, 9}) != 1 {
		t.Fatal(s)
	}
	if s.Erase([]int{6, 8}) != 2 {
		t.Fatal(s)
	}
	if s.Erase([]int{0, 1}) != 2 {
		t.Fatal(s)
	}
	if !s.Equal([]int{2, 3, 4, 5}) {
		t.Fatal(s.Items())
	}

	clone := s.Clone()
	if !s.Equal(clone.Items()) {
		t.Fatal(clone.Value())
	}
	s.Erase(5)
	if s.Equal(clone.Items()) {
		t.Fatal(clone.Value(), s.Value())
	}
}

func TestUnion(t *testing.T) {
	arr1 := []int{0, 2, 4, 5}
	arr2 := []int{1, 2, 3, 5, 6}

	it3 := goset.Union(goset.Ints(arr1), goset.Ints(arr2))
	except := []int{0, 1, 2, 3, 4, 5, 6}
	if !goset.Equal(it3, goset.Ints(except)) {
		t.Fatal(it3.Value())
	}
}

func TestIntersection(t *testing.T) {
	arr1 := []int{0, 1, 1, 2, 2, 4, 5}
	arr2 := []int{1, 1, 2, 2, 3, 5, 6}
	it3 := goset.Intersection(goset.Ints(arr1), goset.Ints(arr2))
	except := []int{1, 2, 5}

	if !goset.Equal(it3, goset.Ints(except)) {
		t.Fatal(it3.Value())
	}
}

func TestDifference(t *testing.T) {
	arr1 := []int{0, 2, 4, 5}
	arr2 := []int{1, 2, 3, 5, 6}
	it3 := goset.Difference(goset.Ints(arr1), goset.Ints(arr2))
	except := []int{0, 1, 3, 4, 6}
	if it3.Len() != len(except) {
		t.Fatal(it3)
	}

	if !goset.Equal(it3, goset.Ints(except)) {
		t.Fatal(it3.Value())
	}
}
