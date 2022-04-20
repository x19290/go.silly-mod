package silly

import (
	"github.com/stretchr/testify/assert"
	. "github.com/x19290/go.silly-mod/1testdata"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, Version, private())
}
