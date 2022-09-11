package internals

type Blob struct {
	ID string
	*Object
}

func NewBlob(repo *Repo, data []byte) *Blob {
	object := &Object{ObjectType: "blob", Data: data}
	id, _ := object.Hash()
	blob := Blob{
		Object: object,
		ID:     id,
	}
	return &blob
}
