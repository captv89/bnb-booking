package forms

import (
	"log"
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}


func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}


// New creates a new form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	c := r.Form.Get(field)
	if c == "" {
		f.Errors.Add(field, "This field cannot be empty")
		log.Println(field)
		return false
	}
	return true
}