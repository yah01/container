package container

import (
	"golang.org/x/exp/constraints"
)

const (
	LeafNodeType = iota
	InternalNodeType
)

type NodeType int

type BTreeNode[K constraints.Ordered, V any] interface {
	Insert(K, V)
	Get(K) *V
	Len() int
	Cap() int
	NodeType() NodeType
}

type LeafNode[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
}

func NewLeafNode[K constraints.Ordered, V any](fanout int) *LeafNode[K, V] {
	return &LeafNode[K, V]{
		keys:   make([]K, 0, fanout-1),
		values: make([]V, 0, fanout-1),
	}
}

func (node *LeafNode[K, V]) Insert(key K, value V) {
	if len(node.keys) == 0 {
		node.keys = append(node.keys, key)
		node.values = append(node.values, value)
		return
	}

	if node.Len() == node.Cap() {
		panic("insert into a full leaf node")
	}

	i := Search(node.keys, key)
	node.keys = append(node.keys[:i+1], node.keys[i:]...)
	node.keys[i] = key

	node.values = append(node.values[:i+1], node.values[i:]...)
	node.values[i] = value
}

func (node *LeafNode[K, V]) Get(key K) *V {
	if node.Len() == 0 {
		return nil
	}

	i := Search(node.keys, key)
	if node.keys[i] != key {
		return nil
	}

	return &node.values[i]
}

func (node *LeafNode[K, V]) Len() int {
	return len(node.keys)
}

func (node *LeafNode[K, V]) Cap() int {
	return cap(node.keys)
}

func (node *LeafNode[K, V]) NodeType() NodeType {
	return LeafNodeType
}

func (node *LeafNode[K, V]) Split() (*LeafNode[K, V], K, V) {
	var (
		mid      = node.Len() / 2
		midKey   = node.keys[mid]
		midValue = node.values[mid]
	)

	newLeaf := NewLeafNode[K, V](node.Cap() + 1)
	newLeaf.keys = append(newLeaf.keys, node.keys[mid+1:]...)
	newLeaf.values = append(newLeaf.values, node.values[mid+1:]...)

	node.keys = node.keys[:mid]
	node.values = node.values[:mid]

	return newLeaf, midKey, midValue
}

type InternalNode[K constraints.Ordered, V any] struct {
	*LeafNode[K, V]
	children []BTreeNode[K, V]
}

func (node *InternalNode[K, V]) Insert(key K, value V) {
	if node.Len() < node.Cap() {
		node.LeafNode.Insert(key, value)
		return
	}

	// i := Search(node.keys, key)

}

func (node *InternalNode[K, V]) findChild(key K) BTreeNode[K, V] {
	i := Search(node.keys, key)
	return node.children[i]
}

type BTree[K constraints.Ordered, V any] struct {
	root BTreeNode[K, V]
}

func NewBTree[K constraints.Ordered, V any](fanout int) *BTree[K, V] {
	return &BTree[K, V]{
		root: NewLeafNode[K, V](fanout),
	}
}

func (tree *BTree[K, V]) Get(key K) *V {
	node := tree.root
	for ; node.NodeType() != LeafNodeType; node = node.(*InternalNode[K, V]).findChild(key) {
		value := node.Get(key)
		if value != nil {
			return value
		}
	}

	return node.Get(key)
}

func (tree *BTree[K, V]) Insert(key K, value V) {
	path := tree.leafPath(tree.root, key)
	leaf := path[len(path)-1].(*LeafNode[K, V])

	if leaf.Len() < leaf.Cap() {
		leaf.Insert(key, value)
		return
	}

	// newLeaf,midKey,midValue := leaf.Split()

	// for i := len(path) - 1; i >= 0; i-- {
	// 	node := path[i]
	// 	if node.Len() < node.Cap() {
	// 		node.Insert(key, value)
	// 		return
	// 	}
	// }

	// return node.Get(key)
}

func (tree *BTree[K, V]) leafPath(node BTreeNode[K, V], key K) []BTreeNode[K, V] {
	nodes := make([]BTreeNode[K, V], 0)

	for ; node.NodeType() != LeafNodeType; node = node.(*InternalNode[K, V]).findChild(key) {
		nodes = append(nodes, node)
	}

	nodes = append(nodes, node)

	return nodes
}
