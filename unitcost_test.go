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

func TestUnitCostListWithOptionalParams(t *testing.T) {
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
	_, err := client.UnitCosts.List(context.TODO(), freee.UnitCostListParams{
		CompanyID: 0,
		Limit:     freee.Int(1),
		Offset:    freee.Int(0),
	})
	if err != nil {
		var apierr *freee.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
