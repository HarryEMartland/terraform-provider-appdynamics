package appdynamics

import (
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccDataSourceDashboardWidget_basic(t *testing.T) {
	resourceName := "data.appdynamics_dashboard_widget.test"

	testDashboardWidget := client.DashboardWidget{
		Type: "TIMESERIES_GRAPH",
	}

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDashboardWidgetConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceAwsArn(resourceName),
					resource.TestCheckResourceAttr(resourceName, "type", testDashboardWidget.Type),
				),
			},
		},
	})
}

func testAccDataSourceAwsArn(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		return nil
	}
}

func testAccDataSourceDashboardWidgetConfig() string {
	return fmt.Sprintf(`
%s
data "appdynamics_dashboard_widget" "test" {
}
`, configureDashboardConfig())
}
