package browser

import "syscall/js"

// htmldoc is the root node of the HTML document.
type htmldoc struct {
	document js.Value
}

func (htmldoc) ActiveElement()          {}
func (htmldoc) AddEventListener()       {}
func (htmldoc) AdoptNode()              {}
func (htmldoc) Anchors()                {}
func (htmldoc) Applets()                {}
func (htmldoc) BaseURI()                {}
func (htmldoc) Body()                   {}
func (htmldoc) Close()                  {}
func (htmldoc) Cookie()                 {}
func (htmldoc) Charset()                {}
func (htmldoc) CharacterSet()           {}
func (htmldoc) CreateAttribute()        {}
func (htmldoc) CreateComment()          {}
func (htmldoc) CreateDocumentFragment() {}
func (htmldoc) CreateElement()          {}
func (htmldoc) CreateEvent()            {}
func (htmldoc) CreateTextNode()         {}
func (htmldoc) DefaultView()            {}
func (htmldoc) DesignMode()             {}
func (htmldoc) Doctype()                {}
func (htmldoc) DocumentElement()        {}
func (htmldoc) DocumentMode()           {}
func (htmldoc) DocumentURI()            {}
func (htmldoc) Domain()                 {}
func (htmldoc) DomConfig()              {}
func (htmldoc) Embeds()                 {}
func (htmldoc) ExecCommand()            {}
func (htmldoc) Forms()                  {}
func (d htmldoc) GetElementById(id string) element {
	return element{el: d.document.Call("getElementById", id)}
}
func (htmldoc) GetElementsByClassName() {}
func (htmldoc) GetElementsByName()      {}
func (htmldoc) GetElementsByTagName()   {}
func (htmldoc) HasFocus()               {}
func (htmldoc) Head()                   {}
func (htmldoc) Images()                 {}
func (htmldoc) Implementation()         {}
func (htmldoc) ImportNode()             {}
func (htmldoc) InputEncoding()          {}
func (htmldoc) LastModified()           {}
func (htmldoc) Links()                  {}
func (htmldoc) Normalize()              {}
func (htmldoc) NormalizeDocument()      {}
func (htmldoc) Open()                   {}
func (htmldoc) QuerySelector()          {}
func (htmldoc) QuerySelectorAll()       {}
func (htmldoc) ReadyState()             {}
func (htmldoc) Referrer()               {}
func (htmldoc) RemoveEventListener()    {}
func (htmldoc) RenameNode()             {}
func (htmldoc) Scripts()                {}
func (htmldoc) StrictErrorChecking()    {}
func (htmldoc) Title()                  {}
func (htmldoc) URL()                    {}
func (htmldoc) Write()                  {}
func (htmldoc) Writeln()                {}
