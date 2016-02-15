package errtypes

// TODO(danck) description
type BadRequest struct{ Msg string }

func (e BadRequest) Error() string { return e.Msg }

// TODO(danck) description
type NotFound struct{ error }

// TODO(danck) description
type NotImplemented struct{ Msg string }

func (e NotImplemented) Error() string { return e.Msg }
