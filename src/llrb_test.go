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

func TestCases(t *testing.T) {
	tree := &Tree{}
	tree.Insert(IntItem(1))
	tree.Insert(IntItem(1))
	//fmt.Printf("L=%d\n", tree.Len())
	if !tree.Has(IntItem(1)) {
		t.Errorf("expecting to find key=1")
	}

	tree.Delete(IntItem(1))
	//fmt.Printf("L=%d\n", tree.Len())
	if tree.Has(IntItem(1)) {
		t.Errorf("not expecting to find key=1")
	}

	tree.Delete(IntItem(1))
	//fmt.Printf("L=%d\n", tree.Len())
	if tree.Has(IntItem(1)) {
		t.Errorf("not expecting to find key=1")
	}
}

func TestReverseInsertOrder(t *testing.T) {
	tree := &Tree{}
	n := 100
	for i := 0; i < n; i++ {
		tree.Insert(IntItem(n-i))
	}
	c := tree.Iter()
	for j, item := 1, <-c; item != nil; j, item = j+1, <-c {
		if int(item.(IntItem)) != j {
			t.Fatalf("bad order")
		}
	}
}

func TestRandomInsertOrder(t *testing.T) {
	tree := &Tree{}
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Insert(IntItem(perm[i]))
	}
	c := tree.Iter()
	for j, item := 0, <-c; item != nil; j, item = j+1, <-c {
		if int(item.(IntItem)) != j {
			t.Fatalf("bad order")
		}
	}
}

func TestRandomInsertSequentialDelete(t *testing.T) {
	tree := &Tree{}
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Insert(IntItem(perm[i]))
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
		tree.Insert(IntItem(perm[i]))
	}
	tree.Delete(IntItem(200))
	tree.Delete(IntItem(-2))
	for i := 0; i < n; i++ {
		tree.Delete(IntItem(i))
	}
	tree.Delete(IntItem(200))
	tree.Delete(IntItem(-2))
}

func TestRandomInsertPartialDeleteOrder(t *testing.T) {
	tree := &Tree{}
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Insert(IntItem(perm[i]))
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
		tree.Insert(IntItem(perm[i]))
	}
	avg, _ := tree.HeightStats()
	expAvg := math.Log2(float64(n)) - 1.5
	if math.Fabs(avg - expAvg) >= 2.0 {
		t.Errorf("too much deviation from expected average height")
	}
}

func BenchmarkInsert(b *testing.B) {
	tree := &Tree{}
	for i := 0; i < b.N; i++ {
		tree.Insert(IntItem(b.N-i))
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	tree := &Tree{}
	for i := 0; i < b.N; i++ {
		tree.Insert(IntItem(b.N-i))
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
		tree.Insert(IntItem(b.N-i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}
