package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccDataSourceAppdService_basic(t *testing.T) {
	resourceName := "data.appdynamics_appd_service.test"

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppdServiceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceAwsArn(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tier_id"),
				),
			},
		},
	})
}

func testAccDataSourceAppdServiceConfig() string {
	return fmt.Sprintf(`
%s

data "appdynamics_appd_service" "test" {
     application_name = "TST-Client-Facing"
     tier_name = "airtime-purchase"
}
`, configureDashboardConfig())
}
