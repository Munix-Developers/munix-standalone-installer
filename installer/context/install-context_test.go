package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueRead(t *testing.T) {
	a := assert.New(t)

	context := New()
	context.SetVar("variable", "value")

	a.Equal("value", context.GetVar("variable"))
}
