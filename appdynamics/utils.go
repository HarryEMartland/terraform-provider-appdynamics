package appdynamics

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"time"
)

func arrayToString(a []string) string {
	result := "["

	for _, s := range a {
		result += fmt.Sprintf("\"%s\",", s)
	}

	result += "]"
	return result
}

func castArrayToStringArray(array []interface{}) []string {
	s := make([]string, len(array))
	for i, v := range array {
		s[i] = fmt.Sprint(v)
	}
	return s
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

func RetryCheck(check func(state *terraform.State) error) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		backOff := backoff.NewExponentialBackOff()
		backOff.MaxElapsedTime = 10 * time.Second

		err := backoff.Retry(func() error {
			err := check(state)
			if err != nil {
				fmt.Printf("retry function failed %s\n", err)
			}
			return err
		}, backOff)

		return err
	}
}
