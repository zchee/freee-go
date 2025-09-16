// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee_test

import (
	"context"
	"os"
	"testing"

	"github.com/zchee/freee-go"
	"github.com/zchee/freee-go/internal/testutil"
	"github.com/zchee/freee-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Users.GetMe(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ID)
}
