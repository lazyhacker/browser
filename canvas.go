package browser

import (
	"fmt"
	"image"
	"strings"
	"syscall/js"
)

const ElementCanvas = "CANVAS"

// HTML canvas
type Canvas struct {
	canvas Element
	ctx    *js.Value
}

// Sets the context object if it hasn't been set.
func (c *Canvas) context() js.Value {
	if c.ctx == nil {
		a := c.canvas.Context("2d")
		c.ctx = &a
	}
	return *c.ctx
}

// ToCanvas casts an element to a Canvas object.
func (e Element) ToCanvas() (Canvas, error) {

	if strings.ToUpper(e.el.Get("tagName").String()) != ElementCanvas {
		return Canvas{}, fmt.Errorf("mismatched element type. %s != %s", e.el.Get("tagName").String(), ElementCanvas)
	}

	return Canvas{
		canvas: e,
	}, nil
}

// Arc creates a arc/curve (used to create circles, or parts of circles).
func (c *Canvas) Arc(x, y, radius, startAngle, endAngle float64, counterclockwise bool) {
	ctx := c.context()
	ctx.Call("arc",
		js.ValueOf(x),
		js.ValueOf(y),
		js.ValueOf(radius),
		js.ValueOf(startAngle),
		js.ValueOf(endAngle),
		js.ValueOf(counterclockwise))

}

// ArcTo creates an arc/cuve between two tangents.
func (c *Canvas) ArcTo(x1, y1, x2, y2, r float64) {
	c.context().Call("arcTo", js.ValueOf(x1), js.ValueOf(y1), js.ValueOf(x2), js.ValueOf(y2), js.ValueOf(r))
}

// BeginPath begins a path or resets the current path.
func (c *Canvas) BeginPath() {
	ctx := c.context()
	ctx.Call("beginPath")
}

// Stroke Draws the path defined.
func (c *Canvas) Stroke() {

	ctx := c.context()
	ctx.Call("stroke")

}

// BezierCurveTo creates a cubic Bezier curve.
func (c *Canvas) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {

	ctx := c.context()
	ctx.Call("bezierCurveTo",
		js.ValueOf(cp1x),
		js.ValueOf(cp1y),
		js.ValueOf(cp2x),
		js.ValueOf(cp2y),
		js.ValueOf(x),
		js.ValueOf(y))

}

// CreateLinearGradient creates a linear gradient to use on canvans content.
func (c *Canvas) CreateLinearGradient(x0, y0, x1, y1 float64) {
	c.context().Call("createLinearGradient", js.ValueOf(x0), js.ValueOf(y0), js.ValueOf(x1), js.ValueOf(y1))
}

// ClearRect clears the specified pixels within a given rectangle.
func (c *Canvas) ClearRect(x, y, w, h float64) {
	c.context().Call("clearRect", js.ValueOf(x), js.ValueOf(y), js.ValueOf(w), js.ValueOf(h))
}

// Clip clips a region of any shape and size from the original canvas.
func (c *Canvas) Clip() {
	c.context().Call("clip")
}

// ClosePath creates a path from the current point back to the starting point.
func (c *Canvas) ClosePath() {
	c.context().Call("closePath")
}

func (c *Canvas) CreateEvent() {
	c.context().Call("createEvent")
}

// Fills the current drawing(path).
func (c *Canvas) Fill() {
	c.context().Call("fill")
}

// Font sets the canvas font (e.g. "30px Arial").
func (c *Canvas) Font(font string) {
	c.context().Set("font", js.ValueOf(font))
}

// TextAlign sets the text alignment.
// start (default) The text starts at the specified position.
// end The text ends at the specified position.
// center The center of the text is placed at the specified position.
// left The text starts at the specified position.
// right The text ends at the specified position.
func (c *Canvas) TextAlign(v string) {

	c.ctx.Set("fillAlign", js.ValueOf(v))
}

// LineTo Adds a new point and create a line from that point to the last specified point in the canvas.
func (c *Canvas) LineTo(x, y float64) {
	c.context().Call("lineTo", js.ValueOf(x), js.ValueOf(y))
}

// MoveTo moves the path to a specific point in the canvas without creating the line.
func (c *Canvas) MoveTo(x, y float64) {
	c.context().Call("moveTo", js.ValueOf(x), js.ValueOf(y))
}

//QuadraticCurveTo creates a quadratic Bezier curve.
func (c *Canvas) QuadraticCurveTo(x1, y1, x2, y2 float64) {

	c.context().Call("quadraticCurveTo", js.ValueOf(x1), js.ValueOf(y1), js.ValueOf(x2), js.ValueOf(y2))

}

// Returns true if the specified poitn is in the current path, otherwise false.
func (c *Canvas) IsPointInPath(x, y float64) bool {
	return c.context().Call("isPointInPath", js.ValueOf(x), js.ValueOf(y)).Bool()
}

func (c *Canvas) Rect(x, y, w, h float64)                        {}
func (c *Canvas) Restore()                                       {}
func (c *Canvas) Rotate(angle float64)                           {}
func (c *Canvas) Save()                                          {}
func (c *Canvas) Scale(x, y float64)                             {}
func (c *Canvas) ShadowBlur()                                    {}
func (c *Canvas) ShadowColor()                                   {}
func (c *Canvas) ShadowOffsetX()                                 {}
func (c *Canvas) ShadowOffsetY()                                 {}
func (c *Canvas) StrokeRect(x, y, w, h float64)                  {}
func (c *Canvas) StrokeStyle(value ...interface{})               {}
func (c *Canvas) StrokeText()                                    {}
func (c *Canvas) Translate(x, y float64)                         {}
func (c *Canvas) GlobalAlpha(alpha float64)                      {}
func (c *Canvas) GlobalCompositeOperation()                      {}
func (c *Canvas) LineCap()                                       {}
func (c *Canvas) MiterLimit()                                    {}
func (c *Canvas) TextBaseline()                                  {}
func (c *Canvas) CreatePattern()                                 {}
func (c *Canvas) CreateRadialGradient()                          {}
func (c *Canvas) DrawImage(image interface{}, coords ...float64) {}
func (c *Canvas) FillRect(x, y, w, h float64)                    {}
func (c *Canvas) FillStyle(value ...interface{})                 {}
func (c *Canvas) PutImageData(img *image.RGBA, x, y int)         {}
func (c *Canvas) LineWidth(width float64)                        {}
func (c *Canvas) AddColorStop()                                  {}
func (c *Canvas) FillText(str string, x, y float64)              {}

//func (c *Canvas) Context()                                       {}
// Width returns the width of an imageData object.
//func (c *Canvas) Width() int     {}
