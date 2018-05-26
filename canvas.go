package browser

import (
	"fmt"
	"image"
	"strings"
	"syscall/js"
)

const ElementCanvas = "CANVAS"

type canvas struct {
	canvas element
	ctx    *js.Value
}

func (e element) ToCanvas() (canvas, error) {

	if strings.ToUpper(e.el.Get("tagName").String()) != ElementCanvas {
		return canvas{}, fmt.Errorf("mismatched element type. %s != %s", e.el.Get("tagName").String(), ElementCanvas)
	}

	return canvas{
		canvas: e,
	}, nil
}

func (c *canvas) Arc(x, y, radius, startAngle, endAngle float64, counterclockwise bool) {
	ctx := c.context()
	ctx.Call("arc",
		js.ValueOf(x),
		js.ValueOf(y),
		js.ValueOf(radius),
		js.ValueOf(startAngle),
		js.ValueOf(endAngle),
		js.ValueOf(counterclockwise))

}

func (c *canvas) BeginPath() {
	ctx := c.context()
	ctx.Call("beginPath")
}

func (c *canvas) Stroke() {

	ctx := c.context()
	ctx.Call("stroke")

}

func (c *canvas) context() js.Value {
	if c.ctx == nil {
		a := c.canvas.Context("2d")
		c.ctx = &a
	}
	return *c.ctx
}

func (c *canvas) BezierCurveTo(x1, y1, x2, y2, x3, y3 float64) {}
func (c *canvas) ClearLinearGradient()                         {}
func (c *canvas) ClearRect(x, y, w, h float64)                 {}
func (c *canvas) Clip()                                        {}
func (c *canvas) ClosePath()                                   {}

//func (c *canvas) Context()                                       {}
func (c *canvas) CreateEvent()                                   {}
func (c *canvas) CreatePattern()                                 {}
func (c *canvas) CreateRadialGradient()                          {}
func (c *canvas) DrawImage(image interface{}, coords ...float64) {}
func (c *canvas) Fill()                                          {}
func (c *canvas) FillRect(x, y, w, h float64)                    {}
func (c *canvas) FillStyle(value ...interface{})                 {}
func (c *canvas) FillText(str string, x, y float64)              {}
func (c *canvas) Font(font interface{}, size float64)            {}
func (c *canvas) GlobalAlpha(alpha float64)                      {}
func (c *canvas) GlobalCompositeOperation()                      {}
func (c *canvas) LineCap()                                       {}
func (c *canvas) LineTo(x, y float64)                            {}
func (c *canvas) LineWidth(width float64)                        {}
func (c *canvas) MoveTo(x, y float64)                            {}
func (c *canvas) PutImageData(img *image.RGBA, x, y int)         {}
func (c *canvas) QuadraticCurveTo(x1, y1, x2, y2 float64)        {}
func (c *canvas) Rect(x, y, w, h float64)                        {}
func (c *canvas) Restore()                                       {}
func (c *canvas) Rotate(angle float64)                           {}
func (c *canvas) Save()                                          {}
func (c *canvas) Scale(x, y float64)                             {}
func (c *canvas) ShadowBlur()                                    {}
func (c *canvas) ShadowColor()                                   {}
func (c *canvas) ShadowOffsetX()                                 {}
func (c *canvas) ShadowOffsetY()                                 {}
func (c *canvas) StrokeRect(x, y, w, h float64)                  {}
func (c *canvas) StrokeStyle(value ...interface{})               {}
func (c *canvas) StrokeText()                                    {}
func (c *canvas) Translate(x, y float64)                         {}

//func (c *canvas) Width() int                         {}
func (c *canvas) AddColorStop()  {}
func (c *canvas) IsPointInPath() {}
func (c *canvas) MiterLimit()    {}
func (c *canvas) TextBaseline()  {}
