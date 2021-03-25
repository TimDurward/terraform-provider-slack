package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccDataSourceSlackUser_Basic(t *testing.T) {
	resourcesConfig := `
resource "slack_user" "person" {
  name = "Jane"
  real_name = "Jane Doe"
  email = "janedoe@gmail.com"
}
`
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: resourcesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.slack_user.person", "name"),
					resource.TestCheckResourceAttrSet("data.slack_user.person", "real_name"),
					resource.TestCheckResourceAttrSet("data.slack_user.person", "email"),
					resource.TestCheckResourceAttr("data.slack_user.person", "name", "Jane"),
					resource.TestCheckResourceAttr("data.slack_user.person", "real_name", "Jane Doe"),
					resource.TestCheckResourceAttr("data.slack_user.person", "email", "janedoe@gmail.com"),
				),
			},
		},
	})
}