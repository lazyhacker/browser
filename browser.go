// +build js,wasm
package browser

import (
	"syscall/js"
)

var global = js.Global

type window struct {
	Console  debug
	Document htmldoc
}

// GetWindow returns the main browser window object.
func Window() window {
	return window{
		Console:  debug{console: global.Get("console")},
		Document: htmldoc{document: global.Get("document")},
	}
}

// Alert shows an alert box in the browser with an OK button.
func (window) Alert(m string) {
	global.Call("alert", m)
}
