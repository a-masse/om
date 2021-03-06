// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type RegenerateCertificatesService struct {
	RegenerateCertificatesStub        func() error
	regenerateCertificatesMutex       sync.RWMutex
	regenerateCertificatesArgsForCall []struct {
	}
	regenerateCertificatesReturns struct {
		result1 error
	}
	regenerateCertificatesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RegenerateCertificatesService) RegenerateCertificates() error {
	fake.regenerateCertificatesMutex.Lock()
	ret, specificReturn := fake.regenerateCertificatesReturnsOnCall[len(fake.regenerateCertificatesArgsForCall)]
	fake.regenerateCertificatesArgsForCall = append(fake.regenerateCertificatesArgsForCall, struct {
	}{})
	fake.recordInvocation("RegenerateCertificates", []interface{}{})
	fake.regenerateCertificatesMutex.Unlock()
	if fake.RegenerateCertificatesStub != nil {
		return fake.RegenerateCertificatesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.regenerateCertificatesReturns
	return fakeReturns.result1
}

func (fake *RegenerateCertificatesService) RegenerateCertificatesCallCount() int {
	fake.regenerateCertificatesMutex.RLock()
	defer fake.regenerateCertificatesMutex.RUnlock()
	return len(fake.regenerateCertificatesArgsForCall)
}

func (fake *RegenerateCertificatesService) RegenerateCertificatesCalls(stub func() error) {
	fake.regenerateCertificatesMutex.Lock()
	defer fake.regenerateCertificatesMutex.Unlock()
	fake.RegenerateCertificatesStub = stub
}

func (fake *RegenerateCertificatesService) RegenerateCertificatesReturns(result1 error) {
	fake.regenerateCertificatesMutex.Lock()
	defer fake.regenerateCertificatesMutex.Unlock()
	fake.RegenerateCertificatesStub = nil
	fake.regenerateCertificatesReturns = struct {
		result1 error
	}{result1}
}

func (fake *RegenerateCertificatesService) RegenerateCertificatesReturnsOnCall(i int, result1 error) {
	fake.regenerateCertificatesMutex.Lock()
	defer fake.regenerateCertificatesMutex.Unlock()
	fake.RegenerateCertificatesStub = nil
	if fake.regenerateCertificatesReturnsOnCall == nil {
		fake.regenerateCertificatesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.regenerateCertificatesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *RegenerateCertificatesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.regenerateCertificatesMutex.RLock()
	defer fake.regenerateCertificatesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RegenerateCertificatesService) recordInvocation(key string, args []interface{}) {
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
