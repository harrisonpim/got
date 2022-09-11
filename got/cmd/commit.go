package cmd

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
	"strconv"

	"github.com/harrisonpim/got/util"
)

func Commit(path string) error {
	repo, err := NewRepo(path)
	if err != nil {
		return err
	}
	fileList, err := repo.ListFiles()
	if err != nil {
		return err
	}
	for _, file := range fileList {
		// read each file, hash the data, and store it as a blob object
		data, err := os.ReadFile(file.Path)
		if err != nil {
			return err
		}
		sha1hash, hashedData := HashObject(&Object{objectType: "blob", data: data})

		// write the hashed data to a path in ".got/objects" according to the git spec
		objectPath, err := util.NewPath(filepath.Join(repo.ObjectDirectory, string(sha1hash[0:2]), string(sha1hash[2:])))
		if err := WriteBinaryData(objectPath, hashedData); err != nil {
			return err
		}
	}
	return nil
}

type Object struct {
	objectType string
	data       []byte
}

func HashObject(obj *Object) (string, []byte) {
	header := []byte(obj.objectType + " " + strconv.Itoa(len(obj.data)) + "\x00")
	data := append(header, obj.data...)

	h := sha1.New()
	h.Write(data)
	sha1hash := hex.EncodeToString(h.Sum(nil))

	return sha1hash, data
}

func WriteBinaryData(path *util.Path, data []byte) error {
	var compressed bytes.Buffer
	w := zlib.NewWriter(&compressed)
	w.Write(data)
	w.Close()
	if !path.Exists {
		if err := os.MkdirAll(path.Directory, os.ModePerm); err != nil {
			return err
		}
	}
	if err := os.WriteFile(path.Path, compressed.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}
