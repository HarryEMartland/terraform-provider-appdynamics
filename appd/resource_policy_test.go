package appd

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
	"testing"
)

func TestAccAppDAction_basicAllHealthRulesPolicy(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_policy.all_health_rules_email_on_call"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allHealthRulesPolicy(name, []string{"HEALTH_RULE_OPEN_CRITICAL"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "action_name", "07421365896"),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "health_rule_event_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_rule_scope_type", "ALL_HEALTH_RULES"),
					RetryCheck(CheckPolicyExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckPolicyDoesNotExist(resourceName)),
	})
}

func TestAccAppDAction_updateAllHealthRulesPolicy(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_policy.all_health_rules_email_on_call"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allHealthRulesPolicy(name, []string{"HEALTH_RULE_OPEN_CRITICAL"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "action_name", "07421365896"),
					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
					resource.TestCheckResourceAttr(resourceName, "health_rule_event_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_rule_scope_type", "ALL_HEALTH_RULES"),
					RetryCheck(CheckPolicyExists(resourceName)),
				),
			},
			{
				Config: allHealthRulesPolicy(name, []string{"HEALTH_RULE_OPEN_CRITICAL", "HEALTH_RULE_OPEN_WARNING"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health_rule_event_types.#", "2"),
					RetryCheck(CheckPolicyExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckPolicyDoesNotExist(resourceName)),
	})
}

func TestAccAppDAction_basicSpecificHealthRulesPolicy(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_policy.all_health_rules_email_on_call"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificHealthRulesPolicy(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "health_rules.#", "1"),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckPolicyDoesNotExist(resourceName)),
	})
}

func TestAccAppDAction_updateSpecificHealthRulesPolicy(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_policy.all_health_rules_email_on_call"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificHealthRulesPolicy(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "health_rules.#", "1"),
				),
			},
			{
				Config: updatedSpecificHealthRulesPolicy(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "health_rules.#", "2"),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckPolicyDoesNotExist(resourceName)),
	})
}

func CheckPolicyDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetPolicy(id, applicationIdI)
		if err == nil {
			return fmt.Errorf("found: %d", id)
		}

		return nil
	}
}

func CheckPolicyExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetPolicy(id, applicationIdI)
		if err != nil {
			return err
		}

		return nil
	}
}

func allHealthRulesPolicy(name string, eventTypes []string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_action" "test_action" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "07421365896"
					}
					resource "appd_policy" "all_health_rules_email_on_call" {
					  name = "%s"
					  application_id = var.application_id
					  action_name = appd_action.test_action.phone_number
					  action_type = appd_action.test_action.action_type
					  health_rule_event_types = %s
					  health_rule_scope_type = "ALL_HEALTH_RULES"
					}
`, configureConfig(), name, arrayToString(eventTypes))
}

func specificHealthRulesPolicy(name string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_health_rule" "test_health_rule" {
					  name = "My Single Metring Health Rule"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition="GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
					resource "appd_action" "test_action" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "07421365896"
					}
					resource "appd_policy" "all_health_rules_email_on_call" {
					  name = "%s"
					  application_id = var.application_id
					  action_name = appd_action.test_action.phone_number
					  action_type = appd_action.test_action.action_type
					  health_rule_event_types = ["HEALTH_RULE_OPEN_CRITICAL"]
					  health_rule_scope_type = "SPECIFIC_HEALTH_RULES"
					  health_rules = [appd_health_rule.test_health_rule.name]
					}
`, configureConfig(), name)
}

func updatedSpecificHealthRulesPolicy(name string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_health_rule" "test_health_rule" {
					  name = "My Single Metring Health Rule"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition="GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
					resource "appd_health_rule" "test_health_rule_2" {
					  name = "My Single Metring Health Rule 2"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition="GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
					resource "appd_action" "test_action" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "07421365896"
					}
					resource "appd_policy" "all_health_rules_email_on_call" {
					  name = "%s"
					  application_id = var.application_id
					  action_name = appd_action.test_action.phone_number
					  action_type = appd_action.test_action.action_type
					  health_rule_event_types = ["HEALTH_RULE_OPEN_CRITICAL"]
					  health_rule_scope_type = "SPECIFIC_HEALTH_RULES"
					  health_rules = [appd_health_rule.test_health_rule.name, appd_health_rule.test_health_rule_2.name]
					}
`, configureConfig(), name)
}
