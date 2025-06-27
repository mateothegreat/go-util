package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPick(t *testing.T) {
	nilled := (*string)(nil)
	empty := ""
	filled := "filled"

	assert.Equal(t, &filled, PickHasValue(nilled, &empty, &filled))
}
