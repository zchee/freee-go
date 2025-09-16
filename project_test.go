// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/freee-go"
	"github.com/stainless-sdks/freee-go/internal/testutil"
	"github.com/stainless-sdks/freee-go/option"
)

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.New(context.TODO(), freee.ProjectNewParams{
		Code:                 "code",
		CompanyID:            1,
		FromDate:             "from_date",
		Name:                 "name",
		PmBudgetsCost:        4000,
		ThruDate:             "thru_date",
		AssignmentURLEnabled: freee.Bool(true),
		ColorID:              freee.Int(3),
		ContractorIDs:        []int64{10},
		Description:          freee.String("description"),
		ManagerPersonID:      freee.Int(10),
		Members: []freee.ProjectNewParamsMember{{
			BudgetsCost:         2000,
			PersonID:            11,
			UnitCostID:          3,
			UseStandardUnitCost: freee.Bool(true),
		}},
		OrdererIDs:         []int64{20},
		PublishToEmployee:  freee.Bool(true),
		SalesOrderStatusID: freee.Int(2),
	})
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectGet(t *testing.T) {
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
	_, err := client.Projects.Get(
		context.TODO(),
		0,
		freee.ProjectGetParams{
			CompanyID: 0,
		},
	)
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.List(context.TODO(), freee.ProjectListParams{
		CompanyID:         0,
		ContractorIDs:     []int64{0},
		Limit:             freee.Int(1),
		ManagerIDs:        []int64{0},
		Offset:            freee.Int(0),
		OperationalStatus: freee.ProjectListParamsOperationalStatusPlanning,
		OrdererIDs:        []int64{0},
	})
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
