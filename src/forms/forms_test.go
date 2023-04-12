package forms

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	testData.Add("SADU", "SAD")
	testForm = New(testData)

	sad := testForm.Get("SADU")

	if sad != "SAD" {
		t.Error("Bad forming form")
	}
}

func TestRequired(t *testing.T) {
	testForm.Required(expectedField)

	if len(testForm.Errors) > 0 {
		t.Error("Missing required validation")
	}
}

func TestHas(t *testing.T) {
	hasTest := testForm.Has(expectedField)

	if !hasTest {
		t.Error("Not has expected field")
	}
}

func TestValid(t *testing.T) {
	testForm.Add("FAIL", "")

	testForm.Required("FAIL")

	isValid := testForm.Valid()

	if isValid {
		t.Error("Expected error validation")
	}
}

func TestMinLength(t *testing.T) {
	fakeBody := strings.NewReader("sad=sad")

	req := httptest.NewRequest("POST", "/fake", fakeBody)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err := req.ParseForm()
	if err != nil {
		t.Error("error parsing the form")
	}

	sad := req.Form.Get("sad")

	if sad != "sad" {
		t.Error("bad form formation")
	}

	testForm.Add("sad", sad)

	lenOk := testForm.MinLength("sad", 4, req)

	if lenOk {
		t.Error("expect error")
	}

}

func TestIsEmail(t *testing.T) {
	testForm.Add("email", "sa")

	validation := testForm.IsEmail("email")

	if validation {
		t.Error("expect a validation error on email")
	}

}
