// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"context"
	"sync"

	"github.com/kyma-incubator/compass/components/system-broker/pkg/director"
	"github.com/kyma-incubator/compass/components/system-broker/pkg/types"
)

type FakeBundleCredentialsDeleteRequester struct {
	RequestBundleInstanceCredentialsDeletionStub        func(context.Context, *director.BundleInstanceAuthDeletionInput) (*director.BundleInstanceAuthDeletionOutput, error)
	requestBundleInstanceCredentialsDeletionMutex       sync.RWMutex
	requestBundleInstanceCredentialsDeletionArgsForCall []struct {
		arg1 context.Context
		arg2 *director.BundleInstanceAuthDeletionInput
	}
	requestBundleInstanceCredentialsDeletionReturns struct {
		result1 *director.BundleInstanceAuthDeletionOutput
		result2 error
	}
	requestBundleInstanceCredentialsDeletionReturnsOnCall map[int]struct {
		result1 *director.BundleInstanceAuthDeletionOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletion(arg1 context.Context, arg2 *director.BundleInstanceAuthDeletionInput) (*director.BundleInstanceAuthDeletionOutput, error) {
	fake.requestBundleInstanceCredentialsDeletionMutex.Lock()
	ret, specificReturn := fake.requestBundleInstanceCredentialsDeletionReturnsOnCall[len(fake.requestBundleInstanceCredentialsDeletionArgsForCall)]
	fake.requestBundleInstanceCredentialsDeletionArgsForCall = append(fake.requestBundleInstanceCredentialsDeletionArgsForCall, struct {
		arg1 context.Context
		arg2 *director.BundleInstanceAuthDeletionInput
	}{arg1, arg2})
	stub := fake.RequestBundleInstanceCredentialsDeletionStub
	fakeReturns := fake.requestBundleInstanceCredentialsDeletionReturns
	fake.recordInvocation("RequestBundleInstanceCredentialsDeletion", []interface{}{arg1, arg2})
	fake.requestBundleInstanceCredentialsDeletionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletionCallCount() int {
	fake.requestBundleInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.RUnlock()
	return len(fake.requestBundleInstanceCredentialsDeletionArgsForCall)
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletionCalls(stub func(context.Context, *director.BundleInstanceAuthDeletionInput) (*director.BundleInstanceAuthDeletionOutput, error)) {
	fake.requestBundleInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestBundleInstanceCredentialsDeletionStub = stub
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletionArgsForCall(i int) (context.Context, *director.BundleInstanceAuthDeletionInput) {
	fake.requestBundleInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.RUnlock()
	argsForCall := fake.requestBundleInstanceCredentialsDeletionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletionReturns(result1 *director.BundleInstanceAuthDeletionOutput, result2 error) {
	fake.requestBundleInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestBundleInstanceCredentialsDeletionStub = nil
	fake.requestBundleInstanceCredentialsDeletionReturns = struct {
		result1 *director.BundleInstanceAuthDeletionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleCredentialsDeleteRequester) RequestBundleInstanceCredentialsDeletionReturnsOnCall(i int, result1 *director.BundleInstanceAuthDeletionOutput, result2 error) {
	fake.requestBundleInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestBundleInstanceCredentialsDeletionStub = nil
	if fake.requestBundleInstanceCredentialsDeletionReturnsOnCall == nil {
		fake.requestBundleInstanceCredentialsDeletionReturnsOnCall = make(map[int]struct {
			result1 *director.BundleInstanceAuthDeletionOutput
			result2 error
		})
	}
	fake.requestBundleInstanceCredentialsDeletionReturnsOnCall[i] = struct {
		result1 *director.BundleInstanceAuthDeletionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleCredentialsDeleteRequester) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestBundleInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestBundleInstanceCredentialsDeletionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBundleCredentialsDeleteRequester) recordInvocation(key string, args []interface{}) {
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

var _ types.BundleCredentialsDeleteRequester = new(FakeBundleCredentialsDeleteRequester)
