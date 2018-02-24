package goset_test

import (
	"testing"

	"github.com/jettyu/goset"
)

func TestReflect(t *testing.T) {
	uintitems := goset.UintItemsCreator([]uint{5, 1, 2, 4, 2, 6, 4, 3})
	s := goset.NewSet(uintitems)
	t.Log(s.Value().([]uint))
	if !goset.Equal(s.Items(), goset.UintItemsCreator([]uint{1, 2, 3, 4, 5, 6})) {
		t.Fatal(s.Value())
	}
}

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

func TestReflectStruct1(t *testing.T) {
	type reflectUser struct {
		ID  string
		Age int
	}
	reflectUserItemsCreator := goset.ReflectItemsCreator(
		func(s1, s2 interface{}) bool {
			u1 := s1.(reflectUser)
			u2 := s2.(reflectUser)
			return u1.Age < u2.Age
		},
		func(i, j int, slice interface{}) {
			arr := slice.([]reflectUser)
			arr[i], arr[j] = arr[j], arr[i]
		},
		func(s1, s2 interface{}) bool {
			return s1.(reflectUser).ID == s2.(reflectUser).ID
		},
	)
	users := []reflectUser{
		{"a", 1},
		{"b", 10},
		{"d", 1},
		{"c", 5},
		{"b", 10},
		{"e", 2},
	}
	// equal by id
	items1 := goset.NewSet(reflectUserItemsCreator(users)).Items().(goset.ReflectItems)
	items1 = items1.WithFunc(nil,
		func(s1, s2 interface{}) bool {
			u, ok := s2.(reflectUser)
			if ok {
				return s1.(reflectUser).ID == u.ID
			}
			return s1.(reflectUser).ID == s2.(string)
		})
	idItems := goset.StringsItemsCreator([]string{"a", "d", "e", "c", "b"})
	if !goset.Equal(items1, idItems) {
		t.Fatal(items1.(goset.ReflectValue).Value())
	}
	items1 = items1.WithFunc(func(s1, s2 interface{}) bool {
		u, ok := s2.(reflectUser)
		if ok {
			return s1.(reflectUser).ID < u.ID
		}
		return s1.(reflectUser).ID < s2.(string)
	}, nil)
	// has by id
	if !goset.NewSet(items1).Has("c", 0) {
		t.Fatal(items1.Value())
	}
	get, ok := goset.NewSet(items1).Get("c")
	if !ok || get.(reflectUser).ID != "c" {
		t.Fatal(get, ok)
	}
	// erase by id
	s := goset.NewSet(items1)
	s.Erase(goset.StringsItemsCreator([]string{"c"}))
	t.Log(s.Value())
	// intersection by id
	idItems = goset.StringsItemsCreator([]string{"b", "d", "c"})
	insItems := goset.Intersection(items1, idItems)
	if !goset.Equal(insItems, goset.StringsItemsCreator([]string{"b", "d"})) {
		t.Fatal(insItems.(goset.ReflectItems).Value())
	}
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goset.NewSet(goset.IntItemsCreator(
			[]int{1, 5, 2, 3, 3, 4})).Value()
	}
}

func BenchmarkInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goset.Ints([]int{1, 5, 2, 3, 3, 4}).Items()
	}
}

func BenchmarkEqual(b *testing.B) {
	arr := []int{1, 5, 2, 3, 3, 4}
	exp := []int{1, 2, 3, 4, 5}
	items := goset.Ints(arr).Items()
	expItems := goset.IntSlice(exp)
	for i := 0; i < b.N; i++ {
		if !goset.Equal(items, expItems) {
			b.Fatal(items)
		}
	}
}

func BenchmarkReflectEqual(b *testing.B) {
	arr := []int{1, 5, 2, 3, 3, 4}
	exp := []int{1, 2, 3, 4, 5}
	items := goset.NewSet(goset.IntItemsCreator(arr)).Items()
	expItems := goset.IntItemsCreator(exp)
	for i := 0; i < b.N; i++ {
		if !goset.Equal(items, expItems) {
			b.Fatal(items)
		}
	}
}
