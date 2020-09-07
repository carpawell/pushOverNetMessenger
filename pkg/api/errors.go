package api

import "errors"

var EmptyMessage = errors.New("empty message")
var FromParameterNotFound = errors.New("'from' parameter is necessary")
