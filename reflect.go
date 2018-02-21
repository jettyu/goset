package goset

import (
	"reflect"
)

// SliceElement ...
type SliceElement interface {
	Element
	Data() interface{}
}

// SliceItems ...
type SliceItems interface {
	Items
	Data() interface{}
}

// ItemsCreator ...
var ItemsCreator = func(lessFunc func(s1, s2 interface{}) bool,
	swapFunc func(i, j int, src interface{}),
	equalFunc func(s1, s2 interface{}) bool,
) func(slice interface{}) Items {
	return func(slice interface{}) Items {
		rv := reflect.ValueOf(slice)
		if equalFunc == nil {
			equalFunc = func(s1, s2 interface{}) bool {
				return reflect.DeepEqual(s1, s2)
			}
		}
		return sliceItems{
			rv:        rv,
			lessFunc:  lessFunc,
			equalFunc: equalFunc,
			swapFunc:  swapFunc,
		}
	}
}

var (
	// IntItemsCreator ...
	IntItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int) < s2.(int)
		}, func(i, j int, src interface{}) {
			arr := src.([]int)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Int64ItemsCreator ...
	Int64ItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int64) < s2.(int64)
		}, func(i, j int, src interface{}) {
			arr := src.([]int64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// UintItemsCreator ...
	UintItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint) < s2.(uint)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Uint64ItemsCreator ...
	Uint64ItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint64) < s2.(uint64)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Float32ItemsCreator ...
	Float32ItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(float32) < s2.(float32)
		}, func(i, j int, src interface{}) {
			arr := src.([]float32)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Float64ItemsCreator ...
	Float64ItemsCreator = ItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(float64) < s2.(float64)
		}, func(i, j int, src interface{}) {
			arr := src.([]float64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
)

// SliceElement ...
type sliceElement struct {
	v         interface{}
	lessFunc  func(s1, s2 interface{}) bool
	equalFunc func(s1, s2 interface{}) bool
}

func (p sliceElement) Data() interface{} {
	return p.v
}

func (p sliceElement) Less(e Element) bool  { return p.lessFunc(p.v, e.(sliceElement).v) }
func (p sliceElement) Equal(e Element) bool { return p.equalFunc(p.v, e.(sliceElement).v) }

// sliceItems ...
type sliceItems struct {
	rv        reflect.Value
	lessFunc  func(s1, s2 interface{}) bool
	equalFunc func(s1, s2 interface{}) bool
	swapFunc  func(i, j int, src interface{})
}

func (p sliceItems) Data() interface{} {
	return p.rv.Interface()
}

var _ Items = sliceItems{}

func (p sliceItems) Len() int { return p.rv.Len() }

func (p sliceItems) Less(i, j int) bool {
	return p.Elem(i).Less(p.Elem(j))
}

func (p sliceItems) Swap(i, j int) {
	p.swapFunc(i, j, p.rv.Interface())
}

// Elem ...
func (p sliceItems) Elem(i int) Element {
	return sliceElement{
		v:         p.rv.Index(i).Interface(),
		lessFunc:  p.lessFunc,
		equalFunc: p.equalFunc,
	}
}

// SetElem ...
func (p sliceItems) SetElem(e Element, pos int) {
	p.rv.Index(pos).Set(reflect.ValueOf(e.(sliceElement).v))
}

// Move ...
func (p sliceItems) Move(dstPos, srcPos, n int) {
	reflect.Copy(p.rv.Slice(dstPos, dstPos+n), p.rv.Slice(srcPos, srcPos+n))
}

// Append ...
func (p sliceItems) Append(e ...Element) Items {
	for _, v := range e {
		p.rv = reflect.Append(p.rv, reflect.ValueOf(v.(sliceElement).v))
	}

	return p
}

// Truncate ...
func (p sliceItems) Truncate(n int) Items {
	p.rv = p.rv.Slice(0, n)
	return p
}
