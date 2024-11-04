package main

import "github.com/stelmanjones/otrenderer/models"

type GPIFParser struct {
	masterTrackAutomations           map[int][]models.Automation
	automationsPerTrackIdAndBarIndex map[string]map[int][]models.Automation
	tracksMapping                    []string
}

type GPIF struct {
	GPVersion  string            `xml:"GPVersion"`
	GPRevision models.GPRevision `xml:"GPRevision"`
	Encoding   struct {
		EncodingDescription string `xml:"EncodingDescription"`
	} `xml:"Encoding"`
	Score       models.Score       `xml:"Score"`
	MasterTrack models.MasterTrack `xml:"MasterTrack"`
	Tracks      []models.Track     `xml:"Tracks"`
}
