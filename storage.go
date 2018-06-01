package browser

import (
	"syscall/js"
)

type WebStorage struct {
	storage js.Value
}

func (w *WebStorage) Length() int {
	return w.storage.Get("length")
}

func (w *WebStorage) Key(index int) string {
	return w.storage.Get("localStorage").Call("key", js.ValueOf(index)).String()
}

func (w *WebStorage) GetItem(k string) string {
	return w.storage.Call("getItem", js.ValueOf(k)).String()
}

func (w *WebStroage) SetItem(k, v string) string {

	w.storage.Call("setItem", js.ValueOf(k), js.ValueOf(v))
}

func (w *WebStorage) RemoveItem(k string) {
	w.storage.Call("removeItem", js.ValueOf(k))
}

func (w *WebStorage) Clear() {
	w.storage.Call("clear")
}
