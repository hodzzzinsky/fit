package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"log"
	"math"
	"os"
)

func run() {

	//words := getInput()
	//length := len(words)
	length := 4

	bounds := pixel.R(0, 0, 1024, 768)
	cells := createGrid(length, bounds)

	log.Println("cells count", length)

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: bounds,
		VSync:  true, // refreshs windows with monitor rate
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//imd := imdraw.New(nil)
	//imd.Push(pixel.V(100, 100), pixel.V(500, 100), pixel.V(100, 600), pixel.V(600, 600))
	//imd.Rectangle(0)

	imd := imdraw.New(nil)
	for _, cell := range cells {

		fmt.Println(cell)
		imd.Color = pixel.RGB(51, 0, 0)
		imd.Push(cell.Rect.Min, cell.Rect.Max)
		imd.Rectangle(1)
	}

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Draw(win)
		win.Update()
	}
}

type Cell struct {
	Rect  pixel.Rect
	Color pixel.RGBA
}

func main() {
	//rMap := initMap(words)
	//SimplePrint(rMap)

	pixelgl.Run(run)
}

func calculateGrid(length int, winWidth, winHeight float64) (cols, rows int, cellSize pixel.Vec) {

	aspect := winWidth / winHeight

	bestScore := math.MaxFloat64

	for testRows := 1; testRows <= length; testRows++ {
		testCols := int(math.Ceil(float64(length)) / float64(testRows))

		gridAspect := float64(testCols) / float64(testRows)
		aspectDiff := math.Abs(gridAspect - aspect)
		emptyCells := testCols*testRows - length

		score := aspectDiff + float64(emptyCells)/float64(length)*0.5

		if score < bestScore {
			bestScore = score
			cols = testCols
			rows = testRows
		}
	}

	cellWidth := winWidth / float64(cols)
	cellHeight := winHeight / float64(rows)

	if cellWidth > winHeight*aspect {
		cellWidth = cellHeight * aspect
	} else {
		cellHeight = cellWidth / aspect
	}

	return cols, rows, pixel.V(cellWidth, cellHeight)
}

func createGrid(length int, winBounds pixel.Rect) []Cell {
	cells := make([]Cell, 0, length)

	gridSize := int(math.Ceil(float64(length)))
	fmt.Println(gridSize)

	cellHeight := winBounds.H() / float64(gridSize)
	cellWidth := winBounds.W() / float64(gridSize)
	fmt.Println("cellH", cellWidth)
	fmt.Println("cellW", cellWidth)

	for y := 0; y < gridSize; y++ {

		for x := 0; x < gridSize; x++ {

			if len(cells) >= length {
				break
			}

			posX := float64(x) * cellWidth
			posY := winBounds.H() - float64(y+1)*cellHeight

			rect := pixel.R(
				posX,
				posY,
				posX+cellWidth,
				posY+cellHeight,
			)

			cc := float64(x) * 15
			cell := Cell{rect, pixel.RGB(cc, cc, cc)}
			cells = append(cells, cell)
		}
	}
	return cells
}

func initMap(words []string) map[string]*popWord {
	rMap := make(map[string]*popWord)

	for _, w := range words {
		pW := rMap[w]
		if pW == nil {
			rMap[w] = &popWord{w, 1, 1}
		} else {

			pW.repeteNum += 1
			rMap[w] = pW
		}
	}

	for _, v := range rMap {
		devision := float32(v.repeteNum) / float32(len(words))
		rating := devision * 100.0
		v.popRating = rating
	}

	return rMap
}

func SimplePrint(rMap map[string]*popWord) {
	for _, v := range rMap {
		rating := ""
		for i := 0; i < v.repeteNum; i++ {
			rating += "-"
		}
		fmt.Printf("%s %f\n", v.name+rating, v.popRating)
	}
}

func getInput() []string {

	words := []string{}

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strVal := scanner.Text()
		if isWordValid(strVal) {
			words = append(words, strVal)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func isWordValid(w string) bool {
	return len(w) < 15
}

type popWord struct {
	name      string
	repeteNum int
	popRating float32
}

func (p popWord) String() string {
	return fmt.Sprintf("{name : \"%s\", repeteNum : %d}", p.name, p.repeteNum)
}
