// Code generated by counterfeiter. DO NOT EDIT.
package inputs

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	ValidateStub        func(inputs map[string]*model.Data, params map[string]*model.Param) map[string][]error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		inputs map[string]*model.Data
		params map[string]*model.Param
	}
	validateReturns struct {
		result1 map[string][]error
	}
	validateReturnsOnCall map[int]struct {
		result1 map[string][]error
	}
	InterpretStub        func(inputArgs map[string]string, inputParams map[string]*model.Param, scope map[string]*model.Data) (map[string]*model.Data, []error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		inputArgs   map[string]string
		inputParams map[string]*model.Param
		scope       map[string]*model.Data
	}
	interpretReturns struct {
		result1 map[string]*model.Data
		result2 []error
	}
	interpretReturnsOnCall map[int]struct {
		result1 map[string]*model.Data
		result2 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Validate(inputs map[string]*model.Data, params map[string]*model.Param) map[string][]error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		inputs map[string]*model.Data
		params map[string]*model.Param
	}{inputs, params})
	fake.recordInvocation("Validate", []interface{}{inputs, params})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(inputs, params)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *Fake) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *Fake) ValidateArgsForCall(i int) (map[string]*model.Data, map[string]*model.Param) {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].inputs, fake.validateArgsForCall[i].params
}

func (fake *Fake) ValidateReturns(result1 map[string][]error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 map[string][]error
	}{result1}
}

func (fake *Fake) ValidateReturnsOnCall(i int, result1 map[string][]error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 map[string][]error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 map[string][]error
	}{result1}
}

func (fake *Fake) Interpret(inputArgs map[string]string, inputParams map[string]*model.Param, scope map[string]*model.Data) (map[string]*model.Data, []error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		inputArgs   map[string]string
		inputParams map[string]*model.Param
		scope       map[string]*model.Data
	}{inputArgs, inputParams, scope})
	fake.recordInvocation("Interpret", []interface{}{inputArgs, inputParams, scope})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(inputArgs, inputParams, scope)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *Fake) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *Fake) InterpretArgsForCall(i int) (map[string]string, map[string]*model.Param, map[string]*model.Data) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].inputArgs, fake.interpretArgsForCall[i].inputParams, fake.interpretArgsForCall[i].scope
}

func (fake *Fake) InterpretReturns(result1 map[string]*model.Data, result2 []error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 map[string]*model.Data
		result2 []error
	}{result1, result2}
}

func (fake *Fake) InterpretReturnsOnCall(i int, result1 map[string]*model.Data, result2 []error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 map[string]*model.Data
			result2 []error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 map[string]*model.Data
		result2 []error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Inputs = new(Fake)
