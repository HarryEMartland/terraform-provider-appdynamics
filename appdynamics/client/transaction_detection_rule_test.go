package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldGetListOfTransactionRule(t *testing.T) {

	secret := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	queryHandler := func(w http.ResponseWriter, r *http.Request) {
		file, _ := ioutil.ReadFile("./transaction_rules_response.json")
		assert.Equal(t, "/controller/restui/transactionConfigProto/getRules/1234", r.URL.Path, "path should be correct")
		assert.Equal(t, fmt.Sprintf("Bearer %s", secret), r.Header.Get("Authorization"), "request should be authenticated")
		w.Write(file)
	}

	ts := httptest.NewServer(http.HandlerFunc(queryHandler))

	client := AppDClient{
		BaseUrl: ts.URL,
		Secret:  secret,
	}

	configs, err := client.GetTransactionDetectionRules(1234)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 2, len(configs.RuleScopeSummaryMappings), "should contain two items")
	assert.Equal(t, "TX_MATCH_RULE", configs.RuleScopeSummaryMappings[0].Rule.Type)

	assert.Equal(t, "87f83904-d9bd-11ea-87d0-0242ac130003", configs.RuleScopeSummaryMappings[0].Rule.Summary.Id)
	assert.Equal(t, "com.appdynamics.platform.services.configuration.impl.persistenceapi.model.ConfigRuleEntity", configs.RuleScopeSummaryMappings[0].Rule.Summary.Type)
	assert.Equal(t, "7d2d1f44-d9bd-11ea-87d0-0242ac130003", configs.RuleScopeSummaryMappings[0].Rule.Summary.AccountId)
	assert.Equal(t, "test", configs.RuleScopeSummaryMappings[0].Rule.Summary.Name)
	assert.Equal(t, "test description", configs.RuleScopeSummaryMappings[0].Rule.Summary.Description)
	assert.Equal(t, 1596904293456, *configs.RuleScopeSummaryMappings[0].Rule.Summary.CreatedOn)
	assert.Equal(t, 1596913107229, *configs.RuleScopeSummaryMappings[0].Rule.Summary.UpdatedOn)

	assert.Equal(t, true, configs.RuleScopeSummaryMappings[0].Rule.Enabled)
	assert.Equal(t, 40, configs.RuleScopeSummaryMappings[0].Rule.Priority)
	assert.Equal(t, 3, *configs.RuleScopeSummaryMappings[0].Rule.Version)
	assert.Equal(t, "NODE_JS_SERVER", configs.RuleScopeSummaryMappings[0].Rule.AgentType)

	assert.Equal(t, "CUSTOM", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.Type)

	assert.Equal(t, "INCLUDE", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.Type)
	assert.Equal(t, "NODEJS_WEB", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.TxEntryPointType)

	assert.Equal(t, 1, len(configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions))

	assert.Equal(t, "HTTP", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions[0].Type)

	assert.Equal(t, "GET", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.HttpMethod)

	assert.Equal(t, "IS_IN_LIST", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.Type)
	assert.Equal(t, 1, len(configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.MatchStrings))
	assert.Equal(t, "/root,/health", configs.RuleScopeSummaryMappings[0].Rule.TxMatchRule.TxCustomRule.MatchConditions[0].HttpMatch.Uri.MatchStrings[0])

}

func TestShouldFindTransactionRule(t *testing.T) {

	secret := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	queryHandler := func(w http.ResponseWriter, r *http.Request) {
		file, _ := ioutil.ReadFile("./transaction_rules_response.json")
		assert.Equal(t, "/controller/restui/transactionConfigProto/getRules/1234", r.URL.Path, "path should be correct")
		assert.Equal(t, fmt.Sprintf("Bearer %s", secret), r.Header.Get("Authorization"), "request should be authenticated")
		w.Write(file)
	}

	ts := httptest.NewServer(http.HandlerFunc(queryHandler))

	client := AppDClient{
		BaseUrl: ts.URL,
		Secret:  secret,
	}

	rule, found, err := client.GetTransactionDetectionRule(1234, "2dfb495f-6d82-424c-9639-d745263ee9ca")

	if err != nil {
		t.Error(err)
		return
	}

	assert.True(t, found)
	assert.NotNil(t, rule)
	assert.Equal(t, "2dfb495f-6d82-424c-9639-d745263ee9ca", rule.Rule.Summary.Id)
}

