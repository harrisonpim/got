package internals

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

type Object struct {
	ObjectType string
	Data       []byte
}

func (object *Object) Hash() (string, []byte) {
	header := []byte(object.ObjectType + " " + strconv.Itoa(len(object.Data)) + "\x00")
	decoratedData := append(header, object.Data...)

	h := sha1.New()
	h.Write(decoratedData)
	sha1hash := hex.EncodeToString(h.Sum(nil))

	return sha1hash, decoratedData
}
