package goset_test

import (
	"reflect"
	"testing"

	"github.com/jettyu/goset"
)

func TestReflect(t *testing.T) {
	uintitems := goset.UintItemsCreator([]uint{5, 1, 2, 4, 2, 6, 4, 3})
	s := goset.NewSet(uintitems)
	t.Log(s.Value().([]uint))
	if !goset.Equal(s, goset.Uints([]uint{1, 2, 3, 4, 5, 6})) {
		t.Fatal(s.Value())
	}
	clone := s.Clone()
	if !s.Equal(clone.Value()) {
		t.Fatal(clone.Value())
	}
	s.Erase(uint(5))
	if s.Equal(clone.Value()) {
		t.Fatal(clone.Value(), s.Value())
	}
	t.Log(s.Value(), clone.Value())
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
	if !goset.Equal(userSet, goset.NewSet(reflectUserItemsCreator([]reflectUser{
		{"a", 1},
		{"d", 1},
		{"e", 2},
		{"c", 5},
		{"b", 10},
	}))) {
		t.Fatal(userSet.Value())
	}
	// has {"c",5}
	if !userSet.Has(reflectUser{"c", 5}, 0) {
		t.Fatal(userSet.Value())
	}
	// erase {"c",5}
	if userSet.Erase(reflectUser{"c", 5}) != 1 {
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
	idItems := goset.Strings([]string{"a", "d", "e", "c", "b"}, true)
	if !goset.Equal(goset.NewSet(items1), idItems) {
		t.Fatal(items1.Value())
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
	idItems = goset.Strings([]string{"b", "d", "c"})
	insItems := goset.Intersection(s, idItems)
	if !goset.Equal(insItems, goset.Strings([]string{"b", "d"})) {
		t.Fatal(insItems.Value())
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

func BenchmarkSwap(b *testing.B) {
	creator := goset.ReflectItemsCreator(nil, nil, nil)
	arr := []int{1, 2}
	items := creator(arr)
	for i := 0; i < b.N; i++ {
		items.Swap(0, 1)
	}
}

func BenchmarkSwap1(b *testing.B) {
	creator := goset.ReflectItemsCreator(nil, func(i, j int, src interface{}) {
		s := src.([]int)
		s[i], s[j] = s[j], s[i]
	}, nil)
	arr := []int{1, 2}
	items := creator(arr)
	for i := 0; i < b.N; i++ {
		items.Swap(0, 1)
	}
}
func BenchmarkEqualFunc(b *testing.B) {
	f := func(i, j interface{}) bool { return reflect.DeepEqual(i, j) }
	for i := 0; i < b.N; i++ {
		f(0, 0)
	}
}

func BenchmarkEqualFunc1(b *testing.B) {
	f := func(i, j interface{}) bool {
		return i.(int) == j.(int)
	}
	for i := 0; i < b.N; i++ {
		f(0, 0)
	}
}
