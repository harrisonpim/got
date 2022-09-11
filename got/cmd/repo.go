package cmd

import (
	"os"
	"path/filepath"

	"github.com/harrisonpim/got/util"
)

type Repo struct {
	RootDirectory   string
	GotDirectory    string
	ObjectDirectory string
	RefsDirectory   string
}

func NewRepo(path string) (*Repo, error) {
	root, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &Repo{
		RootDirectory:   root,
		GotDirectory:    filepath.Join(root, ".got"),
		ObjectDirectory: filepath.Join(root, ".got", "objects"),
		RefsDirectory:   filepath.Join(root, ".got", "refs"),
	}, nil
}

func (repo Repo) ListFiles() ([]util.Path, error) {
	fileList := []util.Path{}
	err := filepath.Walk(repo.RootDirectory,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			p, err := util.NewPath(path)
			pathIsInGotDirectory := false
			for _, component := range p.Components {
				if component == ".got" {
					pathIsInGotDirectory = true
				}
			}
			if !((p.IsDirectory) || (pathIsInGotDirectory)) {
				fileList = append(fileList, *p)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}
