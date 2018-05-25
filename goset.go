package goset

import (
	"reflect"
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

// ItemsCloner ...
type ItemsCloner interface {
	Clone() Items
}

// ItemsClone ...
func ItemsClone(items Items) Items {
	cloner, ok := items.(ItemsCloner)
	if ok {
		return cloner.Clone()
	}
	v := reflect.ValueOf(items)
	rv := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	reflect.Copy(rv, v)
	return rv.Interface().(Items)
}

// Set ...
type Set interface {
	Has(v interface{}, pos int) bool
	Insert(v ...interface{}) int
	Erase(v ...interface{}) int
	Len() int
	Items() Items
	Value() interface{}
	Search(v interface{}, pos int) int
	Equal(slice interface{}) bool
	Get(v interface{}) (data interface{}, ok bool)
	Clone() Set
}

// Equal ...
func Equal(s1, s2 Set) bool {
	rit, ok := s1.(*reflectSet)
	if ok {
		return rit.Equal(s2.Value())
	}
	if s1.Len() != s2.Len() {
		return false
	}
	it1 := s1.Items()
	it2 := s2.Items()
	for i := 0; i < it1.Len(); i++ {
		if !it1.Elem(i).Equal(it2.Elem(i)) {
			return false
		}
	}
	return true
}

// Union ...
func Union(s1, s2 Set) Set {
	s := s1.Clone()
	s.Insert(s2.Value())
	return s
}

// Intersection ...
func Intersection(s1, s2 Set) Set {
	rit, ok := s1.(*reflectSet)
	if ok {
		return NewSet(rit.items.intersection(s2.(*reflectSet).items), true)
	}
	it1 := s1.Items()
	it2 := s2.Items()
	dst := it1.Truncate(0)
	if it1.Len() == 0 || it2.Len() == 0 {
		return NewSet(dst, true)
	}
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

	return NewSet(dst, true)
}

// Difference ...
func Difference(s1, s2 Set) Set {
	s := Union(s1, s2)
	s.Erase(Intersection(s1, s2).Value())
	return s
}
