package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccDataSourceDashboardWidget_basic(t *testing.T) {
	resourceName := "data.appdynamics_dashboard_widget.test"
	sampleWidget, _ := os.ReadFile("./widgets/cpu_widget_small.json")

	resource.ParallelTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDashboardWidgetConfig(string(sampleWidget)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "json"),
					resource.TestCheckResourceAttrSet(resourceName, "widget_json"),
				),
			},
		},
	})
}

func testAccDataSourceDashboardWidgetConfig(sampleWidget string) string {
	return fmt.Sprintf(`
%s

data "appdynamics_dashboard_widget" "test" {
 json = jsonencode(%s)
}
`, configureDashboardConfig(), sampleWidget)
}
