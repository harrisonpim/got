package cmd

import (
	"fmt"
	"os"
)

func InitNewRepo(path string) (*Repo, error) {
	repo, err := NewRepo(path)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Initialising an empty got directory at: %q\n", repo.GotDirectory)
	os.MkdirAll(repo.ObjectDirectory, os.ModePerm)
	os.MkdirAll(repo.RefsDirectory, os.ModePerm)
	return repo, nil
}
