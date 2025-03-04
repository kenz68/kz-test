package yours

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomthing(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var object interface{}
	assert.Nil(t, object, "object should be nil")

	// assert for not nil (good when you expect something not nil)
	object = "Something"
	if assert.NotNil(t, object) {
		// now we know that object is not nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object)
	}
}
