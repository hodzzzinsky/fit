package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faifasje/pixel/pixelgl"
	"log"
	"os"
)

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	//words := getInput()
	//rMap := initMap(words)
	//SimplePrint(rMap)
	pixelgl.Run(run)
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
