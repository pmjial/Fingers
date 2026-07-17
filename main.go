package main

import (
	"Fingers/db"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fingers")

	w.SetContent(widget.NewLabel("Hello fyne"))
	w.ShowAndRun()

	//fmt.Println(ConvertN(0.9, 3))
}

func ConvertN(submit db.Length, n int) []string {
	comps := make([]string, n)
	for i := 0; i < cap(comps); i++ {
		key := db.RandKey()
		comp := db.Get(key)
		comps[i] = fmt.Sprintf("%.2fm is ~%.2f %vs\n", submit, submit/comp, key)
	}
	return comps
}
