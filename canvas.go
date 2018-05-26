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

func (e Element) ToCanvas() (Canvas, error) {

	if strings.ToUpper(e.el.Get("tagName").String()) != ElementCanvas {
		return Canvas{}, fmt.Errorf("mismatched element type. %s != %s", e.el.Get("tagName").String(), ElementCanvas)
	}

	return Canvas{
		canvas: e,
	}, nil
}

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

func (c *Canvas) BeginPath() {
	ctx := c.context()
	ctx.Call("beginPath")
}

func (c *Canvas) Stroke() {

	ctx := c.context()
	ctx.Call("stroke")

}

func (c *Canvas) context() js.Value {
	if c.ctx == nil {
		a := c.canvas.Context("2d")
		c.ctx = &a
	}
	return *c.ctx
}

func (c *Canvas) BezierCurveTo(x1, y1, x2, y2, x3, y3 float64) {}
func (c *Canvas) ClearLinearGradient()                         {}
func (c *Canvas) ClearRect(x, y, w, h float64)                 {}
func (c *Canvas) Clip()                                        {}
func (c *Canvas) ClosePath()                                   {}

//func (c *Canvas) Context()                                       {}
func (c *Canvas) CreateEvent()                                   {}
func (c *Canvas) CreatePattern()                                 {}
func (c *Canvas) CreateRadialGradient()                          {}
func (c *Canvas) DrawImage(image interface{}, coords ...float64) {}
func (c *Canvas) Fill()                                          {}
func (c *Canvas) FillRect(x, y, w, h float64)                    {}
func (c *Canvas) FillStyle(value ...interface{})                 {}
func (c *Canvas) FillText(str string, x, y float64)              {}
func (c *Canvas) Font(font interface{}, size float64)            {}
func (c *Canvas) GlobalAlpha(alpha float64)                      {}
func (c *Canvas) GlobalCompositeOperation()                      {}
func (c *Canvas) LineCap()                                       {}
func (c *Canvas) LineTo(x, y float64)                            {}
func (c *Canvas) LineWidth(width float64)                        {}
func (c *Canvas) MoveTo(x, y float64)                            {}
func (c *Canvas) PutImageData(img *image.RGBA, x, y int)         {}
func (c *Canvas) QuadraticCurveTo(x1, y1, x2, y2 float64)        {}
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

//func (c *Canvas) Width() int                         {}
func (c *Canvas) AddColorStop()  {}
func (c *Canvas) IsPointInPath() {}
func (c *Canvas) MiterLimit()    {}
func (c *Canvas) TextBaseline()  {}
