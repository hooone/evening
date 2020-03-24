package bus

import (
	"context"
	"errors"
	"reflect"
)

// HandlerFunc defines a handler function interface.
type HandlerFunc interface{}

// CtxHandlerFunc defines a context handler function.
type CtxHandlerFunc func()

// Msg defines a message interface.
type Msg interface{}

// ErrHandlerNotFound defines an error if a handler is not found
var ErrHandlerNotFound = errors.New("handler not found")

// Bus type defines the bus interface structure
type Bus interface {
	Dispatch(msg Msg) error
	DispatchCtx(ctx context.Context, msg Msg) error

	AddHandler(handler HandlerFunc)
	AddHandlerCtx(handler HandlerFunc)
}

// InProcBus defines the bus structure
type InProcBus struct {
	handlers        map[string]HandlerFunc
	handlersWithCtx map[string]HandlerFunc
	listeners       map[string][]HandlerFunc
}

// temp stuff, not sure how to handle bus instance, and init yet
var globalBus = New()

// New initialize the bus
func New() Bus {
	bus := &InProcBus{}
	bus.handlers = make(map[string]HandlerFunc)
	bus.handlersWithCtx = make(map[string]HandlerFunc)
	bus.listeners = make(map[string][]HandlerFunc)

	return bus
}

// GetBus Want to get rid of global bus
func GetBus() Bus {
	return globalBus
}

// DispatchCtx function dispatch a message to the bus context.
func (b *InProcBus) DispatchCtx(ctx context.Context, msg Msg) error {
	var msgName = reflect.TypeOf(msg).Elem().Name()

	var handler = b.handlersWithCtx[msgName]
	if handler == nil {
		return ErrHandlerNotFound
	}

	var params = []reflect.Value{}
	params = append(params, reflect.ValueOf(ctx))
	params = append(params, reflect.ValueOf(msg))

	ret := reflect.ValueOf(handler).Call(params)
	if len(ret) == 0 {
		return nil
	}
	err := ret[0].Interface()
	if err == nil {
		return nil
	}
	return err.(error)
}

// Dispatch function dispatch a message to the bus.
func (b *InProcBus) Dispatch(msg Msg) error {
	var msgName = reflect.TypeOf(msg).Elem().Name()

	var handler = b.handlersWithCtx[msgName]
	withCtx := true

	if handler == nil {
		withCtx = false
		handler = b.handlers[msgName]
	}

	if handler == nil {
		return ErrHandlerNotFound
	}

	var params = []reflect.Value{}
	if withCtx {
		params = append(params, reflect.ValueOf(context.Background()))
	}
	params = append(params, reflect.ValueOf(msg))

	ret := reflect.ValueOf(handler).Call(params)
	if len(ret) == 0 {
		return nil
	}
	err := ret[0].Interface()
	if err == nil {
		return nil
	}
	return err.(error)
}

//AddHandler register handler
func (b *InProcBus) AddHandler(handler HandlerFunc) {
	handlerType := reflect.TypeOf(handler)
	queryTypeName := handlerType.In(0).Elem().Name()
	b.handlers[queryTypeName] = handler
}

//AddHandlerCtx register handler with ctx
func (b *InProcBus) AddHandlerCtx(handler HandlerFunc) {
	handlerType := reflect.TypeOf(handler)
	queryTypeName := handlerType.In(1).Elem().Name()
	b.handlersWithCtx[queryTypeName] = handler
}

// AddHandler attach a handler function to the global bus
// Package level function
// implName no use
func AddHandler(implName string, handler HandlerFunc) {
	globalBus.AddHandler(handler)
}

// AddHandlerCtx attach a handler function to the global bus context
// Package level functions
// implName no use
func AddHandlerCtx(implName string, handler HandlerFunc) {
	globalBus.AddHandlerCtx(handler)
}

// Dispatch function dispatch a message to the global bus.
func Dispatch(msg Msg) error {
	return globalBus.Dispatch(msg)
}

// DispatchCtx function dispatch a message to the global bus context.
func DispatchCtx(ctx context.Context, msg Msg) error {
	return globalBus.DispatchCtx(ctx, msg)
}

// ClearBusHandlers close all handlers of the global bus
func ClearBusHandlers() {
	globalBus = New()
}
