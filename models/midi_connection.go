package models

type MidiConnection struct {
	Port                    string `xml:"Port"`
	PrimaryChannel          string `xml:"PrimaryChannel"`
	SecondaryChannel        string `xml:"SecondaryChannel"`
	ForeOneChannelPerString string `xml:"ForeOneChannelPerString"`
}
