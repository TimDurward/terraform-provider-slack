package slack

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider *schema.Provider

)

func init() {
	testAccProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v:= os.Getenv("SLACK_API_TOKEN"); v == "" {
		t.Fatal("SLACK_API_TOKEN must be set for acceptance tests")
	}
}
