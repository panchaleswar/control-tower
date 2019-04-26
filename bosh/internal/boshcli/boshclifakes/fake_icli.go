// Code generated by counterfeiter. DO NOT EDIT.
package boshclifakes

import (
	"io"
	"sync"

	"github.com/EngineerBetter/control-tower/bosh/internal/boshcli"
)

type FakeICLI struct {
	CreateEnvStub        func(boshcli.Store, boshcli.IAASEnvironment, string, string, string, string, map[string]string) error
	createEnvMutex       sync.RWMutex
	createEnvArgsForCall []struct {
		arg1 boshcli.Store
		arg2 boshcli.IAASEnvironment
		arg3 string
		arg4 string
		arg5 string
		arg6 string
		arg7 map[string]string
	}
	createEnvReturns struct {
		result1 error
	}
	createEnvReturnsOnCall map[int]struct {
		result1 error
	}
	LocksStub        func(boshcli.IAASEnvironment, string, string, string) ([]byte, error)
	locksMutex       sync.RWMutex
	locksArgsForCall []struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}
	locksReturns struct {
		result1 []byte
		result2 error
	}
	locksReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RecreateStub        func(boshcli.IAASEnvironment, string, string, string) error
	recreateMutex       sync.RWMutex
	recreateArgsForCall []struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}
	recreateReturns struct {
		result1 error
	}
	recreateReturnsOnCall map[int]struct {
		result1 error
	}
	RunAuthenticatedCommandStub        func(string, string, string, string, bool, io.Writer, ...string) error
	runAuthenticatedCommandMutex       sync.RWMutex
	runAuthenticatedCommandArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
		arg6 io.Writer
		arg7 []string
	}
	runAuthenticatedCommandReturns struct {
		result1 error
	}
	runAuthenticatedCommandReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateCloudConfigStub        func(boshcli.IAASEnvironment, string, string, string) error
	updateCloudConfigMutex       sync.RWMutex
	updateCloudConfigArgsForCall []struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}
	updateCloudConfigReturns struct {
		result1 error
	}
	updateCloudConfigReturnsOnCall map[int]struct {
		result1 error
	}
	UploadConcourseStemcellStub        func(boshcli.IAASEnvironment, string, string, string) error
	uploadConcourseStemcellMutex       sync.RWMutex
	uploadConcourseStemcellArgsForCall []struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}
	uploadConcourseStemcellReturns struct {
		result1 error
	}
	uploadConcourseStemcellReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeICLI) CreateEnv(arg1 boshcli.Store, arg2 boshcli.IAASEnvironment, arg3 string, arg4 string, arg5 string, arg6 string, arg7 map[string]string) error {
	fake.createEnvMutex.Lock()
	ret, specificReturn := fake.createEnvReturnsOnCall[len(fake.createEnvArgsForCall)]
	fake.createEnvArgsForCall = append(fake.createEnvArgsForCall, struct {
		arg1 boshcli.Store
		arg2 boshcli.IAASEnvironment
		arg3 string
		arg4 string
		arg5 string
		arg6 string
		arg7 map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("CreateEnv", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.createEnvMutex.Unlock()
	if fake.CreateEnvStub != nil {
		return fake.CreateEnvStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createEnvReturns
	return fakeReturns.result1
}

func (fake *FakeICLI) CreateEnvCallCount() int {
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	return len(fake.createEnvArgsForCall)
}

func (fake *FakeICLI) CreateEnvCalls(stub func(boshcli.Store, boshcli.IAASEnvironment, string, string, string, string, map[string]string) error) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = stub
}

func (fake *FakeICLI) CreateEnvArgsForCall(i int) (boshcli.Store, boshcli.IAASEnvironment, string, string, string, string, map[string]string) {
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	argsForCall := fake.createEnvArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeICLI) CreateEnvReturns(result1 error) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = nil
	fake.createEnvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) CreateEnvReturnsOnCall(i int, result1 error) {
	fake.createEnvMutex.Lock()
	defer fake.createEnvMutex.Unlock()
	fake.CreateEnvStub = nil
	if fake.createEnvReturnsOnCall == nil {
		fake.createEnvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createEnvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) Locks(arg1 boshcli.IAASEnvironment, arg2 string, arg3 string, arg4 string) ([]byte, error) {
	fake.locksMutex.Lock()
	ret, specificReturn := fake.locksReturnsOnCall[len(fake.locksArgsForCall)]
	fake.locksArgsForCall = append(fake.locksArgsForCall, struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Locks", []interface{}{arg1, arg2, arg3, arg4})
	fake.locksMutex.Unlock()
	if fake.LocksStub != nil {
		return fake.LocksStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.locksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeICLI) LocksCallCount() int {
	fake.locksMutex.RLock()
	defer fake.locksMutex.RUnlock()
	return len(fake.locksArgsForCall)
}

func (fake *FakeICLI) LocksCalls(stub func(boshcli.IAASEnvironment, string, string, string) ([]byte, error)) {
	fake.locksMutex.Lock()
	defer fake.locksMutex.Unlock()
	fake.LocksStub = stub
}

func (fake *FakeICLI) LocksArgsForCall(i int) (boshcli.IAASEnvironment, string, string, string) {
	fake.locksMutex.RLock()
	defer fake.locksMutex.RUnlock()
	argsForCall := fake.locksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeICLI) LocksReturns(result1 []byte, result2 error) {
	fake.locksMutex.Lock()
	defer fake.locksMutex.Unlock()
	fake.LocksStub = nil
	fake.locksReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeICLI) LocksReturnsOnCall(i int, result1 []byte, result2 error) {
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

func (fake *FakeICLI) Recreate(arg1 boshcli.IAASEnvironment, arg2 string, arg3 string, arg4 string) error {
	fake.recreateMutex.Lock()
	ret, specificReturn := fake.recreateReturnsOnCall[len(fake.recreateArgsForCall)]
	fake.recreateArgsForCall = append(fake.recreateArgsForCall, struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Recreate", []interface{}{arg1, arg2, arg3, arg4})
	fake.recreateMutex.Unlock()
	if fake.RecreateStub != nil {
		return fake.RecreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.recreateReturns
	return fakeReturns.result1
}

func (fake *FakeICLI) RecreateCallCount() int {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	return len(fake.recreateArgsForCall)
}

func (fake *FakeICLI) RecreateCalls(stub func(boshcli.IAASEnvironment, string, string, string) error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = stub
}

func (fake *FakeICLI) RecreateArgsForCall(i int) (boshcli.IAASEnvironment, string, string, string) {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	argsForCall := fake.recreateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeICLI) RecreateReturns(result1 error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = nil
	fake.recreateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) RecreateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeICLI) RunAuthenticatedCommand(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool, arg6 io.Writer, arg7 ...string) error {
	fake.runAuthenticatedCommandMutex.Lock()
	ret, specificReturn := fake.runAuthenticatedCommandReturnsOnCall[len(fake.runAuthenticatedCommandArgsForCall)]
	fake.runAuthenticatedCommandArgsForCall = append(fake.runAuthenticatedCommandArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
		arg6 io.Writer
		arg7 []string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("RunAuthenticatedCommand", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.runAuthenticatedCommandMutex.Unlock()
	if fake.RunAuthenticatedCommandStub != nil {
		return fake.RunAuthenticatedCommandStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runAuthenticatedCommandReturns
	return fakeReturns.result1
}

func (fake *FakeICLI) RunAuthenticatedCommandCallCount() int {
	fake.runAuthenticatedCommandMutex.RLock()
	defer fake.runAuthenticatedCommandMutex.RUnlock()
	return len(fake.runAuthenticatedCommandArgsForCall)
}

func (fake *FakeICLI) RunAuthenticatedCommandCalls(stub func(string, string, string, string, bool, io.Writer, ...string) error) {
	fake.runAuthenticatedCommandMutex.Lock()
	defer fake.runAuthenticatedCommandMutex.Unlock()
	fake.RunAuthenticatedCommandStub = stub
}

func (fake *FakeICLI) RunAuthenticatedCommandArgsForCall(i int) (string, string, string, string, bool, io.Writer, []string) {
	fake.runAuthenticatedCommandMutex.RLock()
	defer fake.runAuthenticatedCommandMutex.RUnlock()
	argsForCall := fake.runAuthenticatedCommandArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeICLI) RunAuthenticatedCommandReturns(result1 error) {
	fake.runAuthenticatedCommandMutex.Lock()
	defer fake.runAuthenticatedCommandMutex.Unlock()
	fake.RunAuthenticatedCommandStub = nil
	fake.runAuthenticatedCommandReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) RunAuthenticatedCommandReturnsOnCall(i int, result1 error) {
	fake.runAuthenticatedCommandMutex.Lock()
	defer fake.runAuthenticatedCommandMutex.Unlock()
	fake.RunAuthenticatedCommandStub = nil
	if fake.runAuthenticatedCommandReturnsOnCall == nil {
		fake.runAuthenticatedCommandReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runAuthenticatedCommandReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) UpdateCloudConfig(arg1 boshcli.IAASEnvironment, arg2 string, arg3 string, arg4 string) error {
	fake.updateCloudConfigMutex.Lock()
	ret, specificReturn := fake.updateCloudConfigReturnsOnCall[len(fake.updateCloudConfigArgsForCall)]
	fake.updateCloudConfigArgsForCall = append(fake.updateCloudConfigArgsForCall, struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UpdateCloudConfig", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateCloudConfigMutex.Unlock()
	if fake.UpdateCloudConfigStub != nil {
		return fake.UpdateCloudConfigStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateCloudConfigReturns
	return fakeReturns.result1
}

func (fake *FakeICLI) UpdateCloudConfigCallCount() int {
	fake.updateCloudConfigMutex.RLock()
	defer fake.updateCloudConfigMutex.RUnlock()
	return len(fake.updateCloudConfigArgsForCall)
}

func (fake *FakeICLI) UpdateCloudConfigCalls(stub func(boshcli.IAASEnvironment, string, string, string) error) {
	fake.updateCloudConfigMutex.Lock()
	defer fake.updateCloudConfigMutex.Unlock()
	fake.UpdateCloudConfigStub = stub
}

func (fake *FakeICLI) UpdateCloudConfigArgsForCall(i int) (boshcli.IAASEnvironment, string, string, string) {
	fake.updateCloudConfigMutex.RLock()
	defer fake.updateCloudConfigMutex.RUnlock()
	argsForCall := fake.updateCloudConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeICLI) UpdateCloudConfigReturns(result1 error) {
	fake.updateCloudConfigMutex.Lock()
	defer fake.updateCloudConfigMutex.Unlock()
	fake.UpdateCloudConfigStub = nil
	fake.updateCloudConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) UpdateCloudConfigReturnsOnCall(i int, result1 error) {
	fake.updateCloudConfigMutex.Lock()
	defer fake.updateCloudConfigMutex.Unlock()
	fake.UpdateCloudConfigStub = nil
	if fake.updateCloudConfigReturnsOnCall == nil {
		fake.updateCloudConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateCloudConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) UploadConcourseStemcell(arg1 boshcli.IAASEnvironment, arg2 string, arg3 string, arg4 string) error {
	fake.uploadConcourseStemcellMutex.Lock()
	ret, specificReturn := fake.uploadConcourseStemcellReturnsOnCall[len(fake.uploadConcourseStemcellArgsForCall)]
	fake.uploadConcourseStemcellArgsForCall = append(fake.uploadConcourseStemcellArgsForCall, struct {
		arg1 boshcli.IAASEnvironment
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UploadConcourseStemcell", []interface{}{arg1, arg2, arg3, arg4})
	fake.uploadConcourseStemcellMutex.Unlock()
	if fake.UploadConcourseStemcellStub != nil {
		return fake.UploadConcourseStemcellStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uploadConcourseStemcellReturns
	return fakeReturns.result1
}

func (fake *FakeICLI) UploadConcourseStemcellCallCount() int {
	fake.uploadConcourseStemcellMutex.RLock()
	defer fake.uploadConcourseStemcellMutex.RUnlock()
	return len(fake.uploadConcourseStemcellArgsForCall)
}

func (fake *FakeICLI) UploadConcourseStemcellCalls(stub func(boshcli.IAASEnvironment, string, string, string) error) {
	fake.uploadConcourseStemcellMutex.Lock()
	defer fake.uploadConcourseStemcellMutex.Unlock()
	fake.UploadConcourseStemcellStub = stub
}

func (fake *FakeICLI) UploadConcourseStemcellArgsForCall(i int) (boshcli.IAASEnvironment, string, string, string) {
	fake.uploadConcourseStemcellMutex.RLock()
	defer fake.uploadConcourseStemcellMutex.RUnlock()
	argsForCall := fake.uploadConcourseStemcellArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeICLI) UploadConcourseStemcellReturns(result1 error) {
	fake.uploadConcourseStemcellMutex.Lock()
	defer fake.uploadConcourseStemcellMutex.Unlock()
	fake.UploadConcourseStemcellStub = nil
	fake.uploadConcourseStemcellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) UploadConcourseStemcellReturnsOnCall(i int, result1 error) {
	fake.uploadConcourseStemcellMutex.Lock()
	defer fake.uploadConcourseStemcellMutex.Unlock()
	fake.UploadConcourseStemcellStub = nil
	if fake.uploadConcourseStemcellReturnsOnCall == nil {
		fake.uploadConcourseStemcellReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadConcourseStemcellReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICLI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEnvMutex.RLock()
	defer fake.createEnvMutex.RUnlock()
	fake.locksMutex.RLock()
	defer fake.locksMutex.RUnlock()
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	fake.runAuthenticatedCommandMutex.RLock()
	defer fake.runAuthenticatedCommandMutex.RUnlock()
	fake.updateCloudConfigMutex.RLock()
	defer fake.updateCloudConfigMutex.RUnlock()
	fake.uploadConcourseStemcellMutex.RLock()
	defer fake.uploadConcourseStemcellMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeICLI) recordInvocation(key string, args []interface{}) {
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

var _ boshcli.ICLI = new(FakeICLI)
