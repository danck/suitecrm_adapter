package errtypes

// BadRequest is used if the incoming payload doesn't comply with the
// specification
type BadRequest struct{ Msg string }

func (e BadRequest) Error() string { return e.Msg }

// NotFound is used if a request is syntactically correct, but the requested
// resource does not exist
type NotFound struct{ error }

// NotImplemented is used if calls a function that isn't implemented
// A function here is specified by the HTTP verb and prefix (i.e. PUT /orders)
type NotImplemented struct{ Msg string }

func (e NotImplemented) Error() string { return e.Msg }
