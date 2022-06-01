package main

import (
	"fmt"
	"strconv"
	"strings"
)

var digitMapping = map[string]string{
	"0": "011.101.011",
	"1": "000.000.011",
	"2": "001.111.010",
	"3": "000.111.011",
	"4": "010.010.011",
	"5": "010.111.001",
	"6": "011.111.001",
	"7": "000.100.011",
	"8": "011.111.011",
	"9": "010.111.011"}

func readDigits(digits string) string {

	splited := strings.Split(digits, "\n")
	line1 := splited[0]
	line2 := splited[1]
	line3 := splited[2]

	var digit string
	var nextDigits string

	var translateDigit string
	for i := 0; i <= 2; i++ {
		for _, line := range []string{line1, line2, line3} {
			if string(line[i]) != " " {
				translateDigit += "1"
			} else {
				translateDigit += "0"
			}
		}
		translateDigit += "."
	}
	translateDigit = translateDigit[:len(translateDigit)-1]

	for k, v := range digitMapping {
		if translateDigit == v {
			digit = k
		}
	}
	if digit == "" {
		digit = "?"
	}

	if len(line1) > 3 {
		nextDigits = line1[3:] + "\n" + line2[3:] + "\n" + line3[3:]
	} else {
		nextDigits = ""
	}

	if nextDigits == "" {
		return digit
	} else {
		return digit + readDigits(nextDigits)
	}
}

func checksum(convertedDigits string) bool {
	sum := 0
	for i := 0; i <= 8; i++ {
		s, _ := strconv.Atoi(string(convertedDigits[i]))
		sum += int(s) * (9 - i)
	}
	fmt.Println(sum)
	if sum%11 == 0 {
		return true
	} else {
		return false
	}
}

func handleDigits(digits string) string {
	s := readDigits(digits)
	if strings.Contains(s, "?") {
		return s + " ILL"
	}
	if !checksum(s) {
		return s + " ERR"
	}
	return s
}

func main() {
	ex0 := " _  _  _  _  _  _  _  _  _ \n| || || || || || || || || |\n|_||_||_||_||_||_||_||_||_|"
	fmt.Println(handleDigits(ex0))

	ex1 := "                           \n  |  |  |  |  |  |  |  |  |\n  |  |  |  |  |  |  |  |  |"
	fmt.Println(handleDigits(ex1))

	ex2 := " _  _  _  _  _  _  _  _  _ \n _| _| _| _| _| _| _| _| _|\n|_ |_ |_ |_ |_ |_ |_ |_ |_ "
	fmt.Println(handleDigits(ex2))

	ex3 := " _  _  _  _  _  _  _  _  _ \n _| _| _| _| _| _| _| _| _|\n _| _| _| _| _| _| _| _| _|"
	fmt.Println(handleDigits(ex3))

	ex4 := "    _  _     _  _  _  _  _ \n  | _| _||_||_ |_   ||_||_|\n  ||_  _|  | _||_|  ||_| _|"
	fmt.Println(handleDigits(ex4))

	ex5 := "    _  _     _  _  _  _  _ \n  | _| _||_| _ |_   ||_||_|\n  ||_  _|  | _||_|  ||_| _ "
	fmt.Println(handleDigits(ex5))
}
