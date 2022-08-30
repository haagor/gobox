package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    initWorldRandom(8, 16)
    for i := 1; i < 3; i++ {
        nextGen("current_gen.txt")
        fmt.Println("---")
    }
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

func initWorldRandom(x, y int) {
    f, err := os.Create("current_gen.txt")
    check(err)
    defer f.Close()

    for i := 0; i < x; i++ {
        line := ""
        for j := 0; j < y; j++ {
            if rand.Intn(2) == 0 {
                line += "."
            } else {
                line += "*"
            }
        }
        line += "\n"

        _, err = f.WriteString(line)
        check(err)
    }
}

func nextGen(file string) {
    nextGen := ""

    world := convertWorldFile(file)
    for x, line := range world {
        for y, _ := range line {
            nextGen += evalCell(world, x, y)
        }
        nextGen += "\n"
    }

    f, err := os.OpenFile("current_gen.txt", os.O_RDWR, 0644)
    check(err)
    defer f.Close()
    _, err = f.WriteString(nextGen)
    check(err)

    fmt.Println(nextGen)
}

func convertWorldFile(file string) (res [][]string) {
    f, err := os.Open(file)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)

    i := 0
    for scanner.Scan() {
        line := strings.TrimSuffix(scanner.Text(), "\n")
        res = append(res, []string{})
        for _, s := range line {
            res[i] = append(res[i], string(s))
        }
        i++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return res
}

func evalCell(w [][]string, x int, y int) string {
    liveCell := 0

    if x > 0 && y > 0 {
        if w[x-1][y-1] == "*" {
            liveCell++
        }
    }
    if x > 0 {
        if w[x-1][y] == "*" {
            liveCell++
        }
    }
    if x > 0 && y < len(w[0])-2 {
        if w[x-1][y+1] == "*" {
            liveCell++
        }
    }
    if y > 0 {
        if w[x][y-1] == "*" {
            liveCell++
        }
    }
    if y < len(w[0])-2 {
        if w[x][y+1] == "*" {
            liveCell++
        }
    }
    if x < len(w[0][0])-2 && y > 0 {
        if w[x+1][y-1] == "*" {
            liveCell++
        }
    }
    if x < len(w[0][0])-2 {
        if w[x+1][y] == "*" {
            liveCell++
        }
    }
    if x < len(w[0][0])-2 && y < len(w[0])-2 {
        if w[x+1][y+1] == "*" {
            liveCell++
        }
    }

    if w[x][y] == "." {
        if liveCell == 3 {
            return "*"
        } else {
            return "."
        }
    } else {
        if liveCell < 2 {
            return "."
        } else if liveCell < 4 {
            return "*"
        } else {
            return "."
        }
    }
}
