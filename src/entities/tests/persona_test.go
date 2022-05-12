package tests


import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/haagor/gobox/src/entities"
)

func TestGiveMoney(t *testing.T) {
    p1 := entities.NewPersona()
    p1.Inventory.Money = 10
    p2 := entities.NewPersona()

    p1.GiveMoney(p2, 10)

    assert.Equal(t, p1.Inventory.Money, 0)
    assert.Equal(t, p2.Inventory.Money, 10)

}