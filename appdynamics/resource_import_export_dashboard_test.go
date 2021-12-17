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
	dashboardTemplate, _ := os.ReadFile("./templates/sampleTemplate3.json")
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

func configureDashboardConfig2() string {
	return fmt.Sprintf(`
					provider "appdynamics" {
					  secret = "%s"
					  controller_base_url = "%s"
					  dashboard_client_name = "%s"
					  dashboard_client_password = "%s"
					}

					variable "scope_id" {
					  type = string
					  default = "%s"
					}
					
					variable "application_id" {
					  type = number
					  default = %s
					}`, os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), os.Getenv("APPD_DASHBOARD_USER"), os.Getenv("APPD_DASHBOARD_PASSWORD"), os.Getenv("APPD_SCOPE_ID"), os.Getenv("APPD_APPLICATION_ID"))
}

func basicImportExportDashboardTemplate(dashboardTemplate string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_import_export_dashboard" "test_basic" {
						json = jsonencode(%s)
					}
`, configureDashboardConfig2(), dashboardTemplate)
}
