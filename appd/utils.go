package appd

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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

func validateList(validValues []string) func(val interface{}, key string) (warns []string, errs []error) {
	return func(val interface{}, key string) (warns []string, errs []error) {
		strVal := val.(string)

		if !contains(validValues, strVal) {
			errs = append(errs, fmt.Errorf("%s is not a valid value for %s %v", strVal, key, validValues))
		}

		return
	}
}
