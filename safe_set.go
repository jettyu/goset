package goset

import "sync"

type safeSet struct {
	set Set
	sync.RWMutex
}

var _ Set = (*safeSet)(nil)

// NewSafeSet ...
func NewSafeSet(set Set) Set {
	return &safeSet{
		set: set,
	}
}

func (p *safeSet) Has(v interface{}, pos int) bool {
	p.RLock()
	ok := p.set.Has(v, pos)
	p.RUnlock()
	return ok
}

func (p *safeSet) Insert(v ...interface{}) int {
	p.Lock()
	n := p.set.Insert(v...)
	p.Unlock()
	return n
}

func (p *safeSet) Erase(v ...interface{}) int {
	p.Lock()
	n := p.set.Erase(v...)
	p.Unlock()
	return n
}

func (p *safeSet) Items() Items {
	p.RLock()
	it := p.set.Items()
	p.RUnlock()
	return it
}

func (p *safeSet) Value() interface{} {
	p.RLock()
	v := p.set.Value()
	p.RUnlock()
	return v
}

func (p *safeSet) Search(v interface{}, pos int) int {
	p.RLock()
	i := p.set.Search(v, pos)
	p.RUnlock()
	return i
}

func (p *safeSet) Equal(v Items) bool {
	p.RLock()
	ok := p.set.Equal(v)
	p.RUnlock()
	return ok
}

func (p *safeSet) Get(v interface{}) (data interface{}, ok bool) {
	p.RLock()
	data, ok = p.set.Get(v)
	p.RUnlock()
	return
}

func (p *safeSet) Len() int {
	p.RLock()
	n := p.set.Len()
	p.RUnlock()
	return n
}

func (p *safeSet) Clone() Set {
	p.RLock()
	s := p.set.Clone()
	p.RUnlock()
	return s
}
