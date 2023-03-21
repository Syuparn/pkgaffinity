// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package domain

import (
	"sync"
)

// Ensure, that AntiAffinityGroupRuleRepositoryMock does implement AntiAffinityGroupRuleRepository.
// If this is not the case, regenerate this file with moq.
var _ AntiAffinityGroupRuleRepository = &AntiAffinityGroupRuleRepositoryMock{}

// AntiAffinityGroupRuleRepositoryMock is a mock implementation of AntiAffinityGroupRuleRepository.
//
//	func TestSomethingThatUsesAntiAffinityGroupRuleRepository(t *testing.T) {
//
//		// make and configure a mocked AntiAffinityGroupRuleRepository
//		mockedAntiAffinityGroupRuleRepository := &AntiAffinityGroupRuleRepositoryMock{
//			ListByPathFunc: func(path Path) ([]*AntiAffinityGroupRule, error) {
//				panic("mock out the ListByPath method")
//			},
//		}
//
//		// use mockedAntiAffinityGroupRuleRepository in code that requires AntiAffinityGroupRuleRepository
//		// and then make assertions.
//
//	}
type AntiAffinityGroupRuleRepositoryMock struct {
	// ListByPathFunc mocks the ListByPath method.
	ListByPathFunc func(path Path) ([]*AntiAffinityGroupRule, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListByPath holds details about calls to the ListByPath method.
		ListByPath []struct {
			// Path is the path argument value.
			Path Path
		}
	}
	lockListByPath sync.RWMutex
}

// ListByPath calls ListByPathFunc.
func (mock *AntiAffinityGroupRuleRepositoryMock) ListByPath(path Path) ([]*AntiAffinityGroupRule, error) {
	if mock.ListByPathFunc == nil {
		panic("AntiAffinityGroupRuleRepositoryMock.ListByPathFunc: method is nil but AntiAffinityGroupRuleRepository.ListByPath was just called")
	}
	callInfo := struct {
		Path Path
	}{
		Path: path,
	}
	mock.lockListByPath.Lock()
	mock.calls.ListByPath = append(mock.calls.ListByPath, callInfo)
	mock.lockListByPath.Unlock()
	return mock.ListByPathFunc(path)
}

// ListByPathCalls gets all the calls that were made to ListByPath.
// Check the length with:
//
//	len(mockedAntiAffinityGroupRuleRepository.ListByPathCalls())
func (mock *AntiAffinityGroupRuleRepositoryMock) ListByPathCalls() []struct {
	Path Path
} {
	var calls []struct {
		Path Path
	}
	mock.lockListByPath.RLock()
	calls = mock.calls.ListByPath
	mock.lockListByPath.RUnlock()
	return calls
}