package models

type MasterTrack struct {
	Tracks      string       `xml:"Tracks"`
	Automations []Automation `xml:"Automations"`
	RSE         struct {
		Master struct {
			Effect struct {
				ByPass     string `xml:"ByPass"`
				Parameters string `xml:"Parameters"`
			} `xml:"Effect"`
		} `xml:"Master"`
	} `xml:"RSE"`
}
