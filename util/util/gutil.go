// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gutil provides utility functions.
package util

import (
	"fmt"
)

const (
	dumpIndent = `    `
)

// Throw throws out an exception, which can be caught be TryCatch or recover.
func Throw(exception interface{}) {
	if exception != nil {
		panic(exception)
	}
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
func Try(try func()) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok {
				err = v
			}
		}
	}()
	try()
	return
}

// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs ans passes the exception as an error.
func TryCatch(try func(), catch ...func(exception error)) {
	defer func() {
		if exception := recover(); exception != nil && len(catch) > 0 {
			if v, ok := exception.(error); ok {
				catch[0](v)
			} else {
				catch[0](fmt.Errorf(`%+v`, exception))
			}
		}
	}()
	try()
}
