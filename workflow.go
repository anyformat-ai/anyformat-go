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
	"github.com/anyformat-ai/anyformat-go/shared/constant"
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
	// Per-classifier-node verdicts. Empty when the workflow has no classifier.
	Classifications []WorkflowGetFileResultsResponseClassification `json:"classifications"`
	// **Deprecated** — use `extractions` instead. Extracted fields keyed by field
	// name, populated only for linear workflows (single extract node, no splitter).
	// `null` for split workflows; read `extractions[]` instead.
	//
	// Deprecated: deprecated
	Extraction map[string]WorkflowGetFileResultsResponseExtractionUnion `json:"extraction" api:"nullable"`
	// Flat list of extraction datapoints. Linear workflows produce one entry with
	// `split_name=null` and `partition=null`. Split workflows produce one entry per
	// (split, partition). Empty when no extraction has run yet.
	Extractions []WorkflowGetFileResultsResponseExtraction `json:"extractions"`
	// Parsed markdown for a file.
	Parse WorkflowGetFileResultsResponseParse `json:"parse" api:"nullable"`
	// Splitter output: category-level geometry with optional partitions. Empty when
	// the workflow has no splitter.
	Splits []WorkflowGetFileResultsResponseSplit `json:"splits"`
	// Link to the AnyFormat dashboard for human review of this collection's results.
	// `null` if the dashboard URL cannot be constructed (e.g. no files in the
	// collection, or the deployment has no frontend URL configured).
	VerificationURL string `json:"verification_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionID    respjson.Field
		Classifications respjson.Field
		Extraction      respjson.Field
		Extractions     respjson.Field
		Parse           respjson.Field
		Splits          respjson.Field
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

// One classifier verdict for the collection.
type WorkflowGetFileResultsResponseClassification struct {
	// The category the document was classified as.
	Category string `json:"category" api:"required"`
	// 0-100 model confidence in the verdict.
	Confidence float64 `json:"confidence" api:"required"`
	// Free-form evidence text (the snippets the classifier cited). `null` when none
	// captured.
	Evidence string `json:"evidence" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Confidence  respjson.Field
		Evidence    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseClassification) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseClassification) UnmarshalJSON(data []byte) error {
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

// One unit of extracted data. For linear (parse->extract) workflows there is
// exactly one entry with `split_name=null` and `partition=null`. For split
// workflows there is one entry per (split, partition) pair; join with `splits[]`
// by `split_name` to look up geometry.
type WorkflowGetFileResultsResponseExtraction struct {
	// Extracted fields keyed by field name. Same shape as the legacy top-level
	// `extraction`.
	Fields map[string]WorkflowGetFileResultsResponseExtractionFieldUnion `json:"fields" api:"required"`
	// The partition value within the split. `null` when the split has no partitions.
	Partition string `json:"partition" api:"nullable"`
	// The split category this extraction belongs to. `null` for linear workflows.
	SplitName string `json:"split_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Partition   respjson.Field
		SplitName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseExtraction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseExtraction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowGetFileResultsResponseExtractionFieldUnion contains all possible
// properties and values from
// [WorkflowGetFileResultsResponseExtractionFieldExtractedField],
// [[]map[string]WorkflowGetFileResultsResponseExtractionFieldArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfWorkflowGetFileResultsResponseExtractionFieldArrayItemMap]
type WorkflowGetFileResultsResponseExtractionFieldUnion struct {
	// This field will be present if the value is a
	// [[]map[string]WorkflowGetFileResultsResponseExtractionFieldArrayItem] instead of
	// an object.
	OfMapOfWorkflowGetFileResultsResponseExtractionFieldArrayItemMap []map[string]WorkflowGetFileResultsResponseExtractionFieldArrayItem `json:",inline"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionFieldExtractedField].
	Value any `json:"value"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionFieldExtractedField].
	Confidence float64 `json:"confidence"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionFieldExtractedField].
	Evidence []WorkflowGetFileResultsResponseExtractionFieldExtractedFieldEvidence `json:"evidence"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionFieldExtractedField].
	ValueOverride any `json:"value_override"`
	// This field is from variant
	// [WorkflowGetFileResultsResponseExtractionFieldExtractedField].
	VerificationStatus string `json:"verification_status"`
	JSON               struct {
		OfMapOfWorkflowGetFileResultsResponseExtractionFieldArrayItemMap respjson.Field
		Value                                                            respjson.Field
		Confidence                                                       respjson.Field
		Evidence                                                         respjson.Field
		ValueOverride                                                    respjson.Field
		VerificationStatus                                               respjson.Field
		raw                                                              string
	} `json:"-"`
}

func (u WorkflowGetFileResultsResponseExtractionFieldUnion) AsExtractedField() (v WorkflowGetFileResultsResponseExtractionFieldExtractedField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowGetFileResultsResponseExtractionFieldUnion) AsMapOfWorkflowGetFileResultsResponseExtractionFieldArrayItemMap() (v []map[string]WorkflowGetFileResultsResponseExtractionFieldArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowGetFileResultsResponseExtractionFieldUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowGetFileResultsResponseExtractionFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One extracted field's value, confidence, and supporting evidence.
type WorkflowGetFileResultsResponseExtractionFieldExtractedField struct {
	// The extracted value. Type depends on the field's `data_type` (string, number,
	// date, etc.). `null` when extraction could not produce a value.
	Value any `json:"value" api:"required"`
	// Model confidence in the extracted value, on a 0-100 scale. `null` when the
	// backend did not produce a confidence (e.g. manual entry).
	Confidence float64 `json:"confidence" api:"nullable"`
	// Source-text snippets the model used to derive this value.
	Evidence []WorkflowGetFileResultsResponseExtractionFieldExtractedFieldEvidence `json:"evidence"`
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
func (r WorkflowGetFileResultsResponseExtractionFieldExtractedField) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowGetFileResultsResponseExtractionFieldExtractedField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A snippet of source text supporting an extracted value, with the page it came
// from.
type WorkflowGetFileResultsResponseExtractionFieldExtractedFieldEvidence struct {
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
func (r WorkflowGetFileResultsResponseExtractionFieldExtractedFieldEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowGetFileResultsResponseExtractionFieldExtractedFieldEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One extracted field's value, confidence, and supporting evidence.
type WorkflowGetFileResultsResponseExtractionFieldArrayItem struct {
	// The extracted value. Type depends on the field's `data_type` (string, number,
	// date, etc.). `null` when extraction could not produce a value.
	Value any `json:"value" api:"required"`
	// Model confidence in the extracted value, on a 0-100 scale. `null` when the
	// backend did not produce a confidence (e.g. manual entry).
	Confidence float64 `json:"confidence" api:"nullable"`
	// Source-text snippets the model used to derive this value.
	Evidence []WorkflowGetFileResultsResponseExtractionFieldArrayItemEvidence `json:"evidence"`
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
func (r WorkflowGetFileResultsResponseExtractionFieldArrayItem) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseExtractionFieldArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A snippet of source text supporting an extracted value, with the page it came
// from.
type WorkflowGetFileResultsResponseExtractionFieldArrayItemEvidence struct {
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
func (r WorkflowGetFileResultsResponseExtractionFieldArrayItemEvidence) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowGetFileResultsResponseExtractionFieldArrayItemEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsed markdown for a file.
type WorkflowGetFileResultsResponseParse struct {
	// Document content rendered as structured markdown (with `<DOCUMENT>` /
	// `<section>` tags, embedded images for the `visual` variant). `null` if parsing
	// failed.
	Markdown string `json:"markdown" api:"required"`
	// Structured per-block representation of the parsed document — derived from
	// `markdown` at retrieval time. One entry per `<section>` in document order, with
	// type-specific structured data (`rows` for tables, `image_base64` for pictures)
	// surfaced as first-class fields so consumers don't have to HTML-parse.
	Blocks []WorkflowGetFileResultsResponseParseBlock `json:"blocks"`
	// Document-level YOLO layout confidence on a 0-100 scale, char-weighted mean
	// across all blocks. `null` if no annotated sections.
	LayoutConfidence float64 `json:"layout_confidence" api:"nullable"`
	// Document-level parse confidence on a 0-100 scale, char-weighted mean of
	// per-block LLM logprob scores. `null` when no blocks have logprob-based
	// confidence.
	ParseConfidence float64 `json:"parse_confidence" api:"nullable"`
	// Plain markdown text with structural tags stripped — `<DOCUMENT>`, `<section>`,
	// `<img>`, and `<figure-content>` wrappers removed, leaving the human-readable
	// content only. Useful when feeding the parsed output into an LLM or a search
	// index that doesn't need the block-level metadata. `null` if `markdown` is null.
	Text string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown         respjson.Field
		Blocks           respjson.Field
		LayoutConfidence respjson.Field
		ParseConfidence  respjson.Field
		Text             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseParse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseParse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One semantic block of a parsed document — a structured alternative to
// pattern-matching against `<section>` tags inside `markdown`.
//
// All blocks expose the common fields (`id`, `type`, `page`, `bbox`, `confidence`,
// `content`). Type-specific structured data lives in the optional fields (`rows`
// for tables, `image_base64` for pictures). Consumers can switch on `type` to
// access the per-type fields, or treat `content` as the universal fallback.
type WorkflowGetFileResultsResponseParseBlock struct {
	// Stable block identifier in the form `p<page>_b<index>`.
	ID string `json:"id" api:"required"`
	// Normalised bounding box in [0, 1] page coordinates with keys
	// `x0`/`y0`/`x1`/`y1`.
	Bbox map[string]float64 `json:"bbox" api:"required"`
	// Raw section body — markdown for text/title blocks, HTML for tables,
	// `<figure-content>` for pictures.
	Content string `json:"content" api:"required"`
	// 0-100 YOLO layout detection confidence for this block.
	LayoutConfidence float64 `json:"layout_confidence" api:"required"`
	// 1-indexed page number this block belongs to.
	Page int64 `json:"page" api:"required"`
	// Semantic type: `text`, `title`, `section-header`, `table`, `picture`, `other`.
	Type string `json:"type" api:"required"`
	// Hyperlinks found in the content via `[text](uri)` markdown syntax.
	Hyperlinks []WorkflowGetFileResultsResponseParseBlockHyperlink `json:"hyperlinks"`
	// Inline base64-encoded cropped image for `type=picture` blocks when the response
	// was assembled from the visual markdown variant. `null` for non-picture blocks or
	// when the raw variant was used.
	ImageBase64 string `json:"image_base64" api:"nullable"`
	// 0-100 parse confidence calibrated from LLM logprobs. `null` when logprobs were
	// unavailable (e.g. text-bytes strategy).
	ParseConfidence float64 `json:"parse_confidence" api:"nullable"`
	// 2D array of table cells for `type=table` blocks — each cell is
	// `{cell_id, text}`. `null` for non-table blocks.
	Rows [][]map[string]string `json:"rows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Bbox             respjson.Field
		Content          respjson.Field
		LayoutConfidence respjson.Field
		Page             respjson.Field
		Type             respjson.Field
		Hyperlinks       respjson.Field
		ImageBase64      respjson.Field
		ParseConfidence  respjson.Field
		Rows             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseParseBlock) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseParseBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hyperlink found inside a block's content.
