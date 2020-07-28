package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func (c *AppDClient) CreateHealthRule(healthRule *HealthRule, applicationId int) (*HealthRule, error) {

	resp, err := req.Post(c.createHealthRulesUrl(applicationId), c.createAuthHeader(), req.BodyJSON(&healthRule))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()

		return nil, errors.New(fmt.Sprintf("Error creating Health Rule: %d, %s, %s", resp.Response().StatusCode, respString))
	}

	updated := HealthRule{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) UpdateHealthRule(healthRule *HealthRule, applicationId int) (*HealthRule, error) {

	resp, err := req.Put(c.createHealthRuleUrl(healthRule.ID, applicationId), c.createAuthHeader(), req.BodyJSON(&healthRule))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Health Rule: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := HealthRule{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) DeleteHealthRule(applicationId int, healthRuleId int) error {

	_, err := req.Delete(c.createHealthRuleUrl(healthRuleId, applicationId), c.createAuthHeader())
	if err != nil {
		return err
	}

	return nil
}

func (c *AppDClient) GetHealthRule(healthRuleId int, applicationId int) (*HealthRule, error) {

	resp, err := req.Get(c.createHealthRuleUrl(healthRuleId, applicationId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting Health Rule: %d, %s, %s", resp.Response().StatusCode, c.createHealthRuleUrl(healthRuleId, applicationId), respString))
	}

	updated := HealthRule{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) createHealthRulesUrl(applicationId int) string {
	return fmt.Sprintf("%s/%s", c.createUrl(applicationId), "health-rules")
}

func (c *AppDClient) createHealthRuleUrl(healthRuleId int, applicationId int) string {
	return fmt.Sprintf("%s/%d", c.createHealthRulesUrl(applicationId), healthRuleId)
}

type HealthRule struct {
	ID                      int        `json:"id"`
	Name                    string     `json:"name"`
	Enabled                 bool       `json:"enabled"`
	UseDataFromLastNMinutes int        `json:"useDataFromLastNMinutes"`
	WaitTimeAfterViolation  int        `json:"waitTimeAfterViolation"`
	Affects                 *Affects   `json:"affects"`
	Criterias               *Criterias `json:"evalCriterias"`
}

type Criterias struct {
	Critical *Criteria `json:"criticalCriteria"`
	Warning  *Criteria `json:"warningCriteria"`
}

type Criteria struct {
	ConditionAggregationType string       `json:"conditionAggregationType"`
	Conditions               []*Condition `json:"conditions"`
}

type Condition struct {
	Name                   string      `json:"name"`
	ShortName              string      `json:"shortName"`
	EvaluateToTrueOnNoData bool        `json:"evaluateToTrueOnNoData"`
	EvalDetail             *EvalDetail `json:"evalDetail"`
}

type EvalDetail struct {
	EvalDetailType          string            `json:"evalDetailType"`
	MetricAggregateFunction string            `json:"metricAggregateFunction"`
	MetricPath              string            `json:"metricPath"`
	MetricEvalDetail        *MetricEvalDetail `json:"metricEvalDetail"`
}

type MetricEvalDetail struct {
	MetricEvalDetailType string  `json:"metricEvalDetailType"`
	BaselineCondition    *string `json:"baselineCondition"`
	BaselineName         *string `json:"baselineName"`
	BaselineUnit         *string `json:"baselineUnit"`
	CompareValue         float64 `json:"compareValue"`
	CompareCondition     *string `json:"compareCondition"`
}

type Affects struct {
	AffectedEntityType           string       `json:"affectedEntityType"`
	AffectedBusinessTransactions *Transaction `json:"affectedBusinessTransactions"`
}

type Transaction struct {
	BusinessTransactionScope string         `json:"businessTransactionScope"`
	BusinessTransactions     *[]interface{} `json:"businessTransactions"`
}
