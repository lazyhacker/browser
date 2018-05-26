package browser

import "syscall/js"

// Debug provides access to the browsers Debugging console.
type Debug struct {
	console js.Value
}

// Clear clears the console.
func (d *Debug) Clear() {
	d.console.Call("clear")
}

// Count  logs the number of times this count has been called.
func (d *Debug) Count() int {
	return d.console.Call("count").Int()
}

// Error writes an error message to the console.
func (d *Debug) Error(m string) {
	d.console.Call("error", m)
}

// Info writes an information message to the console.
func (d *Debug) Info(m string) {
	d.console.Call("info", m)

}

// Warn writes an warning message to the console.
func (d *Debug) Warn(m string) {
	d.console.Call("warn", m)
}

// Trace writes the stack trace to the console.
func (d *Debug) Trace() {
	d.console.Call("trace")
}