type WorkflowGetFileResultsResponseParseBlockHyperlink struct {
	// The display text of the link.
	Text string `json:"text" api:"required"`
	// The link target (URL, mailto:, etc.).
	Uri string `json:"uri" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseParseBlockHyperlink) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseParseBlockHyperlink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level split: which pages of which files fall under it, plus any
// partitions inside it. Extraction data lives under `extractions[]` — join by
// `split_name`.
type WorkflowGetFileResultsResponseSplit struct {
	// 0-100 aggregate confidence (min across partitions).
	Confidence int64 `json:"confidence" api:"required"`
	// Per-file page lists, union of all partitions.
	Files []WorkflowGetFileResultsResponseSplitFile `json:"files" api:"required"`
	// The split's category name.
	Name       string                                         `json:"name" api:"required"`
	Partitions []WorkflowGetFileResultsResponseSplitPartition `json:"partitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Files       respjson.Field
		Name        respjson.Field
		Partitions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseSplit) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file's contribution of pages to a split or partition. 1-indexed.
type WorkflowGetFileResultsResponseSplitFile struct {
	// The file's UUID.
	FileID string `json:"file_id" api:"required"`
	// The file's display name.
	FileName string `json:"file_name" api:"required"`
	// 1-indexed page numbers from this file.
	Pages []int64 `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		FileName    respjson.Field
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseSplitFile) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseSplitFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A partition value within a split (e.g. `1234-5678` under `Account Holdings`).
type WorkflowGetFileResultsResponseSplitPartition struct {
	// 0-100 minimum confidence across the partition's ranges.
	Confidence int64                                              `json:"confidence" api:"required"`
	Files      []WorkflowGetFileResultsResponseSplitPartitionFile `json:"files" api:"required"`
	// The partition value (free-form string).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Files       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseSplitPartition) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseSplitPartition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file's contribution of pages to a split or partition. 1-indexed.
type WorkflowGetFileResultsResponseSplitPartitionFile struct {
	// The file's UUID.
	FileID string `json:"file_id" api:"required"`
	// The file's display name.
	FileName string `json:"file_name" api:"required"`
	// 1-indexed page numbers from this file.
	Pages []int64 `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		FileName    respjson.Field
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetFileResultsResponseSplitPartitionFile) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetFileResultsResponseSplitPartitionFile) UnmarshalJSON(data []byte) error {
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
	// Field definitions. Each entry's shape is determined by its `data_type`.
	Fields []WorkflowNewParamsFieldUnion `json:"fields,omitzero" api:"required"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsFieldUnion struct {
	OfString      *WorkflowNewParamsFieldString      `json:",omitzero,inline"`
	OfInteger     *WorkflowNewParamsFieldInteger     `json:",omitzero,inline"`
	OfFloat       *WorkflowNewParamsFieldFloat       `json:",omitzero,inline"`
	OfBoolean     *WorkflowNewParamsFieldBoolean     `json:",omitzero,inline"`
	OfDate        *WorkflowNewParamsFieldDate        `json:",omitzero,inline"`
	OfDatetime    *WorkflowNewParamsFieldDatetime    `json:",omitzero,inline"`
	OfEnum        *WorkflowNewParamsFieldEnum        `json:",omitzero,inline"`
	OfMultiSelect *WorkflowNewParamsFieldMultiSelect `json:",omitzero,inline"`
	OfObject      *WorkflowNewParamsFieldObject      `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfInteger,
		u.OfFloat,
		u.OfBoolean,
		u.OfDate,
		u.OfDatetime,
		u.OfEnum,
		u.OfMultiSelect,
		u.OfObject)
}
func (u *WorkflowNewParamsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WorkflowNewParamsFieldUnion](
		"data_type",
		apijson.Discriminator[WorkflowNewParamsFieldString]("string"),
		apijson.Discriminator[WorkflowNewParamsFieldInteger]("integer"),
		apijson.Discriminator[WorkflowNewParamsFieldFloat]("float"),
		apijson.Discriminator[WorkflowNewParamsFieldBoolean]("boolean"),
		apijson.Discriminator[WorkflowNewParamsFieldDate]("date"),
		apijson.Discriminator[WorkflowNewParamsFieldDatetime]("datetime"),
		apijson.Discriminator[WorkflowNewParamsFieldEnum]("enum"),
		apijson.Discriminator[WorkflowNewParamsFieldMultiSelect]("multi_select"),
		apijson.Discriminator[WorkflowNewParamsFieldObject]("object"),
	)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldString struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsFieldString) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldInteger struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsFieldInteger) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldInteger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldInteger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldFloat struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsFieldFloat) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldFloat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldFloat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldBoolean struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsFieldBoolean) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldDate struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsFieldDate) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldDatetime struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsFieldDatetime) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldDatetime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsFieldEnum struct {
	// Free-form description shown to the extraction model.
	Description string                                 `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsFieldEnumEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsFieldEnum) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldEnum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldEnum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsFieldEnumEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsFieldEnumEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldEnumEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldEnumEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsFieldMultiSelect struct {
	// Free-form description shown to the extraction model.
	Description string                                        `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsFieldMultiSelectEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsFieldMultiSelect) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldMultiSelect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldMultiSelect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsFieldMultiSelectEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsFieldMultiSelectEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldMultiSelectEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldMultiSelectEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name, NestedFields are required.
