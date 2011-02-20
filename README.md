# GoLLRB

GoLLRB is a Left-Leaning Red-Black (LLRB) implementation of 2-3 balanced binary
search trees in Google Go.

## Overview

As of this writing and to the best of the author's knowledge, Google
Go still does not have a balanced binary search tree (BBST) data structure.
These data structures are quite useful in a variety of cases. A BBST maintains
elements in sorted order under dynamic updates (inserts and deletes) and can
support various order-specific queries. Furthermore, in practice one often
implements other common data structures like Priority Queues, using BBST's.

2-3 trees (a type of BBST's), as well as the runtime-similar 2-3-4 trees, are
the de facto standard BBST algoritms found in implementations of Python, Java,
and other libraries. The LLRB method of implementing 2-3 trees is a recent
improvement over the traditional implementation. The LLRB approach was
discovered relatively recently (in 2008) by Robert Sedgewick of Princeton
University.

GoLLRB is a Go implementation of LLRB 2-3 trees.

## Maturity

As of recently, GoLLRB has been used in some pretty heavy-weight machine
learning tasks over many gigabytes of data. I consider it to be in stable,
perhaps even production-grade, shape. There are no known bugs.

## Installation

With a healthy Google Go installed, simply run `goinstall github.com/petar/GoLLRB/llrb`

## Example
    
	package main

	import (
		"fmt"
		"github.com/petar/GoLLRB/llrb"
	)

	type IntItem int

	func (item IntItem) LessThan(other interface{}) bool {
		return int(item) < int(other.(IntItem))
	}

	func main() {
		var tree llrb.Tree
		tree.InsertOrReplace(IntItem(1))
		tree.InsertOrReplace(IntItem(2))
		tree.InsertOrReplace(IntItem(3))
		tree.InsertOrReplace(IntItem(4))
		tree.DeleteMin()
		tree.Delete(IntItem(4))
		c := tree.Iter()
		for {
			u := <-c
			if u == nil {
				break
			}
			fmt.Printf("%d\n", int(u.(IntItem)))
		}
	}

## About

GoLLRB was written by [Petar Maymounkov](http://pdos.csail.mit.edu/~petar/). 

Follow me on [Twitter @maymounkov](http://www.twitter.com/maymounkov)!

