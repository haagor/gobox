package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)

type persona struct {
    id uuid.UUID
    inventory inventory
}

func newPersona() *persona {
    id := uuid.NewV4()
    p := persona{id: id}
    p.inventory = *newInventory()
    return &p
}

func main() {

    p := newPersona()
    fmt.Println(p.id)
}