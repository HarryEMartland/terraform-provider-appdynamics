package appdynamics

import (
	"errors"
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAccAppDTransactionRule_basicRule(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_transaction_detection_rule.test_rule"
	agentType := "NODE_JS_SERVER"
	description := "Health rule created in automated acceptance tests for terraform"
	enabled := "true"
	entryPointType := "NODEJS_WEB"
	method := "GET"
	matchType := "EQUALS"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: transactionRule(name, agentType, description, enabled, 36, entryPointType, matchType, []string{"/test"}, method),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "agent_type", agentType),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "enabled", enabled),
					resource.TestCheckResourceAttr(resourceName, "entry_point_type", entryPointType),
					resource.TestCheckResourceAttr(resourceName, "http_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_method", method),
					resource.TestCheckResourceAttr(resourceName, "http_uri_match_type", matchType),
					RetryCheck(CheckTransactionRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckTransactionRuleDoesNotExist(resourceName)),
	})
}


func TestAccAppDTransactionRule_multipleUris(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_transaction_detection_rule.test_rule"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: basicTransactionRuleMultiple(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "http_uris.#", "2"),
					RetryCheck(CheckTransactionRule(resourceName, func(rule *client.RuleScope) {
						assert.Equal(t, "/health,/user", rule.Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.MatchStrings[0])
					})),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckTransactionRuleDoesNotExist(resourceName)),
	})
}

func CheckTransactionRule(resourceName string, callback func(scope *client.RuleScope)) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		rule, found, err := appDClient.GetTransactionDetectionRule(applicationIdI, resourceState.Primary.ID)
		if err != nil {
			return err
		}
		if !found {
			return errors.New("transaction rule not found")
		}

		callback(rule)
		return nil
	}
}

func CheckTransactionRuleDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		_, found, err := appDClient.GetTransactionDetectionRule(applicationIdI, resourceState.Primary.ID)
		if err != nil {
			return fmt.Errorf("error finding transaction: %s", resourceState.Primary.ID)
		}
		if found {
			return fmt.Errorf("transaction rule found: %s", resourceState.Primary.ID)
		}

		return nil
	}
}

func CheckTransactionRuleExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		_, found, err := appDClient.GetTransactionDetectionRule(applicationIdI, resourceState.Primary.ID)
		if err != nil {
			return err
		}
		if !found {
			return errors.New("transaction rule not found")
		}

		return nil
	}
}

func transactionRule(name string, agentType string, description string, enabled string, priority int, entryPointType string, uriMatchType string, uris []string, httpMethod string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_transaction_detection_rule" "test_rule" {
					  application_id = var.application_id
					  name = "%s"
					  agent_type = "%s"
					  account_id = "%s"
					  description = "%s"
					  enabled = %s
					  priority = %d
					  entry_point_type = "%s"
					  http_uri_match_type = "%s"
					  http_uris = ["%s"]
					  http_method = "%s"
					}
`, configureConfig(), name, agentType, accountId, description, enabled, priority, entryPointType, uriMatchType, strings.Join(uris, "\",\""), httpMethod)

}

func basicTransactionRuleMultiple(name string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_transaction_detection_rule" "test_rule" {
					  application_id = var.application_id
					  name = "%s"
					  agent_type = "NODE_JS_SERVER"
					  account_id = "%s"
					  description = "Health rule created in automated acceptance tests for terraform"
					  enabled = true
					  priority = 36
					  entry_point_type = "NODEJS_WEB"
					  http_uri_match_type = "EQUALS"
					  http_uris = ["/health", "/user"]
					  http_method = "GET"
					}
`, configureConfig(), name, accountId)
}

//todo finish documentation
//todo scope
