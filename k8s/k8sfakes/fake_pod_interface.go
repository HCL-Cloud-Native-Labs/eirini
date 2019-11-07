// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	v1a "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	v1b "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type FakePodInterface struct {
	BindStub        func(*v1a.Binding) error
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		arg1 *v1a.Binding
	}
	bindReturns struct {
		result1 error
	}
	bindReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(*v1a.Pod) (*v1a.Pod, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *v1a.Pod
	}
	createReturns struct {
		result1 *v1a.Pod
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1a.Pod
		result2 error
	}
	DeleteStub        func(string, *v1b.DeleteOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 *v1b.DeleteOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCollectionStub        func(*v1b.DeleteOptions, v1b.ListOptions) error
	deleteCollectionMutex       sync.RWMutex
	deleteCollectionArgsForCall []struct {
		arg1 *v1b.DeleteOptions
		arg2 v1b.ListOptions
	}
	deleteCollectionReturns struct {
		result1 error
	}
	deleteCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	EvictStub        func(*v1beta1.Eviction) error
	evictMutex       sync.RWMutex
	evictArgsForCall []struct {
		arg1 *v1beta1.Eviction
	}
	evictReturns struct {
		result1 error
	}
	evictReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string, v1b.GetOptions) (*v1a.Pod, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 v1b.GetOptions
	}
	getReturns struct {
		result1 *v1a.Pod
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1a.Pod
		result2 error
	}
	GetEphemeralContainersStub        func(string, v1b.GetOptions) (*v1a.EphemeralContainers, error)
	getEphemeralContainersMutex       sync.RWMutex
	getEphemeralContainersArgsForCall []struct {
		arg1 string
		arg2 v1b.GetOptions
	}
	getEphemeralContainersReturns struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}
	getEphemeralContainersReturnsOnCall map[int]struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}
	GetLogsStub        func(string, *v1a.PodLogOptions) *rest.Request
	getLogsMutex       sync.RWMutex
	getLogsArgsForCall []struct {
		arg1 string
		arg2 *v1a.PodLogOptions
	}
	getLogsReturns struct {
		result1 *rest.Request
	}
	getLogsReturnsOnCall map[int]struct {
		result1 *rest.Request
	}
	ListStub        func(v1b.ListOptions) (*v1a.PodList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 v1b.ListOptions
	}
	listReturns struct {
		result1 *v1a.PodList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1a.PodList
		result2 error
	}
	PatchStub        func(string, types.PatchType, []byte, ...string) (*v1a.Pod, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 string
		arg2 types.PatchType
		arg3 []byte
		arg4 []string
	}
	patchReturns struct {
		result1 *v1a.Pod
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *v1a.Pod
		result2 error
	}
	UpdateStub        func(*v1a.Pod) (*v1a.Pod, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 *v1a.Pod
	}
	updateReturns struct {
		result1 *v1a.Pod
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1a.Pod
		result2 error
	}
	UpdateEphemeralContainersStub        func(string, *v1a.EphemeralContainers) (*v1a.EphemeralContainers, error)
	updateEphemeralContainersMutex       sync.RWMutex
	updateEphemeralContainersArgsForCall []struct {
		arg1 string
		arg2 *v1a.EphemeralContainers
	}
	updateEphemeralContainersReturns struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}
	updateEphemeralContainersReturnsOnCall map[int]struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}
	UpdateStatusStub        func(*v1a.Pod) (*v1a.Pod, error)
	updateStatusMutex       sync.RWMutex
	updateStatusArgsForCall []struct {
		arg1 *v1a.Pod
	}
	updateStatusReturns struct {
		result1 *v1a.Pod
		result2 error
	}
	updateStatusReturnsOnCall map[int]struct {
		result1 *v1a.Pod
		result2 error
	}
	WatchStub        func(v1b.ListOptions) (watch.Interface, error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 v1b.ListOptions
	}
	watchReturns struct {
		result1 watch.Interface
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 watch.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePodInterface) Bind(arg1 *v1a.Binding) error {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		arg1 *v1a.Binding
	}{arg1})
	fake.recordInvocation("Bind", []interface{}{arg1})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bindReturns
	return fakeReturns.result1
}

