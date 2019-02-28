// Code generated by counterfeiter. DO NOT EDIT.
package flyfakes

import (
	"sync"

	"github.com/EngineerBetter/control-tower/config"
	"github.com/EngineerBetter/control-tower/fly"
)

type FakeIClient struct {
	CanConnectStub        func() (bool, error)
	canConnectMutex       sync.RWMutex
	canConnectArgsForCall []struct {
	}
	canConnectReturns struct {
		result1 bool
		result2 error
	}
	canConnectReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
	}
	cleanupReturns struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	SetDefaultPipelineStub        func(config.Config, bool) error
	setDefaultPipelineMutex       sync.RWMutex
	setDefaultPipelineArgsForCall []struct {
		arg1 config.Config
		arg2 bool
	}
	setDefaultPipelineReturns struct {
		result1 error
	}
	setDefaultPipelineReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIClient) CanConnect() (bool, error) {
	fake.canConnectMutex.Lock()
	ret, specificReturn := fake.canConnectReturnsOnCall[len(fake.canConnectArgsForCall)]
	fake.canConnectArgsForCall = append(fake.canConnectArgsForCall, struct {
	}{})
	fake.recordInvocation("CanConnect", []interface{}{})
	fake.canConnectMutex.Unlock()
	if fake.CanConnectStub != nil {
		return fake.CanConnectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.canConnectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIClient) CanConnectCallCount() int {
	fake.canConnectMutex.RLock()
	defer fake.canConnectMutex.RUnlock()
	return len(fake.canConnectArgsForCall)
}

func (fake *FakeIClient) CanConnectCalls(stub func() (bool, error)) {
	fake.canConnectMutex.Lock()
	defer fake.canConnectMutex.Unlock()
	fake.CanConnectStub = stub
}

func (fake *FakeIClient) CanConnectReturns(result1 bool, result2 error) {
	fake.canConnectMutex.Lock()
	defer fake.canConnectMutex.Unlock()
	fake.CanConnectStub = nil
	fake.canConnectReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) CanConnectReturnsOnCall(i int, result1 bool, result2 error) {
	fake.canConnectMutex.Lock()
	defer fake.canConnectMutex.Unlock()
	fake.CanConnectStub = nil
	if fake.canConnectReturnsOnCall == nil {
		fake.canConnectReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.canConnectReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Cleanup() error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
	}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cleanupReturns
	return fakeReturns.result1
}

func (fake *FakeIClient) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeIClient) CleanupCalls(stub func() error) {
	fake.cleanupMutex.Lock()
	defer fake.cleanupMutex.Unlock()
	fake.CleanupStub = stub
}

func (fake *FakeIClient) CleanupReturns(result1 error) {
	fake.cleanupMutex.Lock()
	defer fake.cleanupMutex.Unlock()
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) CleanupReturnsOnCall(i int, result1 error) {
	fake.cleanupMutex.Lock()
	defer fake.cleanupMutex.Unlock()
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) SetDefaultPipeline(arg1 config.Config, arg2 bool) error {
	fake.setDefaultPipelineMutex.Lock()
	ret, specificReturn := fake.setDefaultPipelineReturnsOnCall[len(fake.setDefaultPipelineArgsForCall)]
	fake.setDefaultPipelineArgsForCall = append(fake.setDefaultPipelineArgsForCall, struct {
		arg1 config.Config
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("SetDefaultPipeline", []interface{}{arg1, arg2})
	fake.setDefaultPipelineMutex.Unlock()
	if fake.SetDefaultPipelineStub != nil {
		return fake.SetDefaultPipelineStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setDefaultPipelineReturns
	return fakeReturns.result1
}

func (fake *FakeIClient) SetDefaultPipelineCallCount() int {
	fake.setDefaultPipelineMutex.RLock()
	defer fake.setDefaultPipelineMutex.RUnlock()
	return len(fake.setDefaultPipelineArgsForCall)
}

func (fake *FakeIClient) SetDefaultPipelineCalls(stub func(config.Config, bool) error) {
	fake.setDefaultPipelineMutex.Lock()
	defer fake.setDefaultPipelineMutex.Unlock()
	fake.SetDefaultPipelineStub = stub
}

func (fake *FakeIClient) SetDefaultPipelineArgsForCall(i int) (config.Config, bool) {
	fake.setDefaultPipelineMutex.RLock()
	defer fake.setDefaultPipelineMutex.RUnlock()
	argsForCall := fake.setDefaultPipelineArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIClient) SetDefaultPipelineReturns(result1 error) {
	fake.setDefaultPipelineMutex.Lock()
	defer fake.setDefaultPipelineMutex.Unlock()
	fake.SetDefaultPipelineStub = nil
	fake.setDefaultPipelineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) SetDefaultPipelineReturnsOnCall(i int, result1 error) {
	fake.setDefaultPipelineMutex.Lock()
	defer fake.setDefaultPipelineMutex.Unlock()
	fake.SetDefaultPipelineStub = nil
	if fake.setDefaultPipelineReturnsOnCall == nil {
		fake.setDefaultPipelineReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setDefaultPipelineReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.canConnectMutex.RLock()
	defer fake.canConnectMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.setDefaultPipelineMutex.RLock()
	defer fake.setDefaultPipelineMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIClient) recordInvocation(key string, args []interface{}) {
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

var _ fly.IClient = new(FakeIClient)
