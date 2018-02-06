package goset_test

import (
	"log"
	"testing"

	"github.com/jettyu/goset"
)

func TestInts(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	s := goset.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
	arr, ok := s.Items().(goset.IntSlice)
	if !ok {
		t.Fatal(s)
	}
	if len(arr) != 7 {
		t.Fatal(s, arr)
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
	if s.Insert(goset.IntSlice{1, 5, 7, 8}) != 2 {
		t.Fatal(s)
	}
	// 删除中间，末尾混淆
	if s.Erase(goset.IntSlice{7, 9}) != 1 {
		t.Fatal(s)
	}
	// 删除中间和末尾
	if s.Erase(goset.IntSlice{6, 8}) != 2 {
		t.Fatal(s)
	}
	// 删除开头
	if s.Erase(goset.IntSlice{0, 1}) != 2 {
		t.Fatal(s)
	}
	for i, v := range s.Items().(goset.IntSlice) {
		if i+2 != v {
			t.Fatal(arr)
		}
	}
}

func TestUnion(t *testing.T) {
	it1 := goset.IntSlice{0, 2, 4, 5}
	it2 := goset.IntSlice{1, 2, 3, 5, 6}
	it3 := goset.Union(it1, it2)
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
	it1 := goset.IntSlice{0, 2, 4, 5}
	it2 := goset.IntSlice{1, 2, 3, 5, 6}
	it3 := goset.Intersection(it1, it2)
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
	it1 := goset.IntSlice{0, 2, 4, 5}
	it2 := goset.IntSlice{1, 2, 3, 5, 6}
	it3 := goset.Difference(it1, it2)
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
