package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"regexp"
	"testing"
)

func TestAccDataSourceAppdService_basic(t *testing.T) {
	resourceName := "data.appdynamics_appd_service.test"

	applicationName := "TST-Client-Facing"
	tierName := "airtime-purchase"

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppdServiceConfig(applicationName, tierName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tier_id"),
				),
			},
		},
	})
}

func TestAccDataSourceAppdService_notExist(t *testing.T) {
	properApplicationName := "TST-Client-Facing"
	wrongApplicationName := "not exist"
	wrongTierName := "error please"

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				ExpectError: regexp.MustCompile(`Error getting application: 400`),
				Config:      testAccDataSourceAppdServiceConfig(wrongApplicationName, wrongTierName),
			},
			{
				ExpectError: regexp.MustCompile(`Error getting tiers: 400`),
				Config:      testAccDataSourceAppdServiceConfig(properApplicationName, wrongTierName),
			},
		},
	})
}

func testAccDataSourceAppdServiceConfig(applicationName string, tierName string) string {
	return fmt.Sprintf(`
%s

data "appdynamics_appd_service" "test" {
     application_name = "%s"
	 tier_name = "%s"
}
`, configureDashboardConfig(), applicationName, tierName)
}
