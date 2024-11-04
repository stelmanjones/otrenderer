package models

type GPRevision struct {
	Required    string `xml:"required,attr"`
	Recommended string `xml:"recommended,attr"`
	Value       string `xml:",chardata"`
}
