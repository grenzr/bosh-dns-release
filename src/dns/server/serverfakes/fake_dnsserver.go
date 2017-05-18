// Code generated by counterfeiter. DO NOT EDIT.
package serverfakes

import (
	"sync"

	"github.com/cloudfoundry/dns-release/src/dns/server"
)

type FakeDNSServer struct {
	ListenAndServeStub        func() error
	listenAndServeMutex       sync.RWMutex
	listenAndServeArgsForCall []struct{}
	listenAndServeReturns     struct {
		result1 error
	}
	listenAndServeReturnsOnCall map[int]struct {
		result1 error
	}
	ShutdownStub        func() error
	shutdownMutex       sync.RWMutex
	shutdownArgsForCall []struct{}
	shutdownReturns     struct {
		result1 error
	}
	shutdownReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDNSServer) ListenAndServe() error {
	fake.listenAndServeMutex.Lock()
	ret, specificReturn := fake.listenAndServeReturnsOnCall[len(fake.listenAndServeArgsForCall)]
	fake.listenAndServeArgsForCall = append(fake.listenAndServeArgsForCall, struct{}{})
	fake.recordInvocation("ListenAndServe", []interface{}{})
	fake.listenAndServeMutex.Unlock()
	if fake.ListenAndServeStub != nil {
		return fake.ListenAndServeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listenAndServeReturns.result1
}

func (fake *FakeDNSServer) ListenAndServeCallCount() int {
	fake.listenAndServeMutex.RLock()
	defer fake.listenAndServeMutex.RUnlock()
	return len(fake.listenAndServeArgsForCall)
}

func (fake *FakeDNSServer) ListenAndServeReturns(result1 error) {
	fake.ListenAndServeStub = nil
	fake.listenAndServeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSServer) ListenAndServeReturnsOnCall(i int, result1 error) {
	fake.ListenAndServeStub = nil
	if fake.listenAndServeReturnsOnCall == nil {
		fake.listenAndServeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listenAndServeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSServer) Shutdown() error {
	fake.shutdownMutex.Lock()
	ret, specificReturn := fake.shutdownReturnsOnCall[len(fake.shutdownArgsForCall)]
	fake.shutdownArgsForCall = append(fake.shutdownArgsForCall, struct{}{})
	fake.recordInvocation("Shutdown", []interface{}{})
	fake.shutdownMutex.Unlock()
	if fake.ShutdownStub != nil {
		return fake.ShutdownStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shutdownReturns.result1
}

func (fake *FakeDNSServer) ShutdownCallCount() int {
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	return len(fake.shutdownArgsForCall)
}

func (fake *FakeDNSServer) ShutdownReturns(result1 error) {
	fake.ShutdownStub = nil
	fake.shutdownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSServer) ShutdownReturnsOnCall(i int, result1 error) {
	fake.ShutdownStub = nil
	if fake.shutdownReturnsOnCall == nil {
		fake.shutdownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.shutdownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDNSServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listenAndServeMutex.RLock()
	defer fake.listenAndServeMutex.RUnlock()
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDNSServer) recordInvocation(key string, args []interface{}) {
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

var _ server.DNSServer = new(FakeDNSServer)
