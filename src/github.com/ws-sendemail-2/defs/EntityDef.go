package defs

type Email struct {
	From       string `form:"from"`
	To         string `form:"to"`
	Cc         string `form:"cc"`
	Subject    string `form:"subject"`
	IsHtml     bool   `form:"isHtml"`
	Body       string `form:"body"`
	HasAttach  bool	  `form:"hasAttach"`
	Attachment string `form:"attachment"`
}
