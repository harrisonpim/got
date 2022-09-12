package internals

import (
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	ID, Name string
}

type Tree struct {
	*Object
	Entries []Entry
}

func NewTree(entries []Entry) *Tree {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	var b strings.Builder
	for _, entry := range entries {
		fmt.Fprintf(&b, "100644 %s\t%s\n", entry.ID, entry.Name)
	}

	tree := Tree{
		Object:  &Object{ObjectType: "tree", Data: []byte(b.String())},
		Entries: entries,
	}

	return &tree
}

func (tree *Tree) String() string {
	return string(tree.Data)
}
