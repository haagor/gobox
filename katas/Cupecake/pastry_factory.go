package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pastry interface {
	getPrice() string
	getTopping() string
}

type cupecake struct {
}

func (c *cupecake) getPrice() string {
	r := fmt.Sprintf("%.2f$ for \"%s\"", 1.0, c.getTopping())
	return r
}

func (c *cupecake) getTopping() string {
	return "cupecake"
}

type cookie struct {
}

func (c *cookie) getPrice() string {
	r := fmt.Sprintf("%.2f$ for \"%s\"", 2.0, c.getTopping())
	return r
}

func (c *cookie) getTopping() string {
	return "cookie"
}

// Topping section
// Chocolate <<<
type chocolateTopping struct {
	pastry pastry
}

func (t *chocolateTopping) getPrice() string {
	s := strings.Split(t.pastry.getPrice(), " for ")
	f, _ := strconv.ParseFloat(s[0][:len(s[0])-1], 32)
	p := f + 0.1

	r := fmt.Sprintf("%.2f$ for \"%s\"", p, t.getTopping())
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

// Nuts <<<
type nutsTopping struct {
	pastry pastry
}

func (t *nutsTopping) getPrice() string {
	s := strings.Split(t.pastry.getPrice(), " for ")
	f, _ := strconv.ParseFloat(s[0][:len(s[0])-1], 32)
	p := f + 0.2

	r := fmt.Sprintf("%.2f$ for \"%s\"", p, t.getTopping())
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

// Candy <<<
type candyTopping struct {
	pastry pastry
}

func (t *candyTopping) getPrice() string {
	s := strings.Split(t.pastry.getPrice(), " for ")
	f, _ := strconv.ParseFloat(s[0][:len(s[0])-1], 32)
	p := f + 0.3

	r := fmt.Sprintf("%.2f$ for \"%s\"", p, t.getTopping())
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
	fmt.Println(cupecakeWithChocolate.getPrice())
	fmt.Println(cupecakeWithChocolate.getTopping())

	cupecakeWithChocolateAndNuts := &nutsTopping{
		pastry: cupecakeWithChocolate,
	}
	fmt.Println(cupecakeWithChocolateAndNuts.getPrice())
	fmt.Println(cupecakeWithChocolateAndNuts.getTopping())

	cupecakeWithChocolateAndNutsAndCandy := &candyTopping{
		pastry: cupecakeWithChocolateAndNuts,
	}
	fmt.Println(cupecakeWithChocolateAndNutsAndCandy.getPrice())
	fmt.Println(cupecakeWithChocolateAndNutsAndCandy.getTopping())

	cookie := &cookie{}
	cookieWithChocolate := &chocolateTopping{
		pastry: cookie,
	}
	fmt.Println(cookieWithChocolate.getPrice())
	fmt.Println(cookieWithChocolate.getTopping())

}
