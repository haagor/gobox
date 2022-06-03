package main

import (
	"fmt"
	"strings"
)

type pastry interface {
	getPrice() float32
	getTopping() string
}

type cupecake struct {
}

func (c *cupecake) getPrice() float32 {
	return 1.0
}

func (c *cupecake) getTopping() string {
	return "cupecake"
}

type cookie struct {
}

func (c *cookie) getPrice() float32 {
	return 2.0
}

func (c *cookie) getTopping() string {
	return "cookie"
}

type chocolateTopping struct {
	pastry pastry
}

func (t *chocolateTopping) getPrice() float32 {
	r := t.pastry.getPrice() + 0.1
	return r
}

func (t *chocolateTopping) getTopping() string {
	var l string
	if strings.Contains(t.pastry.getTopping(), "with") {
		l = "and"

	} else {
		l = "with"
	}
	r := fmt.Sprintf("%s %s %s", t.pastry.getTopping(), l, "chocolate")
	return r
}

type nutsTopping struct {
	pastry pastry
}

func (t *nutsTopping) getPrice() float32 {
	r := t.pastry.getPrice() + 0.2
	return r
}

func (t *nutsTopping) getTopping() string {
	var l string
	if strings.Contains(t.pastry.getTopping(), "with") {
		l = "and"

	} else {
		l = "with"
	}
	r := fmt.Sprintf("%s %s %s", t.pastry.getTopping(), l, "nuts")
	return r
}

type candyTopping struct {
	pastry pastry
}

func (t *candyTopping) getPrice() float32 {
	r := t.pastry.getPrice() + 0.3
	return r
}

func (t *candyTopping) getTopping() string {
	var l string
	if strings.Contains(t.pastry.getTopping(), "with") {
		l = "and"

	} else {
		l = "with"
	}
	r := fmt.Sprintf("%s %s %s", t.pastry.getTopping(), l, "candy")
	return r
}

func main() {

	cupecake := &cupecake{}
	cupecakeWithChocolate := &chocolateTopping{
		pastry: cupecake,
	}
	cupecakeWithChocolateAndNuts := &nutsTopping{
		pastry: cupecakeWithChocolate,
	}
	fmt.Println(cupecakeWithChocolateAndNuts.getPrice())
	fmt.Println(cupecakeWithChocolateAndNuts.getTopping())

}
