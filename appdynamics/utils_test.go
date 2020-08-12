package appdynamics

import (
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
)

var appDClient client.AppDClient
var applicationIdI int
var applicationIdS string
var httpActionTemplateName string
var bt1 string
var bt2 string
var tier1 string
var tier2 string
var accountId string
var scopeId string

func init() {
	_, acceptanceTest := os.LookupEnv("TF_ACC")
	if !acceptanceTest {
		return
	}

	applicationId, err := strconv.Atoi(os.Getenv("APPD_APPLICATION_ID"))
	if err != nil {
		log.Fatal(fmt.Sprintf("error parsing application id: %s", os.Getenv("APPD_APPLICATION_ID")))
	}
	appDClient = client.AppDClient{
		BaseUrl: os.Getenv("APPD_CONTROLLER_BASE_URL"),
		Secret:  os.Getenv("APPD_SECRET"),
	}
	applicationIdI = applicationId
	applicationIdS = os.Getenv("APPD_APPLICATION_ID")
	httpActionTemplateName = os.Getenv("APPD_HTTP_ACTION_TEMPLATE_NAME")
	bt1 = os.Getenv("APPD_BT1")
	bt2 = os.Getenv("APPD_BT2")
	tier1 = os.Getenv("APPD_TIER1")
	tier2 = os.Getenv("APPD_TIER2")
	accountId = os.Getenv("APPD_ACCOUNT_ID")
	scopeId = os.Getenv("APPD_SCOPE_ID")
}

func TestValidateListShouldReturnNoErrsOrWarnsWhenIsValid(t *testing.T) {

	warns, errs := validateList([]string{"Valid"})("Valid", "test_key")

	if len(warns) > 0 {
		t.Errorf("warns should be empty, got %d", len(warns))
	}

	if len(errs) > 0 {
		t.Errorf("errs should be empty, got %d", len(warns))
	}

}

func TestValidateListShouldReturnNoErrsOrWarnsWhenIsValidList(t *testing.T) {

	warns, errs := validateList([]string{"Another Valid", "Valid"})("Valid", "test_key")

	if len(warns) > 0 {
		t.Errorf("warns should be empty, got %d", len(warns))
	}

	if len(errs) > 0 {
		t.Errorf("errs should be empty, got %d", len(warns))
	}

}

func TestValidateListShouldReturnErrorWhenNotValid(t *testing.T) {

	warns, errs := validateList([]string{"Valid"})("Not Valid", "test_key")

	if len(warns) > 0 {
		t.Errorf("warns should be empty, got %d", len(warns))
	}

	if len(errs) != 1 {
		t.Errorf("one error should be returned, got %d", len(warns))
	}

	if errs[0].Error() != "Not Valid is not a valid value for test_key [Valid]" {
		t.Errorf("error message did did not match, got %s", errs[0].Error())
	}

}

func TestValidateListShouldReturnErrorWhenNotValidList(t *testing.T) {

	warns, errs := validateList([]string{"Another Valid", "Valid"})("Not Valid", "test_key")

	if len(warns) > 0 {
		t.Errorf("warns should be empty, got %d", len(warns))
	}

	if len(errs) != 1 {
		t.Errorf("one error should be returned, got %d", len(warns))
	}

	if errs[0].Error() != "Not Valid is not a valid value for test_key [Another Valid Valid]" {
		t.Errorf("error message did did not match, got %s", errs[0].Error())
	}

}

func TestMapToStringSingle(t *testing.T) {
	assert.Equal(t, "{k1: \"v1\",}", mapToString(map[string]string{"k1": "v1"}), "map should be correctly formatted")
}

func TestMapToStringMultiple(t *testing.T) {

	toString := mapToString(map[string]string{"k1": "v1", "k2": "v2"})
	//have to test like this as arrays are nondeterministic
	assert.Contains(t, toString, "k1: \"v1\"", "contains first key")
	assert.Contains(t, toString, "k2: \"v2\"", "contains second key")
}

func TestArrayToStringEmpty(t *testing.T) {
	var strings []string
	assert.Equal(t, "[]", arrayToString(strings))
}

func TestArrayToStringSingle(t *testing.T) {
	strings := []string{"test"}
	assert.Equal(t, "[\"test\",]", arrayToString(strings))
}

func TestArrayToStringMultiple(t *testing.T) {
	strings := []string{"test", "test2"}
	assert.Equal(t, "[\"test\",\"test2\",]", arrayToString(strings))
}
