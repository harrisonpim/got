package internals

import (
	"fmt"
	"strings"
)

type Commit struct {
	ID, Message string
	*Object
	Tree
	Author
}

func NewCommit(tree *Tree, author *Author, message string) *Commit {
	lines := []string{
		fmt.Sprintf("tree %s", tree.ID),
		fmt.Sprintf("author %s", author.String()),
		fmt.Sprintf("committer %s", author.String()),
		"",
		message,
	}
	data := strings.Join(lines, "\n")
	object := &Object{ObjectType: "commit", Data: []byte(data)}
	id, _ := object.Hash()
	return &Commit{
		ID:      id,
		Tree:    *tree,
		Author:  *author,
		Message: message,
		Object:  object,
	}
}

func (commit *Commit) String() string {
	return string(commit.Data)
}
