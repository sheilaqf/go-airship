package airship

// AirShipError is an error that is returned from the API
// in form of JSN response.
type AirshipError struct {
	// Details ...
	Details AirshipErrorDetails `json:"details,omitempty"`
	// Message ...
	Message string `json:"error"`
	// Code ...
	Code int `json:"error_code,omitempty"`
	// OK ...
	OK bool `json:"ok"`
}

// AirshipErrorDetails ...
type AirshipErrorDetails struct {
	// Message ...
	Message string `json:"error,omitempty"`
	// Location ...
	Location AirshipErrorDetailsLocation `json:"location,omitempty"`
	// Path ...
	Path string `json:"path,omitempty"`
}

// AirshipErrorDetailsLocation ...
type AirshipErrorDetailsLocation struct {
	// Column ...
	Column int `json:"column,omitempty"`
	// Column ...
	Line int `json:"line,omitempty"`
}

// Error is returning a string to confirm to the error interface
func (l AirshipError) Error() string {
	return l.Message
}
