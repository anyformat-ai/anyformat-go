// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"bytes"
	"context"
	"encoding/json"
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

// Create a new extraction workflow.
//
// Workflows define what data to extract from documents. After creating a workflow,
// configure its extraction fields in the
// [AnyFormat dashboard](https://app.anyformat.ai).
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "v2/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single workflow by its ID, including its configured extraction
// fields.
func (r *WorkflowService) Get(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *Workflow, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all workflows in your organization with pagination.
//
// Workflows can be filtered by status and sorted by any field.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *WorkflowListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "v2/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a workflow and all associated file collections and extraction results.
//
// This action is irreversible.
func (r *WorkflowService) Delete(ctx context.Context, workflowID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/workflows/%s/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Upload one or more files to a workflow, creating a new file collection.
//
// Use this when you want to upload files without immediately running extraction.
// To upload and extract in one step, use `POST /v2/workflows/{workflow_id}/run/`
// instead.
//
// Supported file types: PDF, PNG, JPG, TIFF, TXT, DOCX, XLSX, CSV, and more.
func (r *WorkflowService) NewFile(ctx context.Context, workflowID string, body WorkflowNewFileParams, opts ...option.RequestOption) (res *WorkflowNewFileResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/files/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the extraction results for a file collection.
//
// Returns the structured data extracted from each file, including field values,
// confidence scores, and source evidence (text excerpts and page numbers). Also
// includes a `verification_url` linking to the AnyFormat dashboard for human
// review.
//
// Possible non-200 responses:
//
//   - **412 `PRECONDITION_FAILED`** — extraction still in progress; retry with
//     backoff.
//   - **422 `EXTRACTION_FAILED`** — extraction did not complete successfully;
//     terminal. Polling will not transition the collection out of this state.
//     Possible next steps: review the document, retry the upload, or open the
//     collection in the AnyFormat dashboard for more context.
//   - **422 `EXTRACTION_CANCELLED`** — extraction was cancelled; terminal. Possible
//     next steps: review the document, retry the upload, or open the collection in
//     the AnyFormat dashboard.
//
// Use webhooks (`extraction.completed` event) to avoid polling.
func (r *WorkflowService) GetFileResults(ctx context.Context, collectionID string, query WorkflowGetFileResultsParams, opts ...option.RequestOption) (res *WorkflowGetFileResultsResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
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
//
// A file collection groups one or more uploaded files together. Each collection
// has a status indicating the extraction progress: `pending`, `processing`,
// `completed`, or `failed`.
func (r *WorkflowService) ListFiles(ctx context.Context, workflowID string, query WorkflowListFilesParams, opts ...option.RequestOption) (res *WorkflowListFilesResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/files/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List all extraction runs for a workflow with pagination.
//
// Each run corresponds to a file collection that was processed by the workflow.
// Use the run's `id` (collection UUID) with
// `GET /v2/workflows/{workflow_id}/files/{id}/results/` to fetch detailed results.
func (r *WorkflowService) ListRuns(ctx context.Context, workflowID string, query WorkflowListRunsParams, opts ...option.RequestOption) (res *WorkflowListRunsResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/runs/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Upload a file and immediately run the extraction workflow on it.
//
// This is the primary endpoint for document extraction. It creates a file
// collection, uploads the file, and starts extraction in one step. The response
// includes a collection `id` that you can use to poll for results via
// `GET /v2/workflows/{workflow_id}/files/{collection_id}/results/`.
//
// Provide the file as a binary upload in the `file` field, or send raw text in the
// `text` field for text-only extraction.
func (r *WorkflowService) Run(ctx context.Context, workflowID string, body WorkflowRunParams, opts ...option.RequestOption) (res *WorkflowRunResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/run/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload a file to a workflow without running extraction.
//
// Use this when you want to stage files for later processing. For
// upload-and-extract in one step, use `POST /v2/workflows/{workflow_id}/run/`
// instead.
func (r *WorkflowService) Upload(ctx context.Context, workflowID string, body WorkflowUploadParams, opts ...option.RequestOption) (res *WorkflowUploadResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/workflows/%s/upload/", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A workflow defines the extraction template — what fields to extract from
// documents, their types, and validation rules.
type Workflow struct {
	// Unique identifier of the workflow (UUID).
	ID string `json:"id" api:"required"`
	// Human-readable name of the workflow.
	Name string `json:"name" api:"required"`
	// Timestamp when the workflow was created (ISO 8601).
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Optional description of what this workflow extracts.
	Description string `json:"description" api:"nullable"`
	// List of extraction field definitions configured for this workflow. `null` if not
	// yet configured.
	Fields []map[string]any `json:"fields" api:"nullable"`
	// Timestamp when the workflow was last modified (ISO 8601).
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
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

// Paginated list of workflows.
type WorkflowListResponse struct {
	// Total number of workflows matching the query.
	Count int64 `json:"count" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Number of results per page.
	PageSize int64 `json:"page_size" api:"required"`
	// List of workflows for the current page.
	Results []Workflow `json:"results" api:"required"`
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

// Response from creating a file collection. Contains the collection ID and the
// status of each uploaded file.
type WorkflowNewFileResponse struct {
	// Unique identifier of the newly created file collection.
	ID string `json:"id" api:"required"`
	// List of files included in the collection, with their upload status.
	Files []WorkflowNewFileResponseFile `json:"files" api:"required"`
	// The UUID of the workflow this collection belongs to.
	WorkflowID string `json:"workflow_id" api:"required"`
	// Human-readable name for the collection.
	Name string `json:"name" api:"nullable"`
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

// A single file within a collection, showing its name and upload status.
type WorkflowNewFileResponseFile struct {
	// Name of the uploaded file.
	Filename string `json:"filename" api:"required"`
	// Upload status: `uploaded` or `failed`.
	Status string `json:"status" api:"required"`
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

// Canonical response shape for the file-collection results endpoint.
//
// Returned with HTTP 200 once processing completes. Returns 412 while processing
// is in progress; poll until 200, or use webhooks.
type WorkflowGetFileResultsResponse struct {
	// The file collection's UUID. Same value as the `id` returned by
	// `POST /v2/workflows/{wid}/run/`.
	CollectionID string `json:"collection_id" api:"required"`
	// Extracted fields keyed by field name. `null` for parse-only workflows. Always
	// present in the response. Each value is either a scalar field (`ExtractedField`)
	// or a list of object-field rows (`list[dict[str, ExtractedField]]`) for compound
	// fields like line items.
	Extraction map[string]WorkflowGetFileResultsResponseExtractionUnion `json:"extraction" api:"nullable"`
	// Parsed markdown for a file.
	Parse WorkflowGetFileResultsResponseParse `json:"parse" api:"nullable"`
	// Link to the AnyFormat dashboard for human review of this collection's results.
	// `null` if the dashboard URL cannot be constructed (e.g. no files in the
	// collection, or the deployment has no frontend URL configured).
	VerificationURL string `json:"verification_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionID    respjson.Field
		Extraction      respjson.Field
		Parse           respjson.Field
		VerificationURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowGetFileResultsResponseExtractionUnion contains all possible properties
// and values from [WorkflowGetFileResultsResponseExtractionExtractedField],
// [[]map[string]WorkflowGetFileResultsResponseExtractionArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfWorkflowGetFileResultsResponseExtractionArrayItemMap]
type WorkflowGetFileResultsResponseExtractionUnion struct {
	// This field will be present if the value is a
	// [[]map[string]WorkflowGetFileResultsResponseExtractionArrayItem] instead of an
	// object.
	OfMapOfWorkflowGetFileResultsResponseExtractionArrayItemMap []map[string]WorkflowGetFileResultsResponseExtractionArrayItem `json:",inline"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionExtractedField].
	Value any `json:"value"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionExtractedField].
	Confidence float64 `json:"confidence"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionExtractedField].
	Evidence []WorkflowGetFileResultsResponseExtractionExtractedFieldEvidence `json:"evidence"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionExtractedField].
	ValueOverride any `json:"value_override"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionExtractedField].
	VerificationStatus string `json:"verification_status"`
	JSON               struct {
		OfMapOfWorkflowGetFileResultsResponseExtractionArrayItemMap respjson.Field
		Value                                                       respjson.Field
		Confidence                                                  respjson.Field
		Evidence                                                    respjson.Field
		ValueOverride                                               respjson.Field
		VerificationStatus                                          respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (u WorkflowGetFileResultsResponseExtractionUnion) AsExtractedField() (v WorkflowGetFileResultsResponseExtractionExtractedField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowGetFileResultsResponseExtractionUnion) AsMapOfWorkflowGetFileResultsResponseExtractionArrayItemMap() (v []map[string]WorkflowGetFileResultsResponseExtractionArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowGetFileResultsResponseExtractionUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowGetFileResultsResponseExtractionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One extracted field's value, confidence, and supporting evidence.
type WorkflowGetFileResultsResponseExtractionExtractedField struct {
	// The extracted value. Type depends on the field's `data_type` (string, number,
	// date, etc.). `null` when extraction could not produce a value.
	Value any `json:"value" api:"required"`
	// Model confidence in the extracted value, on a 0-100 scale. `null` when the
	// backend did not produce a confidence (e.g. manual entry).
	Confidence float64 `json:"confidence" api:"nullable"`
	// Source-text snippets the model used to derive this value.
	Evidence []WorkflowGetFileResultsResponseExtractionExtractedFieldEvidence `json:"evidence"`
	// A human-supplied override of the extracted `value`, if one was set during
	// verification. `null` when no override exists.
	ValueOverride any `json:"value_override"`
	// Verification state for this datapoint (e.g. `not_verified`, `verified`). `null`
	// when not yet reviewed.
	VerificationStatus string `json:"verification_status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value              respjson.Field
		Confidence         respjson.Field
		Evidence           respjson.Field
		ValueOverride      respjson.Field
		VerificationStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseExtractionExtractedField) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseExtractionExtractedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A snippet of source text supporting an extracted value, with the page it came
// from.
type WorkflowGetFileResultsResponseExtractionExtractedFieldEvidence struct {
	// 1-indexed page number where the snippet was found.
	PageNumber int64 `json:"page_number" api:"required"`
	// The exact source-text snippet that supports the extracted value.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseExtractionExtractedFieldEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowGetFileResultsResponseExtractionExtractedFieldEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One extracted field's value, confidence, and supporting evidence.
type WorkflowGetFileResultsResponseExtractionArrayItem struct {
	// The extracted value. Type depends on the field's `data_type` (string, number,
	// date, etc.). `null` when extraction could not produce a value.
	Value any `json:"value" api:"required"`
	// Model confidence in the extracted value, on a 0-100 scale. `null` when the
	// backend did not produce a confidence (e.g. manual entry).
	Confidence float64 `json:"confidence" api:"nullable"`
	// Source-text snippets the model used to derive this value.
	Evidence []WorkflowGetFileResultsResponseExtractionArrayItemEvidence `json:"evidence"`
	// A human-supplied override of the extracted `value`, if one was set during
	// verification. `null` when no override exists.
	ValueOverride any `json:"value_override"`
	// Verification state for this datapoint (e.g. `not_verified`, `verified`). `null`
	// when not yet reviewed.
	VerificationStatus string `json:"verification_status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value              respjson.Field
		Confidence         respjson.Field
		Evidence           respjson.Field
		ValueOverride      respjson.Field
		VerificationStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseExtractionArrayItem) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseExtractionArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A snippet of source text supporting an extracted value, with the page it came
// from.
type WorkflowGetFileResultsResponseExtractionArrayItemEvidence struct {
	// 1-indexed page number where the snippet was found.
	PageNumber int64 `json:"page_number" api:"required"`
	// The exact source-text snippet that supports the extracted value.
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseExtractionArrayItemEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowGetFileResultsResponseExtractionArrayItemEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsed markdown for a file.
type WorkflowGetFileResultsResponseParse struct {
	// Document content rendered as structured markdown (with `<DOCUMENT>` /
	// `<section>` tags, embedded images for the `visual` variant). `null` if parsing
	// failed.
	Markdown string `json:"markdown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseParse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseParse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListFilesResponse struct {
	// Total number of items matching the query.
	Count int64 `json:"count" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Number of results per page.
	PageSize int64 `json:"page_size" api:"required"`
	// List of items for the current page.
	Results []WorkflowListFilesResponseResult `json:"results" api:"required"`
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

// A file collection entry in list responses. Each collection groups one or more
// uploaded files and tracks their extraction status.
type WorkflowListFilesResponseResult struct {
	// Unique identifier of the file collection.
	ID string `json:"id" api:"required"`
	// Processing status: `pending`, `queued`, `in_progress`, `processed`, `error`, or
	// `cancelled`.
	Status string `json:"status" api:"required"`
	// Timestamp when the collection was created (ISO 8601).
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Human-readable name for the collection.
	Name string `json:"name" api:"nullable"`
	// Timestamp when the collection was last updated (ISO 8601).
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

// Paginated list of workflow runs.
type WorkflowListRunsResponse struct {
	// Total number of runs for this workflow.
	Count int64 `json:"count" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Number of results per page.
	PageSize int64 `json:"page_size" api:"required"`
	// List of runs for the current page.
	Results []WorkflowListRunsResponseResult `json:"results" api:"required"`
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

// An extraction run entry, representing one execution of a workflow on a file
// collection.
type WorkflowListRunsResponseResult struct {
	// The collection UUID for this run. Use this ID with
	// `GET /v2/workflows/{workflow_id}/files/{id}/results/` to fetch results.
	ID string `json:"id" api:"required"`
	// Processing status: `pending`, `queued`, `in_progress`, `processed`, `error`, or
	// `cancelled`.
	Status string `json:"status" api:"required"`
	// Timestamp when the run started (ISO 8601).
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Timestamp when the run status was last updated (ISO 8601).
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

// Response after triggering a workflow run. Contains the collection ID to use for
// polling extraction results.
type WorkflowRunResponse struct {
	// The collection UUID for this run. Use this ID to poll for results via
	// `GET /v2/workflows/{workflow_id}/files/{id}/results/`.
	ID string `json:"id" api:"required"`
	// Initial status of the run, typically `success` (meaning the run was accepted,
	// not that extraction is complete).
	Status string `json:"status" api:"required"`
	// The UUID of the workflow that was executed.
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

// Confirmation that a file was uploaded successfully without triggering
// extraction.
type WorkflowUploadResponse struct {
	// Upload result: `uploaded` on success.
	Status string `json:"status" api:"required"`
	// Name of the uploaded file.
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

type WorkflowNewParams struct {
	// Field definitions
	Fields []map[string]any `json:"fields,omitzero" api:"required"`
	// Workflow name
	Name string `json:"name" api:"required"`
	// Workflow description
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
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
