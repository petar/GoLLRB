// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A Left-Leaning Red-Black (LLRB) implementation of 2-3 balanced binary search trees

package llrb

// |Tree| is a Left-Leaning Red-Black (LLRB) implementation of 2-3 trees, based on:
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
// |Tree| has an associative interface, i.e. duplicate key are not allowed.
// The zero-value of a |Tree| represents a ready-for-use tree.
type Tree struct {
	count int
	root  *node
}

// An |Item| represents an object that can be inserted in |Tree|. It acts as a
// key via the method |LessThan|, which induces a full ordering on |Item|s. It is
// also a value, as the user can attach any data to it.
type Item interface {

	// |LessThan| returns true, if and only if |this| is STRICTLY less than |other|
	LessThan(other interface{}) bool
}

// Init resets (empties) the tree
func (t *Tree) Init() {
	t.root = nil
}

func (t *Tree) Len() int { return t.count }

// |Has| returns true if the tree contains an element
// whose |LessThan| order equals that of |key|.
func (t *Tree) Has(key Item) bool {
	return t.Get(key) != nil
}

// |Get| retrieves an element from the tree whose |LessThan| order
// equals that of |key|.
func (t *Tree) Get(key Item) Item {
	return get(t.root, key)
}

func get(h *node, item Item) Item {
	if h == nil {
		return nil
	}
	if item.LessThan(h.item) {
		return get(h.left, item)
	}
	if h.item.LessThan(item) {
		return get(h.right, item)
	}
	return h.item
}

// |Min| returns the minimum element in the tree.
func (t *Tree) Min() Item {
	return min(t.root)
}

func min(h *node) Item {
	if h == nil {
		return nil
	}
	if h.left == nil {
		return h.item
	}
	return min(h.left)
}

// |Max| returns the maximum element in the tree.
func (t *Tree) Max() Item {
	return max(t.root)
}

func max(h *node) Item {
	if h == nil {
		return nil
	}
	if h.right == nil {
		return h.item
	}
	return max(h.right)
}

// |Insert| inserts a new element in the tree, or replaces
// an existing one of identical |LessThan| order.
// If a replacement occurred, the replaced item is returned.
func (t *Tree) InsertOrReplace(item Item) Item {
	if item == nil {
		panic("inserting nil item")
	}
	var replaced Item
	t.root, replaced = insert(t.root, item)
	t.root.black = true
	if replaced == nil {
		t.count++
	}
	return replaced
}

func insert(h *node, item Item) (*node, Item) {
	if h == nil {
		return newNode(item), nil
	}

	// PLACEHOLDER: 2-3-4 tree (see comment below)

	var replaced Item
	if item.LessThan(h.item) {
		h.left, replaced = insert(h.left, item)
	} else if h.item.LessThan(item) {
		h.right, replaced = insert(h.right, item)
	} else {
		replaced, h.item = h.item, item
	}

	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}

	// PETAR: added |h.left != nil|
	if h.left != nil && isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}

	// When the next 3 lines of code are here, the LLRB behaves
	// like a 2-3 tree. If they are moved to the 2-3-4 placeholder above,
	// the LLRB tree behaves like a 2-3-4 tree.
	if isRed(h.left) && isRed(h.right) {
		flip(h)
	}

	return h, replaced
}

// |DeleteMin| deletes the minimum element in the tree and returns the
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

// deleteMin code for LLRB 2-3 trees
func deleteMin(h *node) (*node, Item) {
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

// |DeleteMax| deletes the maximum element in the tree and returns
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

func deleteMax(h *node) (*node, Item) {
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

// |Delete| deletes an item from the tree whose key equals |key|.
// The deleted item is return, otherwise nil is returned.
func (t *Tree) Delete(key Item) Item {
	var deleted Item
	t.root, deleted = delete(t.root, key)
	if t.root != nil {
		t.root.black = true
	}
	if deleted != nil {
		t.count--
	}
	return deleted
}

func delete(h *node, item Item) (*node, Item) {
	var deleted Item
	if h == nil {
		return nil, nil
	}
	if item.LessThan(h.item) {
		if h.left == nil { // item not present. Nothing to delete
			return h, nil
		}
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left, deleted = delete(h.left, item)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		// If |item| equals |h.item| and no right children at |h|
		if !h.item.LessThan(item) && h.right == nil {
			return nil, h.item
		}
		// PETAR: Added |h.right != nil| below
		if h.right != nil && !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		// If |item| equals |h.item|, and (from above) |h.right != nil|
		if !h.item.LessThan(item) {
			var subDeleted Item
			h.right, subDeleted = deleteMin(h.right)
			if subDeleted == nil {
				panic("logic")
			}
			deleted, h.item = h.item, subDeleted
		} else { // Else, |item| is bigger than |h.item|
			h.right, deleted = delete(h.right, item)
		}
	}

	return fixUp(h), deleted
}

// |Iter| returns a chan that iterates through all elements in the
// tree in ascending order.
func (t *Tree) Iter() <-chan Item {
	c := make(chan Item)
	go func() {
		iterateInOrder(t.root, c)
		close(c)
	}()
	return c
}

func iterateInOrder(h *node, c chan<- Item) {
	if h == nil {
		return
	}
	iterateInOrder(h.left, c)
	c <- h.item
	iterateInOrder(h.right, c)
}

func iteratePreOrder(h *node, c chan<- Item) {
	if h == nil {
		return
	}
	c <- h.item
	iteratePreOrder(h.left, c)
	iteratePreOrder(h.right, c)
}

func iteratePostOrder(h *node, c chan<- Item) {
	if h == nil {
		return
	}
	iteratePostOrder(h.left, c)
	iteratePostOrder(h.right, c)
	c <- h.item
}

type node struct {
	item        Item
	left, right *node // Pointers to left and right child nodes
	black       bool  // If set, the color of the link (incoming from the parent) is black
	// In the LLRB, new nodes are always red, hence the zero-value for node
}

func newNode(item Item) *node { return &node{item: item} }

func isRed(h *node) bool {
	if h == nil {
		return false
	}
	return !h.black
}

func rotateLeft(h *node) *node {
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

func rotateRight(h *node) *node {
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

// Left and right children must be present
func flip(h *node) {
	h.black = !h.black
	h.left.black = !h.left.black
	h.right.black = !h.right.black
}

// Left and right children must be present
func moveRedLeft(h *node) *node {
	flip(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flip(h)
	}
	return h
}

// Left and right children must be present
func moveRedRight(h *node) *node {
	flip(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flip(h)
	}
	return h
}

func fixUp(h *node) *node {
	if isRed(h.right) {
		h = rotateLeft(h)
	}

	// PETAR: added |h.left != nil|
	if h.left != nil && isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}

	if isRed(h.left) && isRed(h.right) {
		flip(h)
	}

	return h
}
