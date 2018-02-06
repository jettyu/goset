package goset

// Ints ...
func Ints(v []int) Set {
	return NewSet(IntSlice(v))
}

// IntElement ...
type IntElement int

var _ Element = IntElement(0)

// Less ...
func (p IntElement) Less(v Element) bool { return p < v.(IntElement) }

// Equal ...
func (p IntElement) Equal(v Element) bool { return p == v.(IntElement) }

// IntSlice ...
type IntSlice []int

var _ Items = IntSlice{}

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Elem ...
func (p IntSlice) Elem(i int) Element { return IntElement(p[i]) }

// SetElem ...
func (p IntSlice) SetElem(v Element, pos int) { p[pos] = int(v.(IntElement)) }

// Move ...
func (p IntSlice) Move(dstPos, srcPos, n int) { copy(p[dstPos:dstPos+n], p[srcPos:srcPos+n]) }

// Append ...
func (p IntSlice) Append(arr ...Element) Items {
	for _, v := range arr {
		p = append(p, int(v.(IntElement)))
	}
	return p
}

// Truncate ...
func (p IntSlice) Truncate(n int) Items { return p[:n] }
