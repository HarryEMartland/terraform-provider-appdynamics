package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccDataSourceAppdService_basic(t *testing.T) {
	resourceName := "data.appdynamics_dashboard_widget.test"

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppdServiceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceAwsArn(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "json"),
					resource.TestCheckResourceAttrSet(resourceName, "widget_json"),
				),
			},
		},
	})
}

//func testAccDataSourceAwsArn(name string) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		_, ok := s.RootModule().Resources[name]
//		if !ok {
//			return fmt.Errorf("root module has no resource called %s", name)
//		}
//
//		return nil
//	}
//}
//
func testAccDataSourceAppdServiceConfig() string {
	return fmt.Sprintf(`
%s

data "appdynamics_appd_service" "test" {
     application_name = "TEST"
     tier_name = "TEST"
}
`, configureDashboardConfig())
}
