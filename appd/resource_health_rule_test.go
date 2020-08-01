package appd

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
	"testing"
)

func TestAccAppDHealthRule_basicSingleMetricAllBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_health_rule.test_all_bts"
	aggregationFunction := "VALUE"
	detailType := "SINGLE_METRIC"
	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	metric := "95th Percentile Response Time (ms)"
	condition := "GREATER_THAN_SPECIFIC_VALUE"
	warnValue := "1"
	criticalValue := "2"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(name, aggregationFunction, detailType, entityType, metric, condition, warnValue, criticalValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "metric_aggregation_function", aggregationFunction),
					resource.TestCheckResourceAttr(resourceName, "eval_detail_type", detailType),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", "ALL_BUSINESS_TRANSACTIONS"),
					resource.TestCheckResourceAttr(resourceName, "metric_eval_detail_type", "SPECIFIC_TYPE"), //bug in api?
					resource.TestCheckResourceAttr(resourceName, "metric_path", metric),
					resource.TestCheckResourceAttr(resourceName, "compare_condition", condition),
					resource.TestCheckResourceAttr(resourceName, "warn_compare_value", warnValue),
					resource.TestCheckResourceAttr(resourceName, "critical_compare_value", criticalValue),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_updateSingleMetricAllBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appd_health_rule.test_all_bts"
	aggregationFunction := "VALUE"
	detailType := "SINGLE_METRIC"
	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	metric := "95th Percentile Response Time (ms)"
	condition := "GREATER_THAN_SPECIFIC_VALUE"
	warnValue := "1"
	criticalValue := "2"

	updatedAggregationFunction := "SUM"
	updatedMetric := "95th Percentile Response Time (ms)"
	updatedCondition := "LESS_THAN_SPECIFIC_VALUE"
	updatedWarnValue := "3"
	updatedCriticalValue := "4"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(name, aggregationFunction, detailType, entityType, metric, condition, warnValue, criticalValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "metric_aggregation_function", aggregationFunction),
					resource.TestCheckResourceAttr(resourceName, "eval_detail_type", detailType),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", "ALL_BUSINESS_TRANSACTIONS"),
					resource.TestCheckResourceAttr(resourceName, "metric_eval_detail_type", "SPECIFIC_TYPE"), //bug in api?
					resource.TestCheckResourceAttr(resourceName, "metric_path", metric),
					resource.TestCheckResourceAttr(resourceName, "compare_condition", condition),
					resource.TestCheckResourceAttr(resourceName, "warn_compare_value", warnValue),
					resource.TestCheckResourceAttr(resourceName, "critical_compare_value", criticalValue),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: allBTsHealthRule(name, updatedAggregationFunction, detailType, entityType, updatedMetric, updatedCondition, updatedWarnValue, updatedCriticalValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "metric_aggregation_function", updatedAggregationFunction),
					resource.TestCheckResourceAttr(resourceName, "eval_detail_type", detailType),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", "ALL_BUSINESS_TRANSACTIONS"),
					resource.TestCheckResourceAttr(resourceName, "metric_eval_detail_type", "SPECIFIC_TYPE"), //bug in api?
					resource.TestCheckResourceAttr(resourceName, "metric_path", updatedMetric),
					resource.TestCheckResourceAttr(resourceName, "compare_condition", updatedCondition),
					resource.TestCheckResourceAttr(resourceName, "warn_compare_value", updatedWarnValue),
					resource.TestCheckResourceAttr(resourceName, "critical_compare_value", updatedCriticalValue),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSpecificBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	bts := []string{bt1}

	resourceName := "appd_health_rule.test_specific_bts"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificBTsHealthRule(name, bts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_updateSpecificBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	bts := []string{bt1}
	updatedBts := []string{bt1, bt2}

	resourceName := "appd_health_rule.test_specific_bts"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificBTsHealthRule(name, bts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: specificBTsHealthRule(name, updatedBts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSpecificTiers(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	tiers := []string{tier1}

	resourceName := "appd_health_rule.test_specific_tiers"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificTiersHealthRule(name, tiers),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_updateSpecificTiers(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	tiers := []string{tier1}
	updatedBts := []string{tier1, tier2}

	resourceName := "appd_health_rule.test_specific_tiers"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appd": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificTiersHealthRule(name, tiers),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: specificTiersHealthRule(name, updatedBts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func CheckHealthRuleExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetHealthRule(id, applicationIdI)
		if err != nil {
			return err
		}

		return nil
	}
}

func CheckHealthRuleDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetHealthRule(id, applicationIdI)
		if err == nil {
			return fmt.Errorf("health rule found: %d", id)
		}

		return nil
	}
}

func allBTsHealthRule(name string, aggregationFunction string, detailType string, entityType string, metric string, compareCondition string, warnValue string, criticalValue string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_health_rule" "test_all_bts" {
					  name = "%s"
					  application_id = var.application_id
					  metric_aggregation_function = "%s"
					  eval_detail_type = "%s"
					  affected_entity_type = "%s"
					  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "%s"
					  compare_condition="%s"
					  warn_compare_value = %s
					  critical_compare_value = %s
					}
`, configureConfig(), name, aggregationFunction, detailType, entityType, metric, compareCondition, warnValue, criticalValue)
}

func specificBTsHealthRule(name string, bts []string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_health_rule" "test_specific_bts" {
					  name = "%s"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "SPECIFIC_BUSINESS_TRANSACTIONS"
					  business_transactions = %s
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
`, configureConfig(), name, arrayToString(bts))
}

func specificTiersHealthRule(name string, tiers []string) string {
	return fmt.Sprintf(`
					%s
					resource "appd_health_rule" "test_specific_tiers" {
					  name = "%s"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS"
					  specific_tiers = %s
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
`, configureConfig(), name, arrayToString(tiers))
}
