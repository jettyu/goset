# sort set

sort set for go

## example

### ints

```go
    func TestInts(t *testing.T) {
        s := goset.Ints([]int{2, 6, 4, 5, 4, 2, 3, 0, 1})
        if !s.Equal([]int{0, 1, 2, 3, 4, 5, 6}) {
            t.Fatal(s.Items())
        }
        if !s.Has(0, 0) {
            t.Fatal(s)
        }
        if s.Has(0, 1) {
            t.Fatal(s)
        }
        if !s.Has(3, 2) {
            t.Fatal(s)
        }
        if s.Has(10, 0) {
            t.Fatal(s)
        }
        if s.Insert([]int{1, 5, 7, 8}) != 2 {
            t.Fatal(s)
        }

        if s.Erase([]int{7, 9}) != 1 {
            t.Fatal(s)
        }
        if s.Erase([]int{6, 8}) != 2 {
            t.Fatal(s)
        }
        if s.Erase([]int{0, 1}) != 2 {
            t.Fatal(s)
        }
        if !s.Equal([]int{2, 3, 4, 5}) {
            t.Fatal(s.Items())
        }

        clone := s.Clone()
        if !s.Equal(clone.Items()) {
            t.Fatal(clone.Value())
        }
        s.Erase(5)
        if s.Equal(clone.Items()) {
            t.Fatal(clone.Value(), s.Value())
        }
    }
```

### reflect

```go
func TestReflectStruct(t *testing.T) {
    type reflectUser struct {
        Name string
        Age  int
    }
    reflectUserItemsCreator := goset.ReflectItemsCreator(
        func(s1, s2 interface{}) bool {
            u1 := s1.(reflectUser)
            u2 := s2.(reflectUser)
            if u1.Age == u2.Age {
                return u1.Name < u2.Name
            }
            return u1.Age < u2.Age
        },
        func(i, j int, slice interface{}) {
            arr := slice.([]reflectUser)
            arr[i], arr[j] = arr[j], arr[i]
        },
        func(s1, s2 interface{}) bool {
            u1 := s1.(reflectUser)
            u2 := s2.(reflectUser)
            return u1.Name == u2.Name && u1.Age == u2.Age
        },
    )
    items1 := reflectUserItemsCreator([]reflectUser{
        {"a", 1},
        {"b", 10},
        {"d", 1},
        {"c", 5},
        {"b", 10},
        {"e", 2},
    })
    userSet := goset.NewSet(items1)
    // [{a 1} {d 1} {e 2} {c 5} {b 10}]
    t.Log(userSet.Value().([]reflectUser))
    if !goset.Equal(userSet.Items(), reflectUserItemsCreator([]reflectUser{
        {"a", 1},
        {"d", 1},
        {"e", 2},
        {"c", 5},
        {"b", 10},
    })) {
        t.Fatal(userSet.Value())
    }
    // has {"c",5}
    if !userSet.Has(reflectUser{"c", 5}, 0) {
        t.Fatal(userSet.Value())
    }
    // erase {"c",5}
    if userSet.Erase(reflectUserItemsCreator([]reflectUser{
        {"c", 5},
    })) != 1 {
        t.Fatal(userSet.Value())
    }
    // not has {"c", 5}
    if userSet.Has(reflectUser{"c", 5}, 0) {
        t.Fatal(userSet.Value())
    }
}
```

## more example

see the test file