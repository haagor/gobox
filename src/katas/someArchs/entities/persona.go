package entities

import (
    "github.com/satori/go.uuid"
)

type Persona struct {
    Id uuid.UUID
    Inventory Inventory
}

func NewPersona() *Persona {
    id := uuid.NewV4()
    p := Persona{Id: id}
    p.Inventory = *NewInventory()
    return &p
}

func (p *Persona) GiveMoney(receiver *Persona, money int) {
    p.Inventory.Money -= money
    receiver.Inventory.Money += money
}