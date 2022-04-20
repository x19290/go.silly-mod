package naked

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNaked(t *testing.T) {
	assert.Equal(t, version, Version)
}
