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

type mrsClaus struct {
	tm              toyMachine
	pipeElfDelivery chan gift
}

func (mc *mrsClaus) handleToyMachineDelivery() {
	var standByGift []gift
	go func() {
		for g := range mc.tm.pipeDelivery {
			standByGift = append(standByGift, g)
		}
	}()

	go func() {
		for {
			if len(standByGift) > 0 {
				mc.pipeElfDelivery <- standByGift[0]
				standByGift = standByGift[1:]
			}
		}
	}()
}

func main() {
	var done sync.WaitGroup

	pd := make(chan gift, 100)
	tm := toyMachine{pd}

	ped := make(chan gift, 1)
	claus := mrsClaus{tm, ped}
	go claus.handleToyMachineDelivery()

	done.Add(1)
	go func() {
		tm.createGift(10)
		done.Done()
	}()

	e1 := elf{"Tingle", claus.pipeElfDelivery}
	e2 := elf{"Bob", claus.pipeElfDelivery}
	e3 := elf{"Anna", claus.pipeElfDelivery}
	go e1.handleGift()
	go e2.handleGift()
	go e3.handleGift()

	done.Wait()
	time.Sleep(2000 * time.Millisecond)
}
