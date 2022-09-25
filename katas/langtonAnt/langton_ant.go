package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"

    "github.com/faiface/pixel"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
)

var sizeCoord = 80
var world = initWorld(sizeCoord, sizeCoord)

type ant struct {
    id        int
    x         int
    y         int
    direction int // 0 North, 1 East, 2 South, 3 West
    step      int
}

func (a *ant) nextStep(imd *imdraw.IMDraw) {
    cell, err := strconv.Atoi(world[a.y][a.x])
    check(err)

    a.rotate(cell)
    a.switchCell(imd, cell)
    a.goforward(cell)
    a.step += 1
    fmt.Println(a.step)
}

func (a *ant) rotate(cell int) {
    if cell == 0 {
        a.direction = (a.direction + 1) % 4
    } else {
        if a.direction == 0 {
            a.direction = 3
        } else {
            a.direction -= 1
        }
    }
}

func (a *ant) switchCell(imd *imdraw.IMDraw, cell int) {
    if cell == 1 {
        world[a.y][a.x] = "0"
        drawCell(imd, float64(a.x), float64(a.y), 0)
    } else if cell == 0 {
        world[a.y][a.x] = "1"
        drawCell(imd, float64(a.x), float64(a.y), 1)
    }
}

func (a *ant) goforward(cell int) {
    if a.direction == 0 && a.y > 0 {
        a.y++
    } else if a.direction == 1 && a.x < sizeCoord {
        a.x++
    } else if a.direction == 2 && a.y < sizeCoord {
        a.y--
    } else if a.direction == 3 && a.x > 0 {
        a.x--
    }
}

func (a *ant) draw(imd *imdraw.IMDraw) {
    imd.Color = colornames.Red
    imd.EndShape = imdraw.RoundEndShape

    xDraw := float64(float64(a.x)*10 + 5)
    yDraw := float64(float64(a.y)*10 + 5)

    imd.Push(pixel.V(xDraw, yDraw), pixel.V(xDraw, yDraw))
    imd.Polygon(10)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    pixelgl.Run(run)
}

func run() {
    sizeGrid := float64(sizeCoord * 10)
    cfg := pixelgl.WindowConfig{
        Title:  "LangtonAnt",
        Bounds: pixel.R(0, 0, sizeGrid, sizeGrid),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    win.Clear(colornames.Aliceblue)

    imd := imdraw.New(nil)
    drawGrid(imd, sizeGrid)
    imd.Draw(win)
    win.Update()

    ant1 := ant{id: 1, x: sizeCoord / 2, y: sizeCoord / 2, direction: 0, step: 0}
    ant1.draw(imd)
    imd.Draw(win)
    win.Update()
    time.Sleep(1 * time.Second)

    for !win.Closed() {
        ant1.nextStep(imd)

        if ant1.step%1000 == 0 {
            imd.Draw(win)
            win.Update()
        }
    }
}

func drawCell(imd *imdraw.IMDraw, x, y float64, color int) {
    if color == 0 {
        imd.Color = colornames.White
    } else {
        imd.Color = colornames.Black
    }

    xDraw := float64(x*10 + 5)
    yDraw := float64(y*10 + 5)

    imd.EndShape = imdraw.RoundEndShape
    imd.Push(pixel.V(xDraw, yDraw), pixel.V(xDraw, yDraw))
    imd.Polygon(10)
}

func drawGrid(imd *imdraw.IMDraw, size float64) {
    imd.Color = colornames.Black
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

func initWorld(x, y int) (res [][]string) {
    for i := 0; i < y; i++ {
        line := []string{}
        for j := 0; j < x; j++ {
            line = append(line, "0")
        }

        res = append(res, line)
    }
    return res
}
