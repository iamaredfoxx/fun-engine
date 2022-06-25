package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	assert.NotNil(t, Route(), "Route can not be nil")
}
