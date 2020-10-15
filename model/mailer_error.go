package model

// maximum number of attempts to connect to the dialer
var MaxConnectionAttempts = 5

// this is to provide more info on MaxConnectionAttempts
type ErrorMaxConnectionAttempts struct {
	ownError			error
}

// now it becomes an error automatically
func (e *ErrorMaxConnectionAttempts) Error() string {
	var prefix = "maximum num. of attempts exceeded"
	if e.ownError != nil {
		return prefix + " | " + e.ownError.Error()
	}

	return prefix
}
