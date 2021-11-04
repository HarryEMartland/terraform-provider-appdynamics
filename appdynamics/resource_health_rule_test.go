package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
	"strings"
	"testing"
	//"github.com/k0kubun/pp"
)

func TestAccAppDHealthRule_basicSingleMetricAllBtsMultipleCrit(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_health_rule.test_all_bts_multiple_criteria"

	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	businessTransactionScope := "ALL_BUSINESS_TRANSACTIONS"

	criticalConditionAggregationType := "ANY"
/*
	var criticalCriteria []map[string]interface{}
	criticalCriteria = nil
*/
	criticalCriteria  := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 2.4,
		},
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "Average CPU Used (ms)",
			"metric_eval_detail_type": "BASELINE_TYPE",
			"baseline_name": "All data - Last 15 days",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
			"compare_value": 7.5,
		},
	}

	warningConditionAggregationType := "ALL"
	warningCriteria := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 6.6,
		},
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "Average CPU Used (ms)",
			"metric_eval_detail_type": "BASELINE_TYPE",
			"baseline_name": "All data - Last 15 days",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
			"compare_value": 7.5,
		},
	}

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(strings.Split(resourceName, ".")[1], name, entityType, businessTransactionScope, criticalConditionAggregationType, criticalCriteria, warningConditionAggregationType, warningCriteria),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", businessTransactionScope),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.1.baseline_condition", "WITHIN_BASELINE"),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.compare_value", "2.4"),
					resource.TestCheckResourceAttr(resourceName, "warning_criteria.0.compare_condition", "GREATER_THAN_SPECIFIC_VALUE"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSingleMetricAllBtsSingleCrit(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_health_rule.test_all_bts_single_criteria"

	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	businessTransactionScope := "ALL_BUSINESS_TRANSACTIONS"

	criticalConditionAggregationType := "ANY"
	criticalCriteria := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 1.8,
		},
	}

	warningConditionAggregationType := "ALL"
	var warningCriteria []map[string]interface{}
	warningCriteria = nil

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(strings.Split(resourceName, ".")[1], name, entityType, businessTransactionScope, criticalConditionAggregationType, criticalCriteria, warningConditionAggregationType, warningCriteria),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", businessTransactionScope),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.metric_eval_detail_type", "SPECIFIC_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.compare_value", "1.8"),
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

func prepareHealthRuleCondition(level string, crit map[string]interface{}) string {

	criteriaBlock := fmt.Sprintf(`
            %s_criteria {
                name = "%s"
                shortname = "%s"
                evaluate_to_true_on_no_data = %s
                eval_detail_type = "%s"
                metric_aggregation_function = "%s"
                metric_path = "%s"
                metric_eval_detail_type = "%s"`,
		level,
		crit["name"].(string),
		crit["shortname"].(string),
		strconv.FormatBool(crit["evaluate_to_true_on_no_data"].(bool)),
		crit["eval_detail_type"].(string),
		crit["metric_aggregation_function"].(string),
		crit["metric_path"].(string),
		crit["metric_eval_detail_type"].(string))

	if crit["baseline_name"] != nil {
		criteriaBlock += fmt.Sprintf(`
				baseline_condition = "%s"
                baseline_name = "%s"
                baseline_unit = "%s"`,
			crit["baseline_condition"].(string),
			crit["baseline_name"].(string),
			crit["baseline_unit"].(string))
	}

	if crit["compare_condition"] != nil {
		criteriaBlock += fmt.Sprintf(`
			compare_condition = "%s"`, crit["compare_condition"].(string))
	}

	criteriaBlock += fmt.Sprintf(`
		compare_value = %f
            }`, crit["compare_value"].(float64))

	return criteriaBlock
}

func allBTsHealthRule(resourceName string, name string, entityType string, businessTransactionScope string, criticalConditionAggregationType string, criticalCriteria []map[string]interface{}, warningConditionAggregationType string, warningCriteria []map[string]interface{}) string {

	var criticalCriteriaData []string
	var warningCriteriaData []string

	for _, crit := range criticalCriteria {
		criticalCriteriaData = append(criticalCriteriaData, prepareHealthRuleCondition("critical", crit))
	}

	for _, crit := range warningCriteria {
		warningCriteriaData = append(warningCriteriaData, prepareHealthRuleCondition("warning", crit))
	}

	criteriaData := fmt.Sprintf(`
%s

resource "appdynamics_health_rule" "%s" {
	name = "%s"
	application_id = var.application_id
	affected_entity_type = "%s"
	business_transaction_scope = "%s"
	critical_condition_aggregation_type = "%s"
	warning_condition_aggregation_type = "%s"
	%s
	%s
}`, configureConfig(), resourceName, name, entityType, businessTransactionScope, criticalConditionAggregationType, warningConditionAggregationType, strings.Join(criticalCriteriaData,"\n"), strings.Join(warningCriteriaData, "\n"))

	return criteriaData
}