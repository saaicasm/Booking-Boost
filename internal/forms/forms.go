package forms

import (
	"net/http"
	"net/url"
)

// Creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}


func (f *Form) Has(field string, r http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}