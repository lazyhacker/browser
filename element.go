package browser

import (
	"syscall/js"
)

// ElementTagName is the tag name of the element.
type ElementTagName string

// Element is a HTML element object.
type Element struct {
	el js.Value
}

// Context returns the element's contex value.
func (e *Element) Context(ctx string) js.Value {
	return e.el.Call("getContext", ctx)
}

// AddEventListener adds a callback function to an element's browser event.
func (e Element) AddEventListener(evt string, cb js.Callback) {
	e.el.Call("addEventListener", js.ValueOf(evt), js.ValueOf(cb))
}

func (Element) AccessKey()         {}
func (Element) AppendChild()       {}
func (Element) Attributes()        {}
func (Element) Blur()              {}
func (Element) ChildElementCount() {}
func (Element) ChildNodes()        {}
func (Element) Children()          {}
func (Element) ClassList()         {}
func (Element) ClassName()         {}
func (Element) Click()             {}
func (e *Element) ClientHeight() int {

	return e.el.Get("clientHeight").Int()
}
func (Element) ClientLeft() {}
func (Element) ClientTop()  {}
func (e *Element) ClientWidth() int {

	return e.el.Get("clientWidth").Int()

}
func (Element) CloneNode()               {}
func (Element) CompareDocumentPosition() {}
func (Element) Contains()                {}
func (Element) ContentEditable()         {}
func (Element) Dir()                     {}
func (Element) FirstChild()              {}
func (Element) FirstElementChild()       {}
func (Element) Focus()                   {}
func (Element) GetAttribute()            {}
func (Element) GetAttributeNode()        {}
func (Element) GetElementsByClassName()  {}
func (Element) GetElementsByTagName()    {}
func (Element) HasAttribute()            {}
func (Element) HasAttributes()           {}
func (Element) HasChildNodes()           {}
func (Element) Id()                      {}
func (Element) InnerHTML()               {}
func (Element) InnerText()               {}
func (Element) InsertAdjacentElement()   {}
func (Element) InsertAdjacentHTML()      {}
func (Element) InsertAdjacentText()      {}
func (Element) InsertBefore()            {}
func (Element) IsContentEditable()       {}
func (Element) IsDefaultNamespace()      {}
func (Element) IsEqualNode()             {}
func (Element) IsSameNode()              {}
func (Element) IsSupported()             {}
func (Element) Lang()                    {}
func (Element) LastChild()               {}
func (Element) LastElementChild()        {}
func (Element) NamespaceURI()            {}
func (Element) NextSibling()             {}
func (Element) NextElementSibling()      {}
func (Element) NodeName()                {}
func (Element) NodeType()                {}
func (Element) NodeValue()               {}
func (Element) Normalize()               {}
func (Element) OffsetHeight()            {}
func (Element) OffsetWidth()             {}
func (Element) OffsetLeft()              {}
func (Element) OffsetParent()            {}
func (Element) OffsetTop()               {}
func (Element) OwnerDocument()           {}
func (Element) ParentNode()              {}
func (Element) ParentElement()           {}
func (Element) PreviousSibling()         {}
func (Element) PreviousElementSibling()  {}
func (Element) QuerySelector()           {}
func (Element) QuerySelectorAll()        {}
func (Element) RemoveAttribute()         {}
func (Element) RemoveAttributeNode()     {}
func (Element) RemoveChild()             {}
func (Element) RemoveEventListener()     {}
func (Element) ReplaceChild()            {}
func (Element) ScrollHeight()            {}
func (Element) ScrollIntoView()          {}
func (Element) ScrollLeft()              {}
func (Element) ScrollTop()               {}
func (Element) ScrollWidth()             {}
func (e Element) SetAttribute(a string, k interface{}) {

	e.el.Call("setAttribute", js.ValueOf(a), js.ValueOf(k))

}
func (e Element) SetProperty(a string, k interface{}) {

	e.el.Set(a, js.ValueOf(k))

}

func (Element) SetAttributeNode() {}
func (Element) Style()            {}
func (Element) TabIndex()         {}
func (Element) TagName()          {}
func (Element) TextContent()      {}
func (Element) Title()            {}
func (Element) ToString()         {}
