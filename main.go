package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	
	myApp := app.New()

	myWindow := myApp.NewWindow("app test")
	myWindow.Resize(fyne.NewSize(400, 350))
	
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("file",
			fyne.NewMenuItem("New", nil),
			fyne.NewMenuItem("Close", func() {
				myWindow.Close()
			}),
		),
	)

	myWindow.SetMainMenu(mainMenu)

	label1 := widget.NewLabel("Choose")

	label2 := widget.NewLabel("Window")

	button1 := widget.NewButton("Show Dialog", func() {
		dialog.ShowConfirm("Dialog", "Confirm Dialog", func(r bool) {
			if r {
				label1.SetText("Yes Clicked!")
			} else {
				label1.SetText("No Clicked!")
			}
		},myWindow)
	})

	button2 := widget.NewButton("NewWindow", func()	{NewWindow(myApp)})

	tabs := container.NewAppTabs(
		container.NewTabItem("Dialog", container.NewVBox(label1, button1),),
		container.NewTabItem("Window", container.NewVBox(label2,button2),),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
}

func NewWindow(myApp fyne.App) {

	myWindow2 := myApp.NewWindow("New Window")
	myWindow2.Resize(fyne.NewSize(400,350))

	button3 := widget.NewButton("close this", func() {
		myWindow2.Close()
	})

	myWindow2.SetContent(
		container.NewVBox(
			button3,
		),
	)

	myWindow2.Show()
}

