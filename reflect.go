package goset

import (
	"reflect"
	"sort"
)

// ReflectValue ...
type ReflectValue interface {
	Value() interface{}
}

// ReflectItems ...
type ReflectItems interface {
	ReflectValue
	Items
	// WithFunc, func is invalid when is nil
	WithFunc(lessFunc, equalFunc func(s1, s2 interface{}) bool) ReflectItems
}

var (
	// Int8s ...
	Int8s = func(arr []int8) Set { return NewSet(Int8ItemsCreator(arr)) }
	// Int16s ...
	Int16s = func(arr []int16) Set { return NewSet(Int16ItemsCreator(arr)) }
	// Int32s ...
	Int32s = func(arr []int32) Set { return NewSet(Int32ItemsCreator(arr)) }
	// Int64s ...
	Int64s = func(arr []int64) Set { return NewSet(Int64ItemsCreator(arr)) }
	// Uints ...
	Uints = func(arr []uint) Set { return NewSet(UintItemsCreator(arr)) }
	// Uint8s ...
	Uint8s = func(arr []uint8) Set { return NewSet(Uint8ItemsCreator(arr)) }
	// Uint16s ...
	Uint16s = func(arr []uint16) Set { return NewSet(Uint16ItemsCreator(arr)) }
	// Uint32s ...
	Uint32s = func(arr []uint32) Set { return NewSet(Uint32ItemsCreator(arr)) }
	// Uint64s ...
	Uint64s = func(arr []uint64) Set { return NewSet(Uint64ItemsCreator(arr)) }
	// Float32s ...
	Float32s = func(arr []float32) Set { return NewSet(Float32ItemsCreator(arr)) }
	// Float64s ...
	Float64s = func(arr []float64) Set { return NewSet(Float64ItemsCreator(arr)) }
)

var (
	// ReflectItemsCreator ...
	ReflectItemsCreator = func(lessFunc func(s1, s2 interface{}) bool,
		swapFunc func(i, j int, src interface{}),
		equalFunc func(s1, s2 interface{}) bool,
	) func(slice interface{}) ReflectItems {
		return func(slice interface{}) ReflectItems {
			rv := reflect.ValueOf(slice)
			if equalFunc == nil {
				equalFunc = func(s1, s2 interface{}) bool {
					return reflect.DeepEqual(s1, s2)
				}
			}
			return reflectItems{
				rv:        rv,
				lessFunc:  lessFunc,
				equalFunc: equalFunc,
				swapFunc:  swapFunc,
			}
		}
	}
	// StringsItemsCreator ...
	StringsItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(string) < s2.(string)
		}, func(i, j int, src interface{}) {
			arr := src.([]string)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// IntItemsCreator ...
	IntItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int) < s2.(int)
		}, func(i, j int, src interface{}) {
			arr := src.([]int)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Int8ItemsCreator ...
	Int8ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int8) < s2.(int8)
		}, func(i, j int, src interface{}) {
			arr := src.([]int8)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Int16ItemsCreator ...
	Int16ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int16) < s2.(int16)
		}, func(i, j int, src interface{}) {
			arr := src.([]int16)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Int32ItemsCreator ...
	Int32ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int32) < s2.(int32)
		}, func(i, j int, src interface{}) {
			arr := src.([]int32)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Int64ItemsCreator ...
	Int64ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(int64) < s2.(int64)
		}, func(i, j int, src interface{}) {
			arr := src.([]int64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// UintItemsCreator ...
	UintItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint) < s2.(uint)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Uint8ItemsCreator ...
	Uint8ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint8) < s2.(uint8)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint8)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Uint16ItemsCreator ...
	Uint16ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint16) < s2.(uint16)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint16)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Uint32ItemsCreator ...
	Uint32ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint32) < s2.(uint32)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint32)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Uint64ItemsCreator ...
	Uint64ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(uint64) < s2.(uint64)
		}, func(i, j int, src interface{}) {
			arr := src.([]uint64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Float32ItemsCreator ...
	Float32ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(float32) < s2.(float32)
		}, func(i, j int, src interface{}) {
			arr := src.([]float32)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
	// Float64ItemsCreator ...
	Float64ItemsCreator = ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			return s1.(float64) < s2.(float64)
		}, func(i, j int, src interface{}) {
			arr := src.([]float64)
			arr[i], arr[j] = arr[j], arr[i]
		},
		nil,
	)
)

