// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freeepm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/zchee/freee-go"
	"github.com/zchee/freee-go/internal/testutil"
	"github.com/zchee/freee-go/option"
)

func TestProjectNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := freeepm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Projects.New(context.TODO(), freeepm.ProjectNewParams{
		Code:                 "code",
		CompanyID:            1,
		FromDate:             "from_date",
		Name:                 "name",
		PmBudgetsCost:        4000,
		ThruDate:             "thru_date",
		AssignmentURLEnabled: freeepm.Bool(true),
		ColorID:              freeepm.Int(3),
		ContractorIDs:        []int64{10},
		Description:          freeepm.String("description"),
		ManagerPersonID:      freeepm.Int(10),
		Members: []freeepm.ProjectNewParamsMember{{
			BudgetsCost:         2000,
			PersonID:            11,
			UnitCostID:          3,
			UseStandardUnitCost: freeepm.Bool(true),
		}},
		OrdererIDs:         []int64{20},
		PublishToEmployee:  freeepm.Bool(true),
		SalesOrderStatusID: freeepm.Int(2),
	})
	if err != nil {
		var apierr *freeepm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := freeepm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Projects.Get(
		context.TODO(),
		0,
		freeepm.ProjectGetParams{
			CompanyID: 0,
		},
	)
	if err != nil {
		var apierr *freeepm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := freeepm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Projects.List(context.TODO(), freeepm.ProjectListParams{
		CompanyID:         0,
		ContractorIDs:     []int64{0},
		Limit:             freeepm.Int(1),
		ManagerIDs:        []int64{0},
		Offset:            freeepm.Int(0),
		OperationalStatus: freeepm.ProjectListParamsOperationalStatusPlanning,
		OrdererIDs:        []int64{0},
	})
	if err != nil {
		var apierr *freeepm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
