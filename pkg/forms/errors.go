package forms

// Declaring an errors map
type errors map[string][]string

// Add adds an error to the map
func (e errors) Add(field, message string){
	e[field] = append(e[field], message)
}

// Get gets the errors from the map
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}