func (fake *FakePodInterface) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakePodInterface) BindCalls(stub func(*v1a.Binding) error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = stub
}

func (fake *FakePodInterface) BindArgsForCall(i int) *v1a.Binding {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	argsForCall := fake.bindArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) BindReturns(result1 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) BindReturnsOnCall(i int, result1 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) Create(arg1 *v1a.Pod) (*v1a.Pod, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *v1a.Pod
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakePodInterface) CreateCalls(stub func(*v1a.Pod) (*v1a.Pod, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakePodInterface) CreateArgsForCall(i int) *v1a.Pod {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) CreateReturns(result1 *v1a.Pod, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) CreateReturnsOnCall(i int, result1 *v1a.Pod, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1a.Pod
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) Delete(arg1 string, arg2 *v1b.DeleteOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 *v1b.DeleteOptions
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakePodInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakePodInterface) DeleteCalls(stub func(string, *v1b.DeleteOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakePodInterface) DeleteArgsForCall(i int) (string, *v1b.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) DeleteCollection(arg1 *v1b.DeleteOptions, arg2 v1b.ListOptions) error {
	fake.deleteCollectionMutex.Lock()
	ret, specificReturn := fake.deleteCollectionReturnsOnCall[len(fake.deleteCollectionArgsForCall)]
	fake.deleteCollectionArgsForCall = append(fake.deleteCollectionArgsForCall, struct {
		arg1 *v1b.DeleteOptions
		arg2 v1b.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("DeleteCollection", []interface{}{arg1, arg2})
	fake.deleteCollectionMutex.Unlock()
	if fake.DeleteCollectionStub != nil {
		return fake.DeleteCollectionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCollectionReturns
	return fakeReturns.result1
}

func (fake *FakePodInterface) DeleteCollectionCallCount() int {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	return len(fake.deleteCollectionArgsForCall)
}

func (fake *FakePodInterface) DeleteCollectionCalls(stub func(*v1b.DeleteOptions, v1b.ListOptions) error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = stub
}

func (fake *FakePodInterface) DeleteCollectionArgsForCall(i int) (*v1b.DeleteOptions, v1b.ListOptions) {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	argsForCall := fake.deleteCollectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) DeleteCollectionReturns(result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	fake.deleteCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) DeleteCollectionReturnsOnCall(i int, result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	if fake.deleteCollectionReturnsOnCall == nil {
		fake.deleteCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) Evict(arg1 *v1beta1.Eviction) error {
	fake.evictMutex.Lock()
	ret, specificReturn := fake.evictReturnsOnCall[len(fake.evictArgsForCall)]
	fake.evictArgsForCall = append(fake.evictArgsForCall, struct {
		arg1 *v1beta1.Eviction
	}{arg1})
	fake.recordInvocation("Evict", []interface{}{arg1})
	fake.evictMutex.Unlock()
	if fake.EvictStub != nil {
		return fake.EvictStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.evictReturns
	return fakeReturns.result1
}

func (fake *FakePodInterface) EvictCallCount() int {
	fake.evictMutex.RLock()
	defer fake.evictMutex.RUnlock()
	return len(fake.evictArgsForCall)
}

func (fake *FakePodInterface) EvictCalls(stub func(*v1beta1.Eviction) error) {
	fake.evictMutex.Lock()
	defer fake.evictMutex.Unlock()
	fake.EvictStub = stub
}

func (fake *FakePodInterface) EvictArgsForCall(i int) *v1beta1.Eviction {
	fake.evictMutex.RLock()
	defer fake.evictMutex.RUnlock()
	argsForCall := fake.evictArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) EvictReturns(result1 error) {
	fake.evictMutex.Lock()
	defer fake.evictMutex.Unlock()
	fake.EvictStub = nil
	fake.evictReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) EvictReturnsOnCall(i int, result1 error) {
	fake.evictMutex.Lock()
	defer fake.evictMutex.Unlock()
	fake.EvictStub = nil
	if fake.evictReturnsOnCall == nil {
		fake.evictReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evictReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePodInterface) Get(arg1 string, arg2 v1b.GetOptions) (*v1a.Pod, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 v1b.GetOptions
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakePodInterface) GetCalls(stub func(string, v1b.GetOptions) (*v1a.Pod, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakePodInterface) GetArgsForCall(i int) (string, v1b.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) GetReturns(result1 *v1a.Pod, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) GetReturnsOnCall(i int, result1 *v1a.Pod, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1a.Pod
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) GetEphemeralContainers(arg1 string, arg2 v1b.GetOptions) (*v1a.EphemeralContainers, error) {
	fake.getEphemeralContainersMutex.Lock()
	ret, specificReturn := fake.getEphemeralContainersReturnsOnCall[len(fake.getEphemeralContainersArgsForCall)]
	fake.getEphemeralContainersArgsForCall = append(fake.getEphemeralContainersArgsForCall, struct {
		arg1 string
		arg2 v1b.GetOptions
	}{arg1, arg2})
	fake.recordInvocation("GetEphemeralContainers", []interface{}{arg1, arg2})
	fake.getEphemeralContainersMutex.Unlock()
	if fake.GetEphemeralContainersStub != nil {
		return fake.GetEphemeralContainersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getEphemeralContainersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) GetEphemeralContainersCallCount() int {
	fake.getEphemeralContainersMutex.RLock()
	defer fake.getEphemeralContainersMutex.RUnlock()
	return len(fake.getEphemeralContainersArgsForCall)
}

func (fake *FakePodInterface) GetEphemeralContainersCalls(stub func(string, v1b.GetOptions) (*v1a.EphemeralContainers, error)) {
	fake.getEphemeralContainersMutex.Lock()
	defer fake.getEphemeralContainersMutex.Unlock()
	fake.GetEphemeralContainersStub = stub
}

func (fake *FakePodInterface) GetEphemeralContainersArgsForCall(i int) (string, v1b.GetOptions) {
	fake.getEphemeralContainersMutex.RLock()
	defer fake.getEphemeralContainersMutex.RUnlock()
	argsForCall := fake.getEphemeralContainersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) GetEphemeralContainersReturns(result1 *v1a.EphemeralContainers, result2 error) {
	fake.getEphemeralContainersMutex.Lock()
	defer fake.getEphemeralContainersMutex.Unlock()
	fake.GetEphemeralContainersStub = nil
	fake.getEphemeralContainersReturns = struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) GetEphemeralContainersReturnsOnCall(i int, result1 *v1a.EphemeralContainers, result2 error) {
	fake.getEphemeralContainersMutex.Lock()
	defer fake.getEphemeralContainersMutex.Unlock()
	fake.GetEphemeralContainersStub = nil
	if fake.getEphemeralContainersReturnsOnCall == nil {
		fake.getEphemeralContainersReturnsOnCall = make(map[int]struct {
			result1 *v1a.EphemeralContainers
			result2 error
		})
	}
	fake.getEphemeralContainersReturnsOnCall[i] = struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) GetLogs(arg1 string, arg2 *v1a.PodLogOptions) *rest.Request {
	fake.getLogsMutex.Lock()
	ret, specificReturn := fake.getLogsReturnsOnCall[len(fake.getLogsArgsForCall)]
	fake.getLogsArgsForCall = append(fake.getLogsArgsForCall, struct {
		arg1 string
		arg2 *v1a.PodLogOptions
	}{arg1, arg2})
	fake.recordInvocation("GetLogs", []interface{}{arg1, arg2})
	fake.getLogsMutex.Unlock()
	if fake.GetLogsStub != nil {
		return fake.GetLogsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getLogsReturns
	return fakeReturns.result1
}

func (fake *FakePodInterface) GetLogsCallCount() int {
	fake.getLogsMutex.RLock()
	defer fake.getLogsMutex.RUnlock()
	return len(fake.getLogsArgsForCall)
}

func (fake *FakePodInterface) GetLogsCalls(stub func(string, *v1a.PodLogOptions) *rest.Request) {
	fake.getLogsMutex.Lock()
	defer fake.getLogsMutex.Unlock()
	fake.GetLogsStub = stub
}

func (fake *FakePodInterface) GetLogsArgsForCall(i int) (string, *v1a.PodLogOptions) {
	fake.getLogsMutex.RLock()
	defer fake.getLogsMutex.RUnlock()
	argsForCall := fake.getLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) GetLogsReturns(result1 *rest.Request) {
	fake.getLogsMutex.Lock()
	defer fake.getLogsMutex.Unlock()
	fake.GetLogsStub = nil
	fake.getLogsReturns = struct {
		result1 *rest.Request
	}{result1}
}

func (fake *FakePodInterface) GetLogsReturnsOnCall(i int, result1 *rest.Request) {
	fake.getLogsMutex.Lock()
	defer fake.getLogsMutex.Unlock()
	fake.GetLogsStub = nil
	if fake.getLogsReturnsOnCall == nil {
		fake.getLogsReturnsOnCall = make(map[int]struct {
			result1 *rest.Request
		})
	}
	fake.getLogsReturnsOnCall[i] = struct {
		result1 *rest.Request
	}{result1}
}

func (fake *FakePodInterface) List(arg1 v1b.ListOptions) (*v1a.PodList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 v1b.ListOptions
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePodInterface) ListCalls(stub func(v1b.ListOptions) (*v1a.PodList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakePodInterface) ListArgsForCall(i int) v1b.ListOptions {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) ListReturns(result1 *v1a.PodList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1a.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) ListReturnsOnCall(i int, result1 *v1a.PodList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1a.PodList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1a.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) Patch(arg1 string, arg2 types.PatchType, arg3 []byte, arg4 ...string) (*v1a.Pod, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 string
		arg2 types.PatchType
		arg3 []byte
		arg4 []string
	}{arg1, arg2, arg3Copy, arg4})
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.patchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakePodInterface) PatchCalls(stub func(string, types.PatchType, []byte, ...string) (*v1a.Pod, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakePodInterface) PatchArgsForCall(i int) (string, types.PatchType, []byte, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePodInterface) PatchReturns(result1 *v1a.Pod, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) PatchReturnsOnCall(i int, result1 *v1a.Pod, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *v1a.Pod
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) Update(arg1 *v1a.Pod) (*v1a.Pod, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 *v1a.Pod
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakePodInterface) UpdateCalls(stub func(*v1a.Pod) (*v1a.Pod, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakePodInterface) UpdateArgsForCall(i int) *v1a.Pod {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) UpdateReturns(result1 *v1a.Pod, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) UpdateReturnsOnCall(i int, result1 *v1a.Pod, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1a.Pod
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) UpdateEphemeralContainers(arg1 string, arg2 *v1a.EphemeralContainers) (*v1a.EphemeralContainers, error) {
	fake.updateEphemeralContainersMutex.Lock()
	ret, specificReturn := fake.updateEphemeralContainersReturnsOnCall[len(fake.updateEphemeralContainersArgsForCall)]
	fake.updateEphemeralContainersArgsForCall = append(fake.updateEphemeralContainersArgsForCall, struct {
		arg1 string
		arg2 *v1a.EphemeralContainers
	}{arg1, arg2})
	fake.recordInvocation("UpdateEphemeralContainers", []interface{}{arg1, arg2})
	fake.updateEphemeralContainersMutex.Unlock()
	if fake.UpdateEphemeralContainersStub != nil {
		return fake.UpdateEphemeralContainersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateEphemeralContainersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) UpdateEphemeralContainersCallCount() int {
	fake.updateEphemeralContainersMutex.RLock()
	defer fake.updateEphemeralContainersMutex.RUnlock()
	return len(fake.updateEphemeralContainersArgsForCall)
}

func (fake *FakePodInterface) UpdateEphemeralContainersCalls(stub func(string, *v1a.EphemeralContainers) (*v1a.EphemeralContainers, error)) {
	fake.updateEphemeralContainersMutex.Lock()
	defer fake.updateEphemeralContainersMutex.Unlock()
	fake.UpdateEphemeralContainersStub = stub
}

func (fake *FakePodInterface) UpdateEphemeralContainersArgsForCall(i int) (string, *v1a.EphemeralContainers) {
	fake.updateEphemeralContainersMutex.RLock()
	defer fake.updateEphemeralContainersMutex.RUnlock()
	argsForCall := fake.updateEphemeralContainersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePodInterface) UpdateEphemeralContainersReturns(result1 *v1a.EphemeralContainers, result2 error) {
	fake.updateEphemeralContainersMutex.Lock()
	defer fake.updateEphemeralContainersMutex.Unlock()
	fake.UpdateEphemeralContainersStub = nil
	fake.updateEphemeralContainersReturns = struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) UpdateEphemeralContainersReturnsOnCall(i int, result1 *v1a.EphemeralContainers, result2 error) {
	fake.updateEphemeralContainersMutex.Lock()
	defer fake.updateEphemeralContainersMutex.Unlock()
	fake.UpdateEphemeralContainersStub = nil
	if fake.updateEphemeralContainersReturnsOnCall == nil {
		fake.updateEphemeralContainersReturnsOnCall = make(map[int]struct {
			result1 *v1a.EphemeralContainers
			result2 error
		})
	}
	fake.updateEphemeralContainersReturnsOnCall[i] = struct {
		result1 *v1a.EphemeralContainers
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) UpdateStatus(arg1 *v1a.Pod) (*v1a.Pod, error) {
	fake.updateStatusMutex.Lock()
	ret, specificReturn := fake.updateStatusReturnsOnCall[len(fake.updateStatusArgsForCall)]
	fake.updateStatusArgsForCall = append(fake.updateStatusArgsForCall, struct {
		arg1 *v1a.Pod
	}{arg1})
	fake.recordInvocation("UpdateStatus", []interface{}{arg1})
	fake.updateStatusMutex.Unlock()
	if fake.UpdateStatusStub != nil {
		return fake.UpdateStatusStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) UpdateStatusCallCount() int {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	return len(fake.updateStatusArgsForCall)
}

func (fake *FakePodInterface) UpdateStatusCalls(stub func(*v1a.Pod) (*v1a.Pod, error)) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = stub
}

func (fake *FakePodInterface) UpdateStatusArgsForCall(i int) *v1a.Pod {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	argsForCall := fake.updateStatusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) UpdateStatusReturns(result1 *v1a.Pod, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	fake.updateStatusReturns = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) UpdateStatusReturnsOnCall(i int, result1 *v1a.Pod, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	if fake.updateStatusReturnsOnCall == nil {
		fake.updateStatusReturnsOnCall = make(map[int]struct {
			result1 *v1a.Pod
			result2 error
		})
	}
	fake.updateStatusReturnsOnCall[i] = struct {
		result1 *v1a.Pod
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) Watch(arg1 v1b.ListOptions) (watch.Interface, error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 v1b.ListOptions
	}{arg1})
	fake.recordInvocation("Watch", []interface{}{arg1})
	fake.watchMutex.Unlock()
	if fake.WatchStub != nil {
		return fake.WatchStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.watchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePodInterface) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakePodInterface) WatchCalls(stub func(v1b.ListOptions) (watch.Interface, error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakePodInterface) WatchArgsForCall(i int) v1b.ListOptions {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePodInterface) WatchReturns(result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) WatchReturnsOnCall(i int, result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 watch.Interface
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakePodInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	fake.evictMutex.RLock()
	defer fake.evictMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getEphemeralContainersMutex.RLock()
	defer fake.getEphemeralContainersMutex.RUnlock()
	fake.getLogsMutex.RLock()
	defer fake.getLogsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.updateEphemeralContainersMutex.RLock()
	defer fake.updateEphemeralContainersMutex.RUnlock()
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePodInterface) recordInvocation(key string, args []interface{}) {
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

var _ v1.PodInterface = new(FakePodInterface)
