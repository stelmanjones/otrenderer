package models

type RSE struct {
	SoundbankPatch   string   `xml:"SoundbankPatch"`
	ElementsSettings string   `xml:"ElementsSettings"`
	Pickups          Pickups  `xml:"Pickups"`
	EffectChain      []Effect `xml:"EffectChain"`
	ChannelStrip     struct {
		Parameters string `xml:"Parameters"`
	}
}
