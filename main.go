package main

import (
	"math/rand"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func numbers() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "sudoku",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(numbers)
	app.Run()
}
