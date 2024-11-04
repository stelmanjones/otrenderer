package models

type Score struct {
	Title                     string `xml:"Title"`
	SubTitle                  string `xml:"SubTitle"`
	Artist                    string `xml:"Artist"`
	Album                     string `xml:"Album"`
	Words                     string `xml:"Words"`
	Music                     string `xml:"Music"`
	WordsAndMusic             string `xml:"WordsAndMusic"`
	Copyright                 string `xml:"Copyright"`
	Tabber                    string `xml:"Tabber"`
	Instructions              string `xml:"Instructions"`
	Notices                   string `xml:"Notices"`
	FirstPageHeader           string `xml:"FirstPageHeader"`
	FirstPageFooter           string `xml:"FirstPageFooter"`
	PageHeader                string `xml:"PageHeader"`
	PageFooter                string `xml:"PageFooter"`
	ScoreSystemsDefaultLayout string `xml:"ScoreSystemsDefaultLayout"`
	ScoreSystemsLayout        string `xml:"ScoreSystemsLayout"`
	ScoreZoomPolicy           string `xml:"ScoreZoomPolicy"`
	ScoreZoom                 string `xml:"ScoreZoom"`
	MultiVoice                string `xml:"MultiVoice"`
}
