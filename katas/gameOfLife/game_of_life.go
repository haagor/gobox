package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"

    "github.com/faiface/pixel"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
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

    pixelgl.Run(run)
}

func run() {
    cfg := pixelgl.WindowConfig{
        Title:  "GameOfLife",
        Bounds: pixel.R(0, 0, 800, 800),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    win.Clear(colornames.Aliceblue)

    imd := imdraw.New(nil)
    imd.Color = colornames.Black
    drawR(imd)

    for !win.Closed() {
        imd.Clear()
        win.Clear(colornames.Aliceblue)
        drawR(imd)
        imd.Draw(win)
        win.Update()
        time.Sleep(2 * time.Second)
    }
}

func drawR(imd *imdraw.IMDraw) {
    imd.EndShape = imdraw.RoundEndShape
    x := float64(rand.Intn(800))
    y := float64(rand.Intn(800))
    imd.Push(pixel.V(x, y), pixel.V(x, y))
    imd.Rectangle(10)
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
