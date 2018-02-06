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
	Has(v Element, pos int) bool
	Insert(v Items) int
	Erase(v Items) int
	Items() Items
	Search(v Element, pos int) int
}

// Union ...
func Union(it1, it2 Items) Items {
	s := NewSet(it1)
	s.Insert(it2)
	return s.Items()
}

// Intersection ...
func Intersection(it1, it2 Items) (dst Items) {
	dst = it1.Truncate(0)
	if it1.Len() == 0 || it2.Len() == 0 {
		return
	}

	s1 := NewSet(it1)
	s2 := NewSet(it2)
	it1 = s1.Items()
	it2 = s2.Items()
	pos := 0
	for i := 0; i < it2.Len(); i++ {
		v := it2.Elem(i)
		pos += s1.Search(v, pos)
		if pos == it1.Len() || !it1.Elem(pos).Equal(v) {
			continue
		}
		dst = dst.Append(v)
	}

	return
}

// Difference ...
func Difference(it1, it2 Items) (dst Items) {
	s := NewSet(Union(it1, it2))
	s.Erase(Intersection(it1, it2))
	return s.Items()
}
