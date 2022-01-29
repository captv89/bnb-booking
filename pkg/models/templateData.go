package models

import "github.com/captv89/bnb-booking/pkg/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	BoolMap map[string]bool
	FloatMap map[string]float32
	Data map[string]interface{}
	Token string
	Flash string
	Warning string
	Error string
	Form *forms.Form
}