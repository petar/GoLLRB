// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A Left-Leaning Red-Black (LLRB) implementation of 2-3 balanced binary search trees

package llrb

// Tree{} is a Left-Leaning Red-Black (LLRB) implementation of 2-3 trees, based on:
//
//   http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf
//   http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//   http://www.cs.princeton.edu/~rs/talks/LLRB/Java/RedBlackBST.java
//
//  2-3 trees (and the run-time equivalent 2-3-4 trees) are the de facto standard BST
//  algoritms found in implementations of Python, Java, and other libraries. The LLRB
//  implementation of 2-3 trees is a recent improvement on the traditional implementation,
//  observed and documented by Robert Sedgewick.
//
type Tree struct {
	less  LessFunc
	count int
	root  *Node
}

type Item interface{}

type LessFunc func(a, b interface{}) bool

// New() allocates a new tree
func New(lessfunc LessFunc) *Tree {
	t := &Tree{}
	t.Init(lessfunc)
	return t
}

// Init() resets (empties) the tree
func (t *Tree) Init(lessfunc LessFunc) {
	t.less = lessfunc
	t.root = nil
	t.count = 0
}

func (t *Tree) Len() int { return t.count }

// Has() returns true if the tree contains an element
// whose LessThan() order equals that of @key.
func (t *Tree) Has(key Item) bool {
	return t.Get(key) != nil
}

// Get() retrieves an element from the tree whose LessThan() order
// equals that of @key.
func (t *Tree) Get(key Item) Item {
	return t.get(t.root, key)
}

func (t *Tree) get(h *Node, item Item) Item {
	if h == nil {
		return nil
	}
	if t.less(item, h.item) {
		return t.get(h.left, item)
	}
	if t.less(h.item, item) {
		return t.get(h.right, item)
	}
	return h.item
}

// Min() returns the minimum element in the tree.
func (t *Tree) Min() Item {
	return min(t.root)
}

func min(h *Node) Item {
	if h == nil {
		return nil
	}
	if h.left == nil {
		return h.item
	}
	return min(h.left)
}

// Max() returns the maximum element in the tree.
func (t *Tree) Max() Item {
	return max(t.root)
}

func max(h *Node) Item {
	if h == nil {
		return nil
	}
	if h.right == nil {
		return h.item
	}
	return max(h.right)
}

func (t *Tree) ReplaceOrInsertBulk(items ...Item) {
	for _, i := range items {
		t.ReplaceOrInsert(i)
	}
}

func (t *Tree) InsertNoReplaceBulk(items ...Item) {
	for _, i := range items {
		t.InsertNoReplace(i)
	}
}

// ReplaceOrInsert() inserts @item into the tree. If an existing
// element has the same order, it is removed from the tree and returned.
func (t *Tree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("inserting nil item")
	}
	var replaced Item
	t.root, replaced = t.replaceOrInsert(t.root, item)
	t.root.black = true
	if replaced == nil {
		t.count++
	}
	return replaced
}

// InsertOrReplace() inserts @item into the tree. If an existing
// element has the same order, both elements remain in the tree.
func (t *Tree) InsertNoReplace(item Item) {
	if item == nil {
		panic("inserting nil item")
	}
	t.root = t.insertNoReplace(t.root, item)
	t.root.black = true
	t.count++
}

func (t *Tree) replaceOrInsert(h *Node, item Item) (*Node, Item) {
	if h == nil {
		return newNode(item), nil
	}

	h = walkDownRot23(h)

	var replaced Item
	if t.less(item, h.item) {
		h.left, replaced = t.replaceOrInsert(h.left, item)
	} else if t.less(h.item, item) {
		h.right, replaced = t.replaceOrInsert(h.right, item)
	} else {
		replaced, h.item = h.item, item
	}

	h = walkUpRot23(h)

	return h, replaced
}

func (t *Tree) insertNoReplace(h *Node, item Item) *Node {
	if h == nil {
		return newNode(item)
	}

	h = walkDownRot23(h)

	if t.less(item, h.item) {
		h.left = t.insertNoReplace(h.left, item)
	} else {
		h.right = t.insertNoReplace(h.right, item)
	}

	return walkUpRot23(h)
}

// Rotation driver routines for 2-3 algorithm

func walkDownRot23(h *Node) *Node { return h }

func walkUpRot23(h *Node) *Node {
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}

	// PETAR: added 'h.left != nil'
	if h.left != nil && isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}

	if isRed(h.left) && isRed(h.right) {
		flip(h)
	}

	return h
}

// Rotation driver routines for 2-3-4 algorithm

func walkDownRot234(h *Node) *Node {
	if isRed(h.left) && isRed(h.right) {
		flip(h)
	}

	return h
}

func walkUpRot234(h *Node) *Node {
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}

	// PETAR: added 'h.left != nil'
	if h.left != nil && isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}

	return h
}

// DeleteMin() deletes the minimum element in the tree and returns the
// deleted item or nil otherwise.
func (t *Tree) DeleteMin() Item {
	var deleted Item
	t.root, deleted = deleteMin(t.root)
	if t.root != nil {
		t.root.black = true
	}
	if deleted != nil {
		t.count--
	}
	return deleted
}

// deleteMin() code for LLRB 2-3 trees
func deleteMin(h *Node) (*Node, Item) {
	if h == nil {
		return nil, nil
	}
	if h.left == nil {
		return nil, h.item
	}

	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}

	var deleted Item
	h.left, deleted = deleteMin(h.left)

	return fixUp(h), deleted
}

