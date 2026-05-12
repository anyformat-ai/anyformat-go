// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anyformat_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/anyformat-ai/anyformat-go"
	"github.com/anyformat-ai/anyformat-go/internal/testutil"
	"github.com/anyformat-ai/anyformat-go/option"
)

func TestWorkflowNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.New(context.TODO(), anyformat.WorkflowNewParams{
		Fields: []anyformat.WorkflowNewParamsFieldUnion{{
			OfString: &anyformat.WorkflowNewParamsFieldString{
				Description: "x",
				Name:        "invoice_number",
			},
		}},
		Name:        "Invoice Processing",
		Description: anyformat.String("Extracts invoice number, vendor, total, and line items."),
	})
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.Get(context.TODO(), "workflow_id")
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.List(context.TODO(), anyformat.WorkflowListParams{
		Order:    anyformat.String("order"),
		Page:     anyformat.Int(1),
		PageSize: anyformat.Int(1),
		SortBy:   anyformat.String("sort_by"),
		Status:   anyformat.String("status"),
	})
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Workflows.Delete(context.TODO(), "workflow_id")
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowNewFile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.NewFile(
		context.TODO(),
		"workflow_id",
		anyformat.WorkflowNewFileParams{
			Files: []string{"string"},
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowGetFileResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.GetFileResults(
		context.TODO(),
		"collection_id",
		anyformat.WorkflowGetFileResultsParams{
			WorkflowID: "workflow_id",
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowListFilesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.ListFiles(
		context.TODO(),
		"workflow_id",
		anyformat.WorkflowListFilesParams{
			Page:     anyformat.Int(1),
			PageSize: anyformat.Int(1),
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowListRunsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.ListRuns(
		context.TODO(),
		"workflow_id",
		anyformat.WorkflowListRunsParams{
			Page:     anyformat.Int(1),
			PageSize: anyformat.Int(1),
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowRunWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.Run(
		context.TODO(),
		"workflow_id",
		anyformat.WorkflowRunParams{
			File: anyformat.String("file"),
			Text: anyformat.String("text"),
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowUploadWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anyformat.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workflows.Upload(
		context.TODO(),
		"workflow_id",
		anyformat.WorkflowUploadParams{
			File: anyformat.String("file"),
			Text: anyformat.String("text"),
		},
	)
	if err != nil {
		var apierr *anyformat.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
