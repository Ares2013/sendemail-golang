package defs

type Email struct {
	From       string
	To         string
	Cc         string
	Subject    string
	IsHtml     bool
	Body       string
	HasAttach  bool
	Attachment string
}
