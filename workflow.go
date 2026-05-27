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

// Create a workflow from a strongly-typed graph (atomic).
//
// Provide an explicit list of typed `nodes` (parse / classify / splitter /
// extract) and `edges` between them. The full workflow — fields, nodes, routing —
// is created in a single transaction.
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
	// Document content rendered as structured markdown. Each block is preceded by an
	// empty `<a id="p{page}_b{idx}"></a>` anchor (invisible in any markdown renderer;
	// the id joins to `blocks[].id` and can be used as an in-page link target). Image
	// hydration for picture/figure blocks happens client-side. `null` if parsing
	// failed.
	Markdown string `json:"markdown" api:"required"`
	// Structured per-block representation of the parsed document. One entry per
	// `<a id></a>` anchor in document order, with type-specific structured data
	// (`rows` for tables, `image_base64` for pictures) surfaced as first-class fields
	// so consumers don't have to HTML-parse.
	Blocks []WorkflowGetFileResultsResponseParseBlock `json:"blocks"`
	// Document-level YOLO layout confidence on a 0-100 scale, char-weighted mean
	// across all blocks. `null` if no annotated sections.
	LayoutConfidence float64 `json:"layout_confidence" api:"nullable"`
	// Document-level parse confidence on a 0-100 scale, char-weighted mean of
	// per-block LLM logprob scores. `null` when no blocks have logprob-based
	// confidence.
	ParseConfidence float64 `json:"parse_confidence" api:"nullable"`
	// Plain markdown with structural HTML removed — `<DOCUMENT>` framing, block
	// anchors, `<img>` tags, and `<figure-content>` wrappers stripped. Useful when
	// feeding the parsed output into an LLM or a search index that doesn't need the
	// block-level metadata. `null` if `markdown` is null.
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

// One semantic block of a parsed document — a structured alternative to walking
// `<a id></a>` anchors in `markdown`.
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
	// Inline base64-encoded cropped image for `type=picture` blocks. Currently `null`
	// for all blocks — image hydration is performed client-side by the SDK consumer.
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
	// The workflow version this run was bound to (the latest version at submission
	// time). Lets callers verify which schema produced the results — useful right
	// after an edit.
	VersionID string `json:"version_id" api:"required"`
	// The UUID of the workflow that was executed.
	WorkflowID string `json:"workflow_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		VersionID   respjson.Field
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
	Name        string                       `json:"name" api:"required"`
	Nodes       []WorkflowNewParamsNodeUnion `json:"nodes,omitzero" api:"required"`
	Description param.Opt[string]            `json:"description,omitzero"`
	Edges       []WorkflowNewParamsEdge      `json:"edges,omitzero"`
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
type WorkflowNewParamsNodeUnion struct {
	OfParse    *WorkflowNewParamsNodeParse    `json:",omitzero,inline"`
	OfClassify *WorkflowNewParamsNodeClassify `json:",omitzero,inline"`
	OfSplitter *WorkflowNewParamsNodeSplitter `json:",omitzero,inline"`
	OfExtract  *WorkflowNewParamsNodeExtract  `json:",omitzero,inline"`
	OfValidate *WorkflowNewParamsNodeValidate `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsNodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfParse,
		u.OfClassify,
		u.OfSplitter,
		u.OfExtract,
		u.OfValidate)
}
func (u *WorkflowNewParamsNodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WorkflowNewParamsNodeUnion](
		"type",
		apijson.Discriminator[WorkflowNewParamsNodeParse]("parse"),
		apijson.Discriminator[WorkflowNewParamsNodeClassify]("classify"),
		apijson.Discriminator[WorkflowNewParamsNodeSplitter]("splitter"),
		apijson.Discriminator[WorkflowNewParamsNodeExtract]("extract"),
		apijson.Discriminator[WorkflowNewParamsNodeValidate]("validate"),
	)
}

// The properties ID, Type are required.
type WorkflowNewParamsNodeParse struct {
	// Stable identifier for this node within the graph.
	ID string `json:"id" api:"required"`
	// Free-form hint shown to the parse model to bias output.
	PromptHint        param.Opt[string] `json:"prompt_hint,omitzero"`
	FigureEnhancement param.Opt[bool]   `json:"figure_enhancement,omitzero"`
	// Any of "standard", "agentic".
	Mode string `json:"mode,omitzero"`
	// This field can be elided, and will marshal its zero value as "parse".
	Type constant.Parse `json:"type" default:"parse"`
	paramObj
}

