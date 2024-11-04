package models

type Property struct {
	Fret         string `xml:"Fret"`
	Number       string `xml:"Number"`
	Bitset       string `xml:"Bitset"`
	Pitches      string `xml:"Pitches"`
	Instrument   string `xml:"Instrument"`
	Label        string `xml:"Label"`
	LabelVisible string `xml:"LabelVisible"`
	Items        string `xml:"Items"`
	Name         string `xml:"Name"`
}
