package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Has(field string) bool {
	input := f.Get(field)

	if input == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MinLength(field string, length int, req *http.Request) bool {
	input := req.Form.Get(field)

	if len(input) < length {
		f.Errors.Add(field, fmt.Sprintf("This fields must be at least %d characters long", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) bool {
	input := f.Get(field)

	if !govalidator.IsEmail(input) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}

	return true
}
