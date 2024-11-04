package models

type InstrumentSet struct {
	Name      string    `xml:"Name"`
	Type      string    `xml:"Type"`
	LineCount string    `xml:"LineCount"`
	Elements  []Element `xml:"Elements"`
}
