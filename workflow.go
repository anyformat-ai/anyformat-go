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

	"github.com/anyformat-ai/anyformat-go/internal/apiform"
	"github.com/anyformat-ai/anyformat-go/internal/apijson"
	"github.com/anyformat-ai/anyformat-go/internal/apiquery"
	"github.com/anyformat-ai/anyformat-go/internal/requestconfig"
	"github.com/anyformat-ai/anyformat-go/option"
	"github.com/anyformat-ai/anyformat-go/packages/param"
	"github.com/anyformat-ai/anyformat-go/packages/respjson"
)

// WorkflowService contains methods and other services that help with interacting
// with the anyformat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.options = opts
	return
}

// Create a new workflow.
func (r *WorkflowService) New(ctx context.Context, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get workflow by ID.
func (r *WorkflowService) Get(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List workflows with pagination.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *WorkflowListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete workflow by ID.
func (r *WorkflowService) Delete(ctx context.Context, workflowID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/workflows/%s/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Upload files to a workflow, creating a file collection.
func (r *WorkflowService) NewFile(ctx context.Context, workflowID string, body WorkflowNewFileParams, opts ...option.RequestOption) (res *WorkflowNewFileResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/files/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get processing results for a file collection.
//
// Returns the backend collection results with internal metadata stripped. Returns
// 412 if processing is not yet complete.
func (r *WorkflowService) GetFileResults(ctx context.Context, collectionID string, query WorkflowGetFileResultsParams, opts ...option.RequestOption) (res *WorkflowGetFileResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	if collectionID == "" {
		err = errors.New("missing required collection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/files/%s/results/", url.PathEscape(query.WorkflowID), url.PathEscape(collectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List file collections for a workflow.
func (r *WorkflowService) ListFiles(ctx context.Context, workflowID string, query WorkflowListFilesParams, opts ...option.RequestOption) (res *WorkflowListFilesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/files/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List extraction runs for a workflow, identified by collection UUID.
func (r *WorkflowService) ListRuns(ctx context.Context, workflowID string, query WorkflowListRunsParams, opts ...option.RequestOption) (res *WorkflowListRunsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/runs/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute workflow — returns collection UUID.
func (r *WorkflowService) Run(ctx context.Context, workflowID string, body WorkflowRunParams, opts ...option.RequestOption) (res *WorkflowRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/run/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload file without executing workflow.
func (r *WorkflowService) Upload(ctx context.Context, workflowID string, body WorkflowUploadParams, opts ...option.RequestOption) (res *WorkflowUploadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/upload/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Workflow detail — used for get, create, and list items.
type Workflow struct {
	ID          string           `json:"id" api:"required"`
	Name        string           `json:"name" api:"required"`
	CreatedAt   time.Time        `json:"created_at" api:"nullable" format:"date-time"`
	Description string           `json:"description" api:"nullable"`
	Fields      []map[string]any `json:"fields" api:"nullable"`
	UpdatedAt   time.Time        `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Fields      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GET /workflows/ — paginated workflow list.
type WorkflowListResponse struct {
	Count    int64      `json:"count" api:"required"`
	Page     int64      `json:"page" api:"required"`
	PageSize int64      `json:"page_size" api:"required"`
	Results  []Workflow `json:"results" api:"required"`
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
func (r WorkflowListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from creating a file collection.
type WorkflowNewFileResponse struct {
	ID         string                        `json:"id" api:"required"`
	Files      []WorkflowNewFileResponseFile `json:"files" api:"required"`
	WorkflowID string                        `json:"workflow_id" api:"required"`
	Name       string                        `json:"name" api:"nullable"`
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
func (r WorkflowNewFileResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNewFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single file within a collection.
type WorkflowNewFileResponseFile struct {
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
func (r WorkflowNewFileResponseFile) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNewFileResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowGetFileResultsResponse = any

type WorkflowListFilesResponse struct {
	Count    int64                             `json:"count" api:"required"`
	Page     int64                             `json:"page" api:"required"`
	PageSize int64                             `json:"page_size" api:"required"`
	Results  []WorkflowListFilesResponseResult `json:"results" api:"required"`
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
func (r WorkflowListFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single file collection entry in list responses.
type WorkflowListFilesResponseResult struct {
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
func (r WorkflowListFilesResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListFilesResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GET /workflows/{id}/runs/ — paginated run list.
type WorkflowListRunsResponse struct {
	Count    int64                            `json:"count" api:"required"`
	Page     int64                            `json:"page" api:"required"`
	PageSize int64                            `json:"page_size" api:"required"`
	Results  []WorkflowListRunsResponseResult `json:"results" api:"required"`
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
func (r WorkflowListRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Item in GET /workflows/{id}/runs/ paginated list.
type WorkflowListRunsResponseResult struct {
	ID        string    `json:"id" api:"required"`
	Status    string    `json:"status" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowListRunsResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListRunsResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for workflow run endpoint (v2) — collection UUID as identifier.
type WorkflowRunResponse struct {
	ID         string `json:"id" api:"required"`
	Status     string `json:"status" api:"required"`
	WorkflowID string `json:"workflow_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		WorkflowID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// POST /workflows/{id}/upload/ — upload confirmation.
type WorkflowUploadResponse struct {
	Status   string `json:"status" api:"required"`
	Filename string `json:"filename" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListParams struct {
	Order    param.Opt[string] `query:"order,omitzero" json:"-"`
	SortBy   param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	Status   param.Opt[string] `query:"status,omitzero" json:"-"`
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowNewFileParams struct {
	Files []string `json:"files,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewFileParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type WorkflowGetFileResultsParams struct {
	WorkflowID string `path:"workflow_id" api:"required" json:"-"`
	paramObj
}

type WorkflowListFilesParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListFilesParams]'s query parameters as
// `url.Values`.
func (r WorkflowListFilesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowListRunsParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListRunsParams]'s query parameters as `url.Values`.
func (r WorkflowListRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunParams struct {
	File param.Opt[string] `json:"file,omitzero"`
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WorkflowRunParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type WorkflowUploadParams struct {
	File param.Opt[string] `json:"file,omitzero"`
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WorkflowUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
