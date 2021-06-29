package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/xusworld/flash/codec"
	"github.com/xusworld/flash/config"
	"github.com/xusworld/flash/log"
	"github.com/xusworld/flash/protocol"
)

// tcpServer
type tcpServer struct {
	options    *Options
	ln         net.Listener
	serviceMap map[string]Service
	mutex      sync.RWMutex
}

// NewTcpServer create a tcp server
func NewTcpServer(optionFuncSet []OptionFunc) *tcpServer {
	server := &tcpServer{}

	// set server options
	server.options = &Options{}
	for _, optionFunc := range optionFuncSet {
		optionFunc(server.options)
	}

	// initialize
	server.serviceMap = make(map[string]Service, 0)

	return server
}

// Register
func (ts *tcpServer) Register(service Service) error {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	if ds, ok := service.(*defaultService); ok {
		ts.serviceMap[ds.name] = ds

	}

	return nil
}

// Address tcpServer address
func (ts *tcpServer) Address() net.Addr {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	return ts.ln.Addr()
}

// Serve run the tcpServer
func (ts *tcpServer) Serve() error {
	ln, err := net.Listen(ts.options.Network, ts.options.Address)

	if err != nil {
		fmt.Println("net.Listen error")
		return err
	}

	ts.ln = ln
	return ts.serve()
}

// serve
func (ts *tcpServer) serve() error {
	for {
		conn, err := ts.ln.Accept()
		if err != nil {
			fmt.Printf("flash: tcpServer linster error %s", err)
			return err
		}

		go func() {
			fmt.Println("connect accept. handle")
			err := tcpHandleFunc(ts, conn)
			if err != nil {
				fmt.Printf("flash: tcpHandleFunc() error %s", err)
			}
		}()
	}
}

// tcpHandleFunc
func tcpHandleFunc(server *tcpServer, conn net.Conn) error {
	defer conn.Close()

	buff := make([]byte, ReadBuffSize)
	num, err := conn.Read(buff)

	if err != nil {
		log.Debug(err.Error())
		return err
	}

	fmt.Println("Unmarshal request")
	req := &protocol.Request{}
	marshaler := codec.GetMarshaler(codec.Protobuf)
	if err := marshaler.Unmarshal(buff[:num], req); err != nil {
		fmt.Println("Unmarshal error, err", err)

	}

	fmt.Println("serviceMap", server.serviceMap)
	service := server.serviceMap[req.Header.ServicePath]

	if service == nil {
		return errors.New(
			fmt.Sprintf("%s: service %s not registered", config.Project, req.Header.ServicePath))
	}

	ds := service.(*defaultService)
	method := ds.methodMap[req.Header.ServiceMethod]

	argsType := method.Type.In(2)
	replyType := method.Type.In(3)
	request := reflect.New(argsType.Elem()).Interface()
	response := reflect.New(replyType.Elem()).Interface()

	err = marshaler.Unmarshal(req.Payload, request)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("request", reflect.ValueOf(response))
	_ = service.Call(req.Header.ServiceMethod, context.Background(), request, response)

	fmt.Println("response", reflect.ValueOf(response).Elem())


	buf, _ := marshaler.Marshal(response)
	typedResponse := &protocol.Response{
		Header:  nil,
		RetCode: 0,
		RetMsg:  "Hello,world!",
		Payload: buf,
	}

	fmt.Println("typedResponse", typedResponse)
	respBuff, _ := proto.Marshal(typedResponse)

	_, _ = conn.Write(respBuff)
	return nil
}
