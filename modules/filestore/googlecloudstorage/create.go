package googlecloudstorage

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/spaceuptech/space-cloud/model"
)

// CreateFile creates a file in GCS
func (gcs *GCS) CreateFile(ctx context.Context, project string, req *model.CreateFileRequest, file io.Reader) error {
	bkt := gcs.client.Bucket(project)

	path := strings.Join([]string{req.Path, req.Name}, "/")
	obj := bkt.Object(path)

	w := obj.NewWriter(ctx)
	_, err := io.Copy(w, file)
	return err
}

// CreateDir is a NOOP as Google Cloud Storage directories are virtual and part of a file name
func (gcs *GCS) CreateDir(ctx context.Context, project string, req *model.CreateFileRequest) error {
	return errors.New("Not Implemented")
}
