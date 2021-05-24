package util

import (
	"context"
	"reflect"
)

// MethodChecker
type MethodChecker struct {
	reflect.Method
}

func (mc *MethodChecker) IsExportMethod() bool {
	return mc.PkgPath != ""
}

func (mc *MethodChecker) IsParamNumEqualsToN(n int) bool {
	return mc.Type.NumIn() == n
}

func (mc *MethodChecker) IsFirstParamImplContext() bool {
	// nil  is the zero value of reference types, simply conversion is OK
	actualContextType := (*context.Context)(nil)
	// context.Context is an interface not a reflect.Type, we need to convert it to  reflect.Type
	contextType := reflect.TypeOf(actualContextType).Elem()
	// get function's first input parameter
	firstParam := mc.Type.In(0)

	return firstParam.Implements(contextType)
}

func (mc *MethodChecker) IsNthParamIsPtr(n int) bool {
	return mc.Type.In(n).Kind() == reflect.Ptr
}

func (mc *MethodChecker) IsReturnNumsEqualsToN(n int) bool {
	return mc.Type.NumOut() == n
}
func (mc *MethodChecker) IsFirstReturnValImplError() bool {
	actualErrorType := (*error)(nil)
	errorType := reflect.TypeOf(actualErrorType).Elem()
	firstReturnVal := mc.Type.Out(0)

	return firstReturnVal.Implements(errorType)
}
