package models

type Articulations []Articulation

type Articulation struct {
	Name               string `xml:"Name"`
	StaffLine          string `xml:"StaffLine"`
	Noteheads          string `xml:"Noteheads"`
	TechniquePlacement string `xml:"TechniquePlacement"`
	TechniqueSymbol    string `xml:"TechniqueSymbol"`
	InputMidiNumbers   string `xml:"InputMidiNumbers"`
	OutputRSESound     string `xml:"OutputRSESound"`
	OutputMidiNumber   string `xml:"OutputMidiNumber"`
}
