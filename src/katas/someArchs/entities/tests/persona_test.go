package tests

import (
	"testing"

	"github.com/haagor/gobox/src/entities"
	"github.com/stretchr/testify/assert"
)

func TestGiveMoney(t *testing.T) {
	p1 := entities.NewPersona()
	p1.Inventory.Money = 10
	p2 := entities.NewPersona()

	p1.GiveMoney(p2, 10)

	assert.Equal(t, p1.Inventory.Money, 0)
	assert.Equal(t, p2.Inventory.Money, 10)

}
