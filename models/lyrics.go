package models

type Lyrics []Line

type Line []struct {
	Text   string `xml:"Text"`
	Offset string `xml:"Offset"`
}
