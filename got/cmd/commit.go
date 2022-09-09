package cmd

import (
	"fmt"
	"path/filepath"
)

func Commit(path string) error {
	repo := Repo{
		WorkTree:     path,
		GotDirectory: filepath.Join(path, ".got"),
	}
	fileList, err := repo.ListFiles()
	if err != nil {
		return err
	}
	for _, file := range fileList {
		fmt.Println(file)
	}
	return nil
}
