package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strconv"
	"testing"
)

func TestAccAppDCollector_Create(t *testing.T) {

	fmt.Println("TestAccAppDCollector_Create test started")

	resourceName := "appdynamics_collector.test"
	collectorName := "testAutomationCreate"
	tf := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), collectorName)
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: tf,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", collectorName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckCollectorDoesNotExist(resourceName),
	})

}

func TestAccAppDCollector_Update(t *testing.T) {

	fmt.Println("TestAccAppDCollector_Update test started")

	resourceName := "appdynamics_collector.test"
	collectorName := "testAutomationUpdate1"
	updatedName := "testAutomationUpdate2"
	tf := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), collectorName)
	tfUpdated := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), updatedName)
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: tf,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", collectorName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
			{
				Config: tfUpdated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckCollectorDoesNotExist(resourceName),
	})
}

func configureCollectorTest(appdSecret string, appdControllerUrl string, collectorName string) string {
	tf := fmt.Sprintf(`
provider "appdynamics" {
  secret = "%s"
  controller_base_url = "%s"
}

resource appdynamics_collector test {
	name="%s"
	type="MYSQL"
	hostname="test"
	username="user"
	password="paswd"
	port=3306
	agent_name="test"
}`, appdSecret, appdControllerUrl, collectorName)

	return tf
}

func CheckCollectorDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetCollector(id)
		if err == nil {
			return fmt.Errorf("collector found, but should be removed : %d", id)
		}

		return nil
	}
}

func CheckCollectorExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetCollector(id)
		if err != nil {
			return err
		}

		return nil
	}
}
