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

func TestAccAppDDashboard_basic(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_dashboard.test_basic"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: dashboardBasic("TEST_RAFAL_LIZON"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					//RetryCheck(CheckDashboardDoesNotExist(resourceName)),
				),
			},
		},
		//CheckDestroy: RetryCheck(CheckDashboardDoesNotExist(resourceName)),
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

//func sampleDashboard() string {
//	return `{"name":"abcde",
//             "width":1024,
//             "height":768,
//             "canvasType":
//             "CANVAS_TYPE_GRID",
//             "templateEntityType":"APPLICATION_COMPONENT_NODE",
//             "refreshInterval":120000,
//             "backgroundColor":15856629,
//             "warRoom":false,
//             "template":false,
//             "widgets":[],
//             "version":0,
//             "minutesBeforeAnchorTime":-1,
//             "startTime":-1,
//             "endTime":-1
//            }`
//}

//func importTemplate() string {
//	return `{
//		"schemaVersion": null,
//		"dashboardFormatVersion": "4.0",
//		"name": "TEMPLATE__NAME",
//		"description": null,
//			"properties": null,
//			"templateEntityType": "APPLICATION_COMPONENT_NODE",
//			"associatedEntityTemplates": null,
//			"minutesBeforeAnchorTime": -1,
//			"startDate": null,
//			"endDate": null,
//			"refreshInterval": 120000,
//			"backgroundColor": 15856629,
//			"color" : 15856629,
//			"height": 768,
//			"width": 1024,
//			"canvasType": "CANVAS_TYPE_GRID",
//			"layoutType": "",
//			"widgetTemplates" : null,
//			"warRoom": false,
//			"template": false
//	}`
//}

func dashboardBasic(name string) string {
	return fmt.Sprintf(`
					%s
					data "appdynamics_dashboard_widget" "basic_widget" {
					}

					data "appdynamics_dashboard_widget" "basic_widget2" {
							x = 4
                            y = 4
                            width = 12
					}

					resource "appdynamics_dashboard" "test_basic" {
		  				name = "%s"
                        widgets = [
  									data.appdynamics_dashboard_widget.basic_widget.json,
   									#data.appdynamics_dashboard_widget.basic_widget2.json
						]
					}
`, configureDashboardConfig(), name)
}
