package main

import (
    "github.com/satori/go.uuid"
)

type inventory struct {
    id uuid.UUID
    money int
    fournituresQuantityKg int
}

func newInventory() *inventory {
    id := uuid.NewV4()
    i := inventory{id: id}
    i.money = 0
    i.fournituresQuantityKg = 0
    return &i
}
