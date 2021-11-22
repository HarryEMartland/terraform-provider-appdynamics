package appdynamics

import (
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
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
			"schedule_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Always",
			},
			"use_data_from_last_n_minutes": {
				Type:     schema.TypeInt,
				Default:  30,
				Optional: true,
			},
			"wait_time_after_violation": {
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
			"tier_or_node": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"TIER_AFFECTED_ENTITIES",
					"NODE_AFFECTED_ENTITIES",
				}),
			},
			"type_of_node": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"ALL_NODES",
					"JAVA_NODES",
					"DOT_NET_NODES",
					"PHP_NODES",
				}),
			},
			"affected_tier_scope": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"ALL_TIERS",
					"SPECIFIC_TIERS",
				}),
			},
			"tiers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"affected_node_scope": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"ALL_NODES",
					"SPECIFIC_NODES",
					"NODES_OF_SPECIFIC_TIERS",
					"NODES_MATCHING_PATTERN",
				}),
			},
			"nodes_specific_tiers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nodes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nodes_match": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"STARTS_WITH",
					"ENDS_WITH",
					"CONTAINS",
					"EQUALS",
					"MATCH_REG_EX",
				}),
			},
			"nodes_match_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nodes_match_negation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"business_transaction_scope": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"ALL_BUSINESS_TRANSACTIONS",
					"SPECIFIC_BUSINESS_TRANSACTIONS",
					"BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS",
					"BUSINESS_TRANSACTIONS_MATCHING_PATTERN",
				}),
			},
			"business_transaction_match": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validateList([]string{
					"STARTS_WITH",
					"ENDS_WITH",
					"CONTAINS",
					"EQUALS",
					"MATCH_REG_EX",
				}),
			},
			"business_transaction_match_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"business_transaction_match_negation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"business_transactions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"business_transaction_specific_tiers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"warning_condition_aggregation_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALL",
				ValidateFunc: validateList([]string{
					"ALL",
					"ANY",
				}),
			},
			"warning_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"shortname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"evaluate_to_true_on_no_data": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"eval_detail_type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validateList([]string{
								"SINGLE_METRIC",
								"METRIC_EXPRESSION",
							}),
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
								"BASELINE_TYPE",
								"SPECIFIC_TYPE",
							}),
						},
						"baseline_name": {
							Type:     schema.TypeString,
							Optional: true,
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
						"compare_value": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
			},
			"critical_condition_aggregation_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ALL",
				ValidateFunc: validateList([]string{
					"ALL",
					"ANY",
				}),
			},
			"critical_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"shortname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"evaluate_to_true_on_no_data": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"eval_detail_type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validateList([]string{
								"SINGLE_METRIC",
								"METRIC_EXPRESSION",
							}),
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
								"BASELINE_TYPE",
								"SPECIFIC_TYPE",
							}),
						},
						"baseline_name": {
							Type:     schema.TypeString,
							Optional: true,
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
						"compare_value": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
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

func GetOrNilB(d *schema.ResourceData, key string) *bool {
	value, set := d.GetOkExists(key)
	if set {
		boolValue := value.(bool)
		return &boolValue
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

func defineMetricEvalDetail(metricEvalDetailType *string, baselineCondition *string, baselineName *string, baselineUnit *string, compareValue *float64, compareCondition *string) client.MetricEvalDetail {

	aa := client.MetricEvalDetail{
		MetricEvalDetailType: metricEvalDetailType,
		BaselineCondition:    baselineCondition,
		BaselineName:         baselineName,
		BaselineUnit:         baselineUnit,
		CompareValue:         compareValue,
		CompareCondition:     compareCondition,
	}

	return aa
}

func defineEvalDetail(evalDetailType string, metricAggregationFunction string, metricPath string, metricEvalDetail *client.MetricEvalDetail) client.EvalDetail {
	return client.EvalDetail{
		EvalDetailType:          evalDetailType,
		MetricAggregateFunction: metricAggregationFunction,
		MetricPath:              metricPath,
		MetricEvalDetail:        metricEvalDetail,
	}
}

func defineCondition(name string, shortName string, evaluateToTrueOnNoData bool, evalDetail *client.EvalDetail) client.Condition {
	return client.Condition{
		Name:                   name,
		ShortName:              shortName,
		EvaluateToTrueOnNoData: evaluateToTrueOnNoData,
		EvalDetail:             evalDetail,
	}
}

func conditionsOrNil(conditions []*client.Condition, conditionAggregationType string) *client.Criteria {
	if len(conditions) == 0 {
		return nil
	}

	tmp := client.Criteria{
		ConditionAggregationType: conditionAggregationType,
		Conditions:               conditions,
	}

	return &tmp
}

func createHealthRule(d *schema.ResourceData) client.HealthRule {

	// Criterias
	criticalCriteria := d.Get("critical_criteria").([]interface{})
	warningCriteria := d.Get("warning_criteria").([]interface{})
	criticalConditionAggregationType := d.Get("critical_condition_aggregation_type").(string)
	warningConditionAggregationType := d.Get("warning_condition_aggregation_type").(string)

	var criticalConditions []*client.Condition
	var warningConditions []*client.Condition

	for i, tmpCriteria := range criticalCriteria {

		criteria := tmpCriteria.(map[string]interface{})

		critName := criteria["name"].(string)
		shortname := criteria["shortname"].(string)
		evaluateToTrueOnNoData := criteria["evaluate_to_true_on_no_data"].(bool)
		evalDetailType := criteria["eval_detail_type"].(string)
		metricAggregationFunction := criteria["metric_aggregation_function"].(string)
		metricPath := criteria["metric_path"].(string)
		metricEvalDetailType := criteria["metric_eval_detail_type"].(string)
		baselineCondition := GetOrNilS(d, "critical_criteria."+fmt.Sprint(i)+".baseline_condition")
		baselineName := GetOrNilS(d, "critical_criteria."+fmt.Sprint(i)+".baseline_name")
		baselineUnit := GetOrNilS(d, "critical_criteria."+fmt.Sprint(i)+".baseline_unit")
		compareCondition := GetOrNilS(d, "critical_criteria."+fmt.Sprint(i)+".compare_condition")
		compareValue := criteria["compare_value"].(float64)

		metricEvalDetail := defineMetricEvalDetail(&metricEvalDetailType, baselineCondition, baselineName, baselineUnit, &compareValue, compareCondition)
		evalDetail := defineEvalDetail(evalDetailType, metricAggregationFunction, metricPath, &metricEvalDetail)
		condition := defineCondition(critName, shortname, evaluateToTrueOnNoData, &evalDetail)

		criticalConditions = append(criticalConditions, &condition)
	}

	for i, tmpCriteria := range warningCriteria {

		criteria := tmpCriteria.(map[string]interface{})

		critName := criteria["name"].(string)
		shortname := criteria["shortname"].(string)
		evaluateToTrueOnNoData := criteria["evaluate_to_true_on_no_data"].(bool)
		evalDetailType := criteria["eval_detail_type"].(string)
		metricAggregationFunction := criteria["metric_aggregation_function"].(string)
		metricPath := criteria["metric_path"].(string)
		metricEvalDetailType := criteria["metric_eval_detail_type"].(string)
		baselineCondition := GetOrNilS(d, "warning_criteria."+fmt.Sprint(i)+".baseline_condition")
		baselineName := GetOrNilS(d, "warning_criteria."+fmt.Sprint(i)+".baseline_name")
		baselineUnit := GetOrNilS(d, "warning_criteria."+fmt.Sprint(i)+".baseline_unit")
		compareCondition := GetOrNilS(d, "warning_criteria."+fmt.Sprint(i)+".compare_condition")
		compareValue := criteria["compare_value"].(float64)

		metricEvalDetail := defineMetricEvalDetail(&metricEvalDetailType, baselineCondition, baselineName, baselineUnit, &compareValue, compareCondition)
		evalDetail := defineEvalDetail(evalDetailType, metricAggregationFunction, metricPath, &metricEvalDetail)
		condition := defineCondition(critName, shortname, evaluateToTrueOnNoData, &evalDetail)

		warningConditions = append(warningConditions, &condition)
	}

	criterias := client.Criterias{
		Critical: conditionsOrNil(criticalConditions, criticalConditionAggregationType),
		Warning:  conditionsOrNil(warningConditions, warningConditionAggregationType),
	}

	healthRule := client.HealthRule{
		Name:                    d.Get("name").(string),
		Enabled:                 true,
		ScheduleName:            d.Get("schedule_name").(string),
		UseDataFromLastNMinutes: d.Get("use_data_from_last_n_minutes").(int),
		WaitTimeAfterViolation:  d.Get("wait_time_after_violation").(int),
		Affects: &client.Affects{
			AffectedEntityType: d.Get("affected_entity_type").(string),
			AffectedEntities: &client.AffectedEntities{
				TierOrNode: GetOrNilS(d, "tier_or_node"),
				TypeOfNode: GetOrNilS(d, "type_of_node"),
				AffectedTiers: &client.AffectedTiers{
					AffectedTierScope: GetOrNilS(d, "affected_tier_scope"),
					Tiers:             GetOrNilL(d, "tiers"),
				},
				AffectedNodes: &client.AffectedNodes{
					AffectedNodeScope: GetOrNilS(d, "affected_node_scope"),
					SpecificTiers:     GetOrNilL(d, "nodes_specific_tiers"),
					Nodes:             GetOrNilL(d, "nodes"),
					PatternMatcher: &client.PatternMatcher{
						MatchTo:    GetOrNilS(d, "nodes_match"),
						MatchValue: GetOrNilS(d, "nodes_match_value"),
						ShouldNot:  GetOrNilB(d, "nodes_match_negation"),
					},
				},
			},
			AffectedBusinessTransactions: &client.AffectedBusinessTransactions{
				BusinessTransactionScope: GetOrNilS(d, "business_transaction_scope"),
				BusinessTransactions:     GetOrNilL(d, "business_transactions"),
				SpecificTiers:            GetOrNilL(d, "business_transaction_specific_tiers"),
				PatternMatcher: &client.PatternMatcher{
					MatchTo:    GetOrNilS(d, "business_transaction_match"),
					MatchValue: GetOrNilS(d, "business_transaction_match_value"),
					ShouldNot:  GetOrNilB(d, "business_transaction_match_negation"),
				},
			},
		},
		Criterias: &criterias,
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

func mapConditionToSchema(condition *client.Condition) map[string]interface{} {
	return map[string]interface{}{
		"name":                        condition.Name,
		"shortname":                   condition.ShortName,
		"evaluate_to_true_on_no_data": condition.EvaluateToTrueOnNoData,
		"eval_detail_type":            condition.EvalDetail.EvalDetailType,
		"metric_aggregation_function": condition.EvalDetail.MetricAggregateFunction,
		"metric_path":                 condition.EvalDetail.MetricPath,
		"metric_eval_detail_type":     condition.EvalDetail.MetricEvalDetail.MetricEvalDetailType,
		"baseline_name":               condition.EvalDetail.MetricEvalDetail.BaselineName,
		"baseline_condition":          condition.EvalDetail.MetricEvalDetail.BaselineCondition,
		"baseline_unit":               condition.EvalDetail.MetricEvalDetail.BaselineUnit,
		"compare_condition":           condition.EvalDetail.MetricEvalDetail.CompareCondition,
		"compare_value":               condition.EvalDetail.MetricEvalDetail.CompareValue,
	}
}

func updateHealthRule(d *schema.ResourceData, healthRule *client.HealthRule) {
	//pp.Println(healthRule)

	d.Set("name", healthRule.Name)
	d.Set("schedule_name", healthRule.ScheduleName)
	d.Set("use_data_from_last_n_minutes", healthRule.UseDataFromLastNMinutes)
	d.Set("wait_time_after_violation", healthRule.WaitTimeAfterViolation)

	d.Set("affected_entity_type", healthRule.Affects.AffectedEntityType)

	// AffectedEntities
	if healthRule.Affects.AffectedEntities != nil {
		d.Set("tier_or_node", healthRule.Affects.AffectedEntities.TierOrNode)
		d.Set("type_of_node", healthRule.Affects.AffectedEntities.TypeOfNode)
		if healthRule.Affects.AffectedEntities.AffectedTiers != nil {
			d.Set("affected_tier_scope", healthRule.Affects.AffectedEntities.AffectedTiers.AffectedTierScope)
			d.Set("tiers", healthRule.Affects.AffectedEntities.AffectedTiers.Tiers)
		}
		if healthRule.Affects.AffectedEntities.AffectedNodes != nil {
			d.Set("affected_node_scope", healthRule.Affects.AffectedEntities.AffectedNodes.AffectedNodeScope)
			d.Set("nodes_specific_tiers", healthRule.Affects.AffectedEntities.AffectedNodes.SpecificTiers)
			d.Set("nodes", healthRule.Affects.AffectedEntities.AffectedNodes.Nodes)
			if healthRule.Affects.AffectedEntities.AffectedNodes.PatternMatcher != nil {
				d.Set("nodes_match", healthRule.Affects.AffectedEntities.AffectedNodes.PatternMatcher.MatchTo)
				d.Set("nodes_match_value", healthRule.Affects.AffectedEntities.AffectedNodes.PatternMatcher.MatchValue)
				d.Set("nodes_match_negation", healthRule.Affects.AffectedEntities.AffectedNodes.PatternMatcher.ShouldNot)
			}
		}
	}
	// AffectedBusinessTransactions
	if healthRule.Affects.AffectedBusinessTransactions != nil {
		d.Set("business_transaction_scope", healthRule.Affects.AffectedBusinessTransactions.BusinessTransactionScope)
		d.Set("business_transactions", healthRule.Affects.AffectedBusinessTransactions.BusinessTransactions)
		d.Set("business_transaction_specific_tiers", healthRule.Affects.AffectedBusinessTransactions.SpecificTiers)
		if healthRule.Affects.AffectedBusinessTransactions.PatternMatcher != nil {
			d.Set("business_transaction_match", healthRule.Affects.AffectedBusinessTransactions.PatternMatcher.MatchTo)
			d.Set("business_transaction_match_value", healthRule.Affects.AffectedBusinessTransactions.PatternMatcher.MatchValue)
			d.Set("business_transaction_match_negation", healthRule.Affects.AffectedBusinessTransactions.PatternMatcher.ShouldNot)
		}
	}

	if healthRule.Criterias.Critical != nil {
		var criticals []map[string]interface{}
		for i, _ := range healthRule.Criterias.Critical.Conditions {
			criticals = append(criticals, mapConditionToSchema(healthRule.Criterias.Critical.Conditions[i]))
		}

		err := d.Set("critical_criteria", criticals)
		if err != nil {
			fmt.Println(err)
		}
	}

	if healthRule.Criterias.Warning != nil {
		var warnings []map[string]interface{}
		for i, _ := range healthRule.Criterias.Warning.Conditions {
			warnings = append(warnings, mapConditionToSchema(healthRule.Criterias.Warning.Conditions[i]))
		}

		err := d.Set("warning_criteria", warnings)
		if err != nil {
			fmt.Println(err)
		}
	}
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
