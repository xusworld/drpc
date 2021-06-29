package server

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/xusworld/flash/config"
	"github.com/xusworld/flash/util"
)

// Service interface
type Service interface {
	Register(interface{}) error

	Call(name string, ctx context.Context, req, rsp interface{}) error

	Unregister() error
}

const (
	Running = iota
	Stopped
	Error
)

// defaultService impl
type defaultService struct {
	// defaultService name
	name string

	// pointer receiver
	receiver reflect.Value

	// reflect type
	typ reflect.Type

	// method map
	methodMap map[string]reflect.Method

	// defaultService status
	status int

	// mutex
	mutex sync.Mutex
}

// Register
// TODO not thread-safe
func (s *defaultService) Register(receiver interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.receiver = reflect.ValueOf(receiver)
	s.name = reflect.Indirect(s.receiver).Type().Name()
	s.typ = reflect.TypeOf(receiver)
	s.methodMap = make(map[string]reflect.Method, 0)

	for i := 0; i < s.typ.NumMethod(); i++ {
		typ := s.typ
		methodType := typ.Method(i).Type

		// filter methodMap
		if !util.IsParamNumEqualsToN(methodType, 4) {
			fmt.Println("continue IsParamNumEqualsToN")
			continue
		}

		// function checker
		if !util.IsNthParamImplContext(methodType, 1) {
			fmt.Println("continue IsNthParamImplContext")
			continue
		}

		s.methodMap[typ.Method(i).Name] = typ.Method(i)
	}

	if len(s.methodMap) == 0 {
		return errors.New(config.Project + ": register error")
	}

	s.status = Running
	return nil
}

// Call
// TODO not thread-safe
func (s *defaultService) Call(caller string, ctx context.Context, req, rsp interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	fmt.Println("in")
	if s.status != Running {
		return errors.New(config.Project + ":defaultService is not registered yet")
	}

	function := s.methodMap[caller]
	fmt.Println("Call")
	output := function.Func.Call([]reflect.Value{s.receiver, reflect.ValueOf(ctx),
		reflect.ValueOf(req), reflect.ValueOf(rsp)})

	errInter := output[0].Interface()

	if errInter != nil {
		return errInter.(error)
	}

	return nil
}

// Unregister
// TODO not thread-safe
func (s *defaultService) Unregister() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.status == Running {
		return errors.New(config.Project + ":defaultService is not registered yet")
	}

	s.methodMap = make(map[string]reflect.Method, 0)

	s.status = Stopped
	return nil
}
