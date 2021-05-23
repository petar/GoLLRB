package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

func Print(item llrb.Item) bool {
	i, ok := item.(llrb.Int)
	if !ok {
		return false
	}
	fmt.Println(int(i))
	return true
}

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.DeleteMin()
	tree.Delete(llrb.Int(4))
	tree.AscendGreaterOrEqual(tree.Min(), Print)
}
