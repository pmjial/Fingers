package main

import (
	"Fingers/db"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var (
	units    db.Database = db.Units()
	database db.Database = db.Fingers()

	first string    = units.RandKey()
	unit  db.Length = units[first]
)

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
			list.Length = func() int {
				return len(database) / 5
			}
			list.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
				var key string = database.RandKey()
				// This converts length of [unit]s to the amount of fingers it is equivalent to
				var metric db.Length = length * unit / database[key]

				item.(*widget.Label).SetText(
					fmt.Sprintf("%.4f\t%s", metric, key),
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
		nil,
		nil,
		nil,
		list,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
