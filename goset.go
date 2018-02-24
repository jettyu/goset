package goset

import (
	"sort"
)

// Element ...
type Element interface {
	Less(v Element) bool
	Equal(v Element) bool
}

// Items ..
type Items interface {
	sort.Interface
	Elem(i int) Element
	SetElem(v Element, pos int)
	Move(dstPos, srcPos, n int)
	Append(e ...Element) Items
	Truncate(n int) Items
}

// Set ...
type Set interface {
	Has(v interface{}, pos int) bool
	Insert(v Items) int
	Erase(v Items) int
	Items() Items
	Value() interface{}
	Search(v interface{}, pos int) int
	Equal(Items) bool
	Get(v interface{}) (data interface{}, ok bool)
}

// Equal ...
func Equal(it1, it2 Items) bool {
	rit, ok := it1.(reflectItems)
	if ok {
		return rit.equal(it2)
	}
	if it1.Len() != it2.Len() {
		return false
	}
	for i := 0; i < it1.Len(); i++ {
		if !it1.Elem(i).Equal(it2.Elem(i)) {
			return false
		}
	}
	return true
}

// Union ...
func Union(it1, it2 Items) Items {
	s := NewSet(it1)
	s.Insert(it2)
	return s.Items()
}

// Intersection ...
func Intersection(it1, it2 Items) (dst Items) {
	rit, ok := it1.(reflectItems)
	if ok {
		return rit.intersection(it2)
	}
	dst = it1.Truncate(0)
	if it1.Len() == 0 || it2.Len() == 0 {
		return
	}

	s1 := NewSet(it1)
	s2 := NewSet(it2)
	it1 = s1.Items()
	it2 = s2.Items()
	pos := 0
	for i := 0; i < it2.Len() && pos < it1.Len(); i++ {
		e := it2.Elem(i)
		pos += s1.Search(e, pos)

		if pos == it1.Len() {
			continue
		}
		v := it1.Elem(pos)
		if v.Equal(e) {
			dst = dst.Append(v)
		}
	}

	return
}

// Difference ...
func Difference(it1, it2 Items) (dst Items) {
	s := NewSet(Union(it1, it2))
	s.Erase(Intersection(it1, it2))
	return s.Items()
}
