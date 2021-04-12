// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/eirini/api"
	"code.cloudfoundry.org/eirini/handler"
	"code.cloudfoundry.org/eirini/models/cf"
)

type FakeLRPBifrost struct {
	GetAppStub        func(context.Context, api.LRPIdentifier) (cf.DesiredLRP, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}
	getAppReturns struct {
		result1 cf.DesiredLRP
		result2 error
	}
	getAppReturnsOnCall map[int]struct {
		result1 cf.DesiredLRP
		result2 error
	}
	GetInstancesStub        func(context.Context, api.LRPIdentifier) ([]*cf.Instance, error)
	getInstancesMutex       sync.RWMutex
	getInstancesArgsForCall []struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}
	getInstancesReturns struct {
		result1 []*cf.Instance
		result2 error
	}
	getInstancesReturnsOnCall map[int]struct {
		result1 []*cf.Instance
		result2 error
	}
	ListStub        func(context.Context) ([]cf.DesiredLRPSchedulingInfo, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 []cf.DesiredLRPSchedulingInfo
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []cf.DesiredLRPSchedulingInfo
		result2 error
	}
	StopStub        func(context.Context, api.LRPIdentifier) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	StopInstanceStub        func(context.Context, api.LRPIdentifier, uint) error
	stopInstanceMutex       sync.RWMutex
	stopInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
		arg3 uint
	}
	stopInstanceReturns struct {
		result1 error
	}
	stopInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	TransferStub        func(context.Context, cf.DesireLRPRequest) error
	transferMutex       sync.RWMutex
	transferArgsForCall []struct {
		arg1 context.Context
		arg2 cf.DesireLRPRequest
	}
	transferReturns struct {
		result1 error
	}
	transferReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(context.Context, cf.UpdateDesiredLRPRequest) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 cf.UpdateDesiredLRPRequest
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLRPBifrost) GetApp(arg1 context.Context, arg2 api.LRPIdentifier) (cf.DesiredLRP, error) {
	fake.getAppMutex.Lock()
	ret, specificReturn := fake.getAppReturnsOnCall[len(fake.getAppArgsForCall)]
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}{arg1, arg2})
	stub := fake.GetAppStub
	fakeReturns := fake.getAppReturns
	fake.recordInvocation("GetApp", []interface{}{arg1, arg2})
	fake.getAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPBifrost) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeLRPBifrost) GetAppCalls(stub func(context.Context, api.LRPIdentifier) (cf.DesiredLRP, error)) {
	fake.getAppMutex.Lock()
	defer fake.getAppMutex.Unlock()
	fake.GetAppStub = stub
}

func (fake *FakeLRPBifrost) GetAppArgsForCall(i int) (context.Context, api.LRPIdentifier) {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	argsForCall := fake.getAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPBifrost) GetAppReturns(result1 cf.DesiredLRP, result2 error) {
	fake.getAppMutex.Lock()
	defer fake.getAppMutex.Unlock()
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 cf.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) GetAppReturnsOnCall(i int, result1 cf.DesiredLRP, result2 error) {
	fake.getAppMutex.Lock()
	defer fake.getAppMutex.Unlock()
	fake.GetAppStub = nil
	if fake.getAppReturnsOnCall == nil {
		fake.getAppReturnsOnCall = make(map[int]struct {
			result1 cf.DesiredLRP
			result2 error
		})
	}
	fake.getAppReturnsOnCall[i] = struct {
		result1 cf.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) GetInstances(arg1 context.Context, arg2 api.LRPIdentifier) ([]*cf.Instance, error) {
	fake.getInstancesMutex.Lock()
	ret, specificReturn := fake.getInstancesReturnsOnCall[len(fake.getInstancesArgsForCall)]
	fake.getInstancesArgsForCall = append(fake.getInstancesArgsForCall, struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}{arg1, arg2})
	stub := fake.GetInstancesStub
	fakeReturns := fake.getInstancesReturns
	fake.recordInvocation("GetInstances", []interface{}{arg1, arg2})
	fake.getInstancesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPBifrost) GetInstancesCallCount() int {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return len(fake.getInstancesArgsForCall)
}

func (fake *FakeLRPBifrost) GetInstancesCalls(stub func(context.Context, api.LRPIdentifier) ([]*cf.Instance, error)) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = stub
}

