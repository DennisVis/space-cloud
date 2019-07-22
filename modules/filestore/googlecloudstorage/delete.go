package googlecloudstorage

import (
	"context"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func isDir(path string) bool {
	return filepath.Dir(path) == path
}

// DeleteFile deletes a file from Google Cloud Storage.
// If the provided path is a directory all files within this directory will be deleted.
func (gcs *GCS) DeleteFile(ctx context.Context, project, path string) error {
	bkt := gcs.client.Bucket(project)
	if isDir(path) {
		it := bkt.Objects(ctx, &storage.Query{Prefix: path})
		for {
			objAttrs, err := it.Next()
			if err != nil && err != iterator.Done {
				return err
			}
			if err == iterator.Done {
				return nil
			}

			err = bkt.Object(objAttrs.Name).Delete(ctx)
			if err != nil {
				return err
			}
		}
	}
	obj := bkt.Object(path)
	return obj.Delete(ctx)
}
