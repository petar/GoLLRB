package main

import (
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

type Int int

// implement llrb.Item interface
func (a Int) Less(b llrb.Item) bool { return a < b.(Int) }

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(Int(1))
	tree.ReplaceOrInsert(Int(2))
	tree.ReplaceOrInsert(Int(3))
	tree.ReplaceOrInsert(Int(4))
	tree.DeleteMin()
	tree.Delete(Int(4))
	tree.AscendGreaterOrEqual(tree.Min(), func(item llrb.Item) bool {
		fmt.Printf("%v\n", item)
		return true
	})
}
