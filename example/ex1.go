package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

func printIntItem(i llrb.Item) bool {
	fmt.Printf("%d\n", int(i.(llrb.Int)))
	// return false to terminate the traversal procedure
	return true
}

func main() {
	tree := llrb.New()
	// llrb.Int implemented the Less() method
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.AscendGreaterOrEqual(llrb.Int(0), printIntItem)
	fmt.Printf("-------\n")
	tree.DeleteMin()
	tree.Delete(llrb.Int(4))
	tree.AscendGreaterOrEqual(llrb.Int(0), printIntItem)
}
