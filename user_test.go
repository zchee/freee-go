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

func TestUserGetMe(t *testing.T) {
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
	_, err := client.Users.GetMe(context.TODO())
	if err != nil {
		var apierr *freeepm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
