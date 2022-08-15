package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/cdipaolo/sentiment"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chatbot")

	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	e := widget.NewEntry()

	f := widget.NewForm(
		widget.NewFormItem("Statement", e),
	)

	e.OnChanged = func(text string) {
		f.Refresh()
	}
	e.Validator = func(text string) error {
		analysis := model.SentimentAnalysis(e.Text, sentiment.English)
		if analysis.Score != 1 {
			f.Disable()
		} else {
			f.Enable()
		}
		return nil
	}

	f.OnSubmit = func() {
		fmt.Println("Submitted")
	}

	w.SetContent(f)

	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
