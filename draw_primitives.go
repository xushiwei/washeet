package washeet

import (
	"math"
	"syscall/js"
)

var path2dCtor = js.Global().Get("Path2D")

func drawVertLines(canvasContext *js.Value, xcoords []float64, ylow, yhigh float64, color string) {
	path := path2dCtor.New()
	canvasContext.Set("strokeStyle", color)
	for _, xcoord := range xcoords {
		addLine2Path(&path, xcoord, ylow, xcoord, yhigh)
	}
	canvasContext.Call("stroke", path)
}

func drawHorizLines(canvasContext *js.Value, ycoords []float64, xlow, xhigh float64, color string) {
	path := path2dCtor.New()
	canvasContext.Set("strokeStyle", color)
	for _, ycoord := range ycoords {
		addLine2Path(&path, xlow, ycoord, xhigh, ycoord)
	}
	canvasContext.Call("stroke", path)
}

func drawLine(canvasContext *js.Value, xlow, ylow, xhigh, yhigh float64, color string) {
	path := path2dCtor.New()
	canvasContext.Set("strokeStyle", color)
	addLine2Path(&path, xlow, ylow, xhigh, yhigh)
	canvasContext.Call("stroke", path)
}

// Adds a line to the path, no actual drawing takes place
func addLine2Path(path *js.Value, xlow, ylow, xhigh, yhigh float64) {
	path.Call("moveTo", xlow, ylow)
	path.Call("lineTo", xhigh, yhigh)
}

func noStrokeFillRect(canvasContext *js.Value, xlow, ylow, xhigh, yhigh float64, fillColor string) {
	path := path2dCtor.New()
	canvasContext.Set("fillStyle", fillColor)
	addRect2Path(&path, xlow, ylow, xhigh, yhigh)
	canvasContext.Call("fill", path)
}

func strokeFillRect(canvasContext *js.Value, xlow, ylow, xhigh, yhigh float64, strokeColor, fillColor string) {
	path := path2dCtor.New()
	canvasContext.Set("strokeStyle", strokeColor)
	canvasContext.Set("fillStyle", fillColor)
	addRect2Path(&path, xlow, ylow, xhigh, yhigh)
	canvasContext.Call("stroke", path)
	canvasContext.Call("fill", path)
}

func strokeNoFillRect(canvasContext *js.Value, xlow, ylow, xhigh, yhigh float64, strokeColor string) {
	path := path2dCtor.New()
	canvasContext.Set("strokeStyle", strokeColor)
	addRect2Path(&path, xlow, ylow, xhigh, yhigh)
	canvasContext.Call("stroke", path)
}

// Adds a rect to the path, no actual drawing takes place
func addRect2Path(path *js.Value, xlow, ylow, xhigh, yhigh float64) {
	path.Call("rect", xlow, ylow, xhigh-xlow, yhigh-ylow)
}

func drawText(canvasContext *js.Value, xlow, ylow, xhigh, yhigh, xmax, ymax float64, text string, align TextAlignType) {
	xend, yend := math.Min(xhigh, xmax), math.Min(yhigh, ymax)
	canvasContext.Call("save")
	path := path2dCtor.New()
	addRect2Path(&path, xlow, ylow, xend-1.0, yend-1.0)
	canvasContext.Call("clip", path)

	canvasContext.Set("textBaseline", "bottom")
	startx, starty := xlow, yhigh // yhigh assuming English like language.
	if align == AlignLeft {
		canvasContext.Set("textAlign", "left")
	}
	if align == AlignCenter {
		startx = 0.5 * (xlow + xhigh)
		canvasContext.Set("textAlign", "center")
	} else if align == AlignRight {
		startx = xhigh
		canvasContext.Set("textAlign", "right")
	}

	canvasContext.Call("fillText", text, startx, starty)
	// Kill the clip path
	canvasContext.Call("restore")

	// DEBUG code
	//strokeNoFillRect(canvasContext, xlow, ylow, xhigh, yhigh, "#0000ff")
	//strokeNoFillRect(canvasContext, xlow, ylow, xmax, ymax, "#ff0000")
}

func setFont(canvasContext *js.Value, fontCSS string) {
	canvasContext.Set("font", fontCSS)
}

func setFillColor(canvasContext *js.Value, fillColor string) {
	canvasContext.Set("fillStyle", fillColor)
}

func setStrokeColor(canvasContext *js.Value, strokeColor string) {
	canvasContext.Set("strokeStyle", strokeColor)
}
