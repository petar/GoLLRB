package main

import (
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

func lessInt(a, b interface{}) bool { return a.(int) < b.(int) }

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.DeleteMin()         // Delete 1
	tree.Delete(llrb.Int(4)) // Delete 4

	// Will output:
	// 2 llrb.Int
	// 3 llrb.Int
	tree.AscendGreaterOrEqual(llrb.Int(-1), func(i llrb.Item) bool {
		fmt.Printf("%d %T\n", i, i)
		return true
	})
}
