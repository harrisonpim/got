package internals

import (
	"bytes"
	"compress/zlib"
	"os"
	"path/filepath"
)

type Repo struct {
	RootDirectory, GotDirectory, ObjectDirectory, RefsDirectory string
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

func (repo Repo) ListFiles() ([]Path, error) {
	fileList := []Path{}
	err := filepath.Walk(repo.RootDirectory,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			p, err := NewPath(path)
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

func (repo Repo) WriteObject(object Object) error {
	objectPath, err := NewPath(filepath.Join(repo.ObjectDirectory, string(object.ID[0:2]), string(object.ID[2:])))
	if err != nil {
		return err
	}
	if !objectPath.Exists {
		var compressed bytes.Buffer
		w := zlib.NewWriter(&compressed)
		w.Write(object.DecoratedData)
		w.Close()
		if err := os.MkdirAll(objectPath.Directory, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(objectPath.Path, compressed.Bytes(), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (repo *Repo) UpdateHead(commit Commit) error {
	headPath := filepath.Join(repo.GotDirectory, "HEAD")
	if err := os.WriteFile(headPath, []byte(commit.ID), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (repo *Repo) ReadHead() (string, error) {
	headPath := filepath.Join(repo.GotDirectory, "HEAD")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
