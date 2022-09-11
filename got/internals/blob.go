package internals

type Blob struct {
	Repository *Repo
	*Object
}

func NewBlob(repo *Repo, data []byte) *Blob {
	blob := Blob{
		Repository: repo,
		Object:     &Object{ObjectType: "blob", Data: data},
	}
	return &blob
}

