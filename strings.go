package goset

// Strings ...
func Strings(v []string) Set {
	return NewSet(StringSlice(v))
}

// StringElement ...
type StringElement string

var _ Element = StringElement(0)

// Less ...
func (p StringElement) Less(v Element) bool {
	return p < v.(StringElement)
}

// Equal ...
func (p StringElement) Equal(v Element) bool {
	return p == v.(StringElement)
}

// StringSlice ...
type StringSlice []string

var _ Items = StringSlice{}

// Len ...
func (p StringSlice) Len() int { return len(p) }

// Less ...
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

// Swap ...
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Elem ...
func (p StringSlice) Elem(i int) Element {
	return StringElement(p[i])
}

// SetElem ...
func (p StringSlice) SetElem(v Element, pos int) {
	p[pos] = string(v.(StringElement))
}

// Move ...
func (p StringSlice) Move(dstPos, srcPos, n int) {
	copy(p[dstPos:dstPos+n], p[srcPos:srcPos+n])
}

// Append ...
func (p StringSlice) Append(arr ...Element) Items {
	for _, v := range arr {
		p = append(p, string(v.(StringElement)))
	}
	return p
}

// Truncate ...
func (p StringSlice) Truncate(n int) Items {
	p = p[:n]
	return p
}
