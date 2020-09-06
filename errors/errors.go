package errors

import "reflect"

// APIError is the generic structure for any api error
type APIError struct {
	Name       string `json:"name"`
	AppCode    string `json:"appCode,omitempty"`
	StatusCode int    `json:"statusCode"`
	Detail     string `json:"detail"`
	Pointer    string `json:"pointer"`
}

// New instantiates a new error from the inputted map of values
func New(defaults *APIError, overrides *APIError) *APIError {
	e := APIError{}
	v := reflect.ValueOf(defaults)
	ov := reflect.ValueOf(overrides)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		value := v.Field(i).Interface()
		oValue := reflect.ValueOf(ov.FieldByName(key))
		if !oValue.IsNil() && !oValue.IsZero() {
			value = oValue
		}
		ev := reflect.ValueOf(e)
		ev.FieldByName(key).Set(reflect.ValueOf(value))
		e = ev.Interface().(APIError)
	}

	return &e
}

// NewInternalError instantiates a new internal error instance, populating defaults
func NewInternalError(overrides *APIError) *APIError {
	defaults := &APIError{
		StatusCode: 500,
		Name:       "Internal Error",
		AppCode:    "0",
		Detail:     "Unknown internal error",
		Pointer:    "",
	}

	return New(defaults, overrides)
}

// NewUnauthorizedError instantiates a new unauthorized error instance, populating defaults
func NewUnauthorizedError(overrides *APIError) *APIError {
	defaults := &APIError{
		StatusCode: 401,
		Name:       "Unauthorized",
		AppCode:    "100",
		Detail:     "Credentials provided do not have access to this resource",
		Pointer:    "",
	}

	return New(defaults, overrides)
}
