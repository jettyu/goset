package goset_test

import (
	"testing"

	"github.com/jettyu/goset"
)

func TestReflect(t *testing.T) {
	uintitems := goset.UintItemsCreator([]uint{1, 5, 2, 4, 2, 6, 4, 3})
	s := goset.NewSet(uintitems)
	t.Log(s.Data().([]uint))
	if !goset.Equal(s.Items(), goset.UintItemsCreator([]uint{1, 2, 3, 4, 5, 6})) {
		t.Fatal(s.Data())
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
	t.Log(userSet.Data().([]reflectUser))
	if !goset.Equal(userSet.Items(), reflectUserItemsCreator([]reflectUser{
		{"a", 1},
		{"d", 1},
		{"e", 2},
		{"c", 5},
		{"b", 10},
	})) {
		t.Fatal(userSet.Items())
	}
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goset.NewSet(goset.IntItemsCreator(
			[]int{1, 5, 2, 3, 3, 4})).Data()
	}
}

func BenchmarkInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goset.Ints([]int{1, 5, 2, 3, 3, 4}).Items()
	}
}
