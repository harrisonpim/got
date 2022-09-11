package util

import (
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	Path        string
	Filename    string
	Directory   string
	Components  []string
	Exists      bool
	IsDirectory bool
}

func NewPath(path string) (*Path, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	exists := Exists(absolutePath)

	directory, filename := filepath.Split(absolutePath)
	components := strings.Split(absolutePath, "/")
	isDirectory, _ := IsDirectory(absolutePath)

	newPath := Path{
		Path:        absolutePath,
		Components:  components,
		Directory:   directory,
		Filename:    filename,
		Exists:      exists,
		IsDirectory: isDirectory,
	}
	return &newPath, nil
}

func IsDirectory(path string) (bool, error) {
	fInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (path Path) Parent() (*Path, error) {
	parent := path.Components[:len(path.Components)-1]
	return NewPath(filepath.Join(parent...))
}
