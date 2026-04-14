// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/anyformat-go/internal/apijson"
	"github.com/stainless-sdks/anyformat-go/internal/requestconfig"
	"github.com/stainless-sdks/anyformat-go/option"
	"github.com/stainless-sdks/anyformat-go/packages/param"
	"github.com/stainless-sdks/anyformat-go/packages/respjson"
)

// Webhook subscriptions for async notifications.
//
// WebhookService contains methods and other services that help with interacting
// with the anyformat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	return
}

// Create a new webhook subscription.
//
// Validates URL (HTTPS only) and event types, then forwards to backend service.
// Returns the created webhook with generated secret.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/webhooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List all webhook subscriptions for the authenticated organization.
//
// Returns a list of webhooks (secrets are excluded in list view).
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *[]WebhookListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v2/webhooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook subscription by ID.
//
// Returns 204 on success, 404 if webhook not found, 403 if unauthorized.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/webhooks/%s/", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Response schema for webhook subscription (includes secret)
type WebhookNewResponse struct {
	ID        string   `json:"id" api:"required"`
	CreatedAt string   `json:"created_at" api:"required"`
	Events    []string `json:"events" api:"required"`
	IsActive  bool     `json:"is_active" api:"required"`
	Secret    string   `json:"secret" api:"required"`
	URL       string   `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Events      respjson.Field
		IsActive    respjson.Field
		Secret      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response schema for listing webhooks (excludes secret)
type WebhookListResponse struct {
	ID        string   `json:"id" api:"required"`
	CreatedAt string   `json:"created_at" api:"required"`
	Events    []string `json:"events" api:"required"`
	IsActive  bool     `json:"is_active" api:"required"`
	URL       string   `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Events      respjson.Field
		IsActive    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	URL    string   `json:"url" api:"required" format:"uri"`
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
