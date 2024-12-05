/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package seccompprofilefakes

import (
	"context"
	"sync"

	"github.com/go-logr/logr"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/security-profiles-operator/api/seccompprofile/v1beta1"
	"sigs.k8s.io/security-profiles-operator/api/spod/v1alpha1"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/artifact"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/metrics"
)

type FakeImpl struct {
	ClientGetProfileStub        func(context.Context, client.Client, client.ObjectKey, ...client.GetOption) (*v1beta1.SeccompProfile, error)
	clientGetProfileMutex       sync.RWMutex
	clientGetProfileArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 client.ObjectKey
		arg4 []client.GetOption
	}
	clientGetProfileReturns struct {
		result1 *v1beta1.SeccompProfile
		result2 error
	}
	clientGetProfileReturnsOnCall map[int]struct {
		result1 *v1beta1.SeccompProfile
		result2 error
	}
	GetSPODStub        func(context.Context, client.Client) (*v1alpha1.SecurityProfilesOperatorDaemon, error)
	getSPODMutex       sync.RWMutex
	getSPODArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
	}
	getSPODReturns struct {
		result1 *v1alpha1.SecurityProfilesOperatorDaemon
		result2 error
	}
	getSPODReturnsOnCall map[int]struct {
		result1 *v1alpha1.SecurityProfilesOperatorDaemon
		result2 error
	}
	IncSeccompProfileErrorStub        func(*metrics.Metrics, string)
	incSeccompProfileErrorMutex       sync.RWMutex
	incSeccompProfileErrorArgsForCall []struct {
		arg1 *metrics.Metrics
		arg2 string
	}
	PullStub        func(context.Context, logr.Logger, string, string, string, *v1.Platform, bool) (*artifact.PullResult, error)
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		arg1 context.Context
		arg2 logr.Logger
		arg3 string
		arg4 string
		arg5 string
		arg6 *v1.Platform
		arg7 bool
	}
	pullReturns struct {
		result1 *artifact.PullResult
		result2 error
	}
	pullReturnsOnCall map[int]struct {
		result1 *artifact.PullResult
		result2 error
	}
	PullResultSeccompProfileStub        func(*artifact.PullResult) *v1beta1.SeccompProfile
	pullResultSeccompProfileMutex       sync.RWMutex
	pullResultSeccompProfileArgsForCall []struct {
		arg1 *artifact.PullResult
	}
	pullResultSeccompProfileReturns struct {
		result1 *v1beta1.SeccompProfile
	}
	pullResultSeccompProfileReturnsOnCall map[int]struct {
		result1 *v1beta1.SeccompProfile
	}
	PullResultTypeStub        func(*artifact.PullResult) artifact.PullResultType
	pullResultTypeMutex       sync.RWMutex
	pullResultTypeArgsForCall []struct {
		arg1 *artifact.PullResult
	}
	pullResultTypeReturns struct {
		result1 artifact.PullResultType
	}
	pullResultTypeReturnsOnCall map[int]struct {
		result1 artifact.PullResultType
	}
	RecordEventStub        func(record.EventRecorder, runtime.Object, string, string, string)
	recordEventMutex       sync.RWMutex
	recordEventArgsForCall []struct {
		arg1 record.EventRecorder
		arg2 runtime.Object
		arg3 string
		arg4 string
		arg5 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) ClientGetProfile(arg1 context.Context, arg2 client.Client, arg3 client.ObjectKey, arg4 ...client.GetOption) (*v1beta1.SeccompProfile, error) {
	fake.clientGetProfileMutex.Lock()
	ret, specificReturn := fake.clientGetProfileReturnsOnCall[len(fake.clientGetProfileArgsForCall)]
	fake.clientGetProfileArgsForCall = append(fake.clientGetProfileArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 client.ObjectKey
		arg4 []client.GetOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ClientGetProfileStub
	fakeReturns := fake.clientGetProfileReturns
	fake.recordInvocation("ClientGetProfile", []interface{}{arg1, arg2, arg3, arg4})
	fake.clientGetProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ClientGetProfileCallCount() int {
	fake.clientGetProfileMutex.RLock()
	defer fake.clientGetProfileMutex.RUnlock()
	return len(fake.clientGetProfileArgsForCall)
}

func (fake *FakeImpl) ClientGetProfileCalls(stub func(context.Context, client.Client, client.ObjectKey, ...client.GetOption) (*v1beta1.SeccompProfile, error)) {
	fake.clientGetProfileMutex.Lock()
	defer fake.clientGetProfileMutex.Unlock()
	fake.ClientGetProfileStub = stub
}

func (fake *FakeImpl) ClientGetProfileArgsForCall(i int) (context.Context, client.Client, client.ObjectKey, []client.GetOption) {
	fake.clientGetProfileMutex.RLock()
	defer fake.clientGetProfileMutex.RUnlock()
	argsForCall := fake.clientGetProfileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeImpl) ClientGetProfileReturns(result1 *v1beta1.SeccompProfile, result2 error) {
	fake.clientGetProfileMutex.Lock()
	defer fake.clientGetProfileMutex.Unlock()
	fake.ClientGetProfileStub = nil
	fake.clientGetProfileReturns = struct {
		result1 *v1beta1.SeccompProfile
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ClientGetProfileReturnsOnCall(i int, result1 *v1beta1.SeccompProfile, result2 error) {
	fake.clientGetProfileMutex.Lock()
	defer fake.clientGetProfileMutex.Unlock()
	fake.ClientGetProfileStub = nil
	if fake.clientGetProfileReturnsOnCall == nil {
		fake.clientGetProfileReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.SeccompProfile
			result2 error
		})
	}
	fake.clientGetProfileReturnsOnCall[i] = struct {
		result1 *v1beta1.SeccompProfile
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetSPOD(arg1 context.Context, arg2 client.Client) (*v1alpha1.SecurityProfilesOperatorDaemon, error) {
	fake.getSPODMutex.Lock()
	ret, specificReturn := fake.getSPODReturnsOnCall[len(fake.getSPODArgsForCall)]
	fake.getSPODArgsForCall = append(fake.getSPODArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
	}{arg1, arg2})
	stub := fake.GetSPODStub
	fakeReturns := fake.getSPODReturns
	fake.recordInvocation("GetSPOD", []interface{}{arg1, arg2})
	fake.getSPODMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetSPODCallCount() int {
	fake.getSPODMutex.RLock()
	defer fake.getSPODMutex.RUnlock()
	return len(fake.getSPODArgsForCall)
}

func (fake *FakeImpl) GetSPODCalls(stub func(context.Context, client.Client) (*v1alpha1.SecurityProfilesOperatorDaemon, error)) {
	fake.getSPODMutex.Lock()
	defer fake.getSPODMutex.Unlock()
	fake.GetSPODStub = stub
}

func (fake *FakeImpl) GetSPODArgsForCall(i int) (context.Context, client.Client) {
	fake.getSPODMutex.RLock()
	defer fake.getSPODMutex.RUnlock()
	argsForCall := fake.getSPODArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) GetSPODReturns(result1 *v1alpha1.SecurityProfilesOperatorDaemon, result2 error) {
	fake.getSPODMutex.Lock()
	defer fake.getSPODMutex.Unlock()
	fake.GetSPODStub = nil
	fake.getSPODReturns = struct {
		result1 *v1alpha1.SecurityProfilesOperatorDaemon
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetSPODReturnsOnCall(i int, result1 *v1alpha1.SecurityProfilesOperatorDaemon, result2 error) {
	fake.getSPODMutex.Lock()
	defer fake.getSPODMutex.Unlock()
	fake.GetSPODStub = nil
	if fake.getSPODReturnsOnCall == nil {
		fake.getSPODReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.SecurityProfilesOperatorDaemon
			result2 error
		})
	}
	fake.getSPODReturnsOnCall[i] = struct {
		result1 *v1alpha1.SecurityProfilesOperatorDaemon
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IncSeccompProfileError(arg1 *metrics.Metrics, arg2 string) {
	fake.incSeccompProfileErrorMutex.Lock()
	fake.incSeccompProfileErrorArgsForCall = append(fake.incSeccompProfileErrorArgsForCall, struct {
		arg1 *metrics.Metrics
		arg2 string
	}{arg1, arg2})
	stub := fake.IncSeccompProfileErrorStub
	fake.recordInvocation("IncSeccompProfileError", []interface{}{arg1, arg2})
	fake.incSeccompProfileErrorMutex.Unlock()
	if stub != nil {
		fake.IncSeccompProfileErrorStub(arg1, arg2)
	}
}

func (fake *FakeImpl) IncSeccompProfileErrorCallCount() int {
	fake.incSeccompProfileErrorMutex.RLock()
	defer fake.incSeccompProfileErrorMutex.RUnlock()
	return len(fake.incSeccompProfileErrorArgsForCall)
}

func (fake *FakeImpl) IncSeccompProfileErrorCalls(stub func(*metrics.Metrics, string)) {
	fake.incSeccompProfileErrorMutex.Lock()
	defer fake.incSeccompProfileErrorMutex.Unlock()
	fake.IncSeccompProfileErrorStub = stub
}

func (fake *FakeImpl) IncSeccompProfileErrorArgsForCall(i int) (*metrics.Metrics, string) {
	fake.incSeccompProfileErrorMutex.RLock()
	defer fake.incSeccompProfileErrorMutex.RUnlock()
	argsForCall := fake.incSeccompProfileErrorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) Pull(arg1 context.Context, arg2 logr.Logger, arg3 string, arg4 string, arg5 string, arg6 *v1.Platform, arg7 bool) (*artifact.PullResult, error) {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		arg1 context.Context
		arg2 logr.Logger
		arg3 string
		arg4 string
		arg5 string
		arg6 *v1.Platform
		arg7 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	stub := fake.PullStub
	fakeReturns := fake.pullReturns
	fake.recordInvocation("Pull", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.pullMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeImpl) PullCalls(stub func(context.Context, logr.Logger, string, string, string, *v1.Platform, bool) (*artifact.PullResult, error)) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = stub
}

func (fake *FakeImpl) PullArgsForCall(i int) (context.Context, logr.Logger, string, string, string, *v1.Platform, bool) {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	argsForCall := fake.pullArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeImpl) PullReturns(result1 *artifact.PullResult, result2 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 *artifact.PullResult
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) PullReturnsOnCall(i int, result1 *artifact.PullResult, result2 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 *artifact.PullResult
			result2 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 *artifact.PullResult
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) PullResultSeccompProfile(arg1 *artifact.PullResult) *v1beta1.SeccompProfile {
	fake.pullResultSeccompProfileMutex.Lock()
	ret, specificReturn := fake.pullResultSeccompProfileReturnsOnCall[len(fake.pullResultSeccompProfileArgsForCall)]
	fake.pullResultSeccompProfileArgsForCall = append(fake.pullResultSeccompProfileArgsForCall, struct {
		arg1 *artifact.PullResult
	}{arg1})
	stub := fake.PullResultSeccompProfileStub
	fakeReturns := fake.pullResultSeccompProfileReturns
	fake.recordInvocation("PullResultSeccompProfile", []interface{}{arg1})
	fake.pullResultSeccompProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) PullResultSeccompProfileCallCount() int {
	fake.pullResultSeccompProfileMutex.RLock()
	defer fake.pullResultSeccompProfileMutex.RUnlock()
	return len(fake.pullResultSeccompProfileArgsForCall)
}

func (fake *FakeImpl) PullResultSeccompProfileCalls(stub func(*artifact.PullResult) *v1beta1.SeccompProfile) {
	fake.pullResultSeccompProfileMutex.Lock()
	defer fake.pullResultSeccompProfileMutex.Unlock()
	fake.PullResultSeccompProfileStub = stub
}

func (fake *FakeImpl) PullResultSeccompProfileArgsForCall(i int) *artifact.PullResult {
	fake.pullResultSeccompProfileMutex.RLock()
	defer fake.pullResultSeccompProfileMutex.RUnlock()
	argsForCall := fake.pullResultSeccompProfileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) PullResultSeccompProfileReturns(result1 *v1beta1.SeccompProfile) {
	fake.pullResultSeccompProfileMutex.Lock()
	defer fake.pullResultSeccompProfileMutex.Unlock()
	fake.PullResultSeccompProfileStub = nil
	fake.pullResultSeccompProfileReturns = struct {
		result1 *v1beta1.SeccompProfile
	}{result1}
}

func (fake *FakeImpl) PullResultSeccompProfileReturnsOnCall(i int, result1 *v1beta1.SeccompProfile) {
	fake.pullResultSeccompProfileMutex.Lock()
	defer fake.pullResultSeccompProfileMutex.Unlock()
	fake.PullResultSeccompProfileStub = nil
	if fake.pullResultSeccompProfileReturnsOnCall == nil {
		fake.pullResultSeccompProfileReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.SeccompProfile
		})
	}
	fake.pullResultSeccompProfileReturnsOnCall[i] = struct {
		result1 *v1beta1.SeccompProfile
	}{result1}
}

func (fake *FakeImpl) PullResultType(arg1 *artifact.PullResult) artifact.PullResultType {
	fake.pullResultTypeMutex.Lock()
	ret, specificReturn := fake.pullResultTypeReturnsOnCall[len(fake.pullResultTypeArgsForCall)]
	fake.pullResultTypeArgsForCall = append(fake.pullResultTypeArgsForCall, struct {
		arg1 *artifact.PullResult
	}{arg1})
	stub := fake.PullResultTypeStub
	fakeReturns := fake.pullResultTypeReturns
	fake.recordInvocation("PullResultType", []interface{}{arg1})
	fake.pullResultTypeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) PullResultTypeCallCount() int {
	fake.pullResultTypeMutex.RLock()
	defer fake.pullResultTypeMutex.RUnlock()
	return len(fake.pullResultTypeArgsForCall)
}

func (fake *FakeImpl) PullResultTypeCalls(stub func(*artifact.PullResult) artifact.PullResultType) {
	fake.pullResultTypeMutex.Lock()
	defer fake.pullResultTypeMutex.Unlock()
	fake.PullResultTypeStub = stub
}

func (fake *FakeImpl) PullResultTypeArgsForCall(i int) *artifact.PullResult {
	fake.pullResultTypeMutex.RLock()
	defer fake.pullResultTypeMutex.RUnlock()
	argsForCall := fake.pullResultTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) PullResultTypeReturns(result1 artifact.PullResultType) {
	fake.pullResultTypeMutex.Lock()
	defer fake.pullResultTypeMutex.Unlock()
	fake.PullResultTypeStub = nil
	fake.pullResultTypeReturns = struct {
		result1 artifact.PullResultType
	}{result1}
}

func (fake *FakeImpl) PullResultTypeReturnsOnCall(i int, result1 artifact.PullResultType) {
	fake.pullResultTypeMutex.Lock()
	defer fake.pullResultTypeMutex.Unlock()
	fake.PullResultTypeStub = nil
	if fake.pullResultTypeReturnsOnCall == nil {
		fake.pullResultTypeReturnsOnCall = make(map[int]struct {
			result1 artifact.PullResultType
		})
	}
	fake.pullResultTypeReturnsOnCall[i] = struct {
		result1 artifact.PullResultType
	}{result1}
}

func (fake *FakeImpl) RecordEvent(arg1 record.EventRecorder, arg2 runtime.Object, arg3 string, arg4 string, arg5 string) {
	fake.recordEventMutex.Lock()
	fake.recordEventArgsForCall = append(fake.recordEventArgsForCall, struct {
		arg1 record.EventRecorder
		arg2 runtime.Object
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.RecordEventStub
	fake.recordInvocation("RecordEvent", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.recordEventMutex.Unlock()
	if stub != nil {
		fake.RecordEventStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakeImpl) RecordEventCallCount() int {
	fake.recordEventMutex.RLock()
	defer fake.recordEventMutex.RUnlock()
	return len(fake.recordEventArgsForCall)
}

func (fake *FakeImpl) RecordEventCalls(stub func(record.EventRecorder, runtime.Object, string, string, string)) {
	fake.recordEventMutex.Lock()
	defer fake.recordEventMutex.Unlock()
	fake.RecordEventStub = stub
}

func (fake *FakeImpl) RecordEventArgsForCall(i int) (record.EventRecorder, runtime.Object, string, string, string) {
	fake.recordEventMutex.RLock()
	defer fake.recordEventMutex.RUnlock()
	argsForCall := fake.recordEventArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientGetProfileMutex.RLock()
	defer fake.clientGetProfileMutex.RUnlock()
	fake.getSPODMutex.RLock()
	defer fake.getSPODMutex.RUnlock()
	fake.incSeccompProfileErrorMutex.RLock()
	defer fake.incSeccompProfileErrorMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	fake.pullResultSeccompProfileMutex.RLock()
	defer fake.pullResultSeccompProfileMutex.RUnlock()
	fake.pullResultTypeMutex.RLock()
	defer fake.pullResultTypeMutex.RUnlock()
	fake.recordEventMutex.RLock()
	defer fake.recordEventMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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
