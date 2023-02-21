package main

import (
	"runtime"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	
	a := app.New()

	w := a.NewWindow("app test")
	w.Resize(fyne.NewSize(400, 350))
	
	mm := fyne.NewMainMenu(
		fyne.NewMenu("file",
			fyne.NewMenuItem("New", nil),
			fyne.NewMenuItem("Close", func() {
				w.Close()
			}),
		),
	)

	w.SetMainMenu(mm)

	l1 := widget.NewLabel("link to dialog")

	b1 := widget.NewButton("Show Dialog", func() {
		dialog.ShowConfirm("Dialog", "Confirm Dialog", func(r bool) {
			if r {
				l1.SetText("Yes Clicked!")
			} else {
				l1.SetText("No Clicked!")
			}}, w)
		})

	b2 := widget.NewButton("New Window", func() { intentWindow(a) }) 

	e := widget.NewEntry()

	w.SetContent(
		container.NewVBox(
			b1, b2, e, l1,
		),
	)

	w.ShowAndRun()
}

func intentWindow(a fyne.App) {

	w2 := a.NewWindow("New Window")
	w2.Resize(fyne.NewSize(400,350))

	l2 := widget.NewLabel("This System is")

	b3 := widget.NewButton("Detect System", func() {
		l2.SetText(runtime.GOOS)
	})

	b4 := widget.NewButton("close this", func() {
		w2.Close()
	})

	w2.SetContent(
		container.NewVBox(
			b3, b4, l2,
		),
	)

	w2.Show()
}

	