func (fake *FakeLRPBifrost) GetInstancesArgsForCall(i int) (context.Context, api.LRPIdentifier) {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	argsForCall := fake.getInstancesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPBifrost) GetInstancesReturns(result1 []*cf.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	fake.getInstancesReturns = struct {
		result1 []*cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) GetInstancesReturnsOnCall(i int, result1 []*cf.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	if fake.getInstancesReturnsOnCall == nil {
		fake.getInstancesReturnsOnCall = make(map[int]struct {
			result1 []*cf.Instance
			result2 error
		})
	}
	fake.getInstancesReturnsOnCall[i] = struct {
		result1 []*cf.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) List(arg1 context.Context) ([]cf.DesiredLRPSchedulingInfo, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPBifrost) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeLRPBifrost) ListCalls(stub func(context.Context) ([]cf.DesiredLRPSchedulingInfo, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeLRPBifrost) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLRPBifrost) ListReturns(result1 []cf.DesiredLRPSchedulingInfo, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []cf.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) ListReturnsOnCall(i int, result1 []cf.DesiredLRPSchedulingInfo, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []cf.DesiredLRPSchedulingInfo
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []cf.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPBifrost) Stop(arg1 context.Context, arg2 api.LRPIdentifier) error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
	}{arg1, arg2})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{arg1, arg2})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPBifrost) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeLRPBifrost) StopCalls(stub func(context.Context, api.LRPIdentifier) error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeLRPBifrost) StopArgsForCall(i int) (context.Context, api.LRPIdentifier) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	argsForCall := fake.stopArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPBifrost) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) StopInstance(arg1 context.Context, arg2 api.LRPIdentifier, arg3 uint) error {
	fake.stopInstanceMutex.Lock()
	ret, specificReturn := fake.stopInstanceReturnsOnCall[len(fake.stopInstanceArgsForCall)]
	fake.stopInstanceArgsForCall = append(fake.stopInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 api.LRPIdentifier
		arg3 uint
	}{arg1, arg2, arg3})
	stub := fake.StopInstanceStub
	fakeReturns := fake.stopInstanceReturns
	fake.recordInvocation("StopInstance", []interface{}{arg1, arg2, arg3})
	fake.stopInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPBifrost) StopInstanceCallCount() int {
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	return len(fake.stopInstanceArgsForCall)
}

func (fake *FakeLRPBifrost) StopInstanceCalls(stub func(context.Context, api.LRPIdentifier, uint) error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = stub
}

func (fake *FakeLRPBifrost) StopInstanceArgsForCall(i int) (context.Context, api.LRPIdentifier, uint) {
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	argsForCall := fake.stopInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLRPBifrost) StopInstanceReturns(result1 error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = nil
	fake.stopInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) StopInstanceReturnsOnCall(i int, result1 error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = nil
	if fake.stopInstanceReturnsOnCall == nil {
		fake.stopInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) Transfer(arg1 context.Context, arg2 cf.DesireLRPRequest) error {
	fake.transferMutex.Lock()
	ret, specificReturn := fake.transferReturnsOnCall[len(fake.transferArgsForCall)]
	fake.transferArgsForCall = append(fake.transferArgsForCall, struct {
		arg1 context.Context
		arg2 cf.DesireLRPRequest
	}{arg1, arg2})
	stub := fake.TransferStub
	fakeReturns := fake.transferReturns
	fake.recordInvocation("Transfer", []interface{}{arg1, arg2})
	fake.transferMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPBifrost) TransferCallCount() int {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	return len(fake.transferArgsForCall)
}

func (fake *FakeLRPBifrost) TransferCalls(stub func(context.Context, cf.DesireLRPRequest) error) {
	fake.transferMutex.Lock()
	defer fake.transferMutex.Unlock()
	fake.TransferStub = stub
}

func (fake *FakeLRPBifrost) TransferArgsForCall(i int) (context.Context, cf.DesireLRPRequest) {
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	argsForCall := fake.transferArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPBifrost) TransferReturns(result1 error) {
	fake.transferMutex.Lock()
	defer fake.transferMutex.Unlock()
	fake.TransferStub = nil
	fake.transferReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) TransferReturnsOnCall(i int, result1 error) {
	fake.transferMutex.Lock()
	defer fake.transferMutex.Unlock()
	fake.TransferStub = nil
	if fake.transferReturnsOnCall == nil {
		fake.transferReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.transferReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) Update(arg1 context.Context, arg2 cf.UpdateDesiredLRPRequest) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 cf.UpdateDesiredLRPRequest
	}{arg1, arg2})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPBifrost) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeLRPBifrost) UpdateCalls(stub func(context.Context, cf.UpdateDesiredLRPRequest) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeLRPBifrost) UpdateArgsForCall(i int) (context.Context, cf.UpdateDesiredLRPRequest) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPBifrost) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPBifrost) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	fake.transferMutex.RLock()
	defer fake.transferMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLRPBifrost) recordInvocation(key string, args []interface{}) {
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

var _ handler.LRPBifrost = new(FakeLRPBifrost)
