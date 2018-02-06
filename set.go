package goset

import "sort"

type set struct {
	items Items
}

var _ Set = (*set)(nil)

// NewSet ...
func NewSet(items Items) Set {
	s := &set{items.Truncate(0)}
	s.Insert(items)
	return s
}

// Search ...
func (p *set) Search(v Element, pos int) int {
	return sort.Search(p.items.Len()-pos, func(i int) bool {
		return !p.items.Elem(pos + i).Less(v)
	})
}

func (p *set) Has(v Element, pos int) bool {
	n := p.Search(v, pos)
	if n == p.items.Len() || !p.items.Elem(pos+n).Equal(v) {
		return false
	}
	return true
}

func (p *set) Insert(items Items) (insertNum int) {
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

func (p *set) Erase(items Items) (eraseNum int) {
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
