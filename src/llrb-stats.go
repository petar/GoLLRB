// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package llrb

// |GetHeight| returns an item in the tree with key |key|, and it's height in the tree
func (t *Tree) GetHeight(key Item) (result Item, depth int) {
	return getHeight(t.root, key)
}

func getHeight(h *node, item Item) (Item, int) {
	if h == nil {
		return nil, 0
	}
	if item.LessThan(h.item) {
		result, depth := getHeight(h.left, item)
		return result, depth+1
	}
	if h.item.LessThan(item) {
		result, depth := getHeight(h.right, item)
		return result, depth+1
	}
	return h.item, 0
}

// |HeightStats| returns the average and standard deviation of the height
// of elements in the tree
func (t *Tree) HeightStats() (avg, stddev float64) {
	av := &avgVar{}
	heightStats(t.root, 0, av)
	return av.GetAvg(), av.GetStdDev()
}

func heightStats(h *node, d int, av *avgVar) {
	if h == nil {
		return
	}
	av.Add(float64(d))
	if h.left != nil {
		heightStats(h.left, d+1, av)
	}
	if h.right != nil {
		heightStats(h.right, d+1, av)
	}
}
