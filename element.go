package browser

import (
	"syscall/js"
)

type ElementTagName string

type element struct {
	el js.Value
}

func (e *element) Context(ctx string) js.Value {

	return e.el.Call("getContext", ctx)
}

func (element) AccessKey() {}
func (e element) AddEventListener(evt string, cb js.Callback) {
	e.el.Call("addEventListener", js.ValueOf(evt), js.ValueOf(cb))
}
func (element) AppendChild()             {}
func (element) Attributes()              {}
func (element) Blur()                    {}
func (element) ChildElementCount()       {}
func (element) ChildNodes()              {}
func (element) Children()                {}
func (element) ClassList()               {}
func (element) ClassName()               {}
func (element) Click()                   {}
func (element) ClientHeight()            {}
func (element) ClientLeft()              {}
func (element) ClientTop()               {}
func (element) ClientWidth()             {}
func (element) CloneNode()               {}
func (element) CompareDocumentPosition() {}
func (element) Contains()                {}
func (element) ContentEditable()         {}
func (element) Dir()                     {}
func (element) FirstChild()              {}
func (element) FirstElementChild()       {}
func (element) Focus()                   {}
func (element) GetAttribute()            {}
func (element) GetAttributeNode()        {}
func (element) GetElementsByClassName()  {}
func (element) GetElementsByTagName()    {}
func (element) HasAttribute()            {}
func (element) HasAttributes()           {}
func (element) HasChildNodes()           {}
func (element) Id()                      {}
func (element) InnerHTML()               {}
func (element) InnerText()               {}
func (element) InsertAdjacentElement()   {}
func (element) InsertAdjacentHTML()      {}
func (element) InsertAdjacentText()      {}
func (element) InsertBefore()            {}
func (element) IsContentEditable()       {}
func (element) IsDefaultNamespace()      {}
func (element) IsEqualNode()             {}
func (element) IsSameNode()              {}
func (element) IsSupported()             {}
func (element) Lang()                    {}
func (element) LastChild()               {}
func (element) LastElementChild()        {}
func (element) NamespaceURI()            {}
func (element) NextSibling()             {}
func (element) NextElementSibling()      {}
func (element) NodeName()                {}
func (element) NodeType()                {}
func (element) NodeValue()               {}
func (element) Normalize()               {}
func (element) OffsetHeight()            {}
func (element) OffsetWidth()             {}
func (element) OffsetLeft()              {}
func (element) OffsetParent()            {}
func (element) OffsetTop()               {}
func (element) OwnerDocument()           {}
func (element) ParentNode()              {}
func (element) ParentElement()           {}
func (element) PreviousSibling()         {}
func (element) PreviousElementSibling()  {}
func (element) QuerySelector()           {}
func (element) QuerySelectorAll()        {}
func (element) RemoveAttribute()         {}
func (element) RemoveAttributeNode()     {}
func (element) RemoveChild()             {}
func (element) RemoveEventListener()     {}
func (element) ReplaceChild()            {}
func (element) ScrollHeight()            {}
func (element) ScrollIntoView()          {}
func (element) ScrollLeft()              {}
func (element) ScrollTop()               {}
func (element) ScrollWidth()             {}
func (element) SetAttribute()            {}
func (element) SetAttributeNode()        {}
func (element) Style()                   {}
func (element) TabIndex()                {}
func (element) TagName()                 {}
func (element) TextContent()             {}
func (element) Title()                   {}
func (element) ToString()                {}
