package appd

import (
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceHealthRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthRuleCreate,
		Read:   resourceHealthRuleRead,
		Update: resourceHealthRuleUpdate,
		Delete: resourceHealthRuleDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"evaluation_minutes": {
				Type:     schema.TypeInt,
				Default:  30,
				Optional: true,
			},
			"violation_length_minutes": {
				Type:     schema.TypeInt,
				Default:  5,
				Optional: true,
			},
			"affected_entity_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"OVERALL_APPLICATION_PERFORMANCE",
					"BUSINESS_TRANSACTION_PERFORMANCE",
					"TIER_NODE_TRANSACTION_PERFORMANCE",
					"TIER_NODE_HARDWARE",
					"SERVERS_IN_APPLICATION",
					"BACKENDS",
					"ERRORS",
					"SERVICE_ENDPOINTS",
					"INFORMATION_POINTS",
					"CUSTOM",
					"DATABASES",
					"SERVERS",
				}),
			},
			"business_transaction_scope": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"ALL_BUSINESS_TRANSACTIONS",
					"SPECIFIC_BUSINESS_TRANSACTIONS",
					"BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS",
					"BUSINESS_TRANSACTIONS_MATCHING_PATTERN",
				}),
			},
			"evaluate_to_true_on_no_data": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"warn_compare_value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"critical_compare_value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"eval_detail_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_aggregation_function": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_eval_detail_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"SINGLE_METRIC",
					"METRIC_EXPRESSION",
					"BASELINE_TYPE",
					"SPECIFIC_TYPE",
				}),
			},
			"baseline_condition": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"WITHIN_BASELINE",
					"NOT_WITHIN_BASELINE",
					"GREATER_THAN_BASELINE",
					"LESS_THAN_BASELINE",
				}),
			},
			"baseline_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"baseline_unit": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"STANDARD_DEVIATIONS",
					"PERCENTAGE",
				}),
			},
			"compare_condition": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"GREATER_THAN_SPECIFIC_VALUE",
					"LESS_THAN_SPECIFIC_VALUE",
				}),
			},
			"business_transactions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceHealthRuleCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	healthRule := createHealthRule(d)

	updatedHealthRule, err := appdClient.CreateHealthRule(&healthRule, applicationId)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(updatedHealthRule.ID))

	return resourceHealthRuleRead(d, m)
}

func GetOrNilS(d *schema.ResourceData, key string) *string {
	value, set := d.GetOk(key)
	if set {
		strValue := value.(string)
		return &strValue
	}
	return nil
}

func GetOrNilL(d *schema.ResourceData, key string) *[]interface{} {
	value, set := d.GetOk(key)
	if set {
		listValue := value.(*schema.Set).List()
		return &listValue
	}
	return nil
}

