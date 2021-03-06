package main

import (
	"fmt"
	"github.com/dennisfrancis/washeet"
	"math"
	"syscall/js"
)

type SheetModel struct {
}

// Satisfy SheetDataProvider interface.

func (self *SheetModel) GetDisplayString(column int64, row int64) string {
	if column == 5 && row == 5 {
		return "Hello washeet !"
	}

	return fmt.Sprintf("Cell(%d, %d)", column, row)
}

func (self *SheetModel) GetColumnWidth(column int64) float64 {
	if column == 5 {
		return 230.0
	}
	return 0.0
}

func (self *SheetModel) GetRowHeight(row int64) float64 {
	if row == 5 {
		return 70.0
	}
	return 0.0
}

func (self *SheetModel) TrimToNonEmptyRange(c1, r1, c2, r2 *int64) bool {
	return true
}

// Satisfy SheetModelUpdater interface.

func (self *SheetModel) SetColumnWidth(column int64, width float64) {
}

func (self *SheetModel) SetRowHeight(row int64, height float64) {
}

func (self *SheetModel) SetCellContent(row, column int64, content string) {
}

func main() {

	fmt.Println("Hello washeet !")

	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "washeet")
	container := doc.Call("getElementById", "container")
	width, height := doc.Get("body").Get("clientWidth").Float(), doc.Get("body").Get("clientHeight").Float()
	width, height = math.Floor(width*0.85), math.Floor(height*0.85)
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)
	closeButton := doc.Call("getElementById", "close-button")
	quit := make(chan bool)

	closeHandler := js.NewCallback(func(args []js.Value) {
		quit <- true
	})
	closeButton.Call("addEventListener", "click", closeHandler)

	model := &SheetModel{}
	sheet := washeet.NewSheet(&canvasEl, &container, 0.0, 0.0, width-1.0, height-1.0, model, model)
	sheet.Start()

	<-quit
	closeButton.Call("removeEventListener", "click", closeHandler)
	closeHandler.Release()
	sheet.Stop()
	fmt.Println("sheet.Stop() successfull !")
}
