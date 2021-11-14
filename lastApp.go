package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/widget"

	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	//"github.com/faiface/beep/speaker"
)

//var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool= false

func musicplayer(w fyne.Window){
	os.Setenv("FYNE_THEME" ,"dark")

	go func (msg string)  {
		fmt.Println(msg)
		if streamer == nil{

		}else{
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
		
	}("")
	time.Sleep(time.Second)

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400 ,400))

	img :=canvas.NewImageFromFile("D:\\VS_Code_Go\\musiclogo.jpg")

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func(){
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)

		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause{
				pause=true
				speaker.Lock()
			}else if pause{
				pause= false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
		
	)
	label1 :=widget.NewLabel("Music Player")
	label1.Alignment = fyne.TextAlignCenter

	label :=widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter

	browse :=widget.NewButtonWithIcon("List" , theme.ListIcon(), func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error){
			streamer, format, _= mp3.Decode(uc)
			label1.Text =uc.URI().Name()
			label1.Refresh()
		
		},w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	c := container.NewVBox(layout.NewSpacer(),label1,layout.NewSpacer(),browse,layout.NewSpacer(),label,layout.NewSpacer(),toolbar)
	musicContainer :=container.New(layout.NewMaxLayout(), img ,
	container.NewBorder(panelContent,nil,nil,nil, c))
	
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,musicContainer))
	w.Show()




}