// DeleteMax() deletes the maximum element in the tree and returns
// the deleted item or nil otherwise
func (t *Tree) DeleteMax() Item {
	var deleted Item
	t.root, deleted = deleteMax(t.root)
	if t.root != nil {
		t.root.black = true
	}
	if deleted != nil {
		t.count--
	}
	return deleted
}

func deleteMax(h *Node) (*Node, Item) {
	if h == nil {
		return nil, nil
	}
	if isRed(h.left) {
		h = rotateRight(h)
	}
	if h.right == nil {
		return nil, h.item
	}
	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}
	var deleted Item
	h.right, deleted = deleteMax(h.right)

	return fixUp(h), deleted
}

// Delete() deletes an item from the tree whose key equals @key.
// The deleted item is return, otherwise nil is returned.
func (t *Tree) Delete(key Item) Item {
	var deleted Item
	t.root, deleted = t.delete(t.root, key)
	if t.root != nil {
		t.root.black = true
	}
	if deleted != nil {
		t.count--
	}
	return deleted
}

func (t *Tree) delete(h *Node, item Item) (*Node, Item) {
	var deleted Item
	if h == nil {
		return nil, nil
	}
	if t.less(item, h.item) {
		if h.left == nil { // item not present. Nothing to delete
			return h, nil
		}
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left, deleted = t.delete(h.left, item)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		// If @item equals @h.item and no right children at @h
		if !t.less(h.item, item) && h.right == nil {
			return nil, h.item
		}
		// PETAR: Added 'h.right != nil' below
		if h.right != nil && !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		// If @item equals @h.item, and (from above) 'h.right != nil'
		if !t.less(h.item, item) {
			var subDeleted Item
			h.right, subDeleted = deleteMin(h.right)
			if subDeleted == nil {
				panic("logic")
			}
			deleted, h.item = h.item, subDeleted
		} else { // Else, @item is bigger than @h.item
			h.right, deleted = t.delete(h.right, item)
		}
	}

	return fixUp(h), deleted
}

// IterAscend() returns a chan that iterates through all elements in
// in ascending order.
func (t *Tree) IterAscend() <-chan Item {
	c := make(chan Item)
	go func() {
		iterateInOrder(t.root, c)
		close(c)
	}()
	return c
}

// IterDescend() returns a chan that iterates through all elements
// in descending order.
func (t *Tree) IterDescend() <-chan Item {
	c := make(chan Item)
	go func() {
		iterateInOrderRev(t.root, c)
		close(c)
	}()
	return c
}

// IterRangeInclusive() returns a chan that iterates through all elements E in the
// tree with @lower <= E <= @upper in ascending order.
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
	lessThanLower := t.less(h.item, lower)
	greaterThanUpper := t.less(upper, h.item)
	if !lessThanLower {
		t.iterateRangeInclusive(h.left, c, lower, upper)
	}
	if !lessThanLower && !greaterThanUpper {
		c <- h.item
	}
	if !greaterThanUpper {
		t.iterateRangeInclusive(h.right, c, lower, upper)
	}
}

// IterRange() returns a chan that iterates through all elements E in the
// tree with @lower <= E < @upper in ascending order.
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
	lessThanLower := t.less(h.item, lower)
	lessThanUpper := t.less(h.item, upper)
	if !lessThanLower {
		t.iterateRange(h.left, c, lower, upper)
	}
	if !lessThanLower && lessThanUpper {
		c <- h.item
	}
	if lessThanUpper {
		t.iterateRange(h.right, c, lower, upper)
	}
}

func iterateInOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iterateInOrder(h.left, c)
	c <- h.item
	iterateInOrder(h.right, c)
}

func iterateInOrderRev(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iterateInOrderRev(h.right, c)
	c <- h.item
	iterateInOrderRev(h.left, c)
}

func iteratePreOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	c <- h.item
	iteratePreOrder(h.left, c)
	iteratePreOrder(h.right, c)
}

func iteratePostOrder(h *Node, c chan<- Item) {
	if h == nil {
		return
	}
	iteratePostOrder(h.left, c)
	iteratePostOrder(h.right, c)
	c <- h.item
}

type Node struct {
	item        Item
	left, right *Node // Pointers to left and right child nodes
	black       bool  // If set, the color of the link (incoming from the parent) is black
	// In the LLRB, new nodes are always red, hence the zero-value for node
}

func newNode(item Item) *Node { return &Node{item: item} }

func isRed(h *Node) bool {
	if h == nil {
		return false
	}
	return !h.black
}

func rotateLeft(h *Node) *Node {
	x := h.right
	if x.black {
		panic("rotating a black link")
	}
	h.right = x.left
	x.left = h
	x.black = h.black
	h.black = false
	return x
}

func rotateRight(h *Node) *Node {
	x := h.left
	if x.black {
		panic("rotating a black link")
	}
	h.left = x.right
	x.right = h
	x.black = h.black
	h.black = false
	return x
}

// XXX:
//      - Can flip() ever be called with |h.left == nil| or |h.right == nil|?

// REQUIRE: Left and right children must be present
func flip(h *Node) {
	h.black = !h.black
	h.left.black = !h.left.black
	h.right.black = !h.right.black
}

// REQUIRE: Left and right children must be present
func moveRedLeft(h *Node) *Node {
	flip(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flip(h)
	}
	return h
}

// REQUIRE: Left and right children must be present
func moveRedRight(h *Node) *Node {
	flip(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flip(h)
	}
	return h
}

func fixUp(h *Node) *Node {
	if isRed(h.right) {
		h = rotateLeft(h)
	}

	// PETAR: added 'h.left != nil'
	if h.left != nil && isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}

	if isRed(h.left) && isRed(h.right) {
		flip(h)
	}

	return h
}
