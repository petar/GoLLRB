package llrb

import (
	"reflect"
	"testing"
)

func TestAscendGreaterOrEqual(t *testing.T) {
	tree := New(IntLess)
	tree.InsertNoReplace(4)
	tree.InsertNoReplace(6)
	tree.InsertNoReplace(1)
	tree.InsertNoReplace(3)
	var ary []Item
	tree.AscendGreaterOrEqual(-1, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected := []Item{1,3,4,6}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.AscendGreaterOrEqual(3, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected = []Item{3,4,6}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.AscendGreaterOrEqual(2, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected = []Item{3,4,6}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
}

func TestDescendLessOrEqual(t *testing.T) {
	tree := New(IntLess)
	tree.InsertNoReplace(4)
	tree.InsertNoReplace(6)
	tree.InsertNoReplace(1)
	tree.InsertNoReplace(3)
	var ary []Item
	tree.DescendLessOrEqual(10, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected := []Item{6,4,3,1}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.DescendLessOrEqual(4, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected = []Item{4,3,1}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
	ary = nil
	tree.DescendLessOrEqual(5, func(i Item) bool {
		ary = append(ary, i)
		return true
	});
	expected = []Item{4,3,1}
	if !reflect.DeepEqual(ary, expected) {
		t.Errorf("expected %v but got %v", expected, ary)
	}
}
