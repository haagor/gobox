package entities

import (
    "github.com/satori/go.uuid"
)

type Inventory struct {
    Id uuid.UUID
    Money int
    FournituresQuantityKg int
}

func NewInventory() *Inventory {
    id := uuid.NewV4()
    i := Inventory{Id: id}
    i.Money = 0
    i.FournituresQuantityKg = 0
    return &i
}
