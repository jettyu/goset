package goset

import (
	"reflect"
	"sort"
)

// NewItems ...
func NewItems(v sort.Interface) Items {
	items, ok := v.(Items)
	if ok {
		return items
	}
	return setItems{reflect.ValueOf(v)}
}

type setItems struct {
	rv reflect.Value
}

func (p setItems) Len() int {
	return p.rv.Len()
}

func (p setItems) Less(i, j int) bool {
	return p.rv.Interface().(sort.Interface).Less(i, j)
}

func (p setItems) Swap(i, j int) {
	p.rv.Interface().(sort.Interface).Swap(i, j)
}

func (p setItems) Value() interface{} {
	return p.rv.Interface()
}

func (p setItems) Elem(i int) Element {
	return p.rv.Index(i).Interface().(Element)
}

func (p setItems) SetElem(v Element, pos int) {
	p.rv.Index(pos).Set(reflect.ValueOf(v))
}

func (p setItems) Append(e ...Element) Items {
	for _, v := range e {
		p.rv = reflect.Append(p.rv, reflect.ValueOf(v))
	}
	return p
}

func (p setItems) Truncate(n int) Items {
	p.rv = p.rv.Slice(0, n)
	return p
}

//  ...
func (p setItems) Move(dstPos, srcPos, n int) {
	reflect.Copy(p.rv.Slice(dstPos, dstPos+n), p.rv.Slice(srcPos, srcPos+n))
}

func (p setItems) Clone() Items {
	rv := reflect.MakeSlice(p.rv.Type(), p.rv.Len(), p.rv.Len())
	reflect.Copy(rv, p.rv)
	return setItems{rv}
}
