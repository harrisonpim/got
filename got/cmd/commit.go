package cmd

import (
	"os"

	"github.com/harrisonpim/got/internals"
)

func Commit(path string) error {
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

		// write the data to .got/objects
		if err := repo.WriteObject(*blob.Object); err != nil {
			return err
		}
		entries = append(entries, internals.Entry{
			ID:   blob.ID,
			Name: file.Path,
		})
	}
	tree := internals.NewTree(entries)
	if err := repo.WriteObject(*tree.Object); err != nil {
		return err
	}
	return nil
}
