package conf

import (
	errs "github.com/bdlm/errors"
)

/*
Internal errors
*/
const (
	// ErrInvalidValueType - The defined value type does not match the type requested
	ErrInvalidValueType errs.Code = iota + 1000
)

func init() {
	errs.Codes[ErrInvalidValueType] = errs.ErrCode{
		Int:  "The defined value type does not match the type requested",
		HTTP: 0,
	}
}
