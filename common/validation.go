package common

type ValidationResult struct {
	Form   []FormValidationResult
	Fields []FieldValidationResult
}

type FieldValidationResult struct {
	Field   string
	Invalid bool
	Error   string
	Warning string
}

type FormValidationResult struct {
	Invalid bool
	Error   string
	Warning string
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		Fields: []FieldValidationResult{},
		Form:   []FormValidationResult{},
	}
}

func (vr *ValidationResult) AddFormValidationError(code string) {
	vr.Form = append(vr.Form, FormValidationResult{
		Invalid: true,
		Error:   code,
		Warning: "",
	})
}

func (vr *ValidationResult) AddFieldValidationError(field string, code string) {
	vr.Fields = append(vr.Fields, FieldValidationResult{
		Field:   field,
		Invalid: true,
		Error:   code,
	})
}

func (vr *ValidationResult) IsValid() bool {
	for i := range vr.Form {
		if vr.Form[i].Invalid {
			return false
		}
	}
	for i := range vr.Fields {
		if vr.Fields[i].Invalid {
			return false
		}
	}
	return true
}

func (vr *ValidationResult) ToUserError() error {
	return NewUserErrorWithDetail(ErrFormValidationError, nil, vr)
}

const ErrFormValidationError = "FORM_VALIDATION_ERROR"
const ErrFieldEmpty = "FIELD_EMPTY"
