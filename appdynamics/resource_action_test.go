package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAccAppDAction_basicSMS(t *testing.T) {

	phoneNumber := acctest.RandStringFromCharSet(11, "0123456789")

	resourceName := "appdynamics_action.test_sms"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: smsAction(phoneNumber),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func TestAccAppDAction_updateSMS(t *testing.T) {

	phoneNumber1 := acctest.RandStringFromCharSet(11, "0123456789")
	phoneNumber2 := acctest.RandStringFromCharSet(11, "0123456789")

	resourceName := "appdynamics_action.test_sms"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: smsAction(phoneNumber1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber1),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
			{
				Config: smsAction(phoneNumber2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber2),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func TestAccAppDAction_basicEmail(t *testing.T) {

	email := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	emails := []string{email}

	resourceName := "appdynamics_action.test_email"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: emailAction(emails),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "emails.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("emails.%d", hash(email)), email),
					resource.TestCheckResourceAttr(resourceName, "action_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func TestAccAppDAction_updateEmail(t *testing.T) {

	email := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	email2 := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	emails := []string{email}
	emailsUpdated := []string{email, email2}

	resourceName := "appdynamics_action.test_email"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: emailAction(emails),
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("emails.%d", hash(email)), email),
					resource.TestCheckResourceAttr(resourceName, "emails.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
			{
				Config: emailAction(emailsUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("emails.%d", hash(email)), email),
					//resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("emails.%d", hash(email2)), email2),
					resource.TestCheckResourceAttr(resourceName, "emails.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "action_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func TestAccAppDAction_basicHttp(t *testing.T) {

	name := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_action.test_http_action"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: httpAction(name, httpActionTemplateName, map[string]string{"config": "cValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", httpActionTemplateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func TestAccAppDAction_updateHttp(t *testing.T) {

	name := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	updateName := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_action.test_http_action"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: httpAction(name, httpActionTemplateName, map[string]string{"config": "cValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", httpActionTemplateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					CheckActionExists(resourceName),
				),
			},
			{
				Config: httpAction(updateName, httpActionTemplateName, map[string]string{"config": "cValue", "second": "sValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", httpActionTemplateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.second", "sValue"),
					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
	})
}

func CheckActionDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetAction(id, applicationIdI)
		if err == nil {
			return fmt.Errorf("action found: %d", id)
		}

		return nil
	}
}

func CheckActionExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetAction(id, applicationIdI)
		if err != nil {
			return err
		}

		return nil
	}
}

func smsAction(phoneNumber string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_action" "test_sms" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "%s"
					}
`, configureConfig(), phoneNumber)
}

func emailAction(emails []string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_action" "test_email" {
					  application_id = var.application_id
					  action_type = "EMAIL"
					  emails = ["%s"]
					}
`, configureConfig(), strings.Join(emails, "\",\""))
}

func httpAction(name string, templateName string, variableMap map[string]string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_action" "test_http_action" {
					  application_id = var.application_id
					  name = "%s"
					  action_type = "HTTP_REQUEST"
					  http_request_template_name = "%s"
					  custom_template_variables = %s
					}
`, configureConfig(), name, templateName, mapToString(variableMap))
}

func configureConfig() string {
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
