package appd

import (
	"fmt"
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var appDClient client.AppDClient
var applicationIdI int
var applicationIdS string
var httpActionTemplateName string

func init() {
	_, acceptanceTest := os.LookupEnv("TF_ACC")
	if !acceptanceTest {
		return
	}

	applicationId, err := strconv.Atoi(os.Getenv("APPD_APPLICATION_ID"))
	if err != nil {
		log.Fatal(fmt.Sprintf("error parsing application id: %s", os.Getenv("APPD_APPLICATION_ID")))
	}
	appDClient = client.AppDClient{
		BaseUrl: os.Getenv("APPD_CONTROLLER_BASE_URL"),
		Secret:  os.Getenv("APPD_SECRET"),
	}
	applicationIdI = applicationId
	applicationIdS = os.Getenv("APPD_APPLICATION_ID")
	httpActionTemplateName = os.Getenv("APPD_HTTP_ACTION_TEMPLATE_NAME")
}

func TestAccAppDAction_basicSMS(t *testing.T) {

	phoneNumber := acctest.RandStringFromCharSet(11, "0123456789")

	resourceName := "appd_action.test_sms"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: smsAction(phoneNumber),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestAccAppDAction_updateSMS(t *testing.T) {

	phoneNumber1 := acctest.RandStringFromCharSet(11, "0123456789")
	phoneNumber2 := acctest.RandStringFromCharSet(11, "0123456789")

	resourceName := "appd_action.test_sms"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: smsAction(phoneNumber1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber1),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
			{
				Config: smsAction(phoneNumber2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber2),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestAccAppDAction_basicEmail(t *testing.T) {

	email := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	emails := []string{email}

	resourceName := "appd_action.test_email"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
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
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestAccAppDAction_updateEmail(t *testing.T) {

	email := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	email2 := fmt.Sprintf("%s@example.com", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	emails := []string{email}
	emailsUpdated := []string{email, email2}

	resourceName := "appd_action.test_email"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
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
					CheckActionExists(resourceName, appDClient, applicationIdI),
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
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestAccAppDAction_basicHttp(t *testing.T) {

	name := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	templateName := os.Getenv("APPD_HTTP_ACTION_TEMPLATE_NAME")

	resourceName := "appd_action.test_http_action"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: httpAction(name, templateName, map[string]string{"config": "cValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", templateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestAccAppDAction_updateHttp(t *testing.T) {

	name := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	updateName := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	templateName := os.Getenv("APPD_HTTP_ACTION_TEMPLATE_NAME")

	resourceName := "appd_action.test_http_action"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: httpAction(name, templateName, map[string]string{"config": "cValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", templateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
			{
				Config: httpAction(updateName, templateName, map[string]string{"config": "cValue", "second": "sValue"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "http_request_template_name", templateName),
					resource.TestCheckResourceAttr(resourceName, "action_type", "HTTP_REQUEST"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.config", "cValue"),
					resource.TestCheckResourceAttr(resourceName, "custom_template_variables.second", "sValue"),
					CheckActionExists(resourceName, appDClient, applicationIdI),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
}

func TestMapToStringSingle(t *testing.T) {
	assert.Equal(t, "{k1: \"v1\",}", mapToString(map[string]string{"k1": "v1"}), "map should be correctly formatted")
}

func TestMapToStringMultiple(t *testing.T) {
	assert.Equal(t, "{k1: \"v1\",k2: \"v2\",}", mapToString(map[string]string{"k1": "v1", "k2": "v2"}), "map should be correctly formatted")
}

func hash(s string) int {
	return schema.HashString(s)
}

func CheckActionDoesNotExist(resourceName string, appDClient client.AppDClient, applicationId int) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetAction(id, applicationId)
		if err == nil {
			return fmt.Errorf("action found: %d", id)
		}

		return nil
	}
}

func CheckActionExists(resourceName string, appDClient client.AppDClient, applicationId int) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetAction(id, applicationId)
		if err != nil {
			return err
		}

		return nil
	}
}

func smsAction(phoneNumber string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_action" "test_sms" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "%s"
					}
`, configureConfig(), phoneNumber)
}

func emailAction(emails []string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_action" "test_email" {
					  application_id = var.application_id
					  action_type = "EMAIL"
					  emails = ["%s"]
					}
`, configureConfig(), strings.Join(emails, "\",\""))
}

func mapToString(m map[string]string) string {
	result := "{"

	for key, value := range m {
		result += fmt.Sprintf("%s: \"%s\",", key, value)
	}

	result += "}"
	return result
}

func httpAction(name string, templateName string, variableMap map[string]string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_action" "test_http_action" {
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
					provider "appd" {
					  secret = "%s"
					  controller_base_url = "%s"
					}
					
					variable "application_id" {
					  type = number
					  default = %s
					}`, os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), os.Getenv("APPD_APPLICATION_ID"))
}
