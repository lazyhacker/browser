// package browser provides Go bindings to browser APIs.  This package is
// intended to be used with WebAssembly to abstract Go developers from having
// to deal with writing javascript inside of Go.
// +build js,wasm
package browser

import (
	"syscall/js"
)

type window struct {
	window   js.Value
	Console  debug
	Document htmldoc
}

// GetWindow returns the main browser window object.
func Window() window {
	return window{
		window:   js.Global,
		Console:  debug{console: js.Global.Get("console")},
		Document: htmldoc{document: js.Global.Get("document")},
	}
}

// Alert shows an alert box in the browser with an OK button.
func (w *window) Alert(m string) {
	w.window.Call("alert", m)
}

// Invoke executes the Callback that is passed in.
func Invoke(cb js.Callback) {
	js.ValueOf(cb).Invoke()
}

// NewWindow creates a new browser window.  This is analogous to Javascript's
// window.open().
func NewWindow(url string) window {
	w := js.Global.Call("open", url)

	return window{
		window:   w,
		Console:  debug{console: w.Get("console")},
		Document: htmldoc{document: w.Get("document")},
	}
}

// Blur removes the focus from the window.
func (w *window) Blur() {
	w.window.Call("blur")
}

// Focus sets the focus to the window.
func (w *window) Focus() {
	w.window.Call("focus")
}

// Confirm shows an confirm dialog window with an OK and Cancel button.
func (w *window) Confirm() bool {
	return w.window.Call("confirm").Bool()
}

// Print outputs the entire window to the browser's print dialog.
func (w *window) Print() {
	w.window.Call("print")
}
