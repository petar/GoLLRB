// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package llrb

// IterAscend returns a chan that iterates through all elements in
// in ascending order.
// TODO: This is a deprecated interface for iteration.
func (t *Tree) IterAscend() <-chan Item {
	c := make(chan Item)
	go func() {
		iterateInOrder(t.root, c)
		close(c)
	}()
	return c
}

// IterDescend returns a chan that iterates through all elements
// in descending order.
// TODO: This is a deprecated interface for iteration.
func (t *Tree) IterDescend() <-chan Item {
	c := make(chan Item)
	go func() {
		iterateInOrderRev(t.root, c)
		close(c)
	}()
	return c
}

// IterRangeInclusive returns a chan that iterates through all elements E in the
// tree with lower <= E <= upper in ascending order.
// TODO: This is a deprecated interface for iteration.
func (t *Tree) IterRangeInclusive(lower, upper Item) <-chan Item {
	c := make(chan Item)
	go func() {
		t.iterateRangeInclusive(t.root, c, lower, upper)
		close(c)
	}()
	return c
}

func (t *Tree) iterateRangeInclusive(h *Node, c chan<- Item, lower, upper Item) {
	if h == nil {
		return
	}
	lessThanLower := t.less(h.Item, lower)
	greaterThanUpper := t.less(upper, h.Item)
	if !lessThanLower {
		t.iterateRangeInclusive(h.Left, c, lower, upper)
	}
	if !lessThanLower && !greaterThanUpper {
		c <- h.Item
	}
	if !greaterThanUpper {
		t.iterateRangeInclusive(h.Right, c, lower, upper)
	}
}

// IterRange() returns a chan that iterates through all elements E in the
// tree with lower <= E < upper in ascending order.
// TODO: This is a deprecated interface for iteration.
func (t *Tree) IterRange(lower, upper Item) <-chan Item {
	c := make(chan Item)
	go func() {
		t.iterateRange(t.root, c, lower, upper)
		close(c)
	}()
	return c
}

func (t *Tree) iterateRange(h *Node, c chan<- Item, lower, upper Item) {
	if h == nil {
		return
	}
	lessThanLower := t.less(h.Item, lower)
	lessThanUpper := t.less(h.Item, upper)
	if !lessThanLower {
		t.iterateRange(h.Left, c, lower, upper)
	}
	if !lessThanLower && lessThanUpper {
		c <- h.Item
	}
	if lessThanUpper {
		t.iterateRange(h.Right, c, lower, upper)
	}
}

func iterateInOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iterateInOrder(h.Left, c)
	c <- h.Item
	iterateInOrder(h.Right, c)
}

func iterateInOrderRev(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iterateInOrderRev(h.Right, c)
	c <- h.Item
	iterateInOrderRev(h.Left, c)
}

func iteratePreOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	c <- h.Item
	iteratePreOrder(h.Left, c)
	iteratePreOrder(h.Right, c)
}

func iteratePostOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iteratePostOrder(h.Left, c)
	iteratePostOrder(h.Right, c)
	c <- h.Item
}
