package cmd

import (
	"fmt"
	"os"

	"github.com/harrisonpim/got/internals"
)

func Commit(path string, message string) error {
	repo, err := internals.NewRepo(path)
	if err != nil {
		return err
	}
	fileList, err := repo.ListFiles()
	if err != nil {
		return err
	}

	entries := []internals.Entry{}
	for _, file := range fileList {
		// parse each file in the worktree as a blob
		data, err := os.ReadFile(file.Path)
		if err != nil {
			return err
		}
		blob := internals.NewBlob(repo, data)
		fileInfo, err := os.Stat(file.Path)

		// write the data to .got/objects
		if err := repo.WriteObject(*blob.Object); err != nil {
			return err
		}
		entries = append(entries, internals.Entry{
			ID:   blob.ID,
			Name: file.Path,
			Mode: fileInfo.Mode(),
		})
	}
	tree := internals.NewTree(entries)
	if err := repo.WriteObject(*tree.Object); err != nil {
		return err
	}

	authorName := os.Getenv("GOT_AUTHOR_NAME")
	authorEmail := os.Getenv("GOT_AUTHOR_EMAIL")
	author := internals.NewAuthor(authorName, authorEmail)

	isRoot := ""
	parent, err := repo.ReadHead()
	if err != nil {
		isRoot = "(root-commit) "
	}

	commit := internals.NewCommit(parent, tree, author, message)
	if err := repo.WriteObject(*commit.Object); err != nil {
		return err
	}
	if err := repo.UpdateHead(*commit); err != nil {
		return err
	}
	fmt.Printf("[%s%s] %s", isRoot, commit.ID, message)
	return nil
}
