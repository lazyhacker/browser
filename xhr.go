package browser

import (
	"syscall/js"
)

type XHR struct {
	xhr js.Value
}

func NewXHR() XHR {
	return XHR{xhr: js.Global().Get("XMLHttpRequest").New()}
}

func (x *XHR) Open(url string, async bool) {

	x.xhr.Call("open", "GET", js.ValueOf(url), js.ValueOf(async))

}

func (x *XHR) Send() error {

	println("actual send")
	x.xhr.Call("send", js.Null)

	//println("check status")
	//if x.xhr.Get("status") != js.ValueOf(200) {
	//		return fmt.Errorf("Status code received: %d", x.xhr.Get("status").Int())
	//}

	return nil
}

func (x *XHR) ResponseText() string {
	return x.xhr.Get("responseText").String()
}

func (x *XHR) Status() int {
	return x.xhr.Get("status").Int()
}

func (x *XHR) ResponseType() string {
	return x.xhr.Get("responseType").String()
}

func (x *XHR) ResponseXML() string {
	return x.xhr.Get("responseXML").String()
}

func (x *XHR) StatusString() string {
	return x.xhr.Get("statusText").String()
}
