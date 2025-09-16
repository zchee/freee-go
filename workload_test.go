// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/zchee/freee-go"
	"github.com/zchee/freee-go/internal/testutil"
	"github.com/zchee/freee-go/option"
)

func TestWorkloadNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := freee.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workloads.New(context.TODO(), freee.WorkloadNewParams{
		CompanyID: 1,
		Date:      time.Now(),
		Minutes:   120,
		ProjectID: 100,
		Memo:      freee.String("コーディング"),
		PersonID:  freee.Int(10),
		WorkloadTags: []freee.WorkloadNewParamsWorkloadTag{{
			TagGroupID: 11,
			TagID:      12,
		}},
	})
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkloadListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := freee.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workloads.List(context.TODO(), freee.WorkloadListParams{
		CompanyID:      0,
		YearMonth:      "year_month",
		EmployeesScope: freee.WorkloadListParamsEmployeesScopeAll,
		Limit:          freee.Int(1),
		Offset:         freee.Int(0),
		PersonIDs:      []int64{0},
		TeamIDs:        []int64{0},
	})
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
