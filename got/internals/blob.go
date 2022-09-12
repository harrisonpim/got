package internals

type Blob struct {
	*Object
}

func NewBlob(repo *Repo, data []byte) *Blob {
	return &Blob{
		Object: NewObject("blob", data),
	}
}
