// Code generated by counterfeiter. DO NOT EDIT.
package terraformfakes

import (
	"bytes"
	"sync"

	"github.com/EngineerBetter/control-tower/terraform"
)

type FakeOutputs struct {
	AssertValidStub        func() error
	assertValidMutex       sync.RWMutex
	assertValidArgsForCall []struct {
	}
	assertValidReturns struct {
		result1 error
	}
	assertValidReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string) (string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 string
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	InitStub        func(*bytes.Buffer) error
	initMutex       sync.RWMutex
	initArgsForCall []struct {
		arg1 *bytes.Buffer
	}
	initReturns struct {
		result1 error
	}
	initReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOutputs) AssertValid() error {
	fake.assertValidMutex.Lock()
	ret, specificReturn := fake.assertValidReturnsOnCall[len(fake.assertValidArgsForCall)]
	fake.assertValidArgsForCall = append(fake.assertValidArgsForCall, struct {
	}{})
	fake.recordInvocation("AssertValid", []interface{}{})
	fake.assertValidMutex.Unlock()
	if fake.AssertValidStub != nil {
		return fake.AssertValidStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.assertValidReturns
	return fakeReturns.result1
}

func (fake *FakeOutputs) AssertValidCallCount() int {
	fake.assertValidMutex.RLock()
	defer fake.assertValidMutex.RUnlock()
	return len(fake.assertValidArgsForCall)
}

func (fake *FakeOutputs) AssertValidCalls(stub func() error) {
	fake.assertValidMutex.Lock()
	defer fake.assertValidMutex.Unlock()
	fake.AssertValidStub = stub
}

func (fake *FakeOutputs) AssertValidReturns(result1 error) {
	fake.assertValidMutex.Lock()
	defer fake.assertValidMutex.Unlock()
	fake.AssertValidStub = nil
	fake.assertValidReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputs) AssertValidReturnsOnCall(i int, result1 error) {
	fake.assertValidMutex.Lock()
	defer fake.assertValidMutex.Unlock()
	fake.AssertValidStub = nil
	if fake.assertValidReturnsOnCall == nil {
		fake.assertValidReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.assertValidReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputs) Get(arg1 string) (string, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOutputs) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeOutputs) GetCalls(stub func(string) (string, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeOutputs) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOutputs) GetReturns(result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOutputs) GetReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOutputs) Init(arg1 *bytes.Buffer) error {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
		arg1 *bytes.Buffer
	}{arg1})
	fake.recordInvocation("Init", []interface{}{arg1})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1
}

func (fake *FakeOutputs) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeOutputs) InitCalls(stub func(*bytes.Buffer) error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeOutputs) InitArgsForCall(i int) *bytes.Buffer {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	argsForCall := fake.initArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOutputs) InitReturns(result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputs) InitReturnsOnCall(i int, result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputs) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assertValidMutex.RLock()
	defer fake.assertValidMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOutputs) recordInvocation(key string, args []interface{}) {
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

var _ terraform.Outputs = new(FakeOutputs)
