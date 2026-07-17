package main

import (
	"Fingers/convert"
	"Fingers/db"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var (
	units = map[string]convert.Unit{
		"Meters":      convert.Meter,
		"Centimeters": convert.Cm,
		"Kilometers":  convert.Km,
		"Inches":      convert.Inch,
		"Feet":        convert.Foot,
		"Yards":       convert.Yard,
		"Miles":       convert.Mile,
	}

	first = "Meters"
	unit  = units[first]
)

func main() {
	a := app.New()
	w := a.NewWindow("Fingers")
	w.Resize(fyne.NewSize(600, 400))

	database := db.Get()
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
				return 3
			}
			list.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
				var key string = database.RandKey()
				// This converts length of [unit]s to the amount of fingers it is equivalent to
				var metric db.Length = convert.InMeter(database[key]) * length * unit

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