type reflectSet struct {
	items reflectItems
}

var _ Set = (*reflectSet)(nil)

// Search ...
func (p reflectSet) Search(v interface{}, pos int) int {
	return sort.Search(p.items.Len()-pos, func(i int) bool {
		return !p.items.lessFunc(p.items.elem(pos+i), v)
	})
}

// Has ...
func (p reflectSet) Has(v interface{}, pos int) bool {
	n := p.Search(v, pos)
	if n == p.items.Len() || !p.items.equalFunc(p.items.elem(pos+n), v) {
		return false
	}
	return true
}

func (p *reflectSet) Insert(v ...interface{}) (intsertNum int) {
	for _, arg := range v {
		items, ok := arg.(Items)
		if ok {
			intsertNum += p.InsertItems(items)
		} else {
			intsertNum += p.insertElem(arg)
		}
	}
	return
}

func (p *reflectSet) Erase(v ...interface{}) (eraseNum int) {
	for _, arg := range v {
		items, ok := arg.(Items)
		if ok {
			eraseNum += p.EraseItems(items)
		} else {
			eraseNum += p.eraseElem(arg)
		}
	}
	return
}

func (p *reflectSet) InsertItems(it Items) (insertNum int) {
	items := it.(reflectItems)
	if !sort.IsSorted(items) {
		sort.Sort(items)
	}
	pos := 0
	for i := 0; i < items.Len(); i++ {
		v := items.elem(i)
		if p.items.Len() == 0 {
			p.items = p.items.append(v)
			insertNum++
			pos++
			continue
		}
		pos += p.Search(v, pos)
		n := pos
		if pos < p.items.Len() {
			e := p.items.elem(pos)
			if p.items.equalFunc(e, v) {
				// has v
				continue
			} else if p.items.lessFunc(e, v) {
				// less than v, insert after e
				n++
			}
		} else {
			pos--
		}
		insertNum++
		p.items = p.items.append(v)
		p.items.Move(pos+1, pos, p.items.Len()-(pos+1))
		p.items.setElem(v, n)
	}

	return
}

func (p *reflectSet) EraseItems(it Items) (eraseNum int) {
	items := it.(reflectItems)
	if p.items.Len() == 0 {
		return
	}
	if !sort.IsSorted(items) {
		sort.Sort(items)
	}

	pos := 0
	for i := 0; i < items.Len() && pos < p.items.Len(); i++ {
		v := items.elem(i)
		pos += p.Search(v, pos)
		if pos == p.items.Len() || !p.items.equalFunc(p.items.elem(pos), v) {
			continue
		}
		p.items.Move(pos, pos+1, p.items.Len()-(pos+1))
		p.items = p.items.truncate(p.items.Len() - 1)
		eraseNum++
	}

	return
}

func (p reflectSet) Items() Items {
	return p.items
}

func (p reflectSet) Value() interface{} {
	return p.items.Value()
}

func (p reflectSet) Equal(items Items) bool {
	return p.items.equal(items)
}

func (p reflectSet) Get(id interface{}) (data interface{}, ok bool) {
	pos := p.Search(id, 0)
	if pos == p.items.Len() {
		return
	}
	data = p.items.elem(pos)
	ok = p.items.equalFunc(p.items.elem(pos), id)
	return
}

func (p reflectSet) Len() int {
	return p.items.Len()
}

// InsertElem ...
func (p *reflectSet) insertElem(v interface{}) int {
	if p.items.Len() == 0 {
		p.items = p.items.append(v)
		return 1
	}
	pos := p.Search(v, 0)
	n := pos
	if pos < p.items.Len() {
		e := p.items.elem(pos)
		if p.items.equalFunc(e, v) {
			// has v
			return 0
		} else if p.items.lessFunc(e, v) {
			// less than v, insert after e
			n++
		}
	} else {
		pos--
	}
	p.items = p.items.append(v)
	p.items.Move(pos+1, pos, p.items.Len()-(pos+1))
	p.items.setElem(v, n)
	return 1
}

func (p *reflectSet) eraseElem(v interface{}) int {
	if p.items.Len() == 0 {
		return 0
	}

	pos := p.Search(v, 0)
	if pos == p.items.Len() || !p.items.equalFunc(p.items.elem(pos), v) {
		return 0
	}
	p.items.Move(pos, pos+1, p.items.Len()-(pos+1))
	p.items = p.items.truncate(p.items.Len() - 1)

	return 1
}

