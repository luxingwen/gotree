package gotree

import (
	"fmt"
	"strings"
	"testing"
)

func TestMount(t *testing.T) {
	root := NewTree(BoleNode, "root", nil)
	fg := NewTree(BoleNode, "fg", nil)

	login := NewTree(LeafNode, "login", nil)
	fg.Mount("fg", login)

	login1 := NewTree(BoleNode, "Login1", nil)
	reg := NewTree(LeafNode, "regist", nil)

	xiaomi := NewTree(BoleNode, "xiaomi", nil)
	xiaomi.Mount("xiaomi", login1)

	root.Mount("", fg)
	root.Mount("fg", reg)

	name := NewLeaf("name", nil)
	root.Mount("xiaomi", xiaomi)
	login1.Mount("", name)
	fg.Mount("", name)
	xiaomi.Mount("", name)
	root.Print("")

	r, ok := root.GetNodeByPath("fg/login")
	if ok {
		fmt.Println("get ", r.Name())
	} else {
		fmt.Println("not found ", "fg/login")
	}
}

func TestPath(t *testing.T) {
	root := NewBole("root", nil)
	mount(root, "d/go/src/cellorigin")
	mount(root, "d/go/src/git.kunqi.xyz/fgs/fglog/server/tree")
	mount(root, "d/go/src/github.com/ss")
	mount(root, "d/go/src/github.com/luxingwen")
	root.Print("")
}

func mount(mt Tree, Path string) {
	lgs1 := strings.Split(Path, "/")
	for i, item := range lgs1 {
		if i == len(lgs1)-1 {
			mt = mt.Mount("", NewLeaf(item, nil))
		} else {
			mt = mt.Mount("", NewBole(item, nil))
		}
	}
}
