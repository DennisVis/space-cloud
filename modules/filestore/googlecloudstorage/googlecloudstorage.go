package googlecloudstorage

import (
	"context"

	"cloud.google.com/go/storage"

	"github.com/spaceuptech/space-cloud/utils"
)

// GCS holds the Google Cloud Storage client
type GCS struct {
	client *storage.Client
}

// Init initializes a Google Cloud Storage client
func Init() (*GCS, error) {
	client, err := storage.NewClient(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &GCS{client: client}, nil
}

// GetStoreType returns the file store type
func (gcs *GCS) GetStoreType() utils.FileStoreType {
	return utils.GoogleCloudStorage
}

// Close gracefully closed the GCS filestore module
func (gcs *GCS) Close() error {
	return nil
}
