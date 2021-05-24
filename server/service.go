package server

import (
	"context"
	"reflect"
	"sync"

	"github.com/xusworld/crpc/log"
	"github.com/xusworld/crpc/util"
)

type Service struct {
	name     string                 // name of service
	receiver reflect.Value          // receiver of methods for the service
	rtype    reflect.Type           // type of the receiver
	method   map[string]*methodType // registered methods
}

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type functionType struct {
	sync.Mutex // protects counters
	fn         reflect.Value
	ArgType    reflect.Type
	ReplyType  reflect.Type
}

func (s *Server) RegisterService(receiver interface{}) {
	_, err := s.doRegisterService(receiver)

	if err != nil {
		log.Error("error")
	}
}

func (s *Server) doRegisterService(receiver interface{}) (string, error) {
	service := Service{}
	service.receiver = reflect.ValueOf(receiver)
	service.rtype = reflect.TypeOf(receiver)
	service.name = reflect.Indirect(service.receiver).Type().Name()
	service.method = filterRecevierMethods(service.rtype)

	return service.name, nil
}

func filterRecevierMethods(rtype reflect.Type) map[string]*methodType {
	methods := make(map[string]*methodType, 0)

	// 遍历receiver的所有方法
	for i := 0; i < rtype.NumMethod(); i++ {
		methodChecker := util.MethodChecker{
			Method: rtype.Method(i),
		}

		if !methodChecker.IsParamNumEqualsToN(4) {
			log.Error("")
			continue
		}

		if !methodChecker.IsExportMethod() {
			log.Error("")
			continue
		}

		if !methodChecker.IsFirstParamImplContext() {
			log.Error("")
			continue
		}

		if !methodChecker.IsNthParamIsPtr(2) {
			log.Error("")
			continue
		}

		if !methodChecker.IsFirstReturnValImplError() {
			log.Error("")
			continue
		}

		methods[rtype.Method(i).Name] = &methodType{
			method:    rtype.Method(i),
			ArgType:   rtype.In(2),
			ReplyType: rtype.In(3),
		}
	}
	return methods
}

func (s *Server) UnregisterService() {

}

// call
func (s *Service) call(ctx context.Context, mtype *methodType, argv, reply reflect.Value) (err error) {
	// func with receiver as first argument
	function := mtype.method.Func
	// Invoke the method, providing a new value for the reply
	returnValues := function.Call([]reflect.Value{s.receiver, reflect.ValueOf(ctx), argv, reply})
	// The return value for the method is an error
	errInter := returnValues[0].Interface()

	if errInter != nil {
		return errInter.(error)
	}

	return nil
}
