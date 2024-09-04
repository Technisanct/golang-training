package errors

import (
	goErr "github.com/pkg/errors"
)

var (
	ErrMarshalStruct         = goErr.New("error in marshalling struct")
	ErrUnMarshalResponseBody = goErr.New("error in unmarshalling http response")
)
