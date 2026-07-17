package main

import (
	"log"
	"math/rand/v2"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.pmjlab.com/fingers/internal/data"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const DecimalLimit = 16

var (
	dataset  = mustLoadSeed()
	first    = dataset.Units.Names()[rand.IntN(len(dataset.Units))]
	unit     = dataset.Units[first]
	randKeys []string
)

func mustLoadSeed() data.Dataset {
	dataset, err := data.LoadSeed()
	if err != nil {
		log.Fatalf("load application data: %v", err)
	}
	return dataset
}

func randomizeKeys() {
	randKeys = dataset.Measurements.Names()
	rand.Shuffle(len(randKeys), func(i, j int) {
		randKeys[i], randKeys[j] = randKeys[j], randKeys[i]
	})
	randKeys = randKeys[:len(randKeys)/2]
}

func main() {
	a := app.NewWithID("com.pmjlab.fingers")
	w := a.NewWindow("Fingers")
	w.Resize(fyne.NewSize(600, 400))

	valueColumn := container.NewVBox()
	nameColumn := container.NewVBox()
	clearResults := func() {
		valueColumn.RemoveAll()
		nameColumn.RemoveAll()
	}

	p := message.NewPrinter(language.English)
	var length float64
	refreshResults := func() {
		clearResults()
		for _, key := range randKeys {
			metric := length * unit / dataset.Measurements[key]
			formatted := p.Sprint(number.Decimal(metric))
			if len(formatted) > DecimalLimit {
				formatted = p.Sprint(number.Scientific(metric, number.Scale(2)))
			}

			value := widget.NewLabel(formatted)
			value.Alignment = fyne.TextAlignTrailing
			value.Truncation = fyne.TextTruncateEllipsis
			valueColumn.Add(value)

			description := widget.NewLabel(key)
			description.Truncation = fyne.TextTruncateEllipsis
			nameColumn.Add(description)
		}
	}

	options := &widget.Select{
		Options:  dataset.Units.Names(),
		Selected: first,
		OnChanged: func(key string) {
			unit = dataset.Units[key]
			refreshResults()
		},
	}

	input := &widget.Entry{
		PlaceHolder: "Input number",
		OnChanged: func(value string) {
			length, _ = strconv.ParseFloat(value, 64)
			randomizeKeys()
			refreshResults()
		},
	}

	results := container.NewHSplit(valueColumn, nameColumn)
	results.SetOffset(0.3)

	content := container.NewBorder(
		container.NewVBox(
			widget.NewForm(widget.NewFormItem("Convert:", input)),
			container.NewHBox(
				widget.NewForm(widget.NewFormItem("Select units:", options)),
				widget.NewButton("Reload", func() {
					randomizeKeys()
					refreshResults()
				}),
			),
		),
		nil,
		nil,
		nil,
		container.NewVScroll(results),
	)

	w.SetContent(content)
	w.ShowAndRun()
}
