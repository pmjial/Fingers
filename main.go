package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Depermitto/Fingers/db"
	"slices"
	"strconv"
)

var (
	units    db.Database = db.Units()
	database db.Database = db.Fingers()

	first string    = units.RandKey()
	unit  db.Length = units[first]

	randKeys = db.Keys(database)[:len(database)/5]
)

func randomizeKeys() {
	for i := range randKeys {
		key := database.RandKey()
		for slices.Contains(randKeys, key) {
			key = database.RandKey()
		}
		randKeys[i] = key
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Fingers")
	w.Resize(fyne.NewSize(600, 400))

	list := widget.NewList(
		func() int {
			return 0
		},
		func() fyne.CanvasObject {
			return new(widget.Label)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {},
	)

	options := &widget.Select{
		Options:  db.Keys(units),
		Selected: first,
		OnChanged: func(key string) {
			unit = units[key]
			w.Content().Refresh()
		},
	}

	input := &widget.Entry{
		PlaceHolder: "Input number",
		OnChanged: func(s string) {
			length, _ := strconv.ParseFloat(s, 64)
			amount := len(database) / 5
			randomizeKeys()

			list.Length = func() int {
				return amount
			}
			list.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
				// This converts length of [unit]s to the amount of fingers it is equivalent to
				var metric db.Length = length * unit / database[randKeys[id]]

				item.(*widget.Label).SetText(
					fmt.Sprintf("%.4f\t%s", metric, randKeys[id]),
				)
			}
			w.Content().Refresh()
		},
	}

	content := container.NewBorder(
		widget.NewForm(
			widget.NewFormItem("Convert: ", input),
			widget.NewFormItem("Select units: ", options),
		),
		widget.NewButton("Reload", func() {
			randomizeKeys()
			list.Refresh()
		}),
		nil,
		nil,
		list,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
