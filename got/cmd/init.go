package cmd

import (
	"fmt"
	"os"

	"github.com/harrisonpim/got/internals"
)

func InitNewRepo(path string) (*internals.Repo, error) {
	repo, err := internals.NewRepo(path)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Initialising an empty got directory at: %q\n", repo.GotDirectory)
	os.MkdirAll(repo.ObjectDirectory, os.ModePerm)
	os.MkdirAll(repo.RefsDirectory, os.ModePerm)
	return repo, nil
}
