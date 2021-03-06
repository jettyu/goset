package goset_test

import (
	"fmt"
	"testing"

	"github.com/jettyu/goset"
)

func TestStrings(t *testing.T) {
	s := goset.Strings([]string{"2", "6", "4", "5", "4", "2", "3", "0", "1"})
	arr, ok := s.Value().([]string)
	if !ok {
		t.Fatal(s)
	}
	if len(arr) != 7 {
		t.Fatal(s, arr)
	}
	if !s.Has("0", 0) {
		t.Fatal(s)
	}
	if s.Has("0", 1) {
		t.Fatal(s)
	}
	if !s.Has("3", 2) {
		t.Fatal(s)
	}
	if s.Has("10", 0) {
		t.Fatal(s)
	}
	if s.Insert("1", "5", "7", "8") != 2 {
		t.Fatal(s)
	}
	// 删除中间，末尾混淆
	if s.Erase("7", "9") != 1 {
		t.Fatal(s)
	}
	// 删除中间和末尾
	if s.Erase("6", "8") != 2 {
		t.Fatal(s)
	}
	// 删除开头
	if s.Erase("0", "1") != 2 {
		t.Fatal(s)
	}
	for i, v := range s.Value().([]string) {
		if fmt.Sprint(i+2) != v {
			t.Fatal(arr)
		}
	}
}
