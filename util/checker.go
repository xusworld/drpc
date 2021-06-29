package util

import (
	"context"
	"reflect"
)

// IsParamNumEqualsToN
func IsParamNumEqualsToN(typ reflect.Type, N int) bool {
	return typ.NumIn() == N
}

// IsNthParamImplContext
func IsNthParamImplContext(typ reflect.Type, N int) bool {
	// nil  is the zero value of reference types, simply conversion is OK
	actualContextType := (*context.Context)(nil)
	// context.Context is an interface not a reflect.Type, we need to convert it to  reflect.Type
	contextType := reflect.TypeOf(actualContextType).Elem()
	// get function's first input parameter
	firstParam := typ.In(N)

	return firstParam.Implements(contextType)
}

// IsNthParamIsPtr
func IsNthParamIsPtr(typ reflect.Type, N int) bool {
	return typ.In(N).Kind() == reflect.Ptr
}

// IsReturnNumsEqualsToN
func IsReturnNumEqualsToN(typ reflect.Type, N int) bool {
	return typ.NumOut() == N
}

// IsFirstReturnValImplError
func IsFirstReturnValImplError(typ reflect.Type) bool {
	actualErrorType := (*error)(nil)
	errorType := reflect.TypeOf(actualErrorType).Elem()
	firstReturnVal := typ.Out(0)

	return firstReturnVal.Implements(errorType)
}
