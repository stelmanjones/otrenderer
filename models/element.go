package models

type Element struct {
	Name          string         `xml:"Name"`
	Type          string         `xml:"Type"`
	SoundbankName string         `xml:"SoundbankName"`
	Articulations []Articulation `xml:"Articulations"`
}
