package errors

import (
	"encoding/json"
	"runtime"
	"time"
)

const (
	// RuntimeCallerDeepOfSelf runtime call deep of self
	RuntimeCallerDeepOfSelf = 0
	// RuntimeCallerDeepOfParent runtime call deep of parent
	RuntimeCallerDeepOfParent = 1
	// RuntimeCallerDeepOfGrandparent runtime call deep of grandparent
	RuntimeCallerDeepOfGrandparent = 2
)

// Error traceable errors
type Error interface {
	Wrap(*ErrorEntry)
	Trace() []*ErrorEntry
	error
}

// ErrorEntry detail error info
type ErrorEntry struct {
	FileName string
	FileLine int
	FuncName string
	Time     time.Time
	Desc     string
}

// ErrorInstance map of ErrorEntry
type ErrorInstance struct {
	ErrorTrace []*ErrorEntry
}

// NewError generate ErrorInstance
func NewError() *ErrorInstance {
	errorInstance := new(ErrorInstance)
	return errorInstance
}

// NewErrorEntry generate ErrorEntry
func NewErrorEntry(deep int, desc string) *ErrorEntry {
	funcPtr, file, line, _ := runtime.Caller(deep)
	funcInfo := runtime.FuncForPC(funcPtr)
	errorItem := new(ErrorEntry)
	errorItem.FileName = file
	errorItem.FileLine = line
	errorItem.FuncName = funcInfo.Name()
	errorItem.Time = time.Now()
	errorItem.Desc = desc
	return errorItem
}

// Wrap wrap error info
func (e *ErrorInstance) Wrap(err *ErrorEntry) {
	e.ErrorTrace = append(e.ErrorTrace, err)
}

// Trace trace error
func (e *ErrorInstance) Trace() []*ErrorEntry {
	return e.ErrorTrace
}

// Error input error info
func (e *ErrorInstance) Error() string {
	jsonError, err := json.Marshal(e.ErrorTrace)
	if err != nil {
		return ""
	}
	return string(jsonError)
}

// Trace wrap trace error
func Trace(err error, desc string) error {
	errorItem := NewErrorEntry(RuntimeCallerDeepOfGrandparent, desc)
	switch err := err.(type) {
	case Error:
		err.Wrap(errorItem)
		return err
	default:
		errorInstance := NewError()
		errorInstance.ErrorTrace = append(errorInstance.ErrorTrace, errorItem)
		return errorInstance
	}
}
