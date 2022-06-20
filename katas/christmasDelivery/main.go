package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type gift struct {
	id uuid.UUID
}

type toyMachine struct {
	pipeDelivery chan gift
}

func (tm *toyMachine) createGift(giftAmount int) {
	var g gift
	for i := 1; i <= giftAmount; i++ {
		g.id = uuid.New()
		tm.pipeDelivery <- g
		time.Sleep(200 * time.Millisecond)
		fmt.Println("* 1 gift created *")
	}
	fmt.Println("* Creating gifts finish *")
	close(tm.pipeDelivery)
}

type elf struct {
	name                 string
	pipeDeliveryAssigned chan gift
}

func (e *elf) handleGift() {
	for g := range e.pipeDeliveryAssigned {
		fmt.Printf("\"Gift n° %s handle captain !\"\n", g.id)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var done sync.WaitGroup

	pd := make(chan gift, 100)
	tm := toyMachine{pd}

	done.Add(1)
	go func() {
		tm.createGift(10)
		done.Done()
	}()

	e := elf{"Tingle", tm.pipeDelivery}
	go e.handleGift()

	done.Wait()
}
