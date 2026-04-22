// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/anyformat-ai/anyformat-go/internal/requestconfig"
	"github.com/anyformat-ai/anyformat-go/option"
)

// File collections group uploaded documents and track their extraction progress.
// Upload files, check status, and retrieve extraction results.
//
// FileService contains methods and other services that help with interacting with
// the anyformat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	return
}

// Delete a file collection and all its files permanently.
//
// This removes all uploaded files and any extraction results associated with the
// collection. This action is irreversible.
func (r *FileService) Delete(ctx context.Context, collectionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/files/%s/", url.PathEscape(collectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}
