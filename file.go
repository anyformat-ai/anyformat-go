// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/anyformat-go/internal/apiform"
	"github.com/stainless-sdks/anyformat-go/internal/apijson"
	"github.com/stainless-sdks/anyformat-go/internal/apiquery"
	"github.com/stainless-sdks/anyformat-go/internal/requestconfig"
	"github.com/stainless-sdks/anyformat-go/option"
	"github.com/stainless-sdks/anyformat-go/packages/param"
	"github.com/stainless-sdks/anyformat-go/packages/respjson"
)

// File collection management.
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

// Upload files to a workflow, creating a file collection.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *FileNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/files/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List file collections for a workflow.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *FileListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/files/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a file collection and all its files.
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

// Get extraction results for a file collection.
//
// Returns 412 if the extraction is not yet complete.
func (r *FileService) GetExtractionResults(ctx context.Context, collectionID string, opts ...option.RequestOption) (res *FileGetExtractionResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/files/%s/extraction/", url.PathEscape(collectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response from creating a file collection.
type FileNewResponse struct {
	ID         string                `json:"id" api:"required"`
	Files      []FileNewResponseFile `json:"files" api:"required"`
	WorkflowID string                `json:"workflow_id" api:"required"`
	Name       string                `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Files       respjson.Field
		WorkflowID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single file within a collection.
type FileNewResponseFile struct {
	Filename string `json:"filename" api:"required"`
	Status   string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileNewResponseFile) RawJSON() string { return r.JSON.raw }
func (r *FileNewResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponse struct {
	Count    int64                    `json:"count" api:"required"`
	Page     int64                    `json:"page" api:"required"`
	PageSize int64                    `json:"page_size" api:"required"`
	Results  []FileListResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single file collection entry in list responses.
type FileListResponseResult struct {
	ID        string    `json:"id" api:"required"`
	Status    string    `json:"status" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	Name      string    `json:"name" api:"nullable"`
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileGetExtractionResultsResponse = any

type FileNewParams struct {
	Files      []string `json:"files,omitzero" api:"required"`
	WorkflowID string   `json:"workflow_id" api:"required"`
	paramObj
}

func (r FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type FileListParams struct {
	WorkflowID param.Opt[string] `query:"workflow_id,omitzero" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
