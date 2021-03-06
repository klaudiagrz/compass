// Code generated by counterfeiter. DO NOT EDIT.
package httpfakes

import (
	httpa "net/http"
	"sync"

	"github.com/kyma-incubator/compass/components/system-broker/pkg/http"
)

type FakeClient struct {
	DoStub        func(*httpa.Request) (*httpa.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 *httpa.Request
	}
	doReturns struct {
		result1 *httpa.Response
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 *httpa.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Do(arg1 *httpa.Request) (*httpa.Response, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 *httpa.Request
	}{arg1})
	stub := fake.DoStub
	fakeReturns := fake.doReturns
	fake.recordInvocation("Do", []interface{}{arg1})
	fake.doMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeClient) DoCalls(stub func(*httpa.Request) (*httpa.Response, error)) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *FakeClient) DoArgsForCall(i int) *httpa.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DoReturns(result1 *httpa.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *httpa.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DoReturnsOnCall(i int, result1 *httpa.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 *httpa.Response
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 *httpa.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ http.Client = new(FakeClient)
