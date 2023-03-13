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
		fyne.NewMenu("edit",
			fyne.NewMenuItem("copy", nil),
			fyne.NewMenuItem("cut", nil),
			fyne.NewMenuItem("paste", nil),
		),
	)

	myWindow.SetMainMenu(mainMenu)

	button1 := widget.NewButton("Confirm Dialog", func() {
		dialog.ShowConfirm("Dialog", "Confirm Dialog", nil, myWindow)
	})

	button2 := widget.NewButton("Entry Dialog", func() {
		dialog.ShowEntryDialog("Dialog", "Entry Dialog", nil, myWindow)
	})

	button3 := widget.NewButton("Information", func() {
		dialog.ShowInformation("Dialog", "Infoermation", myWindow)
	})

	button5 := widget.NewButton("NewWindow", func()	{NewWindow(myApp)})

	tabs := container.NewAppTabs(
		container.NewTabItem("Dialog", container.NewVBox(button1, button2, button3),),
		container.NewTabItem("Window", container.NewVBox(button5),),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
}

func NewWindow(myApp fyne.App) {

	myWindow2 := myApp.NewWindow("New Window")
	myWindow2.Resize(fyne.NewSize(400,350))

	button6 := widget.NewButton("close this", func() {
		myWindow2.Close()
	})

	myWindow2.SetContent(
		container.NewVBox(
			button6,
		),
	)

	myWindow2.Show()
}

