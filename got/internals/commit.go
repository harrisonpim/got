package internals

import (
	"fmt"
	"strings"
)

type Commit struct {
	Parent, Message string
	*Object
	Tree
	Author
}

func NewCommit(parent string, tree *Tree, author *Author, message string) *Commit {
	lines := []string{}
	if parent == "" {
		lines = []string{
			fmt.Sprintf("tree %s", tree.ID),
			fmt.Sprintf("author %s", author.String()),
			fmt.Sprintf("committer %s", author.String()),
			"",
			message,
		}
	} else {
		lines = []string{
			fmt.Sprintf("tree %s", tree.ID),
			fmt.Sprintf("parent %s", parent),
			fmt.Sprintf("author %s", author.String()),
			fmt.Sprintf("committer %s", author.String()),
			"",
			message,
		}
	}
	data := strings.Join(lines, "\n")
	return &Commit{
		Tree:    *tree,
		Author:  *author,
		Message: message,
		Object:  NewObject("commit", []byte(data)),
	}
}

func (commit *Commit) String() string {
	return string(commit.Data)
}
