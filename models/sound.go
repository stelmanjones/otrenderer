package models

type Sound struct {
	Name  string `xml:"Name"`
	Label string `xml:"Label"`
	Path  string `xml:"Path"`
	Role  string `xml:"Role"`
	MIDI  MIDI   `xml:"MIDI"`
	RSE   RSE    `xml:"RSE"`
}
