package cmd

import (
	"fmt"
	"path/filepath"
)

type Repo struct {
	WorkTree     string
	GotDirectory string
}

func InitNewRepo(path string) (*Repo, error) {
	path, _ = filepath.Abs(path)
	repo := Repo{
		WorkTree:     path,
		GotDirectory: filepath.Join(path, ".got"),
	}

	fmt.Printf("Creating an empty got repo at path: %q\n", path)
	fmt.Printf("Initialising got directory at: %q\n", repo.GotDirectory)

	return &repo, nil
}
