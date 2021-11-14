//go run Notes.go
package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/internal/widget"
)

var count int = 1

func NotesApp(w fyne.Window){
	//a:= app.New()
	//w:=a.NewWindow("Notes")

	//w.Resize(fyne.NewSize(500 , 500))


	content:=container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Notes"),
			
		),
	)

	content.Add(widget.NewButton("Add New File" , func() {
		content.Add(widget.NewLabel("New File "+strconv.Itoa(count-1)))
		count++
	}))

	input:=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text ...")
	input.Resize(fyne.NewSize(400 , 400))
	//fubctionality of save button
	saveBtn:=widget.NewButton("Save text file" , func() {
		//func NewFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window) *FileDialog
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData:= []byte(input.Text)

				uc.Write(textData)
			},w)
			

		saveFileDialog.SetFileName("New File "+strconv.Itoa(count))
		saveFileDialog.Show()
	})
//opening a file
openBtn:=widget.NewButton("Open text file" , func() {
	//func NewFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window) *FileDialog
	openFileDialog := dialog.NewFileOpen(
		func(ur fyne.URIReadCloser, _ error) {
			readData ,_:= ioutil.ReadAll(ur)

			output:=fyne.NewStaticResource("New File" , readData)
			
			viewText:=widget.NewMultiLineEntry()

			viewText.SetText(string(output.StaticContent))

			w:=fyne.CurrentApp().NewWindow(
				string(output.StaticName))

			w.SetContent(container.NewScroll(viewText))
			w.Resize(fyne.NewSize(400 , 400))
			w.Show()
			
		},w)

		openFileDialog.SetFilter(
			storage.NewExtensionFileFilter([] string{".txt"}))
		openFileDialog.Show()

	})

		
	notesContainer :=
		container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
		)
w.SetContent(container.NewBorder(panelContent,nil,nil,nil,notesContainer),)
	w.Show()
}