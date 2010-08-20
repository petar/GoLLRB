// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package llrb

import (
	"math"
	"rand"
	"testing"
)

type IntItem int

func (item IntItem) LessThan(other interface{}) bool {
	return int(item) < int(other.(IntItem))
}

type StringItem string

func (item StringItem) LessThan(other interface{}) bool {
	return string(item) < string(other.(StringItem))
}

func TestCases(t *testing.T) {
	tree := &Tree{}
	tree.InsertOrReplace(IntItem(1))
	tree.InsertOrReplace(IntItem(1))
	if tree.Len() != 1 {
		t.Errorf("expecting len 1")
	}
	if !tree.Has(IntItem(1)) {
		t.Errorf("expecting to find key=1")
	}

	tree.Delete(IntItem(1))
	if tree.Len() != 0 {
		t.Errorf("expecting len 0")
	}
	if tree.Has(IntItem(1)) {
		t.Errorf("not expecting to find key=1")
	}

	tree.Delete(IntItem(1))
	if tree.Len() != 0 {
		t.Errorf("expecting len 0")
	}
	if tree.Has(IntItem(1)) {
		t.Errorf("not expecting to find key=1")
	}
}

func TestReverseInsertOrder(t *testing.T) {
	tree := &Tree{}
	n := 100
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(n - i))
	}
	c := tree.Iter()
	for j, item := 1, <-c; item != nil; j, item = j+1, <-c {
		if int(item.(IntItem)) != j {
			t.Fatalf("bad order")
		}
	}
}

func TestRange(t *testing.T) {
	tree := &Tree{}
	order := []StringItem{
		StringItem("ab"), StringItem("aba"), StringItem("abc"),
		StringItem("a"), StringItem("aa"), StringItem("aaa"),
		StringItem("b"), StringItem("a-"), StringItem("a!"),
	}
	for _, i := range order {
		tree.InsertOrReplace(i)
	}
	c := tree.IterRange(StringItem("ab"), StringItem("ac"))
	k := 0
	for item := <-c; item != nil; item = <-c {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := string(order[k])
		i2 := string(item.(StringItem))
		if i1 != i2 {
			t.Errorf("expecting %s, got %s", i1, i2)
		}
		k++
	}
}

func TestRandomInsertOrder(t *testing.T) {
	tree := &Tree{}
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	c := tree.Iter()
	for j, item := 0, <-c; item != nil; j, item = j+1, <-c {
		if int(item.(IntItem)) != j {
			t.Fatalf("bad order")
		}
	}
}

func TestRandomReplace(t *testing.T) {
	tree := &Tree{}
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	perm = rand.Perm(n)
	for i := 0; i < n; i++ {
		if replaced := tree.InsertOrReplace(IntItem(perm[i])); replaced == nil || int(replaced.(IntItem)) != perm[i] {
			t.Errorf("error replacing")
		}
	}
}

func TestRandomInsertSequentialDelete(t *testing.T) {
	tree := &Tree{}
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	for i := 0; i < n; i++ {
		tree.Delete(IntItem(i))
	}
}

func TestRandomInsertDeleteNonExistent(t *testing.T) {
	tree := &Tree{}
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	if tree.Delete(IntItem(200)) != nil {
		t.Errorf("deleted non-existent item")
	}
	if tree.Delete(IntItem(-2)) != nil {
		t.Errorf("deleted non-existent item")
	}
	for i := 0; i < n; i++ {
		if u := tree.Delete(IntItem(i)); u == nil || int(u.(IntItem)) != i {
			t.Errorf("delete failed")
		}
	}
	if tree.Delete(IntItem(200)) != nil {
		t.Errorf("deleted non-existent item")
	}
	if tree.Delete(IntItem(-2)) != nil {
		t.Errorf("deleted non-existent item")
	}
}

func TestRandomInsertPartialDeleteOrder(t *testing.T) {
	tree := &Tree{}
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	for i := 1; i < n-1; i++ {
		tree.Delete(IntItem(i))
	}
	c := tree.Iter()
	if int((<-c).(IntItem)) != 0 {
		t.Errorf("expecting 0")
	}
	if int((<-c).(IntItem)) != n-1 {
		t.Errorf("expecting %d", n-1)
	}
}

func TestRandomInsertStats(t *testing.T) {
	tree := &Tree{}
	n := 100000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.InsertOrReplace(IntItem(perm[i]))
	}
	avg, _ := tree.HeightStats()
	expAvg := math.Log2(float64(n)) - 1.5
	if math.Fabs(avg-expAvg) >= 2.0 {
		t.Errorf("too much deviation from expected average height")
	}
}

func BenchmarkInsert(b *testing.B) {
	tree := &Tree{}
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(IntItem(b.N - i))
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	tree := &Tree{}
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(IntItem(b.N - i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(IntItem(i))
	}
}

func BenchmarkDeleteMin(b *testing.B) {
	b.StopTimer()
	tree := &Tree{}
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(IntItem(b.N - i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}