func (r WorkflowNewParamsNodeParse) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeParse
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeParse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsNodeParse](
		"mode", "standard", "agentic",
	)
}

// The properties ID, Categories, Type are required.
type WorkflowNewParamsNodeClassify struct {
	// Stable identifier for this node within the graph.
	ID         string                                  `json:"id" api:"required"`
	Categories []WorkflowNewParamsNodeClassifyCategory `json:"categories,omitzero" api:"required"`
	// Optional prompt prefix for the classifier.
	UserPrompt param.Opt[string] `json:"user_prompt,omitzero"`
	// This field can be elided, and will marshal its zero value as "classify".
	Type constant.Classify `json:"type" default:"classify"`
	paramObj
}

func (r WorkflowNewParamsNodeClassify) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeClassify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Description, Name are required.
type WorkflowNewParamsNodeClassifyCategory struct {
	// Stable category id used as the edge `branch` value when routing.
	ID string `json:"id" api:"required"`
	// Free-form description shown to the LLM.
	Description string `json:"description" api:"required"`
	// Display name shown to the LLM.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeClassifyCategory) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeClassifyCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeClassifyCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Rules, Type are required.
type WorkflowNewParamsNodeSplitter struct {
	// Stable identifier for this node within the graph.
	ID    string                              `json:"id" api:"required"`
	Rules []WorkflowNewParamsNodeSplitterRule `json:"rules,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "splitter".
	Type constant.Splitter `json:"type" default:"splitter"`
	paramObj
}

func (r WorkflowNewParamsNodeSplitter) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeSplitter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeSplitter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Description, Name are required.
type WorkflowNewParamsNodeSplitterRule struct {
	ID           string            `json:"id" api:"required"`
	Description  string            `json:"description" api:"required"`
	Name         string            `json:"name" api:"required"`
	PartitionKey param.Opt[string] `json:"partition_key,omitzero"`
	paramObj
}

func (r WorkflowNewParamsNodeSplitterRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeSplitterRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeSplitterRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, ExtractionSchema, Type are required.
type WorkflowNewParamsNodeExtract struct {
	// Stable identifier for this node within the graph.
	ID string `json:"id" api:"required"`
	// Schema for the fields this node extracts. Required. The backend executor
	// populates this from the ORM field tree before constructing the typed node;
	// nothing else should be able to build an ExtractNode without it.
	ExtractionSchema WorkflowNewParamsNodeExtractExtractionSchema `json:"extraction_schema,omitzero" api:"required"`
	// Free-form hint shown to the smart-lookup matcher.
	LookupSuggestion param.Opt[string] `json:"lookup_suggestion,omitzero"`
	UseImages        param.Opt[bool]   `json:"use_images,omitzero"`
	// Inline lookup-file content for the typed create call. The backend uploads each
	// entry to S3 and stores the resulting URI in `lookup_files`; this field is never
	// persisted in GraphNode.config.
	LookupFileUploads []WorkflowNewParamsNodeExtractLookupFileUpload `json:"lookup_file_uploads,omitzero"`
	// Smart-lookup reference document URIs persisted on the extract node.
	LookupFiles []string `json:"lookup_files,omitzero"`
	// Typed schema of fields the smart-lookup pass should produce. The _backend_
	// derives this from the field tree (FieldWorkflowVersion rows whose source is
	// SMART_LOOKUP) and attaches it to the node before the message hits the wire — the
	// worker then reads it directly off the node, with no DB round-trip. Default
	// empty: a node without smart-lookup fields carries an empty list.
	LookupSchema []WorkflowNewParamsNodeExtractLookupSchemaUnion `json:"lookup_schema,omitzero"`
	// Any of "standard", "agentic".
	Mode string `json:"mode,omitzero"`
	// This field can be elided, and will marshal its zero value as "extract".
	Type constant.Extract `json:"type" default:"extract"`
	paramObj
}

func (r WorkflowNewParamsNodeExtract) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsNodeExtract](
		"mode", "standard", "agentic",
	)
}

// Schema for the fields this node extracts. Required. The backend executor
// populates this from the ORM field tree before constructing the typed node;
// nothing else should be able to build an ExtractNode without it.
//
// The property Fields is required.
type WorkflowNewParamsNodeExtractExtractionSchema struct {
	// Field definitions making up this extract's output.
	Fields []WorkflowNewParamsNodeExtractExtractionSchemaFieldUnion `json:"fields,omitzero" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchema) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldUnion struct {
	OfString      *WorkflowNewParamsNodeExtractExtractionSchemaFieldString      `json:",omitzero,inline"`
	OfInteger     *WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger     `json:",omitzero,inline"`
	OfFloat       *WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat       `json:",omitzero,inline"`
	OfBoolean     *WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean     `json:",omitzero,inline"`
	OfDate        *WorkflowNewParamsNodeExtractExtractionSchemaFieldDate        `json:",omitzero,inline"`
	OfDatetime    *WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime    `json:",omitzero,inline"`
	OfEnum        *WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum        `json:",omitzero,inline"`
	OfMultiSelect *WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect `json:",omitzero,inline"`
	OfObject      *WorkflowNewParamsNodeExtractExtractionSchemaFieldObject      `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsNodeExtractExtractionSchemaFieldUnion) MarshalJSON() ([]byte, error) {
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
func (u *WorkflowNewParamsNodeExtractExtractionSchemaFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WorkflowNewParamsNodeExtractExtractionSchemaFieldUnion](
		"data_type",
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldString]("string"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger]("integer"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat]("float"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean]("boolean"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldDate]("date"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime]("datetime"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum]("enum"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect]("multi_select"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractExtractionSchemaFieldObject]("object"),
	)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldString struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldString) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldInteger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldFloat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldDate struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldDate) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum struct {
	// Free-form description shown to the extraction model.
	Description string                                                            `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractExtractionSchemaFieldEnumEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldEnum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldEnumEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldEnumEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldEnumEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldEnumEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect struct {
	// Free-form description shown to the extraction model.
	Description string                                                                   `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelectEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelectEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelectEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelectEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldMultiSelectEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name, NestedFields are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObject struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name         string                                                                    `json:"name" api:"required"`
	NestedFields []WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldUnion `json:"nested_fields,omitzero" api:"required"`
	Lookup       param.Opt[bool]                                                           `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "object".
	DataType constant.Object `json:"data_type" default:"object"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldUnion struct {
	OfStringField      *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldStringField      `json:",omitzero,inline"`
	OfIntegerField     *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldIntegerField     `json:",omitzero,inline"`
	OfFloatField       *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldFloatField       `json:",omitzero,inline"`
	OfBooleanField     *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldBooleanField     `json:",omitzero,inline"`
	OfDateField        *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDateField        `json:",omitzero,inline"`
	OfDatetimeField    *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDatetimeField    `json:",omitzero,inline"`
	OfEnumField        *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumField        `json:",omitzero,inline"`
	OfMultiSelectField *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectField `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringField,
		u.OfIntegerField,
		u.OfFloatField,
		u.OfBooleanField,
		u.OfDateField,
		u.OfDatetimeField,
		u.OfEnumField,
		u.OfMultiSelectField)
}
func (u *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldStringField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldStringField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldStringField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldStringField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldIntegerField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldIntegerField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldIntegerField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldIntegerField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldFloatField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldFloatField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldFloatField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldFloatField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldBooleanField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldBooleanField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldBooleanField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDateField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDateField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDateField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDatetimeField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDatetimeField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDatetimeField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldDatetimeField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumField struct {
	// Free-form description shown to the extraction model.
	Description string                                                                                  `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumFieldEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumFieldEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumFieldEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumFieldEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldEnumFieldEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectField struct {
	// Free-form description shown to the extraction model.
	Description string                                                                                         `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectFieldEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectFieldEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectFieldEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectFieldEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractExtractionSchemaFieldObjectNestedFieldMultiSelectFieldEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline lookup-file content carried on the typed create call.
//
// The backend reads `filename` + `content` (base64-encoded bytes), uploads the
// file to S3 during workflow create, and stores the resulting URI in
// `ExtractNode.lookup_files`. This field is stripped from the persisted
// `GraphNode.config` — it is create-input only.
//
// The properties Content, Filename are required.
type WorkflowNewParamsNodeExtractLookupFileUpload struct {
	// Base64-encoded file bytes.
	Content  string `json:"content" api:"required"`
	Filename string `json:"filename" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupFileUpload) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupFileUpload
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupFileUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsNodeExtractLookupSchemaUnion struct {
	OfString      *WorkflowNewParamsNodeExtractLookupSchemaString      `json:",omitzero,inline"`
	OfInteger     *WorkflowNewParamsNodeExtractLookupSchemaInteger     `json:",omitzero,inline"`
	OfFloat       *WorkflowNewParamsNodeExtractLookupSchemaFloat       `json:",omitzero,inline"`
	OfBoolean     *WorkflowNewParamsNodeExtractLookupSchemaBoolean     `json:",omitzero,inline"`
	OfDate        *WorkflowNewParamsNodeExtractLookupSchemaDate        `json:",omitzero,inline"`
	OfDatetime    *WorkflowNewParamsNodeExtractLookupSchemaDatetime    `json:",omitzero,inline"`
	OfEnum        *WorkflowNewParamsNodeExtractLookupSchemaEnum        `json:",omitzero,inline"`
	OfMultiSelect *WorkflowNewParamsNodeExtractLookupSchemaMultiSelect `json:",omitzero,inline"`
	OfObject      *WorkflowNewParamsNodeExtractLookupSchemaObject      `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsNodeExtractLookupSchemaUnion) MarshalJSON() ([]byte, error) {
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
func (u *WorkflowNewParamsNodeExtractLookupSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WorkflowNewParamsNodeExtractLookupSchemaUnion](
		"data_type",
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaString]("string"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaInteger]("integer"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaFloat]("float"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaBoolean]("boolean"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaDate]("date"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaDatetime]("datetime"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaEnum]("enum"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaMultiSelect]("multi_select"),
		apijson.Discriminator[WorkflowNewParamsNodeExtractLookupSchemaObject]("object"),
	)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaString struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaString) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaString
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaString) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaInteger struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaInteger) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaInteger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaInteger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaFloat struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaFloat) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaFloat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaFloat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaBoolean struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaBoolean) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaBoolean
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaBoolean) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaDate struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaDate) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaDate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaDatetime struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaDatetime) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaDatetime
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaDatetime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaEnum struct {
	// Free-form description shown to the extraction model.
	Description string                                                   `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractLookupSchemaEnumEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaEnum) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaEnum
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaEnum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaEnumEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaEnumEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaEnumEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaEnumEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaMultiSelect struct {
	// Free-form description shown to the extraction model.
	Description string                                                          `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractLookupSchemaMultiSelectEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaMultiSelect) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaMultiSelect
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaMultiSelect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaMultiSelectEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaMultiSelectEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaMultiSelectEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaMultiSelectEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name, NestedFields are required.
type WorkflowNewParamsNodeExtractLookupSchemaObject struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name         string                                                           `json:"name" api:"required"`
	NestedFields []WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldUnion `json:"nested_fields,omitzero" api:"required"`
	Lookup       param.Opt[bool]                                                  `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "object".
	DataType constant.Object `json:"data_type" default:"object"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldUnion struct {
	OfStringField      *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldStringField      `json:",omitzero,inline"`
	OfIntegerField     *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldIntegerField     `json:",omitzero,inline"`
	OfFloatField       *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldFloatField       `json:",omitzero,inline"`
	OfBooleanField     *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldBooleanField     `json:",omitzero,inline"`
	OfDateField        *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDateField        `json:",omitzero,inline"`
	OfDatetimeField    *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDatetimeField    `json:",omitzero,inline"`
	OfEnumField        *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumField        `json:",omitzero,inline"`
	OfMultiSelectField *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectField `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringField,
		u.OfIntegerField,
		u.OfFloatField,
		u.OfBooleanField,
		u.OfDateField,
		u.OfDatetimeField,
		u.OfEnumField,
		u.OfMultiSelectField)
}
func (u *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldStringField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "string".
	DataType constant.String `json:"data_type" default:"string"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldStringField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldStringField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldStringField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldIntegerField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "integer".
	DataType constant.Integer `json:"data_type" default:"integer"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldIntegerField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldIntegerField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldIntegerField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldFloatField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "float".
	DataType constant.Float `json:"data_type" default:"float"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldFloatField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldFloatField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldFloatField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldBooleanField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "boolean".
	DataType constant.Boolean `json:"data_type" default:"boolean"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldBooleanField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldBooleanField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDateField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "date".
	DataType constant.Date `json:"data_type" default:"date"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDateField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDateField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDatetimeField struct {
	// Free-form description shown to the extraction model.
	Description string `json:"description" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "datetime".
	DataType constant.Datetime `json:"data_type" default:"datetime"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDatetimeField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDatetimeField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldDatetimeField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumField struct {
	// Free-form description shown to the extraction model.
	Description string                                                                         `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumFieldEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "enum".
	DataType constant.Enum `json:"data_type" default:"enum"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumFieldEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumFieldEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumFieldEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldEnumFieldEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Description, EnumOptions, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectField struct {
	// Free-form description shown to the extraction model.
	Description string                                                                                `json:"description" api:"required"`
	EnumOptions []WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectFieldEnumOption `json:"enum_options,omitzero" api:"required"`
	// Field name. Used as the key in the extraction response.
	Name   string          `json:"name" api:"required"`
	Lookup param.Opt[bool] `json:"lookup,omitzero"`
	// This field can be elided, and will marshal its zero value as "multi_select".
	DataType constant.MultiSelect `json:"data_type" default:"multi_select"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectField) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name are required.
type WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectFieldEnumOption struct {
	// Free-form description shown to the model.
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	paramObj
}

func (r WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectFieldEnumOption) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectFieldEnumOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeExtractLookupSchemaObjectNestedFieldMultiSelectFieldEnumOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Rules, Type are required.
type WorkflowNewParamsNodeValidate struct {
	// Stable identifier for this node within the graph.
	ID    string                              `json:"id" api:"required"`
	Rules []WorkflowNewParamsNodeValidateRule `json:"rules,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "validate".
	Type constant.Validate `json:"type" default:"validate"`
	paramObj
}

func (r WorkflowNewParamsNodeValidate) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeValidate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeValidate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Description are required.
type WorkflowNewParamsNodeValidateRule struct {
	// Stable rule id; round-trips through ValidationResult.rule_id.
	ID string `json:"id" api:"required"`
	// Natural-language description shown to the validation model.
	Description string `json:"description" api:"required"`
	// Optional human-readable rule name shown on the rule card in the Studio. Stored
	// verbatim on the GraphNode config and surfaced back through the config endpoint
	// so renames round-trip.
	Name param.Opt[string] `json:"name,omitzero"`
	// Any of "error", "warning".
	Severity string `json:"severity,omitzero"`
	// Persistent ids of fields this rule references.
	SourceFields []string `json:"source_fields,omitzero"`
	paramObj
}

func (r WorkflowNewParamsNodeValidateRule) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNodeValidateRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNodeValidateRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowNewParamsNodeValidateRule](
		"severity", "error", "warning",
	)
}

// A directed edge between two nodes. `branch` carries the source-port label used
// for routing out of `classify` (category id) or `splitter` (rule id) nodes.
//
// The properties Source, Target are required.
type WorkflowNewParamsEdge struct {
	Source string `json:"source" api:"required"`
	Target string `json:"target" api:"required"`
	// Source-port label for branch routing. Required when leaving a classify or
	// splitter node by category/rule.
	Branch param.Opt[string] `json:"branch,omitzero"`
	paramObj
}

func (r WorkflowNewParamsEdge) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsEdge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsEdge) UnmarshalJSON(data []byte) error {
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
