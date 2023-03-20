// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecase

import (
	"sync"
)

// Ensure, that CheckImportsOutputPortMock does implement CheckImportsOutputPort.
// If this is not the case, regenerate this file with moq.
var _ CheckImportsOutputPort = &CheckImportsOutputPortMock{}

// CheckImportsOutputPortMock is a mock implementation of CheckImportsOutputPort.
//
//	func TestSomethingThatUsesCheckImportsOutputPort(t *testing.T) {
//
//		// make and configure a mocked CheckImportsOutputPort
//		mockedCheckImportsOutputPort := &CheckImportsOutputPortMock{
//			PresentFunc: func(checkImportsOutputData *CheckImportsOutputData)  {
//				panic("mock out the Present method")
//			},
//		}
//
//		// use mockedCheckImportsOutputPort in code that requires CheckImportsOutputPort
//		// and then make assertions.
//
//	}
type CheckImportsOutputPortMock struct {
	// PresentFunc mocks the Present method.
	PresentFunc func(checkImportsOutputData *CheckImportsOutputData)

	// calls tracks calls to the methods.
	calls struct {
		// Present holds details about calls to the Present method.
		Present []struct {
			// CheckImportsOutputData is the checkImportsOutputData argument value.
			CheckImportsOutputData *CheckImportsOutputData
		}
	}
	lockPresent sync.RWMutex
}

// Present calls PresentFunc.
func (mock *CheckImportsOutputPortMock) Present(checkImportsOutputData *CheckImportsOutputData) {
	if mock.PresentFunc == nil {
		panic("CheckImportsOutputPortMock.PresentFunc: method is nil but CheckImportsOutputPort.Present was just called")
	}
	callInfo := struct {
		CheckImportsOutputData *CheckImportsOutputData
	}{
		CheckImportsOutputData: checkImportsOutputData,
	}
	mock.lockPresent.Lock()
	mock.calls.Present = append(mock.calls.Present, callInfo)
	mock.lockPresent.Unlock()
	mock.PresentFunc(checkImportsOutputData)
}

// PresentCalls gets all the calls that were made to Present.
// Check the length with:
//
//	len(mockedCheckImportsOutputPort.PresentCalls())
func (mock *CheckImportsOutputPortMock) PresentCalls() []struct {
	CheckImportsOutputData *CheckImportsOutputData
} {
	var calls []struct {
		CheckImportsOutputData *CheckImportsOutputData
	}
	mock.lockPresent.RLock()
	calls = mock.calls.Present
	mock.lockPresent.RUnlock()
	return calls
}
