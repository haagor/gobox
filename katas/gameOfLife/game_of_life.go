package main

import (
    "bufio"
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

var sizeCoord = 80

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    pixelgl.Run(run)
}

func run() {
    initWorldRandom(sizeCoord, sizeCoord)

    sizeGrid := float64(sizeCoord * 10)
    cfg := pixelgl.WindowConfig{
        Title:  "GameOfLife",
        Bounds: pixel.R(0, 0, sizeGrid, sizeGrid),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    win.Clear(colornames.Aliceblue)

    imd := imdraw.New(nil)
    imd.Color = colornames.Black

    for !win.Closed() {
        imd.Clear()
        win.Clear(colornames.Aliceblue)
        drawGrid(imd, sizeGrid)

        nextGen(imd, "current_gen.txt")

        imd.Draw(win)
        win.Update()
        time.Sleep(1 * time.Second)
    }
}

func drawCell(imd *imdraw.IMDraw, x, y float64) {
    imd.EndShape = imdraw.RoundEndShape
    imd.Push(pixel.V(x, y), pixel.V(x, y))
    imd.Circle(5, 0)
}

func drawLife(imd *imdraw.IMDraw, lifes [][]string) {
    for y, line := range lifes {
        for x, cell := range line {
            if cell == "*" {
                xGrid := float64(x*10 + 5)
                yGrid := float64(y*10 + 5)
                drawCell(imd, xGrid, yGrid)
            }
        }
    }
}

func drawGrid(imd *imdraw.IMDraw, size float64) {
    x := float64(10)
    for x < size {
        imd.Push(pixel.V(x, 0), pixel.V(x, size))
        imd.Line(1)
        x += 10
    }
    y := float64(10)
    for y < size {
        imd.Push(pixel.V(0, y), pixel.V(size, y))
        imd.Line(1)
        y += 10
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

func nextGen(imd *imdraw.IMDraw, file string) {
    nextGen := ""

    world := convertWorldFile(file)
    drawLife(imd, world)
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

    //fmt.Println(nextGen)
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
