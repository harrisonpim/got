package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitNewRepo(path string) (*Repo, error) {
	path, _ = filepath.Abs(path)
	repo := Repo{
		WorkTree:     path,
		GotDirectory: filepath.Join(path, ".got"),
	}

	fmt.Printf("Initialising an empty got directory at: %q\n", repo.GotDirectory)
	os.MkdirAll(repo.GotDirectory, os.ModePerm)
	os.MkdirAll(filepath.Join(repo.GotDirectory, "objects"), os.ModePerm)
	os.MkdirAll(filepath.Join(repo.GotDirectory, "refs"), os.ModePerm)

	return &repo, nil
}
