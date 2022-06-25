package funengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHi(t *testing.T) {
	assert.Equal(t,
		"Hi :)",
		SayHi(),
		"World always fun if you enjoyed it ^^",
	)
}
