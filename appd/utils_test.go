package appd

import (
	"fmt"
	"github.com/HarryEMartland/appd-terraform-provider/appd/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func arrayToString(a []string) string {
	result := "["

	for _, s := range a {
		result += fmt.Sprintf("\"%s\",", s)
	}

	result += "]"
	return result
}

func mapToString(m map[string]string) string {
	result := "{"

	for key, value := range m {
		result += fmt.Sprintf("%s: \"%s\",", key, value)
	}

	result += "}"
	return result
}

func hash(s string) int {
	return schema.HashString(s)
}
