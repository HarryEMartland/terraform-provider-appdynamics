package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strconv"
	"testing"
)

func TestAccAppDDashboard_Create(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	cpuWidget, _ := os.ReadFile("./widgets/cpu_widget_small.json")
	resourceName := "appdynamics_dashboard.test_basic"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: dashboardWidthOneWidgets(name, string(cpuWidget)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckDashboardDoesNotExist(resourceName)),
	})
}

func TestAccAppDDashboard_UpdateDashboard(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	name2 := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	resourceName := "appdynamics_dashboard.test_basic"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicDashboard(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: basicDashboard(name2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckDashboardDoesNotExist(resourceName)),
	})
}

func TestAccAppDDashboard_UpdateWidgets(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	cpuWidget, _ := os.ReadFile("./widgets/cpu_widget_small.json")
	//cpuWidget2, _ := os.ReadFile("./widgets/memory_a.json")
	resourceName := "appdynamics_dashboard.test_basic"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicDashboard(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: dashboardWidthOneWidgets(name, string(cpuWidget)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckDashboardDoesNotExist(resourceName)),
	})
}

func CheckDashboardDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetDashboard(id)
		if err == nil {
			return fmt.Errorf("dashboard found: %d", id)
		}

		return nil
	}
}

func configureDashboardConfig() string {
	return fmt.Sprintf(`
					provider "appdynamics" {
					  secret = "%s"
					  controller_base_url = "%s"
					}

					variable "scope_id" {
					  type = string
					  default = "%s"
					}
					
					variable "application_id" {
					  type = number
					  default = %s
					}`, os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), os.Getenv("APPD_SCOPE_ID"), os.Getenv("APPD_APPLICATION_ID"))
}

func dashboardWidthOneWidgets(name string, cpuWidget string) string {
	return fmt.Sprintf(`
					%s
					data "appdynamics_dashboard_widget" "basic_widget" {
							json = jsonencode(%s)
					}

					resource "appdynamics_dashboard" "test_basic" {
						name = "%s"
  						template_entity_type = "APPLICATION_COMPONENT_NODE"
  						minutes_before_anchor_time = -1
  						refresh_interval = 120000
  						background_color = 15856629
  						height = 768
  						width = 1024
  						canvas_type = "CANVAS_TYPE_GRID"
                        widgets = [data.appdynamics_dashboard_widget.basic_widget.widget_json]
					}
`, configureDashboardConfig(), cpuWidget, name)
}

func dashboardWidthTwoWidgets(name string, cpuWidget string, cpu2Widget string) string {
	return fmt.Sprintf(`
					%s

					resource "appdynamics_dashboard" "test_basic" {
					  name = "%s"
					  template_entity_type = "APPLICATION_COMPONENT_NODE"
					  minutes_before_anchor_time = -1
					  refresh_interval = 120000
					  background_color = 15856629
					  height = 768
					  width = 1024
					  canvas_type = "CANVAS_TYPE_GRID"
					  widgets = [jsonencode(%s), jsonencode(%s)]
					}
`, configureDashboardConfig(), name, cpuWidget, cpu2Widget)
}

func basicDashboard(name string) string {
	return fmt.Sprintf(`
					%s

					resource "appdynamics_dashboard" "test_basic" {
					  name = "%s"
					  template_entity_type = "APPLICATION_COMPONENT_NODE"
					  minutes_before_anchor_time = -1
					  refresh_interval = 120000
					  background_color = 15856629
					  height = 768
					  width = 1024
					  canvas_type = "CANVAS_TYPE_GRID"
					  widgets = []	
					}
`, configureDashboardConfig(), name)
}
