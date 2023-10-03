package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Creo la App
	a := app.New()

	// Dimensiones de la ventana
	var ancho float32 = 510.0
	var alto float32 = 535.0

	// Cargo los assets
	iconApp, _ := fyne.LoadResourceFromPath("C:/Users/Dani/Desktop/go-app/icons8-comentarios-96.png")
	iconSend, _ := fyne.LoadResourceFromPath("C:/Users/Dani/Desktop/go-app/icons8-enviado-96.png")
	//iconThemeMoon, _ := fyne.LoadResourceFromPath("C:/Users/Dani/Desktop/go-app/icons8-luna-96.png")
	//iconThemeSun, _ := fyne.LoadResourceFromPath("C:/Users/Dani/Desktop/go-app/icons8-sol-96.png")

	a.Settings().SetTheme(theme.DarkTheme())
	a.SetIcon(iconApp)

	// Creo nueva ventana para la aplicacion
	w := a.NewWindow(" Chat LoRa ")
	w.Resize(fyne.NewSize(ancho, alto))

	// Menu Item
	menuItem1 := fyne.NewMenuItem("New", nil)
	menuItem2 := fyne.NewMenuItem("Edit", nil)
	menuItem3 := fyne.NewMenuItem("Save", nil)
	menuItem4 := fyne.NewMenuItem("Elegir Tema", nil)
	menuItem5 := fyne.NewMenuItem("Cambiar a modo Wifi", nil)
	menuItem6 := fyne.NewMenuItem("Guia de uso", nil)
	menuItem7 := fyne.NewMenuItem("About us", nil)

	menuOp1 := fyne.NewMenu("File", menuItem1, menuItem2, menuItem3)
	menuOp2 := fyne.NewMenu("Settings", menuItem4, menuItem5)
	menuOp3 := fyne.NewMenu("Help", menuItem6, menuItem7)

	menuItem4.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Oscuro", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		fyne.NewMenuItem("Claro", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)
	menu := fyne.NewMainMenu(menuOp1, menuOp2, menuOp3)

	w.SetMainMenu(menu)

	// Widget del input
	textBox := widget.NewEntry()
	textBox.SetPlaceHolder("Envia tu mensaje por LoRa...")
	textBox.Resize(fyne.NewSize(458, 40))
	textBox.Move(fyne.NewPos(1.5, 455))

	// Widget Button Sender
	sendButton := widget.NewButtonWithIcon("", iconSend, nil)
	sendButton.Resize(fyne.NewSize(40, 40))
	sendButton.Move(fyne.NewPos(460, 455))

	/* 	// Widget hyperlink
	   	url, _ := url.Parse("http://192.168.4.1")
	   	myLink := widget.NewHyperlink("Ir al modo Wifi", url)
	   	myLink.Resize(fyne.NewSize(130, 20))
	   	myLink.Move(fyne.NewPos(20, 300)) */

	// Container para Input + Sender Button
	//inputAndSend := container.NewHBox(textBox, sendButton)

	// Window Content
	appBox := container.NewWithoutLayout(
		textBox,
		sendButton,
	)

	w.SetContent(appBox)

	w.ShowAndRun()
}
