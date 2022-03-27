package string_utils

import "errors"

var IndexOut = errors.New("index out")

var BeginSubstringNotFound = errors.New("the begin substring not found")

var EndSubstringNotFound = errors.New("the end substring not found")
