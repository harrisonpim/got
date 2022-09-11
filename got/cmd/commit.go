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
	for _, file := range fileList {
		// read each file, and store the data as a blob
		data, err := os.ReadFile(file.Path)
		if err != nil {
			return err
		}
		blob := internals.NewBlob(repo, data)

		// write the hashed data to a path in ".got/objects" according to the git spec
		if err := repo.WriteObject(*blob.Object); err != nil {
			return err
		}
	}
	return nil
}
