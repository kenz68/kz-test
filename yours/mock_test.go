package yours

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

/*
	Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct {
	mock.Mock
}

// DoSomething is a mocked method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the reak object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

/*
	Actual test funtions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestSomething(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MyMockedObject)

	// set up expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)
}

func targetFuncThatDoesSomethingWithObj(testObj *MyMockedObject) {
	panic("unimplemented")
}
