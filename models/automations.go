package models

type AutomationType int

const (
	TempoAutomation AutomationType = iota
	VolumeAutomation
	InstrumentAutomation
	BalanceAutomation
)

type Automation struct {
	Type     string `xml:"Type"`
	Linear   string `xml:"Linear"`
	Bar      string `xml:"Bar"`
	Position string `xml:"Position"`
	Visible  string `xml:"Visible"`
	Value    string `xml:"Value"`
}
