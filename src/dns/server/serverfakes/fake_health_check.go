// Code generated by counterfeiter. DO NOT EDIT.
package serverfakes

import (
	"sync"

	"github.com/cloudfoundry/dns-release/src/dns/server"
)

type FakeHealthCheck struct {
	IsHealthyStub        func() error
	isHealthyMutex       sync.RWMutex
	isHealthyArgsForCall []struct{}
	isHealthyReturns     struct {
		result1 error
	}
	isHealthyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthCheck) IsHealthy() error {
	fake.isHealthyMutex.Lock()
	ret, specificReturn := fake.isHealthyReturnsOnCall[len(fake.isHealthyArgsForCall)]
	fake.isHealthyArgsForCall = append(fake.isHealthyArgsForCall, struct{}{})
	fake.recordInvocation("IsHealthy", []interface{}{})
	fake.isHealthyMutex.Unlock()
	if fake.IsHealthyStub != nil {
		return fake.IsHealthyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isHealthyReturns.result1
}

func (fake *FakeHealthCheck) IsHealthyCallCount() int {
	fake.isHealthyMutex.RLock()
	defer fake.isHealthyMutex.RUnlock()
	return len(fake.isHealthyArgsForCall)
}

func (fake *FakeHealthCheck) IsHealthyReturns(result1 error) {
	fake.IsHealthyStub = nil
	fake.isHealthyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHealthCheck) IsHealthyReturnsOnCall(i int, result1 error) {
	fake.IsHealthyStub = nil
	if fake.isHealthyReturnsOnCall == nil {
		fake.isHealthyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.isHealthyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHealthCheck) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isHealthyMutex.RLock()
	defer fake.isHealthyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthCheck) recordInvocation(key string, args []interface{}) {
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

var _ server.HealthCheck = new(FakeHealthCheck)
