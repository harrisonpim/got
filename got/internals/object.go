package internals

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

type Object struct {
	ID, ObjectType      string
	Data, DecoratedData []byte
}

func NewObject(objectType string, data []byte) *Object {
	id, decoratedData := hash(objectType, data)
	return &Object{
		Data:          data,
		ObjectType:    objectType,
		ID:            id,
		DecoratedData: decoratedData,
	}
}

func hash(objectType string, data []byte) (string, []byte) {
	header := []byte(objectType + " " + strconv.Itoa(len(data)) + "\x00")
	decoratedData := append(header, data...)

	h := sha1.New()
	h.Write(decoratedData)
	sha1hash := hex.EncodeToString(h.Sum(nil))

	return sha1hash, decoratedData
}
