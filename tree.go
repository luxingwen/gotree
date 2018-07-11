package gotree

import (
	"fmt"
	"strings"
)

const (
	BoleNode = iota
	LeafNode
)

type Tree interface {
	AddChild(Tree)
	Name() string
	SetName(string)
	Mount(root string, r Tree) Tree
	SetParent(Tree)
	GetParent() Tree
	GetNode(string) (Tree, bool)
	GetNodeByPath(path string) (Tree, bool)
	SetObj(interface{})
	GetObj() interface{}
	GetLeafs() []Tree
	Print(string)
}

type tree struct {
	name   string
	parent Tree
}

func (r *tree) AddChild(Tree) {}

func (r *tree) Name() string {
	return r.name
}

func (r *tree) SetName(name string) {
	r.name = name
}

func (r *tree) Mount(root string, c Tree) Tree {
	panic("not impl mount")
}

func (r *tree) SetParent(pre Tree) {
	r.parent = pre
}
func (r *tree) GetParent() Tree {
	return r.parent
}

func (r *tree) GetNode(string) (Tree, bool) {
	return nil, false
}
func (r *tree) Print(string) {}
func (r *tree) GetNodeByPath(path string) (Tree, bool) {
	return nil, false
}
func (r *tree) SetObj(interface{}) {}
func (r *tree) GetObj() interface{} {
	return nil
}

func (r *tree) GetLeafs() (res []Tree) {
	return
}

type Bole struct {
	tree
	leafs []Tree
	obj   interface{}
}

func NewBole(name string, obj interface{}) *Bole {
	b := &Bole{obj: obj, leafs: make([]Tree, 0)}
	b.tree.SetName(name)
	return b
}

func NewTree(kind int, name string, obj interface{}) Tree {
	var t Tree
	switch kind {
	case BoleNode:
		t = NewBole(name, obj)
	case LeafNode:
		t = NewLeaf(name, obj)
	default:
		panic("invalid kind")
	}
	return t
}

func (t *Bole) SetObj(obj interface{}) {
	t.obj = obj
}

func (t *Bole) AddChild(c Tree) {
	c.SetParent(t)
	t.leafs = append(t.leafs, c)
}

func (t *Bole) Mount(root string, r Tree) Tree {
	for _, item := range t.leafs {
		if item.Name() == r.Name() {
			return item
		}
	}
	if root == "" {
		t.AddChild(r)
		return r
	}

	if t.Name() == root {
		t.AddChild(r)
		return r
	}
	mrouter, ok := t.GetNode(root)
	if ok {
		mrouter.AddChild(r)
		return r
	}
	if root == r.Name() {
		t.AddChild(r)
		return r
	}
	mrouter = NewTree(BoleNode, root, nil)
	mrouter.AddChild(r)
	t.AddChild(mrouter)
	return r
}

func (t *Bole) GetNode(root string) (Tree, bool) {
	for _, item := range t.leafs {
		if item.Name() == root {
			return item, true
		}
		if v, ok := item.GetNode(root); ok {
			return v, true
		}
	}
	return nil, false
}

func (t *Bole) Print(pre string) {
	fmt.Println(pre + "+" + t.Name())
	pre += "  "
	for _, item := range t.leafs {
		item.Print(pre)
	}
}

func (t *Bole) GetNodeByPath(path string) (Tree, bool) {
	if path == "" {
		return nil, false
	}
	if path == t.Name() {
		return t, true
	}
	n := strings.Index(path, "/")
	name := path
	if n > 0 {
		name = path[:n]
		path = path[n+1:]
	}
	if name == t.Name() {
		return t, true
	}
	for i, item := range t.leafs {
		if item.Name() == path {
			return item, true
		}
		if item.Name() == name {
			return item.GetNodeByPath(path)
		}
		if i == len(t.leafs)-1 {
			return nil, false
		}
	}
	return nil, false
}

func (t *Bole) GetLeafs() []Tree {
	return t.leafs
}

type Leaf struct {
	tree
	obj interface{}
}

func NewLeaf(name string, obj interface{}) *Leaf {
	leaf := &Leaf{obj: obj}
	leaf.tree.SetName(name)
	return leaf
}

func (leaf *Leaf) GetObj() interface{} {
	return leaf.obj
}

func (leaf *Leaf) SetObj(obj interface{}) {
	leaf.obj = obj
}

func (leaf *Leaf) Print(pre string) {
	fmt.Println(pre + "-" + leaf.Name())
}
