package appdynamics

import (
	"errors"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func resourceTransactionDetectionRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransactionRuleCreate,
		Read:   resourceTransactionRuleRead,
		Update: resourceTransactionRuleUpdate,
		Delete: resourceTransactionRuleDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"agent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"entry_point_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"http_uri_match_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"http_method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"http_uris": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTransactionRuleCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	transactionRule := createTransactionRule(d)

	updatedHealthRule, err := appdClient.CreateTransactionRule(applicationId, transactionRule)
	if err != nil {
		return err
	}
	if len(updatedHealthRule.Successes) != 1 {
		return errors.New("error creating transaction rule, no results returned")
	}

	d.SetId(updatedHealthRule.Successes[0].Summary.Id)

	return resourceTransactionRuleRead(d, m)
}

func createTransactionRule(d *schema.ResourceData) *client.TransactionRule {

	name := d.Get("name").(string)
	agentType := d.Get("agent_type").(string)

	httpUriList := castArrayToStringArray(d.Get("http_uris").(*schema.Set).List())
	transactionRule := &client.TransactionRule{
		Type: "TX_MATCH_RULE",
		Summary: &client.Summary{
			Id:          d.Id(),
			Type:        "com.appdynamics.platform.services.configuration.impl.persistenceapi.model.ConfigRuleEntity",
			AccountId:   d.Get("account_id").(string),
			Name:        name,
			Description: d.Get("description").(string),
		},
		Enabled:   d.Get("enabled").(bool),
		Priority:  d.Get("priority").(int),
		AgentType: agentType,
		TxMatchRule: &client.TxMatchRule{
			AgentType: agentType,
			Type:      "CUSTOM",
			TxCustomRule: &client.TxCustomRule{
				Type:             "INCLUDE",
				TxEntryPointType: d.Get("entry_point_type").(string),
				MatchConditions: []*client.MatchCondition{{
					Type: "HTTP",
					HttpMatch: &client.HttpMatch{
						Uri: &client.Uri{
							Type:         d.Get("http_uri_match_type").(string),
							MatchStrings: []interface{}{strings.Join(httpUriList, ",")},
						},
						HttpMethod: d.Get("http_method").(string),
					},
				}},
			},
		},
	}

	return transactionRule
}

func updateTransactionRule(d *schema.ResourceData, scope client.RuleScope) {
	matchStrings := scope.Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.MatchStrings

	var flattenedMatches []string
	for _, v := range matchStrings {
		split := strings.Split(v.(string), ",")
		for _,m := range split{
			flattenedMatches = append(flattenedMatches, m)
		}
	}

	d.Set("name", scope.Rule.Summary.Name)
	d.Set("agent_type", scope.Rule.AgentType)
	d.Set("account_id", scope.Rule.Summary.AccountId)
	d.Set("description", scope.Rule.Summary.Description)
	d.Set("enabled", scope.Rule.Enabled)
	d.Set("priority", scope.Rule.Priority)
	d.Set("entry_point_type", scope.Rule.TxMatchRule.TxCustomRule.TxEntryPointType)
	d.Set("http_uri_match_type", scope.Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.Type)
	d.Set("http_uris", flattenedMatches)
	d.Set("http_method", scope.Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.HttpMethod)
}

func resourceTransactionRuleRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	rule, found, err := appdClient.GetTransactionRule(applicationId, id)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("could not find transaction rule")
	}

	updateTransactionRule(d, *rule)

	return nil
}

func resourceTransactionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	transactionRule := createTransactionRule(d)

	_, err := appdClient.UpdateTransactionRule(applicationId, transactionRule)
	if err != nil {
		return err
	}

	return resourceActionRead(d, m)
}

func resourceTransactionRuleDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)

	id := d.Id()

	_, err := appdClient.DeleteTransactionRules([]string{id})
	if err != nil {
		return err
	}

	return nil
}
