package appd

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccAppDHealthRule_basicSMS(t *testing.T) {

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
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName, appDClient, applicationIdI),
	})
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
