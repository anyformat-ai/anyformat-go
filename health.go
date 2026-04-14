// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat

import (
	"context"
	"net/http"
	"slices"

	"github.com/anyformat-ai/anyformat-go/internal/requestconfig"
	"github.com/anyformat-ai/anyformat-go/option"
)

// Health checks.
//
// HealthService contains methods and other services that help with interacting
// with the anyformat API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.options = opts
	return
}

// Health check endpoint.
//
// Returns 200 OK if the service is running. No authentication required.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "health/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type HealthCheckResponse = any
