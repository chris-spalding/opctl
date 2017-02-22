// This file was generated by counterfeiter
package vfs

import (
	"os"
	"sync"
)

type Fake struct {
	CreateStub        func(name string) (*os.File, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name string
	}
	createReturns struct {
		result1 *os.File
		result2 error
	}
	MkdirAllStub        func(path string, perm os.FileMode) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		path string
		perm os.FileMode
	}
	mkdirAllReturns struct {
		result1 error
	}
	OpenStub        func(name string) (*os.File, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		name string
	}
	openReturns struct {
		result1 *os.File
		result2 error
	}
	RemoveAllStub        func(path string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		path string
	}
	removeAllReturns struct {
		result1 error
	}
	StatStub        func(name string) (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		name string
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Create(name string) (*os.File, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Create", []interface{}{name})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *Fake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Fake) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name
}

func (fake *Fake) CreateReturns(result1 *os.File, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) MkdirAll(path string, perm os.FileMode) error {
	fake.mkdirAllMutex.Lock()
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		path string
		perm os.FileMode
	}{path, perm})
	fake.recordInvocation("MkdirAll", []interface{}{path, perm})
	fake.mkdirAllMutex.Unlock()
	if fake.MkdirAllStub != nil {
		return fake.MkdirAllStub(path, perm)
	} else {
		return fake.mkdirAllReturns.result1
	}
}

func (fake *Fake) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *Fake) MkdirAllArgsForCall(i int) (string, os.FileMode) {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return fake.mkdirAllArgsForCall[i].path, fake.mkdirAllArgsForCall[i].perm
}

func (fake *Fake) MkdirAllReturns(result1 error) {
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Open(name string) (*os.File, error) {
	fake.openMutex.Lock()
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Open", []interface{}{name})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(name)
	} else {
		return fake.openReturns.result1, fake.openReturns.result2
	}
}

func (fake *Fake) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *Fake) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].name
}

func (fake *Fake) OpenReturns(result1 *os.File, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) RemoveAll(path string) error {
	fake.removeAllMutex.Lock()
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("RemoveAll", []interface{}{path})
	fake.removeAllMutex.Unlock()
	if fake.RemoveAllStub != nil {
		return fake.RemoveAllStub(path)
	} else {
		return fake.removeAllReturns.result1
	}
}

func (fake *Fake) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *Fake) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return fake.removeAllArgsForCall[i].path
}

func (fake *Fake) RemoveAllReturns(result1 error) {
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Stat(name string) (os.FileInfo, error) {
	fake.statMutex.Lock()
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Stat", []interface{}{name})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub(name)
	} else {
		return fake.statReturns.result1, fake.statReturns.result2
	}
}

func (fake *Fake) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *Fake) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.statArgsForCall[i].name
}

func (fake *Fake) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.invocations
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

var _ Vfs = new(Fake)
