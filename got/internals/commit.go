package internals

import (
	"fmt"
	"strings"
)

type Commit struct {
	Message string
	*Object
	Tree
	Author
}

func NewCommit(tree *Tree, author *Author, message string) *Commit {
	lines := []string{
		fmt.Sprintf("Tree: %s", tree.String()),
		fmt.Sprintf("Author: %s", author.String()),
		"",
		message,
	}
	data := strings.Join(lines, "\n")
	return &Commit{
		Tree:    *tree,
		Author:  *author,
		Message: message,
		Object:  &Object{ObjectType: "commit", Data: []byte(data)},
	}
}

func (commit *Commit) String() string {
	return string(commit.Data)
}
