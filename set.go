package goset

import (
	"reflect"
	"sort"
)

type set struct {
	items Items
}

var _ Set = (*set)(nil)

// NewSet ...
func NewSet(items Items) (s Set) {
	rit, ok := items.(reflectItems)
	if ok {
		s = &reflectSet{rit.truncate(0)}
	} else {
		s = &set{items.Truncate(0)}
	}
	s.Insert(items)
	return s
}

// Search ...
func (p *set) Search(v interface{}, pos int) int {
	return sort.Search(p.items.Len()-pos, func(i int) bool {
		return !p.items.Elem(pos + i).Less(v.(Element))
	})
}

func (p *set) Has(v interface{}, pos int) bool {
	n := p.Search(v, pos)
	if n == p.items.Len() || !p.items.Elem(pos+n).Equal(v.(Element)) {
		return false
	}
	return true
}

func (p *set) Insert(v ...interface{}) (intsertNum int) {
	for _, arg := range v {
		items, ok := arg.(Items)
		if ok {
			intsertNum += p.InsertItems(items)
		} else {
			intsertNum += p.insertElem(arg.(Element))
		}
	}
	return
}

func (p *set) Erase(v ...interface{}) (eraseNum int) {
	for _, arg := range v {
		items, ok := arg.(Items)
		if ok {
			eraseNum += p.EraseItems(items)
		} else {
			eraseNum += p.eraseElem(arg.(Element))
		}
	}
	return
}

func (p *set) InsertItems(items Items) (insertNum int) {
	if !sort.IsSorted(items) {
		sort.Sort(items)
	}
	pos := 0
	for i := 0; i < items.Len(); i++ {
		v := items.Elem(i)
		if p.items.Len() == 0 {
			p.items = p.items.Append(v)
			insertNum++
			pos++
			continue
		}
		pos += p.Search(v, pos)
		n := pos
		if pos < p.items.Len() {
			e := p.items.Elem(pos)
			if e.Equal(v) {
				// has v
				continue
			} else if e.Less(v) {
				// less than v, insert after e
				n++
			}
		} else {
			pos--
		}
		insertNum++
		p.items = p.items.Append(v)
		p.items.Move(pos+1, pos, p.items.Len()-(pos+1))
		p.items.SetElem(v, n)
	}

	return
}

func (p *set) EraseItems(items Items) (eraseNum int) {
	if p.items.Len() == 0 {
		return
	}
	if !sort.IsSorted(items) {
		sort.Sort(items)
	}
	pos := 0
	for i := 0; i < items.Len() && pos < p.items.Len(); i++ {
		v := items.Elem(i)
		pos += p.Search(v, pos)
		if pos == p.items.Len() || !p.items.Elem(pos).Equal(v) {
			continue
		}
		p.items.Move(pos, pos+1, p.items.Len()-(pos+1))
		p.items = p.items.Truncate(p.items.Len() - 1)
		eraseNum++
	}

	return
}

func (p set) Items() Items {
	return p.items
}

func (p set) Value() interface{} {
	return p.items
}

func (p set) Equal(items Items) bool {
	return Equal(p.items, items)
}

func (p set) Get(v interface{}) (data interface{}, ok bool) {
	pos := p.Search(v, 0)
	if pos == p.items.Len() {
		return
	}
	data = p.items.Elem(pos)
	ok = p.items.Elem(pos).Equal(v.(Element))
	return
}

func (p set) Len() int {
	return p.items.Len()
}

func (p *set) insertElem(v Element) int {
	if p.items.Len() == 0 {
		p.items = p.items.Append(v)
		return 1
	}
	pos := p.Search(v, 0)
	n := pos
	if pos < p.items.Len() {
		e := p.items.Elem(pos)
		if e.Equal(v) {
			// has v
			return 0
		} else if e.Less(v) {
			// less than v, insert after e
			n++
		}
	} else {
		pos--
	}
	p.items = p.items.Append(v)
	p.items.Move(pos+1, pos, p.items.Len()-(pos+1))
	p.items.SetElem(v, n)

	return 1
}

func (p *set) eraseElem(e Element) int {
	v := e.(Element)

	if p.items.Len() == 0 {
		return 0
	}

	pos := p.Search(v, 0)
	if pos == p.items.Len() || !p.items.Elem(pos).Equal(v) {
		return 0
	}
	p.items.Move(pos, pos+1, p.items.Len()-(pos+1))
	p.items = p.items.Truncate(p.items.Len() - 1)

	return 1
}

func (p set) Clone() Set {
	v := reflect.ValueOf(p.items)
	rv := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	reflect.Copy(rv, v)
	items := rv.Interface().(Items)
	// items := reflect.Zero(reflect.TypeOf(p.items)).Interface().(Items)
	// for i := 0; i < p.items.Len(); i++ {
	// 	items = items.Append(p.items.Elem(i))
	// }
	return &set{items: items}
}