func TestShouldPostNewTransactionConfig(t *testing.T) {

	secret := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	queryHandler := func(w http.ResponseWriter, r *http.Request) {
		file, _ := ioutil.ReadFile("./transaction_rule_response.json")
		assert.Equal(t, "/controller/restui/transactionConfigProto/createRule", r.URL.Path, "path should be correct")
		assert.Equal(t, "", r.URL.Query().Get("scopeId"), "query should have empty scopeId")
		assert.Equal(t, "1234", r.URL.Query().Get("applicationId"), "query contain applicationId")
		assert.Equal(t, fmt.Sprintf("Bearer %s", secret), r.Header.Get("Authorization"), "request should be authenticated")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		expectedBody, _ := ioutil.ReadFile("./transaction_rule_expected_request.json")
		assert.JSONEq(t, string(body), string(expectedBody), "body should match")

		w.Write(file)
	}

	ts := httptest.NewServer(http.HandlerFunc(queryHandler))

	client := AppDClient{
		BaseUrl: ts.URL,
		Secret:  secret,
	}

	rule := createRule()
	result, err := client.CreateTransactionDetectionRule(1234, &rule)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "SUCCESS", result.ResultType)

	assert.Equal(t, "878bdfc0-da7b-11ea-87d0-0242ac130003", result.Successes[0].Summary.Id)
	assert.Equal(t, 1, len(result.Successes))
	assert.Equal(t, "com.appdynamics.platform.services.configuration.impl.persistenceapi.model.ConfigRuleEntity", result.Successes[0].Summary.Type)
	assert.Equal(t, "8c469528-da7b-11ea-87d0-0242ac130003", result.Successes[0].Summary.AccountId)

}

func TestShouldPostUpdatedTransactionConfig(t *testing.T) {

	secret := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	queryHandler := func(w http.ResponseWriter, r *http.Request) {
		file, _ := ioutil.ReadFile("./transaction_rule_response.json")
		assert.Equal(t, "/controller/restui/transactionConfigProto/updateRule", r.URL.Path, "path should be correct")
		assert.Equal(t, "", r.URL.Query().Get("scopeId"), "query should have empty scopeId")
		assert.Equal(t, "1234", r.URL.Query().Get("applicationId"), "query contain applicationId")
		assert.Equal(t, fmt.Sprintf("Bearer %s", secret), r.Header.Get("Authorization"), "request should be authenticated")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		expectedBody, _ := ioutil.ReadFile("./transaction_rule_expected_request.json")
		assert.JSONEq(t, string(body), string(expectedBody), "body should match")

		w.Write(file)
	}

	ts := httptest.NewServer(http.HandlerFunc(queryHandler))

	client := AppDClient{
		BaseUrl: ts.URL,
		Secret:  secret,
	}

	rule := createRule()
	result, err := client.UpdateTransactionDetectionRule(1234, &rule)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "SUCCESS", result.ResultType)

	assert.Equal(t, "878bdfc0-da7b-11ea-87d0-0242ac130003", result.Successes[0].Summary.Id)
	assert.Equal(t, 1, len(result.Successes))
	assert.Equal(t, "com.appdynamics.platform.services.configuration.impl.persistenceapi.model.ConfigRuleEntity", result.Successes[0].Summary.Type)
	assert.Equal(t, "8c469528-da7b-11ea-87d0-0242ac130003", result.Successes[0].Summary.AccountId)

}

func TestShouldDeleteTransactionConfigs(t *testing.T) {

	secret := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	queryHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/controller/restui/transactionConfigProto/deleteRules", r.URL.Path, "path should be correct")
		assert.Equal(t, fmt.Sprintf("Bearer %s", secret), r.Header.Get("Authorization"), "request should be authenticated")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}

		assert.JSONEq(t, string(`["33e1c67a-da88-11ea-87d0-0242ac130003"]`), string(body), "body should match")

		w.Write([]byte(`{"resultType":"SUCCESS"}`))
	}

	ts := httptest.NewServer(http.HandlerFunc(queryHandler))

	client := AppDClient{
		BaseUrl: ts.URL,
		Secret:  secret,
	}

	configs, err := client.DeleteTransactionDetectionRules([]string{"33e1c67a-da88-11ea-87d0-0242ac130003"})
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "SUCCESS", configs.ResultType)
}

func createRule() TransactionRule {
	return TransactionRule{
		Type: "TX_MATCH_RULE",
		Summary: &Summary{
			Id:          "",
			Type:        "com.appdynamics.platform.services.configuration.impl.persistenceapi.model.ConfigRuleEntity",
			AccountId:   "ad14e72e-da84-11ea-87d0-0242ac130003",
			Name:        "test",
			Description: "test description",
			CreatedOn:   nil,
			UpdatedOn:   nil,
		},
		Enabled:   true,
		Priority:  1,
		Version:   nil,
		AgentType: "NODE_JS_SERVER",
		TxMatchRule: &TxMatchRule{
			Type:      "CUSTOM",
			AgentType: "NODE_JS_SERVER",
			TxCustomRule: &TxCustomRule{
				Type:             "INCLUDE",
				TxEntryPointType: "NODEJS_WEB",
				MatchConditions: []*MatchCondition{{
					Type: "HTTP",
					HttpMatch: &HttpMatch{
						Uri: &Uri{
							Type:         "EQUALS",
							MatchStrings: []interface{}{"/bob"},
						},
						HttpMethod: "POST",
					},
				}},
			},
		},
	}
}
