package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccImportExportDashboard_CreateAndUpdate(t *testing.T) {

	resourceName := "appdynamics_import_export_dashboard.test_basic"
	dashboardTemplate, _ := os.ReadFile("./templates/sampleTemplate.json")
	dashboardTemplate2, _ := os.ReadFile("./templates/sampleTemplate2.json")
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicImportExportDashboardTemplate(string(dashboardTemplate)),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			{
				Config: basicImportExportDashboardTemplate(string(dashboardTemplate2)),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
		CheckDestroy: RetryCheck(CheckDashboardDoesNotExist(resourceName)),
	})
}

func basicImportExportDashboardTemplate(dashboardTemplate string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_import_export_dashboard" "test_basic" {
						json = jsonencode(%s)
					}
`, configureConfig(), dashboardTemplate)
}
