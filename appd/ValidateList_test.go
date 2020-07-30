package appd

import "testing"

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
