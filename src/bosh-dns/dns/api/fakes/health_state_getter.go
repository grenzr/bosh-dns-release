// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-dns/dns/server/healthiness"
	"sync"
)

type HealthStateGetter struct {
	HealthStateStringStub        func(ip string) string
	healthStateStringMutex       sync.RWMutex
	healthStateStringArgsForCall []struct {
		ip string
	}
	healthStateStringReturns struct {
		result1 string
	}
	healthStateStringReturnsOnCall map[int]struct {
		result1 string
	}
	HealthStateStub        func(ip string) healthiness.HealthState
	healthStateMutex       sync.RWMutex
	healthStateArgsForCall []struct {
		ip string
	}
	healthStateReturns struct {
		result1 healthiness.HealthState
	}
	healthStateReturnsOnCall map[int]struct {
		result1 healthiness.HealthState
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HealthStateGetter) HealthStateString(ip string) string {
	fake.healthStateStringMutex.Lock()
	ret, specificReturn := fake.healthStateStringReturnsOnCall[len(fake.healthStateStringArgsForCall)]
	fake.healthStateStringArgsForCall = append(fake.healthStateStringArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("HealthStateString", []interface{}{ip})
	fake.healthStateStringMutex.Unlock()
	if fake.HealthStateStringStub != nil {
		return fake.HealthStateStringStub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.healthStateStringReturns.result1
}

func (fake *HealthStateGetter) HealthStateStringCallCount() int {
	fake.healthStateStringMutex.RLock()
	defer fake.healthStateStringMutex.RUnlock()
	return len(fake.healthStateStringArgsForCall)
}

func (fake *HealthStateGetter) HealthStateStringArgsForCall(i int) string {
	fake.healthStateStringMutex.RLock()
	defer fake.healthStateStringMutex.RUnlock()
	return fake.healthStateStringArgsForCall[i].ip
}

func (fake *HealthStateGetter) HealthStateStringReturns(result1 string) {
	fake.HealthStateStringStub = nil
	fake.healthStateStringReturns = struct {
		result1 string
	}{result1}
}

func (fake *HealthStateGetter) HealthStateStringReturnsOnCall(i int, result1 string) {
	fake.HealthStateStringStub = nil
	if fake.healthStateStringReturnsOnCall == nil {
		fake.healthStateStringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.healthStateStringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *HealthStateGetter) HealthState(ip string) healthiness.HealthState {
	fake.healthStateMutex.Lock()
	ret, specificReturn := fake.healthStateReturnsOnCall[len(fake.healthStateArgsForCall)]
	fake.healthStateArgsForCall = append(fake.healthStateArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("HealthState", []interface{}{ip})
	fake.healthStateMutex.Unlock()
	if fake.HealthStateStub != nil {
		return fake.HealthStateStub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.healthStateReturns.result1
}

func (fake *HealthStateGetter) HealthStateCallCount() int {
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return len(fake.healthStateArgsForCall)
}

func (fake *HealthStateGetter) HealthStateArgsForCall(i int) string {
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return fake.healthStateArgsForCall[i].ip
}

func (fake *HealthStateGetter) HealthStateReturns(result1 healthiness.HealthState) {
	fake.HealthStateStub = nil
	fake.healthStateReturns = struct {
		result1 healthiness.HealthState
	}{result1}
}

func (fake *HealthStateGetter) HealthStateReturnsOnCall(i int, result1 healthiness.HealthState) {
	fake.HealthStateStub = nil
	if fake.healthStateReturnsOnCall == nil {
		fake.healthStateReturnsOnCall = make(map[int]struct {
			result1 healthiness.HealthState
		})
	}
	fake.healthStateReturnsOnCall[i] = struct {
		result1 healthiness.HealthState
	}{result1}
}

func (fake *HealthStateGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.healthStateStringMutex.RLock()
	defer fake.healthStateStringMutex.RUnlock()
	fake.healthStateMutex.RLock()
	defer fake.healthStateMutex.RUnlock()
	return fake.invocations
}

func (fake *HealthStateGetter) recordInvocation(key string, args []interface{}) {
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