type WorkflowNewParamsFieldObject struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name         string                                         `json:"name" api:"required"`
	NestedFields []WorkflowNewParamsFieldObjectNestedFieldUnion `json:"nested_fields,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "object".
	DataType constant.Object `json:"data_type" default:"object"`
	paramObj
}

func (r WorkflowNewParamsFieldObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsFieldObjectNestedFieldUnion struct {
	OfStringFieldDef      *WorkflowNewParamsFieldObjectNestedFieldStringFieldDef      `json:",omitzero,inline"`
	OfIntegerFieldDef     *WorkflowNewParamsFieldObjectNestedFieldIntegerFieldDef     `json:",omitzero,inline"`
	OfFloatFieldDef       *WorkflowNewParamsFieldObjectNestedFieldFloatFieldDef       `json:",omitzero,inline"`
	OfBooleanFieldDef     *WorkflowNewParamsFieldObjectNestedFieldBooleanFieldDef     `json:",omitzero,inline"`
	OfDateFieldDef        *WorkflowNewParamsFieldObjectNestedFieldDateFieldDef        `json:",omitzero,inline"`
	OfDatetimeFieldDef    *WorkflowNewParamsFieldObjectNestedFieldDatetimeFieldDef    `json:",omitzero,inline"`
	OfEnumFieldDef        *WorkflowNewParamsFieldObjectNestedFieldEnumFieldDef        `json:",omitzero,inline"`
	OfMultiSelectFieldDef *WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDef `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsFieldObjectNestedFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringFieldDef,
		u.OfIntegerFieldDef,
		u.OfFloatFieldDef,
		u.OfBooleanFieldDef,
		u.OfDateFieldDef,
		u.OfDatetimeFieldDef,
		u.OfEnumFieldDef,
		u.OfMultiSelectFieldDef)
}
func (u *WorkflowNewParamsFieldObjectNestedFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldStringFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldStringFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldStringFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldStringFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldIntegerFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldIntegerFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldIntegerFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldIntegerFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldFloatFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldFloatFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldFloatFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldFloatFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldBooleanFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldBooleanFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldBooleanFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldBooleanFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldDateFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldDateFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldDateFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldDateFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldDatetimeFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldDatetimeFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldDatetimeFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldDatetimeFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldEnumFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string                                                          `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsFieldObjectNestedFieldEnumFieldDefEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldEnumFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldEnumFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldEnumFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldEnumFieldDefEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldEnumFieldDefEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldEnumFieldDefEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldEnumFieldDefEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDef struct {
	// Free-form description shown to the extraction model.
	Description string                                                                 `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDefEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDef) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDef
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDefEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDefEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDefEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsFieldObjectNestedFieldMultiSelectFieldDefEnumOption) UnmarshalJSON(data []byte) error {
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
