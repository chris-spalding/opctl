// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"context"
	"sync"

	"github.com/opctl/sdk-golang/model"
)

type fakeCaller struct {
	CallStub        func(context.Context, string, map[string]*model.Value, *model.SCG, model.DataHandle, *string, string)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]*model.Value
		arg4 *model.SCG
		arg5 model.DataHandle
		arg6 *string
		arg7 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCaller) Call(arg1 context.Context, arg2 string, arg3 map[string]*model.Value, arg4 *model.SCG, arg5 model.DataHandle, arg6 *string, arg7 string) {
	fake.callMutex.Lock()
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]*model.Value
		arg4 *model.SCG
		arg5 model.DataHandle
		arg6 *string
		arg7 string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Call", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		fake.CallStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

func (fake *fakeCaller) CallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeCaller) Calls(stub func(context.Context, string, map[string]*model.Value, *model.SCG, model.DataHandle, *string, string)) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = stub
}

func (fake *fakeCaller) CallArgsForCall(i int) (context.Context, string, map[string]*model.Value, *model.SCG, model.DataHandle, *string, string) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	argsForCall := fake.callArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *fakeCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeCaller) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
