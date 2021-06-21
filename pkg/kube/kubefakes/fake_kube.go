// Code generated by counterfeiter. DO NOT EDIT.
package kubefakes

import (
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/kube"
)

type FakeKube struct {
	ApplyStub        func([]byte, string) ([]byte, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 []byte
		arg2 string
	}
	applyReturns struct {
		result1 []byte
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	FluxPresentStub        func() (bool, error)
	fluxPresentMutex       sync.RWMutex
	fluxPresentArgsForCall []struct {
	}
	fluxPresentReturns struct {
		result1 bool
		result2 error
	}
	fluxPresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetClusterNameStub        func() (string, error)
	getClusterNameMutex       sync.RWMutex
	getClusterNameArgsForCall []struct {
	}
	getClusterNameReturns struct {
		result1 string
		result2 error
	}
	getClusterNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetClusterStatusStub        func() kube.ClusterStatus
	getClusterStatusMutex       sync.RWMutex
	getClusterStatusArgsForCall []struct {
	}
	getClusterStatusReturns struct {
		result1 kube.ClusterStatus
	}
	getClusterStatusReturnsOnCall map[int]struct {
		result1 kube.ClusterStatus
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKube) Apply(arg1 []byte, arg2 string) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 []byte
		arg2 string
	}{arg1Copy, arg2})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1Copy, arg2})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeKube) ApplyCalls(stub func([]byte, string) ([]byte, error)) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeKube) ApplyArgsForCall(i int) ([]byte, string) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) ApplyReturns(result1 []byte, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) ApplyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresent() (bool, error) {
	fake.fluxPresentMutex.Lock()
	ret, specificReturn := fake.fluxPresentReturnsOnCall[len(fake.fluxPresentArgsForCall)]
	fake.fluxPresentArgsForCall = append(fake.fluxPresentArgsForCall, struct {
	}{})
	stub := fake.FluxPresentStub
	fakeReturns := fake.fluxPresentReturns
	fake.recordInvocation("FluxPresent", []interface{}{})
	fake.fluxPresentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) FluxPresentCallCount() int {
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	return len(fake.fluxPresentArgsForCall)
}

func (fake *FakeKube) FluxPresentCalls(stub func() (bool, error)) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = stub
}

func (fake *FakeKube) FluxPresentReturns(result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	fake.fluxPresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	if fake.fluxPresentReturnsOnCall == nil {
		fake.fluxPresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fluxPresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterName() (string, error) {
	fake.getClusterNameMutex.Lock()
	ret, specificReturn := fake.getClusterNameReturnsOnCall[len(fake.getClusterNameArgsForCall)]
	fake.getClusterNameArgsForCall = append(fake.getClusterNameArgsForCall, struct {
	}{})
	stub := fake.GetClusterNameStub
	fakeReturns := fake.getClusterNameReturns
	fake.recordInvocation("GetClusterName", []interface{}{})
	fake.getClusterNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetClusterNameCallCount() int {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	return len(fake.getClusterNameArgsForCall)
}

func (fake *FakeKube) GetClusterNameCalls(stub func() (string, error)) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = stub
}

func (fake *FakeKube) GetClusterNameReturns(result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	fake.getClusterNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	if fake.getClusterNameReturnsOnCall == nil {
		fake.getClusterNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getClusterNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterStatus() kube.ClusterStatus {
	fake.getClusterStatusMutex.Lock()
	ret, specificReturn := fake.getClusterStatusReturnsOnCall[len(fake.getClusterStatusArgsForCall)]
	fake.getClusterStatusArgsForCall = append(fake.getClusterStatusArgsForCall, struct {
	}{})
	stub := fake.GetClusterStatusStub
	fakeReturns := fake.getClusterStatusReturns
	fake.recordInvocation("GetClusterStatus", []interface{}{})
	fake.getClusterStatusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) GetClusterStatusCallCount() int {
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	return len(fake.getClusterStatusArgsForCall)
}

func (fake *FakeKube) GetClusterStatusCalls(stub func() kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = stub
}

func (fake *FakeKube) GetClusterStatusReturns(result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	fake.getClusterStatusReturns = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) GetClusterStatusReturnsOnCall(i int, result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	if fake.getClusterStatusReturnsOnCall == nil {
		fake.getClusterStatusReturnsOnCall = make(map[int]struct {
			result1 kube.ClusterStatus
		})
	}
	fake.getClusterStatusReturnsOnCall[i] = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKube) recordInvocation(key string, args []interface{}) {
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

var _ kube.Kube = new(FakeKube)
