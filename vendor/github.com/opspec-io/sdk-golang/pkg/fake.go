// Code generated by counterfeiter. DO NOT EDIT.
package pkg

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	CreateStub        func(path, pkgName, pkgDescription string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		path           string
		pkgName        string
		pkgDescription string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	ResolveStub        func(pkgRef *PkgRef, lookPaths ...string) (string, bool)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		pkgRef    *PkgRef
		lookPaths []string
	}
	resolveReturns struct {
		result1 string
		result2 bool
	}
	resolveReturnsOnCall map[int]struct {
		result1 string
		result2 bool
	}
	ParseRefStub        func(pkgRef string) (*PkgRef, error)
	parseRefMutex       sync.RWMutex
	parseRefArgsForCall []struct {
		pkgRef string
	}
	parseRefReturns struct {
		result1 *PkgRef
		result2 error
	}
	parseRefReturnsOnCall map[int]struct {
		result1 *PkgRef
		result2 error
	}
	PullStub        func(path string, pkgRef *PkgRef, opts *PullOpts) error
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		path   string
		pkgRef *PkgRef
		opts   *PullOpts
	}
	pullReturns struct {
		result1 error
	}
	pullReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(dirPath string) ([]*model.PkgManifest, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		dirPath string
	}
	listReturns struct {
		result1 []*model.PkgManifest
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*model.PkgManifest
		result2 error
	}
	SetDescriptionStub        func(pkgPath, pkgDescription string) error
	setDescriptionMutex       sync.RWMutex
	setDescriptionArgsForCall []struct {
		pkgPath        string
		pkgDescription string
	}
	setDescriptionReturns struct {
		result1 error
	}
	setDescriptionReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateStub        func(pkgPath string) []error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		pkgPath string
	}
	validateReturns struct {
		result1 []error
	}
	validateReturnsOnCall map[int]struct {
		result1 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Create(path string, pkgName string, pkgDescription string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		path           string
		pkgName        string
		pkgDescription string
	}{path, pkgName, pkgDescription})
	fake.recordInvocation("Create", []interface{}{path, pkgName, pkgDescription})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(path, pkgName, pkgDescription)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *Fake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Fake) CreateArgsForCall(i int) (string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].path, fake.createArgsForCall[i].pkgName, fake.createArgsForCall[i].pkgDescription
}

func (fake *Fake) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Resolve(pkgRef *PkgRef, lookPaths ...string) (string, bool) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		pkgRef    *PkgRef
		lookPaths []string
	}{pkgRef, lookPaths})
	fake.recordInvocation("Resolve", []interface{}{pkgRef, lookPaths})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(pkgRef, lookPaths...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resolveReturns.result1, fake.resolveReturns.result2
}

func (fake *Fake) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *Fake) ResolveArgsForCall(i int) (*PkgRef, []string) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return fake.resolveArgsForCall[i].pkgRef, fake.resolveArgsForCall[i].lookPaths
}

func (fake *Fake) ResolveReturns(result1 string, result2 bool) {
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *Fake) ResolveReturnsOnCall(i int, result1 string, result2 bool) {
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *Fake) ParseRef(pkgRef string) (*PkgRef, error) {
	fake.parseRefMutex.Lock()
	ret, specificReturn := fake.parseRefReturnsOnCall[len(fake.parseRefArgsForCall)]
	fake.parseRefArgsForCall = append(fake.parseRefArgsForCall, struct {
		pkgRef string
	}{pkgRef})
	fake.recordInvocation("ParseRef", []interface{}{pkgRef})
	fake.parseRefMutex.Unlock()
	if fake.ParseRefStub != nil {
		return fake.ParseRefStub(pkgRef)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.parseRefReturns.result1, fake.parseRefReturns.result2
}

func (fake *Fake) ParseRefCallCount() int {
	fake.parseRefMutex.RLock()
	defer fake.parseRefMutex.RUnlock()
	return len(fake.parseRefArgsForCall)
}

func (fake *Fake) ParseRefArgsForCall(i int) string {
	fake.parseRefMutex.RLock()
	defer fake.parseRefMutex.RUnlock()
	return fake.parseRefArgsForCall[i].pkgRef
}

func (fake *Fake) ParseRefReturns(result1 *PkgRef, result2 error) {
	fake.ParseRefStub = nil
	fake.parseRefReturns = struct {
		result1 *PkgRef
		result2 error
	}{result1, result2}
}

func (fake *Fake) ParseRefReturnsOnCall(i int, result1 *PkgRef, result2 error) {
	fake.ParseRefStub = nil
	if fake.parseRefReturnsOnCall == nil {
		fake.parseRefReturnsOnCall = make(map[int]struct {
			result1 *PkgRef
			result2 error
		})
	}
	fake.parseRefReturnsOnCall[i] = struct {
		result1 *PkgRef
		result2 error
	}{result1, result2}
}

func (fake *Fake) Pull(path string, pkgRef *PkgRef, opts *PullOpts) error {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		path   string
		pkgRef *PkgRef
		opts   *PullOpts
	}{path, pkgRef, opts})
	fake.recordInvocation("Pull", []interface{}{path, pkgRef, opts})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(path, pkgRef, opts)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pullReturns.result1
}

