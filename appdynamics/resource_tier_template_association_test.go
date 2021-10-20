package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strconv"
	"testing"
)

func TestAccTierTemplateAssociation_Create(t *testing.T) {
	// Sample true test data
	//applicationId := 2951 // DEV-Client-Facing
	//tierId := 60300 // airtime-purchases
	//templateId := 1762 // HS test

	applicationId := 1
	tierId := 1
	templateId := 1
	resourceName := "appdynamics_tier_template_association.test_basic"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicAssociation(applicationId, tierId, templateId),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			{
				Config: noAssociations(applicationId, tierId),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
		CheckDestroy: RetryCheck(CheckAssociationDoesNotExist(resourceName)),
	})
}

func CheckAssociationDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {
		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		associatedDashboards, err := appDClient.GetAllDashboardTemplatesByTier(id)
		if err == nil {
			return fmt.Errorf("dashboard found: %d", id)
		}
		if len(associatedDashboards) > 0 {
			return fmt.Errorf("association not removed")
		}

		return nil
	}
}

func configureProviderConfig() string {
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

func noAssociations(applicationId int, tierId int) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_tier_template_association" "test_basic" {
					  application_id = %d
					  tier_id = %d
					  template_ids=[]
					}
`, configureProviderConfig(), applicationId, tierId)
}

func basicAssociation(applicationId int, tierId int, templateId int) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_tier_template_association" "test_basic" {
					  application_id = %d
					  tier_id = %d
					  template_ids=[%d]
					}
`, configureProviderConfig(), applicationId, tierId, templateId)
}
