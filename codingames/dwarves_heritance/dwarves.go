package main

import (
	"fmt"
	"strconv"
)

func (t InfluenceTree) asInfluence(a int) (bool, InfluenceTree) {
	if len(t.Trees) == 0 {
		return false, EmptyTree
	}

	for _, influence := range t.Trees {
		if influence.ID == a {
			return true, influence
		}
		if ok, inf := influence.asInfluence(a); ok {
			return true, inf
		}
	}

	return false, EmptyTree
}

// addInfluenceRec computes influence tree and add new influence at the right place.
func (t InfluenceTree) addInfluenceRec(a, b int, previousTree InfluenceTree, rootTree InfluenceTree) (_ InfluenceTree, addedValue bool) {
	if t.ID == a {
		ok, inf := rootTree.asInfluence(b)
		if !ok {
			inf = InfluenceTree{ID: b}
		}

		t.Trees = append(t.Trees, inf)

		return t, true
	}

	if t.ID == b && previousTree.ID != EmptyTree.ID {
		return InfluenceTree{ID: a, Trees: []InfluenceTree{previousTree}}, true
	}

	if len(t.Trees) != 0 {
		addedValue = false

		for i, subTree := range t.Trees {
			passTree := EmptyTree
			if t.ID == RootTree.ID {
				passTree = subTree
			}
			newInfluences, ok := subTree.addInfluenceRec(a, b, passTree, rootTree)
			if ok {
				addedValue = true
				t.Trees[i] = newInfluences
			}
		}

		if addedValue || t.ID != -1 {
			return t, addedValue
		}
	}

	if t.ID == RootTree.ID {
		t.Trees = append(t.Trees, InfluenceTree{ID: a, Trees: []InfluenceTree{{ID: b}}})
		return t, true
	}

	return EmptyTree, false
}

var (
	EmptyTree = InfluenceTree{ID: -9999}
	RootTree  = InfluenceTree{ID: -1}
)

// InfluenceTree stores influence paths
// ID is the influencer
// Trees is a list of influence path representing
//      dwarves influenced by influencer
type InfluenceTree struct {
	ID    int
	Trees []InfluenceTree
}

// NewTree instantiate a new root tree
func NewTree() InfluenceTree {
	return InfluenceTree{ID: RootTree.ID}
}

// String creates a clean string representation for influence tree
func (t InfluenceTree) String() (res []string) {
	id := ""
	if t.ID != RootTree.ID {
		id = strconv.Itoa(t.ID) // converts int id to string
	}

	if len(t.Trees) == 0 {
		return []string{id + "|"} // current tree is a leaf so we have a single path
	}

	for _, subTree := range t.Trees {
		paths := subTree.String()    // get subtrees paths
		for _, path := range paths { // for each sub path, we add current influence
			res = append(res, id+"->"+path)
		}
	}

	return // return computed paths.
}

// LongestPath computes the longest influence path in a tree
func (t InfluenceTree) LongestPath() int {
	if len(t.Trees) == 0 {
		return 1 // leaf is an element so length of path is 1
	}

	maxLen := 0
	for _, subTree := range t.Trees {
		candidate := subTree.LongestPath() // compute subTree longest path
		if candidate > maxLen {            // if new longestPath is longer than current retain max, replace it
			maxLen = candidate
		}
	}

	if t.ID != RootTree.ID { // if t is not root, current root is a path elements so we add 1 to len
		return maxLen + 1
	}

	return maxLen // else return result
}

// AddInfluence add influence to influence tree.
func (t InfluenceTree) AddInfluence(a, b int) InfluenceTree {
	newTree, _ := t.addInfluenceRec(a, b, EmptyTree, t)
	return newTree
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)

	influences := NewTree()

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		influences = influences.AddInfluence(x, y)
	}

	// The number of people involved in the longest succession of influences
	fmt.Println(influences.String())
	fmt.Println(influences.LongestPath())
}
