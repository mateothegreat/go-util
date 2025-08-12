package files

import (
	"testing"

	"github.com/mateothegreat/go-util/paths"
	"github.com/stretchr/testify/assert"
)

func TestExpandPath(t *testing.T) {
	path := "~/test"
	expanded, err := paths.ExpandPath(path)
	assert.Equal(t, "/Users/mateo/test", expanded)
	assert.NoError(t, err)
}
