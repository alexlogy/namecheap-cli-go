package models

type Error struct {
	Number  *string `xml:"Number,attr"`
	Message *string `xml:",chardata"`
}
