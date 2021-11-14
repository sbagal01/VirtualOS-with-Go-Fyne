//go run .\DesktopOs.go .\lastApp.go .\main.go .\Notes.go .\Weather.go

package main

import (
	//"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)
var Myapp fyne.App = app.New()

var myWindow fyne.Window  = Myapp.NewWindow("DesktopOs")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var img fyne.CanvasObject;
var Dskt fyne.Widget

var panelContent *fyne.Container


func main(){
	Myapp.Settings().SetTheme(theme.LightTheme())
	img=canvas.NewImageFromFile("D:\\VS_Code_Go\\DesktopImg.jpg")

	btn1=widget.NewButtonWithIcon("WeatherApp" , theme.InfoIcon() , func() {
		WeatherApp(myWindow)

	})
	btn2=widget.NewButtonWithIcon("Calculator" , theme.ContentAddIcon() , func() {
		ShowCalci()

	})
	btn3=widget.NewButtonWithIcon("Editor" , theme.HomeIcon() , func() {
		NotesApp(myWindow)

	})
	btn4=widget.NewButtonWithIcon("Music" , theme.HomeIcon() , func() {
		musicplayer(myWindow)

	})
	Dskt=widget.NewButtonWithIcon("Desktop" , theme.HomeIcon() , func() {
		myWindow.SetContent(container.NewBorder(panelContent , nil, nil , nil , img))

	})
	panelContent = container.NewVBox(container.NewGridWithColumns(5,Dskt , btn1, btn2, btn3, btn4))


	myWindow.Resize(fyne.NewSize(1280 , 720))
	myWindow.CenterOnScreen()


	myWindow.SetContent(container.NewBorder(panelContent , nil , nil , nil , img),)
	myWindow.ShowAndRun()

}