// Code generated by counterfeiter. DO NOT EDIT.
package boshfakes

import (
	"sync"

	"github.com/EngineerBetter/control-tower/bosh"
)

type FakeIClient struct {
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
	CreateEnvStub        func([]byte, []byte, string) ([]byte, []byte, error)
	createEnvMutex       sync.RWMutex
	createEnvArgsForCall []struct {
		arg1 []byte
		arg2 []byte
		arg3 string
	}
	createEnvReturns struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	createEnvReturnsOnCall map[int]struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	DeleteStub        func([]byte) ([]byte, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []byte
	}
	deleteReturns struct {
		result1 []byte
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	DeployStub        func([]byte, []byte, bool) ([]byte, []byte, error)
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		arg1 []byte
		arg2 []byte
		arg3 bool
	}
	deployReturns struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	deployReturnsOnCall map[int]struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	InstancesStub        func() ([]bosh.Instance, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct {
	}
	instancesReturns struct {
		result1 []bosh.Instance
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []bosh.Instance
		result2 error
	}
	LocksStub        func() ([]byte, error)
	locksMutex       sync.RWMutex
	locksArgsForCall []struct {
	}
	locksReturns struct {
		result1 []byte
		result2 error
	}
	locksReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RecreateStub        func() error
	recreateMutex       sync.RWMutex
	recreateArgsForCall []struct {
	}
	recreateReturns struct {
		result1 error
	}
	recreateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
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

func (fake *FakeIClient) CreateEnv(arg1 []byte, arg2 []byte, arg3 string) ([]byte, []byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createEnvMutex.Lock()
	ret, specificReturn := fake.createEnvReturnsOnCall[len(fake.createEnvArgsForCall)]
	fake.createEnvArgsForCall = append(fake.createEnvArgsForCall, struct {
		arg1 []byte
		arg2 []byte
		arg3 string
	}{arg1Copy, arg2Copy, arg3})
	fake.recordInvocation("CreateEnv", []interface{}{arg1Copy, arg2Copy, arg3})
	fake.createEnvMutex.Unlock()
	if fake.CreateEnvStub != nil {
		return fake.CreateEnvStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createEnvReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeIClient) CreateEnvCallCount() int {
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	return len(fake.createEnvArgsForCall)
}

func (fake *FakeIClient) CreateEnvCalls(stub func([]byte, []byte, string) ([]byte, []byte, error)) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = stub
}

func (fake *FakeIClient) CreateEnvArgsForCall(i int) ([]byte, []byte, string) {
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	argsForCall := fake.createEnvArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIClient) CreateEnvReturns(result1 []byte, result2 []byte, result3 error) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = nil
	fake.createEnvReturns = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIClient) CreateEnvReturnsOnCall(i int, result1 []byte, result2 []byte, result3 error) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = nil
	if fake.createEnvReturnsOnCall == nil {
		fake.createEnvReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 []byte
			result3 error
		})
	}
	fake.createEnvReturnsOnCall[i] = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIClient) Delete(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Delete", []interface{}{arg1Copy})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeIClient) DeleteCalls(stub func([]byte) ([]byte, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeIClient) DeleteArgsForCall(i int) []byte {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIClient) DeleteReturns(result1 []byte, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) DeleteReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Deploy(arg1 []byte, arg2 []byte, arg3 bool) ([]byte, []byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		arg1 []byte
		arg2 []byte
		arg3 bool
	}{arg1Copy, arg2Copy, arg3})
	fake.recordInvocation("Deploy", []interface{}{arg1Copy, arg2Copy, arg3})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.deployReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeIClient) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeIClient) DeployCalls(stub func([]byte, []byte, bool) ([]byte, []byte, error)) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = stub
}

func (fake *FakeIClient) DeployArgsForCall(i int) ([]byte, []byte, bool) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	argsForCall := fake.deployArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIClient) DeployReturns(result1 []byte, result2 []byte, result3 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIClient) DeployReturnsOnCall(i int, result1 []byte, result2 []byte, result3 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 []byte
			result3 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIClient) Instances() ([]bosh.Instance, error) {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct {
	}{})
	fake.recordInvocation("Instances", []interface{}{})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIClient) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeIClient) InstancesCalls(stub func() ([]bosh.Instance, error)) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = stub
}

func (fake *FakeIClient) InstancesReturns(result1 []bosh.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []bosh.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) InstancesReturnsOnCall(i int, result1 []bosh.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []bosh.Instance
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []bosh.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Locks() ([]byte, error) {
	fake.locksMutex.Lock()
	ret, specificReturn := fake.locksReturnsOnCall[len(fake.locksArgsForCall)]
	fake.locksArgsForCall = append(fake.locksArgsForCall, struct {
	}{})
	fake.recordInvocation("Locks", []interface{}{})
	fake.locksMutex.Unlock()
	if fake.LocksStub != nil {
		return fake.LocksStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.locksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIClient) LocksCallCount() int {
	fake.locksMutex.RLock()
	defer fake.locksMutex.RUnlock()
	return len(fake.locksArgsForCall)
}

func (fake *FakeIClient) LocksCalls(stub func() ([]byte, error)) {
	fake.locksMutex.Lock()
	defer fake.locksMutex.Unlock()
	fake.LocksStub = stub
}

func (fake *FakeIClient) LocksReturns(result1 []byte, result2 error) {
	fake.locksMutex.Lock()
	defer fake.locksMutex.Unlock()
	fake.LocksStub = nil
	fake.locksReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) LocksReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.locksMutex.Lock()
	defer fake.locksMutex.Unlock()
	fake.LocksStub = nil
	if fake.locksReturnsOnCall == nil {
		fake.locksReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.locksReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeIClient) Recreate() error {
	fake.recreateMutex.Lock()
	ret, specificReturn := fake.recreateReturnsOnCall[len(fake.recreateArgsForCall)]
	fake.recreateArgsForCall = append(fake.recreateArgsForCall, struct {
	}{})
	fake.recordInvocation("Recreate", []interface{}{})
	fake.recreateMutex.Unlock()
	if fake.RecreateStub != nil {
		return fake.RecreateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.recreateReturns
	return fakeReturns.result1
}

func (fake *FakeIClient) RecreateCallCount() int {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	return len(fake.recreateArgsForCall)
}

func (fake *FakeIClient) RecreateCalls(stub func() error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = stub
}

func (fake *FakeIClient) RecreateReturns(result1 error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = nil
	fake.recreateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) RecreateReturnsOnCall(i int, result1 error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = nil
	if fake.recreateReturnsOnCall == nil {
		fake.recreateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recreateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	fake.locksMutex.RLock()
	defer fake.locksMutex.RUnlock()
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
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

var _ bosh.IClient = new(FakeIClient)
