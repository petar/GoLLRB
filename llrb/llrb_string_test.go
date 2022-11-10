package llrb

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tree := New()
	for i := 0; i != 10; i++ {
		tree.InsertNoReplace(Int(i))
	}

	want := `
     ┌─ 0
   ┌─ 1
   │ └─ 2
─── 3
   │   ┌─ 4
   │ ┌─ 5
   │ │ └─ 6
   └─ 7
     │ ┌─ 8
     └─ 9
`
	want = want[1:]
	got := tree.String()
	if got != want {
		t.Errorf("the output is not expected")
		fmt.Printf("got:\n%s\nwant:\n%s\n", got, want)
	}
}
