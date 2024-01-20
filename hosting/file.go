package hosting

import "errors"

type Encoding string

const (
	TextEncoding   Encoding = "text"
	Base64Encoding Encoding = "base64"
	NoneEncoding   Encoding = "none"
)

type File struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	Content  *string   `json:"content,omitempty"`
	Encoding *Encoding `json:"encoding,omitempty"`
	Size     *int      `json:"size,omitempty"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
}

func (f *File) SetEncoding(encoding string) (*File, error) {
	var enc Encoding
	switch encoding {
	case string(Base64Encoding):
		enc = Base64Encoding
	case string(TextEncoding):
		enc = TextEncoding
	case string(NoneEncoding):
		enc = NoneEncoding
	default:
		return nil, errors.New("unsupported encoding")
	}

	f.Encoding = &enc

	return f, nil
}

func (f *File) GetEncoding() *string {
	if f.Encoding == nil {
		return nil
	}

	enc := string(*f.Encoding)

	return &enc
}

type GetFileOpts struct {
	Ref *string `json:"ref"`
}

type CreateFileOpts struct {
	SHA    *string `json:"sha"`
	Branch *string `json:"branch"`
	Ref    *string `json:"ref"`
	Commit `json:"commit"`
}

type UpdateFileOpts struct {
	SHA    *string `json:"sha"`
	Branch *string `json:"branch"`
	Ref    *string `json:"ref"`
	Commit `json:"commit"`
}

type DeleteFileOpts struct {
	SHA    *string `json:"sha"`
	Branch *string `json:"branch"`
	Ref    *string `json:"ref"`
	Commit `json:"commit"`
}
