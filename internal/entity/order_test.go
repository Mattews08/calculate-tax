package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_If_It_Gets_An_Error(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "Invalid id")
}
