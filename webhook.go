// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/anyformat-ai/anyformat-go/internal/apijson"
	"github.com/anyformat-ai/anyformat-go/internal/requestconfig"
	"github.com/anyformat-ai/anyformat-go/option"
	"github.com/anyformat-ai/anyformat-go/packages/param"
	"github.com/anyformat-ai/anyformat-go/packages/respjson"
)

// Webhook subscriptions for asynchronous event notifications. Get notified when
// extractions complete or fail.
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

// Create a new webhook subscription for your organization.
//
// The webhook URL must use HTTPS. AnyFormat will send POST requests to this URL
// when the subscribed events occur. The response includes a `secret` that you
// should use to verify webhook signatures.
//
// Supported events:
//
// - `extraction.completed` — fired when a file extraction finishes successfully.
// - `extraction.failed` — fired when a file extraction fails.
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "v2/webhooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List all webhook subscriptions for the authenticated organization.
//
// Returns a list of webhooks. Secrets are excluded from the list response for
// security — they are only returned once, when the webhook is created.
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *[]WebhookListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "v2/webhooks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook subscription by ID.
//
// After deletion, AnyFormat will stop sending events to the webhook URL. This
// action is irreversible.
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/webhooks/%s/", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Webhook subscription details including the signing secret. The secret is only
// returned at creation time.
type WebhookNewResponse struct {
	// Unique identifier of the webhook.
	ID string `json:"id" api:"required"`
	// Timestamp when the webhook was created (ISO 8601).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Event types this webhook is subscribed to.
	Events []string `json:"events" api:"required"`
	// Whether the webhook is currently active and receiving events.
	IsActive bool `json:"is_active" api:"required"`
	// Webhook signing secret. Use this to verify that incoming webhook requests are
	// authentic. **Store securely — this value is only shown once at creation time.**
	Secret string `json:"secret" api:"required"`
	// The URL receiving webhook events.
	URL string `json:"url" api:"required" format:"uri"`
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

// Webhook subscription details (secret excluded for security).
type WebhookListResponse struct {
	// Unique identifier of the webhook.
	ID string `json:"id" api:"required"`
	// Timestamp when the webhook was created (ISO 8601).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Event types this webhook is subscribed to.
	Events []string `json:"events" api:"required"`
	// Whether the webhook is currently active.
	IsActive bool `json:"is_active" api:"required"`
	// The URL receiving webhook events.
	URL string `json:"url" api:"required" format:"uri"`
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
	// The HTTPS URL to receive webhook events. Must be publicly accessible.
	URL string `json:"url" api:"required" format:"uri"`
	// List of event types to subscribe to. Available events: `extraction.completed`,
	// `extraction.failed`.
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