func (fake *Fake) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *Fake) PullArgsForCall(i int) (string, *PkgRef, *PullOpts) {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return fake.pullArgsForCall[i].path, fake.pullArgsForCall[i].pkgRef, fake.pullArgsForCall[i].opts
}

func (fake *Fake) PullReturns(result1 error) {
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) PullReturnsOnCall(i int, result1 error) {
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) List(dirPath string) ([]*model.PkgManifest, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		dirPath string
	}{dirPath})
	fake.recordInvocation("List", []interface{}{dirPath})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(dirPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *Fake) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *Fake) ListArgsForCall(i int) string {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].dirPath
}

func (fake *Fake) ListReturns(result1 []*model.PkgManifest, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) ListReturnsOnCall(i int, result1 []*model.PkgManifest, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*model.PkgManifest
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) SetDescription(pkgPath string, pkgDescription string) error {
	fake.setDescriptionMutex.Lock()
	ret, specificReturn := fake.setDescriptionReturnsOnCall[len(fake.setDescriptionArgsForCall)]
	fake.setDescriptionArgsForCall = append(fake.setDescriptionArgsForCall, struct {
		pkgPath        string
		pkgDescription string
	}{pkgPath, pkgDescription})
	fake.recordInvocation("SetDescription", []interface{}{pkgPath, pkgDescription})
	fake.setDescriptionMutex.Unlock()
	if fake.SetDescriptionStub != nil {
		return fake.SetDescriptionStub(pkgPath, pkgDescription)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setDescriptionReturns.result1
}

func (fake *Fake) SetDescriptionCallCount() int {
	fake.setDescriptionMutex.RLock()
	defer fake.setDescriptionMutex.RUnlock()
	return len(fake.setDescriptionArgsForCall)
}

func (fake *Fake) SetDescriptionArgsForCall(i int) (string, string) {
	fake.setDescriptionMutex.RLock()
	defer fake.setDescriptionMutex.RUnlock()
	return fake.setDescriptionArgsForCall[i].pkgPath, fake.setDescriptionArgsForCall[i].pkgDescription
}

func (fake *Fake) SetDescriptionReturns(result1 error) {
	fake.SetDescriptionStub = nil
	fake.setDescriptionReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) SetDescriptionReturnsOnCall(i int, result1 error) {
	fake.SetDescriptionStub = nil
	if fake.setDescriptionReturnsOnCall == nil {
		fake.setDescriptionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setDescriptionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Validate(pkgPath string) []error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		pkgPath string
	}{pkgPath})
	fake.recordInvocation("Validate", []interface{}{pkgPath})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(pkgPath)
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

func (fake *Fake) ValidateArgsForCall(i int) string {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].pkgPath
}

func (fake *Fake) ValidateReturns(result1 []error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) ValidateReturnsOnCall(i int, result1 []error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 []error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	fake.parseRefMutex.RLock()
	defer fake.parseRefMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.setDescriptionMutex.RLock()
	defer fake.setDescriptionMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
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

var _ Pkg = new(Fake)
