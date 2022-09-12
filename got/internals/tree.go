package internals

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"
)

type Entry struct {
	ID, Name string
	Mode     fs.FileMode
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
		fmt.Fprintf(&b, "%s %s\t%s\n", entry.Mode, entry.ID, entry.Name)
	}

	return &Tree{
		Object:  NewObject("tree", []byte(b.String())),
		Entries: entries,
	}
}

func (tree *Tree) String() string {
	return string(tree.Data)
}
