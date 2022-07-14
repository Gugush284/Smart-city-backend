package modeltimetable

import validation "github.com/go-ozzo/ozzo-validation"

// Validation ...
func (t *Timetable) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.IdUser, validation.Required),
		validation.Field(&t.Title, validation.Required),
		validation.Field(&t.Time, validation.Required),
	)
}
