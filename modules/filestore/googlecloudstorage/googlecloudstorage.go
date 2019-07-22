package googlecloudstorage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"github.com/spaceuptech/space-cloud/utils"
)

// GCS holds the Google Cloud Storage client
type GCS struct {
	client *storage.Client
}

// Init initializes a Google Cloud Storage client
func Init(ctx context.Context, opts ...option.ClientOption) (*GCS, error) {
	client, err := storage.NewClient(ctx, opts...)
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
