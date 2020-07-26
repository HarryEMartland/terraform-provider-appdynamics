package appd

import (
	"fmt"
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"execute_actions_in_batch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"action_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {

					validValues := []string{
						"SMS",
						"EMAIL",
						"CUSTOM_EMAIL",
						"THREAD_DUMP",
						"HTTP_REQUEST",
						"CUSTOM",
						"RUN_SCRIPT_ON_NODES",
						"DIAGNOSE_BUSINESS_TRANSACTIONS",
						"CREATE_UPDATE_JIRA",
					}

					strVal := val.(string)

					if !contains(validValues, strVal) {
						errs = append(errs, fmt.Errorf("%s is not a valid value for %s (%v)", strVal, key, validValues))
					}

					return
				},
			},
			"health_rule_event_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"health_rule_scope_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {

					validValues := []string{
						"ALL_HEALTH_RULES",
						"SPECIFIC_HEALTH_RULES",
					}

					strVal := val.(string)

					if !contains(validValues, strVal) {
						errs = append(errs, fmt.Errorf("%s is not a valid value for %s (%v)", strVal, key, validValues))
					}

					return
				},
			},
			"health_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"other_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	policy := createPolicy(d)

	updatedPolicy, err := appdClient.CreatePolicy(&policy, applicationId)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(updatedPolicy.Id))

	return resourcePolicyRead(d, m)
}

func createPolicy(d *schema.ResourceData) client.Policy {

	policy := client.Policy{
		Name:                  d.Get("name").(string),
		Enabled:               d.Get("enabled").(bool),
		ExecuteActionsInBatch: d.Get("execute_actions_in_batch").(bool),
		Actions: []*client.PolicyAction{{
			ActionName: d.Get("action_name").(string),
			ActionType: d.Get("action_type").(string),
		}},
		Events: &client.Events{
			HealthRuleEvents: &client.HealthRuleEvents{
				HealthRuleEventTypes: d.Get("health_rule_event_types").([]interface{}),
				HealthRuleScope: &client.HealthRuleScope{
					HealthRuleScopeType: d.Get("health_rule_scope_type").(string),
					HealthRules:         d.Get("health_rules").([]interface{}),
				},
			},
			OtherEvents: d.Get("other_events").([]interface{}),
		},
	}
	return policy
}

func updatePolicy(d *schema.ResourceData, policy client.Policy) {
	d.Set("name", policy.Name)
	d.Set("enabled", policy.Enabled)
	d.Set("execute_actions_in_batch", policy.ExecuteActionsInBatch)

	action := policy.Actions[0]
	d.Set("action_name", action.ActionName)
	d.Set("action_type", action.ActionType)
	d.Set("health_rule_event_types", policy.Events.HealthRuleEvents.HealthRuleEventTypes)
	d.Set("health_rule_scope_type", policy.Events.HealthRuleEvents.HealthRuleScope.HealthRuleScopeType)
	d.Set("health_rules", policy.Events.HealthRuleEvents.HealthRuleScope.HealthRules)
	d.Set("other_events", policy.Events.OtherEvents)
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	policyId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	policy, err := appdClient.GetPolicy(policyId, applicationId)
	if err != nil {
		return err
	}

	updatePolicy(d, *policy)

	return nil
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	healthRule := createPolicy(d)

	healthRuleId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	healthRule.Id = healthRuleId

	_, err = appdClient.UpdatePolicy(&healthRule, applicationId)
	if err != nil {
		return err
	}

	return resourcePolicyRead(d, m)
}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	policyId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = appdClient.DeletePolicy(applicationId, policyId)
	if err != nil {
		return err
	}

	return nil
}
