package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func (c *AppDClient) CreatePolicy(policy *Policy, applicationId int) (*Policy, error) {

	resp, err := req.Post(c.createPoliciesUrl(applicationId), c.createAuthHeader(), req.BodyJSON(&policy))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Policy: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Policy{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) UpdatePolicy(policy *Policy, applicationId int) (*Policy, error) {

	resp, err := req.Put(c.createPolicyUrl(policy.Id, applicationId), c.createAuthHeader(), req.BodyJSON(&policy))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Policy: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Policy{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) DeletePolicy(applicationId int, policyId int) error {

	resp, err := req.Delete(c.createPolicyUrl(policyId, applicationId), c.createAuthHeader())
	if err != nil {
		return err
	}

	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		errors.New(fmt.Sprintf("Error deleting Policy: %d, %s", resp.Response().StatusCode, respString))
	}

	return nil
}

func (c *AppDClient) GetPolicy(policyId int, applicationId int) (*Policy, error) {

	resp, err := req.Get(c.createPolicyUrl(policyId, applicationId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting Policy: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Policy{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) createPoliciesUrl(applicationId int) string {
	return fmt.Sprintf("%s/%s", c.createUrl(applicationId), "policies")
}

func (c *AppDClient) createPolicyUrl(policyId int, applicationId int) string {
	return fmt.Sprintf("%s/%d", c.createPoliciesUrl(applicationId), policyId)
}

type Policy struct {
	Id                    int             `json:"id"`
	Name                  string          `json:"name"`
	Enabled               bool            `json:"enabled"`
	ExecuteActionsInBatch bool            `json:"executeActionsInBatch"`
	Actions               []*PolicyAction `json:"actions"`
	Events                *Events         `json:"events"`
}

type PolicyAction struct {
	ActionName string `json:"actionName"`
	ActionType string `json:"actionType"`
}

type Events struct {
	HealthRuleEvents *HealthRuleEvents `json:"healthRuleEvents"`
	OtherEvents      []interface{}     `json:"otherEvents"`
}

type HealthRuleEvents struct {
	HealthRuleEventTypes []interface{}    `json:"healthRuleEventTypes"`
	HealthRuleScope      *HealthRuleScope `json:"healthRuleScope"`
}

type HealthRuleScope struct {
	HealthRuleScopeType string        `json:"healthRuleScopeType"`
	HealthRules         []interface{} `json:"healthRules"`
}

type SelectedEntities struct {
	SelectedEntityType string    `json:"selectedEntityType"`
	Entities           []*Entity `json:"entities"`
}

type Entity struct {
	EntityType                   string                        `json:"entityType"`
	SelectedBusinessTransactions *SelectedBusinessTransactions `json:"selectedBusinessTransactions"`
}

type SelectedBusinessTransactions struct {
	BusinessTransactionScope string        `json:"businessTransactionScope"`
	BusinessTransactions     []interface{} `json:"businessTransactions"`
}

type TierOrNode struct {
	TierOrNodeScope string `json:"tierOrNodeScope"`
	TypeOfNode      string `json:"typeOfNode"`
}

type SelectedNodes struct {
	SelectedNodeScope string        `json:"selectedNodeScope"`
	Nodes             []interface{} `json:"nodes"`
}
