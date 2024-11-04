package models

type MIDI struct {
	LSB     string `xml:"LSB"`
	MSB     string `xml:"MSB"`
	Program string `xml:"Program"`
}