func createHealthRule(d *schema.ResourceData) client.HealthRule {

	name := d.Get("name").(string)
	evaluationMinutes := d.Get("evaluation_minutes").(int)
	violationLengthMinutes := d.Get("violation_length_minutes").(int)

	affectedEntityType := d.Get("affected_entity_type").(string)
	businessTransactionScope := d.Get("business_transaction_scope").(string)

	evaluateToTrueOnNoData := d.Get("evaluate_to_true_on_no_data").(bool)

	evalDetailType := d.Get("eval_detail_type").(string)
	metricAggregationFunction := d.Get("metric_aggregation_function").(string)
	metricPath := d.Get("metric_path").(string)
	criticalCompareValue := d.Get("critical_compare_value").(float64)
	compareCondition := GetOrNilS(d, "compare_condition")
	warnCompareValue := d.Get("warn_compare_value").(float64)
	metricEvalDetailType := d.Get("metric_eval_detail_type").(string)

	baselineCondition := GetOrNilS(d, "baseline_condition")
	baselineName := GetOrNilS(d, "baseline_name")
	baselineUnit := GetOrNilS(d, "baseline_unit")

	healthRule := client.HealthRule{
		Name:                    name,
		Enabled:                 true,
		UseDataFromLastNMinutes: evaluationMinutes,
		WaitTimeAfterViolation:  violationLengthMinutes,
		Affects: &client.Affects{
			AffectedEntityType: affectedEntityType,
			AffectedBusinessTransactions: &client.Transaction{
				BusinessTransactionScope: businessTransactionScope,
				BusinessTransactions:     GetOrNilL(d, "business_transactions"),
			},
		},
		Criterias: &client.Criterias{
			Critical: &client.Criteria{
				ConditionAggregationType: "ALL",
				Conditions: []*client.Condition{{
					Name:                   name,
					ShortName:              "A",
					EvaluateToTrueOnNoData: evaluateToTrueOnNoData,
					EvalDetail: &client.EvalDetail{
						EvalDetailType:          evalDetailType,
						MetricAggregateFunction: metricAggregationFunction,
						MetricPath:              metricPath,
						MetricEvalDetail: &client.MetricEvalDetail{
							MetricEvalDetailType: metricEvalDetailType,
							BaselineCondition:    baselineCondition,
							BaselineName:         baselineName,
							BaselineUnit:         baselineUnit,
							CompareValue:         criticalCompareValue,
							CompareCondition:     compareCondition,
						},
					},
				}},
			},
			Warning: &client.Criteria{
				ConditionAggregationType: "ALL",
				Conditions: []*client.Condition{{
					Name:                   name,
					ShortName:              "A",
					EvaluateToTrueOnNoData: evaluateToTrueOnNoData,
					EvalDetail: &client.EvalDetail{
						EvalDetailType:          evalDetailType,
						MetricAggregateFunction: metricAggregationFunction,
						MetricPath:              metricPath,
						MetricEvalDetail: &client.MetricEvalDetail{
							MetricEvalDetailType: metricEvalDetailType,
							BaselineCondition:    baselineCondition,
							BaselineName:         baselineName,
							BaselineUnit:         baselineUnit,
							CompareValue:         warnCompareValue,
							CompareCondition:     compareCondition,
						},
					},
				}},
			},
		},
	}
	return healthRule
}

func resourceHealthRuleRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	healthRuleId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	healthRule, err := appdClient.GetHealthRule(healthRuleId, applicationId) //read back into d
	if err != nil {
		return err
	}

	updateHealthRule(d, healthRule)

	return nil
}

func updateHealthRule(d *schema.ResourceData, healthRule *client.HealthRule) {
	d.Set("name", healthRule.Name)
	d.Set("evaluation_minutes", healthRule.UseDataFromLastNMinutes)
	d.Set("violation_length_minutes", healthRule.WaitTimeAfterViolation)
	d.Set("affected_entity_type", healthRule.Affects.AffectedEntityType)
	d.Set("business_transaction_scope", healthRule.Affects.AffectedBusinessTransactions.BusinessTransactionScope)
	criticalCondition := healthRule.Criterias.Critical.Conditions[0]
	d.Set("evaluate_to_true_on_no_data", criticalCondition.EvaluateToTrueOnNoData)
	d.Set("eval_detail_type", criticalCondition.EvalDetail.EvalDetailType)
	d.Set("metric_aggregation_function", criticalCondition.EvalDetail.MetricAggregateFunction)
	d.Set("metric_path", criticalCondition.EvalDetail.MetricPath)
	d.Set("critical_compare_value", criticalCondition.EvalDetail.MetricEvalDetail.CompareValue)
	d.Set("warn_compare_value", healthRule.Criterias.Warning.Conditions[0].EvalDetail.MetricEvalDetail.CompareValue)
	d.Set("metric_eval_detail_type", criticalCondition.EvalDetail.EvalDetailType)
	d.Set("baseline_condition", criticalCondition.EvalDetail.MetricEvalDetail.BaselineCondition)
	d.Set("baseline_name", criticalCondition.EvalDetail.MetricEvalDetail.BaselineName)
	d.Set("baseline_unit", criticalCondition.EvalDetail.MetricEvalDetail.BaselineUnit)
}

func resourceHealthRuleUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	healthRule := createHealthRule(d)

	healthRuleId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	healthRule.ID = healthRuleId

	_, err = appdClient.UpdateHealthRule(&healthRule, applicationId)
	if err != nil {
		return err
	}

	return resourceHealthRuleRead(d, m)
}

func resourceHealthRuleDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	healthRuleId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = appdClient.DeleteHealthRule(applicationId, healthRuleId)
	if err != nil {
		return err
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
