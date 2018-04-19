package model

type FG map[string]*Page

type Page struct {
	Title    string   `xml:"title"`
	ID       int      `xml:"id"`
	Revision Revision `xml:"revision"`
	Servant  Servant  `xml:"-"`
}
