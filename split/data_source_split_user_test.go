package split

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strings"
	"testing"
)

func TestAccDatasourceSplitUser_Basic(t *testing.T) {
	email := testAccConfig.GetUserEmailorSkip(t)
	emailSplit := strings.Split(email, "@")
	emailFormatted := fmt.Sprintf("%s+%s@%s", emailSplit[0], acctest.RandString(8), emailSplit[1])

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSplitUserDataSource_Basic(emailFormatted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.split_user.foobar", "email", emailFormatted),
					resource.TestCheckResourceAttr(
						"data.split_user.foobar", "2fa", "false"),
					resource.TestCheckResourceAttr(
						"data.split_user.foobar", "status", "PENDING"),
					resource.TestCheckResourceAttr(
						"data.split_user.foobar", "name", strings.Split(emailFormatted, "@")[0]),
				),
			},
		},
	})
}

func testAccCheckSplitUserDataSource_Basic(email string) string {
	return fmt.Sprintf(`
data "split_user" "foobar" {
  email = "%s"
}
`, email)
}
