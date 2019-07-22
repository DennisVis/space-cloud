package googlecloudstorage

import (
	"context"
	"strings"

	"google.golang.org/api/iterator"

	"github.com/spaceuptech/space-cloud/model"
)

// ListDir lists a directory in GCS
func (gcs *GCS) ListDir(
	ctx context.Context,
	project string,
	req *model.ListFilesRequest,
) ([]*model.ListFilesResponse, error) {
	result := []*model.ListFilesResponse{}

	bkt := gcs.client.Bucket(project)

	if req.Type == "all" || req.Type == "dir" {
		it := bkt.Objects(ctx, nil)
		for {
			objAttrs, err := it.Next()
			if err != nil && err != iterator.Done {
				return result, err
			}
			if err == iterator.Done {
				break
			}

			if strings.HasPrefix(objAttrs.Name, req.Path) {
				result = append(result, &model.ListFilesResponse{
					Name: objAttrs.Name,
					Type: req.Type,
				})
			}
		}
	} else {
		result = append(result, &model.ListFilesResponse{
			Name: req.Path,
			Type: req.Type,
		})
	}

	return result, nil
}

// ReadFile reads a file from GCS
func (gcs *GCS) ReadFile(ctx context.Context, project, path string) (*model.File, error) {
	bkt := gcs.client.Bucket(project)
	obj := bkt.Object(path)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	return &model.File{
		File: r,
		Close: func() error {
			return r.Close()
		},
	}, nil
}
