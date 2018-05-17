package goset_test

import (
	"testing"

	"github.com/jettyu/goset"
)

func TestStruct(t *testing.T) {
	users := Users{
		{"a", 1},
		{"b", 10},
		{"d", 1},
		{"c", 5},
		{"b", 10},
		{"e", 2},
	}
	userSet := goset.NewSet(goset.NewItems(users))
	// [{a 1} {d 1} {e 2} {c 5} {b 10}]
	t.Log(userSet.Value().(Users))
	if !goset.Equal(userSet, goset.NewSet(goset.NewItems(Users{
		{"a", 1},
		{"d", 1},
		{"e", 2},
		{"c", 5},
		{"b", 10},
	}), true)) {
		t.Fatal(userSet.Items())
	}
}

// User ...
type User struct {
	Name string
	Age  int
}

var _ goset.Element = User{}

// Less : order by age
func (p User) Less(e goset.Element) bool {
	v := e.(User)
	if p.Age == v.Age {
		return p.Name < v.Name
	}
	return p.Age < v.Age
}

// Equal : name and age all equal
func (p User) Equal(e goset.Element) bool {
	v := e.(User)
	return p.Name == v.Name && p.Age == v.Age
}

// Users ...
type Users []User

var _ goset.Items = goset.NewItems(Users{})

func (p Users) Len() int           { return len(p) }
func (p Users) Less(i, j int) bool { return p[i].Less(p[j]) }
func (p Users) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
