package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func (c *AppDClient) GetTransactionDetectionRules(applicationId int) (*TransactionRules, error) {

	resp, err := req.Get(c.createGetTransactionDetectionRuleUrl(applicationId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting Transaction configs: %d, %s", resp.Response().StatusCode, respString))
	}

	var updated TransactionRules
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) GetTransactionDetectionRule(applicationId int, transactionRuleId string) (*RuleScope, bool, error) {

	rules, err := c.GetTransactionDetectionRules(applicationId)

	if err != nil {
		return nil, false, err
	}

	for _, n := range rules.RuleScopeSummaryMappings {
		if transactionRuleId == n.Rule.Summary.Id {
			return n, true, nil
		}
	}

	return nil, false, nil
}

func (c *AppDClient) CreateTransactionDetectionRule(applicationId int, rule *TransactionRule) (*UpdateResult, error) {

	resp, err := req.Post(c.createCreateTransactionDetectionRuleUrl(), c.createAuthHeader(), req.QueryParam{"scopeId": "", "applicationId": applicationId}, req.BodyJSON(rule))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Transaction rule: %d, %s", resp.Response().StatusCode, respString))
	}

	var updated UpdateResult
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) DeleteTransactionDetectionRules(transactionRuleIds []string) (*UpdateResult, error) {

	resp, err := req.Post(c.createDeleteTransactionDetectionRuleUrl(), c.createAuthHeader(), req.BodyJSON(transactionRuleIds))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error deleting Transaction rule: %d, %s", resp.Response().StatusCode, respString))
	}

	var updated UpdateResult
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) UpdateTransactionDetectionRule(applicationId int, transactionConfig *TransactionRule) (*UpdateResult, error) {

	resp, err := req.Post(c.createUpdateTransactionDetectionRuleUrl(), c.createAuthHeader(), req.QueryParam{"scopeId": "", "applicationId": applicationId}, req.BodyJSON(transactionConfig))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error deleting Transaction rule: %d, %s", resp.Response().StatusCode, respString))
	}

	var updated UpdateResult
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) createGetTransactionDetectionRuleUrl(applicationId int) string {
	return fmt.Sprintf("%s/controller/restui/transactionConfigProto/getRules/%d", c.BaseUrl, applicationId)
}

func (c *AppDClient) createCreateTransactionDetectionRuleUrl() string {
	return fmt.Sprintf("%s/controller/restui/transactionConfigProto/createRule", c.BaseUrl)
}

func (c *AppDClient) createDeleteTransactionDetectionRuleUrl() string {
	return fmt.Sprintf("%s/controller/restui/transactionConfigProto/deleteRules", c.BaseUrl)
}

func (c *AppDClient) createUpdateTransactionDetectionRuleUrl() string {
	return fmt.Sprintf("%s/controller/restui/transactionConfigProto/updateRule", c.BaseUrl)
}

type UpdateResult struct {
	ResultType string     `json:"resultType"`
	Successes  []*Success `json:"successes"`
}

type Success struct {
	Summary *Summary `json:"summary"`
}

type TransactionRules struct {
	RuleScopeSummaryMappings []*RuleScope `json:"ruleScopeSummaryMappings"`
}

type RuleScope struct {
	Rule *TransactionRule `json:"rule"`
}

type TransactionRule struct {
	Type        string       `json:"type"`
	Summary     *Summary     `json:"summary"`
	Enabled     bool         `json:"enabled"`
	Priority    int          `json:"priority"`
	Version     *int         `json:"version,omitempty"`
	AgentType   string       `json:"agentType"`
	TxMatchRule *TxMatchRule `json:"txMatchRule"`
}

type Summary struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	AccountId   string `json:"accountId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   *int   `json:"createdOn,omitempty"`
	UpdatedOn   *int   `json:"updatedOn,omitempty"`
}

type TxMatchRule struct {
	AgentType    string        `json:"agentType"`
	Type         string        `json:"type"`
	TxCustomRule *TxCustomRule `json:"txCustomRule"`
}

type TxCustomRule struct {
	Type             string            `json:"type"`
	TxEntryPointType string            `json:"txEntryPointType"`
	MatchConditions  []*MatchCondition `json:"matchConditions"`
}

type MatchCondition struct {
	Type      string     `json:"type"`
	HttpMatch *HttpMatch `json:"httpMatch"`
}

type HttpMatch struct {
	Uri        *Uri   `json:"uri"`
	HttpMethod string `json:"httpMethod"`
}

type Uri struct {
	Type         string   `json:"type"`
	MatchStrings []interface{} `json:"matchStrings"`
}
