package browser

import (
	"syscall/js"
)

type WebStorage struct {
	storage js.Value
}

func (w *WebStorage) Length() int {
	return w.storage.Get("length").Int()
}

// The objects probably shouldn't be stored/retrieved as strings... okay as a hack for demo purposes.
func (w *WebStorage) Key(index int) string {
	return w.storage.Get("localStorage").Call("key", js.ValueOf(index)).String()
}

func (w *WebStorage) GetItem(k string) string {
	v := w.storage.Call("getItem", js.ValueOf(k))
	if v == js.Null() {
		return ""
	}
	return v.String()
}

func (w *WebStorage) SetItem(k, v string) string {

	return w.storage.Call("setItem", js.ValueOf(k), js.ValueOf(v)).String()
}

func (w *WebStorage) RemoveItem(k string) {
	w.storage.Call("removeItem", js.ValueOf(k))
}

func (w *WebStorage) Clear() {
	w.storage.Call("clear")
}
