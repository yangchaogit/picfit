package payload

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/mholt/binding"
	"github.com/thoas/gostorages"
	"picfit/image"
)

// MultipartPayload represents a multipart upload
type MultipartPayload struct {
	Data *multipart.FileHeader `json:"data"`
}

// FieldMap defines excepted inputs
func (f *MultipartPayload) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Data: "data",
	}
}

// Upload uploads a file to its storage
func (f *MultipartPayload) Upload(storage gostorages.Storage) (*image.ImageFile, error) {
	var fh io.ReadCloser

	fh, err := f.Data.Open()

	if err != nil {
		return nil, err
	}

	defer fh.Close()

	dataBytes := bytes.Buffer{}

	_, err = dataBytes.ReadFrom(fh)

	if err != nil {
		return nil, err
	}

	err = storage.Save(f.Data.Filename, gostorages.NewContentFile(dataBytes.Bytes()))

	if err != nil {
		return nil, err
	}

	return &image.ImageFile{
		Filepath: f.Data.Filename,
		Storage:  storage,
	}, nil
}
