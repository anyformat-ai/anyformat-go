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

// Workflow CRUD, execution, runs, and results.
//
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

// Get workflow results.
func (r *WorkflowService) Results(ctx context.Context, workflowID string, query WorkflowResultsParams, opts ...option.RequestOption) (res *WorkflowResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/results/", url.PathEscape(workflowID))
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

type WorkflowResultsResponse = any

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

type WorkflowResultsParams struct {
	AsLists      param.Opt[string] `query:"as_lists,omitzero" json:"-"`
	OutputFormat param.Opt[string] `query:"output_format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowResultsParams]'s query parameters as `url.Values`.
func (r WorkflowResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunParams struct {
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	File        param.Opt[string] `json:"file,omitzero"`
	FileBase64  param.Opt[string] `json:"file_base64,omitzero"`
	Filename    param.Opt[string] `json:"filename,omitzero"`
	Text        param.Opt[string] `json:"text,omitzero"`
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
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	File        param.Opt[string] `json:"file,omitzero"`
	FileBase64  param.Opt[string] `json:"file_base64,omitzero"`
	Filename    param.Opt[string] `json:"filename,omitzero"`
	Text        param.Opt[string] `json:"text,omitzero"`
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
