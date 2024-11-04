package models

type Track []struct {
	InstrumentSet struct {
		Name      string    `xml:"Name"`
		Type      string    `xml:"Type"`
		LineCount string    `xml:"LineCount"`
		Elements  []Element `xml:"Elements"`
	} `xml:"InstrumentSet"`
	Sounds                 []Sound        `xml:"Sounds"`
	Automations            []Automation   `xml:"Automations"`
	MidiConnection         MidiConnection `xml:"MidiConnection"`
	Transpose              Transpose      `xml:"Transpose"`
	PalmMute               string         `xml:"PalmMute"`
	ForcedSound            string         `xml:"ForcedSound"`
	UseOneChannelPerString string         `xml:"UseOneChannelPerString"`
	IconId                 string         `xml:"IconId"`
	Name                   string         `xml:"Name"`
	SystemsLayout          string         `xml:"SystemsLayout"`
	RSE                    RSE            `xml:"RSE"`
	PlayingStyle           string         `xml:"PlayingStyle"`
	SystemsDefautLayout    string         `xml:"SystemsDefautLayout"`
	Color                  string         `xml:"Color"`
	PlaybackState          string         `xml:"PlaybackState"`
	AudioEngineState       string         `xml:"AudioEngineState"`
	ShortName              string         `xml:"ShortName"`
	Staves                 []Staff        `xml:"Staves"`
	Lyrics                 Lyrics         `xml:"Lyrics"`
}
