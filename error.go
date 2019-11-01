package xmlvalue

import "fmt"

const (
	errParamInvalidParamType = err("invalid parameter type")
	errEmptyString           = err("empty string parameter")
	errNilParameter          = err("nil parameter")
	errNotInitialized        = err("xmlvalue not initialized")
	errNotExist              = err("not exist")
	errOutOrRange            = err("out of range")
)

// err is the internal error type
type err string

// Error meets the interface error
func (e err) Error() string {
	return string(e)
}

func errorf(f string, v ...interface{}) err {
	s := fmt.Sprintf(f, v...)
	return err(s)
}
