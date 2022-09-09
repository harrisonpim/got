package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

type Repo struct {
	WorkTree     string
	GotDirectory string
}

func (repo Repo) ListFiles() ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(repo.WorkTree,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !((path == ".") || (path == "..") || (strings.HasPrefix(path, ".got"))) {
				fileList = append(fileList, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}
