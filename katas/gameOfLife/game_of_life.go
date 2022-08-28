package main

import (
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    initWorld(8, 16)
}

func initWorld(x, y int) {
    f, err := os.Create("current_gen.txt")
    check(err)
    defer f.Close()
    line := strings.Repeat(".", y) + "\n"
    starter := "***" + strings.Repeat(".", y-3) + "\n"

    _, err = f.WriteString(starter)
    check(err)

    for i := 1; i < x; i++ {
        _, err = f.WriteString(line)
        check(err)
    }
}
