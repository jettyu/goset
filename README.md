# sort set

sort set for go

## example

    func TestInts(t *testing.T) {
        s := goset.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
        if !s.Equal(goset.IntSlice([]int{0, 1, 2, 3, 4, 5, 6})) {
            t.Fatal(s.Items())
        }
        if !s.Has(goset.IntElement(0), 0) {
            t.Fatal(s)
        }
        if s.Has(goset.IntElement(0), 1) {
            t.Fatal(s)
        }
        if !s.Has(goset.IntElement(3), 2) {
            t.Fatal(s)
        }
        if s.Has(goset.IntElement(10), 0) {
            t.Fatal(s)
        }
        if s.Insert(goset.IntSlice([]int{1, 5, 7, 8})) != 2 {
            t.Fatal(s)
        }
        if s.Erase(goset.IntSlice([]int{7, 9})) != 1 {
            t.Fatal(s)
        }
        if s.Erase(goset.IntSlice([]int{6, 8})) != 2 {
            t.Fatal(s)
        }
        if s.Erase(goset.IntSlice([]int{0, 1})) != 2 {
            t.Fatal(s)
        }
        if !s.Equal(goset.IntSlice([]int{2, 3, 4, 5})) {
            t.Fatal(s.Items())
        }
    }

## more example

see the test file