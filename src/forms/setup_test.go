package forms

import (
	"net/url"
	"os"
	"testing"
)

var testData = url.Values{}

var testForm *Form

const expectedField = "TEST"

func TestMain(m *testing.M) {
	testData.Add(expectedField, "SAD")
	testForm = New(testData)


	os.Exit(m.Run())
}
