package llrb

type ItemIterator func(i Item) bool

// AscendGreaterOrEqual will call iterator once for each element greater or equal to
// pivot in ascending order. It will stop whenever the iterator returns false.
func (t *Tree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	t.ascendGreaterOrEqual(t.root, pivot, iterator)
}

func (t *Tree) ascendGreaterOrEqual(h *Node, pivot Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if !h.Item.Less(pivot) {
		if !t.ascendGreaterOrEqual(h.Left, pivot, iterator) {
			return false
		}
		if !iterator(h.Item) {
			return false
		}
	}
	return t.ascendGreaterOrEqual(h.Right, pivot, iterator)
}

// DescendLessOrEqual will call iterator once for each element less than the
// pivot in descending order. It will stop whenever the iterator returns false.
func (t *Tree) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
	t.descendLessOrEqual(t.root, pivot, iterator)
}

func (t *Tree) descendLessOrEqual(h *Node, pivot Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if h.Item.Less(pivot) || !pivot.Less(h.Item) {
		if !t.descendLessOrEqual(h.Right, pivot, iterator) {
			return false
		}
		if !iterator(h.Item) {
			return false
		}
	}
	return t.descendLessOrEqual(h.Left, pivot, iterator)
}