// SliceElement ...
type reflectElement struct {
	v         interface{}
	lessFunc  func(s1, s2 interface{}) bool
	equalFunc func(s1, s2 interface{}) bool
}

func (p reflectElement) Less(e Element) bool  { return p.lessFunc(p.v, e.(reflectElement).v) }
func (p reflectElement) Equal(e Element) bool { return p.equalFunc(p.v, e.(reflectElement).v) }

// reflectItems ...
type reflectItems struct {
	rv        reflect.Value
	lessFunc  func(s1, s2 interface{}) bool
	equalFunc func(s1, s2 interface{}) bool
	swapFunc  func(i, j int, src interface{})
}

func (p reflectItems) Value() interface{} {
	return p.rv.Interface()
}

var _ ReflectItems = reflectItems{}

func (p reflectItems) Len() int { return p.rv.Len() }

func (p reflectItems) Less(i, j int) bool {
	return p.lessFunc(p.rv.Index(i).Interface(), p.rv.Index(j).Interface())
}

func (p reflectItems) Swap(i, j int) {
	p.swapFunc(i, j, p.rv.Interface())
}

// Elem ...
func (p reflectItems) Elem(i int) Element {
	return reflectElement{
		v:         p.rv.Index(i).Interface(),
		lessFunc:  p.lessFunc,
		equalFunc: p.equalFunc,
	}
}

// SetElem ...
func (p reflectItems) SetElem(e Element, pos int) {
	p.rv.Index(pos).Set(reflect.ValueOf(e.(reflectElement).v))
}

// Move ...
func (p reflectItems) Move(dstPos, srcPos, n int) {
	reflect.Copy(p.rv.Slice(dstPos, dstPos+n), p.rv.Slice(srcPos, srcPos+n))
}

// Append ...
func (p reflectItems) Append(e ...Element) Items {
	for _, v := range e {
		p.rv = reflect.Append(p.rv, reflect.ValueOf(v.(reflectElement).v))
	}

	return p
}

// Truncate ...
func (p reflectItems) Truncate(n int) Items {
	return p.truncate(n)
}

func (p reflectItems) WithFunc(lessFunc,
	equalFunc func(s1, s2 interface{}) bool) ReflectItems {
	if lessFunc != nil {
		p.lessFunc = lessFunc
	}
	if equalFunc != nil {
		p.equalFunc = equalFunc
	}
	return p
}

func (p reflectItems) elem(i int) interface{}         { return p.rv.Index(i).Interface() }
func (p reflectItems) setElem(e interface{}, pos int) { p.rv.Index(pos).Set(reflect.ValueOf(e)) }
func (p reflectItems) append(e ...interface{}) reflectItems {
	for _, v := range e {
		p.rv = reflect.Append(p.rv, reflect.ValueOf(v))
	}

	return p
}
func (p reflectItems) truncate(n int) reflectItems {
	p.rv = p.rv.Slice(0, n)
	return p
}
func (p reflectItems) equal(it Items) bool {
	items, ok := it.(reflectItems)
	if !ok {
		return false
	}
	if p.Len() != items.Len() {
		return false
	}
	for i := 0; i < p.Len(); i++ {
		if !p.equalFunc(p.rv.Index(i).Interface(), items.rv.Index(i).Interface()) {
			return false
		}
	}
	return true
}

func (p reflectItems) intersection(it Items) (dst reflectItems) {
	dst = p.truncate(0)
	if p.Len() == 0 || it.Len() == 0 {
		return
	}

	s1 := NewSet(p)
	s2 := NewSet(it)
	it1 := s1.Items().(reflectItems)
	it2 := s2.Items().(reflectItems)
	pos := 0
	for i := 0; i < it2.Len() && pos < it1.Len(); i++ {
		e := it2.elem(i)
		pos += s1.Search(e, pos)
		if pos == it1.Len() {
			continue
		}
		v := it1.elem(pos)
		if it1.equalFunc(v, e) {
			dst = dst.append(v)
		}
	}
	return
}

func (p *reflectSet) Clone() Set {
	items := p.items.truncate(0)
	items.rv = reflect.AppendSlice(reflect.Zero(p.items.rv.Type()), p.items.rv)
	return &reflectSet{
		items: items,
	}
}
