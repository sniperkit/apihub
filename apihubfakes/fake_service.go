// This file was generated by counterfeiter
package apihubfakes

import (
	"sync"
	"time"

	"github.com/apihub/apihub"
)

type FakeService struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	InfoStub        func() (apihub.ServiceSpec, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns     struct {
		result1 apihub.ServiceSpec
		result2 error
	}
	AddBackendStub        func(be apihub.BackendInfo) error
	addBackendMutex       sync.RWMutex
	addBackendArgsForCall []struct {
		be apihub.BackendInfo
	}
	addBackendReturns struct {
		result1 error
	}
	RemoveBackendStub        func(be apihub.BackendInfo) error
	removeBackendMutex       sync.RWMutex
	removeBackendArgsForCall []struct {
		be apihub.BackendInfo
	}
	removeBackendReturns struct {
		result1 error
	}
	BackendsStub        func() ([]apihub.BackendInfo, error)
	backendsMutex       sync.RWMutex
	backendsArgsForCall []struct{}
	backendsReturns     struct {
		result1 []apihub.BackendInfo
		result2 error
	}
	SetTimeoutStub        func(time.Duration)
	setTimeoutMutex       sync.RWMutex
	setTimeoutArgsForCall []struct {
		arg1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeService) Handle() string {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	} else {
		return fake.handleReturns.result1
	}
}

func (fake *FakeService) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeService) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeService) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeService) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeService) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Stop() error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeService) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeService) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Info() (apihub.ServiceSpec, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeService) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeService) InfoReturns(result1 apihub.ServiceSpec, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 apihub.ServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeService) AddBackend(be apihub.BackendInfo) error {
	fake.addBackendMutex.Lock()
	fake.addBackendArgsForCall = append(fake.addBackendArgsForCall, struct {
		be apihub.BackendInfo
	}{be})
	fake.recordInvocation("AddBackend", []interface{}{be})
	fake.addBackendMutex.Unlock()
	if fake.AddBackendStub != nil {
		return fake.AddBackendStub(be)
	} else {
		return fake.addBackendReturns.result1
	}
}

func (fake *FakeService) AddBackendCallCount() int {
	fake.addBackendMutex.RLock()
	defer fake.addBackendMutex.RUnlock()
	return len(fake.addBackendArgsForCall)
}

func (fake *FakeService) AddBackendArgsForCall(i int) apihub.BackendInfo {
	fake.addBackendMutex.RLock()
	defer fake.addBackendMutex.RUnlock()
	return fake.addBackendArgsForCall[i].be
}

func (fake *FakeService) AddBackendReturns(result1 error) {
	fake.AddBackendStub = nil
	fake.addBackendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) RemoveBackend(be apihub.BackendInfo) error {
	fake.removeBackendMutex.Lock()
	fake.removeBackendArgsForCall = append(fake.removeBackendArgsForCall, struct {
		be apihub.BackendInfo
	}{be})
	fake.recordInvocation("RemoveBackend", []interface{}{be})
	fake.removeBackendMutex.Unlock()
	if fake.RemoveBackendStub != nil {
		return fake.RemoveBackendStub(be)
	} else {
		return fake.removeBackendReturns.result1
	}
}

func (fake *FakeService) RemoveBackendCallCount() int {
	fake.removeBackendMutex.RLock()
	defer fake.removeBackendMutex.RUnlock()
	return len(fake.removeBackendArgsForCall)
}

func (fake *FakeService) RemoveBackendArgsForCall(i int) apihub.BackendInfo {
	fake.removeBackendMutex.RLock()
	defer fake.removeBackendMutex.RUnlock()
	return fake.removeBackendArgsForCall[i].be
}

func (fake *FakeService) RemoveBackendReturns(result1 error) {
	fake.RemoveBackendStub = nil
	fake.removeBackendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeService) Backends() ([]apihub.BackendInfo, error) {
	fake.backendsMutex.Lock()
	fake.backendsArgsForCall = append(fake.backendsArgsForCall, struct{}{})
	fake.recordInvocation("Backends", []interface{}{})
	fake.backendsMutex.Unlock()
	if fake.BackendsStub != nil {
		return fake.BackendsStub()
	} else {
		return fake.backendsReturns.result1, fake.backendsReturns.result2
	}
}

func (fake *FakeService) BackendsCallCount() int {
	fake.backendsMutex.RLock()
	defer fake.backendsMutex.RUnlock()
	return len(fake.backendsArgsForCall)
}

func (fake *FakeService) BackendsReturns(result1 []apihub.BackendInfo, result2 error) {
	fake.BackendsStub = nil
	fake.backendsReturns = struct {
		result1 []apihub.BackendInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeService) SetTimeout(arg1 time.Duration) {
	fake.setTimeoutMutex.Lock()
	fake.setTimeoutArgsForCall = append(fake.setTimeoutArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("SetTimeout", []interface{}{arg1})
	fake.setTimeoutMutex.Unlock()
	if fake.SetTimeoutStub != nil {
		fake.SetTimeoutStub(arg1)
	}
}

func (fake *FakeService) SetTimeoutCallCount() int {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return len(fake.setTimeoutArgsForCall)
}

func (fake *FakeService) SetTimeoutArgsForCall(i int) time.Duration {
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return fake.setTimeoutArgsForCall[i].arg1
}

func (fake *FakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.addBackendMutex.RLock()
	defer fake.addBackendMutex.RUnlock()
	fake.removeBackendMutex.RLock()
	defer fake.removeBackendMutex.RUnlock()
	fake.backendsMutex.RLock()
	defer fake.backendsMutex.RUnlock()
	fake.setTimeoutMutex.RLock()
	defer fake.setTimeoutMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeService) recordInvocation(key string, args []interface{}) {
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

var _ apihub.Service = new(FakeService)
