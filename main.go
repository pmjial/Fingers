package main

import (
	"slices"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Depermitto/Fingers/db"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var (
	units    db.Database = db.Units()
	database db.Database = db.Fingers()

	first string    = units.RandKey()
	unit  db.Length = units[first]

	randKeys = db.Keys(database)[:len(database)/2]
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

	p := message.NewPrinter(language.English)
	input := &widget.Entry{
		PlaceHolder: "Input number",
		OnChanged: func(s string) {
			length, _ := strconv.ParseFloat(s, 64)
			randomizeKeys()

			list.Length = func() int {
				return len(randKeys)
			}
			list.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
				// This converts length of [unit]s to the amount of fingers it is equivalent to
				var metric db.Length = length * unit / database[randKeys[id]]

				fmtted := p.Sprint(number.Decimal(metric))
				if len(fmtted) > 6 {
					fmtted = p.Sprint(number.Scientific(metric, number.Scale(2)))
				}

				item.(*widget.Label).SetText(p.Sprintf("%-15v%v", fmtted, randKeys[id]))
			}
			w.Content().Refresh()
		},
	}

	content := container.NewBorder(
		container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("Convert: ", input),
			),
			container.NewHBox(
				widget.NewForm(
					widget.NewFormItem("Select units: ", options),
				),
				widget.NewButton("Reload", func() {
					randomizeKeys()
					list.Refresh()
				}),
			),
		),
		nil,
		nil,
		nil,
		list,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
