package main

type MusicFontSymbol rune

/**
 * Lists all music font symbols. The names
 * and values are aligned with the SMuFL standard.
 */
const (
	None MusicFontSymbol = -1

	GClef                    MusicFontSymbol = 0xe050
	CClef                    MusicFontSymbol = 0xe05c
	FClef                    MusicFontSymbol = 0xe062
	UnpitchedPercussionClef1 MusicFontSymbol = 0xe069
	SixStringTabClef         MusicFontSymbol = 0xe06d
	FourStringTabClef        MusicFontSymbol = 0xe06e

	TimeSig0         MusicFontSymbol = 0xe080
	TimeSig1         MusicFontSymbol = 0xe081
	TimeSig2         MusicFontSymbol = 0xe082
	TimeSig3         MusicFontSymbol = 0xe083
	TimeSig4         MusicFontSymbol = 0xe084
	TimeSig5         MusicFontSymbol = 0xe085
	TimeSig6         MusicFontSymbol = 0xe086
	TimeSig7         MusicFontSymbol = 0xe087
	TimeSig8         MusicFontSymbol = 0xe088
	TimeSig9         MusicFontSymbol = 0xe089
	TimeSigCommon    MusicFontSymbol = 0xe08a
	TimeSigCutCommon MusicFontSymbol = 0xe08b

	NoteheadDoubleWholeSquare MusicFontSymbol = 0xe0a1
	NoteheadDoubleWhole       MusicFontSymbol = 0xe0a0
	NoteheadWhole             MusicFontSymbol = 0xe0a2
	NoteheadHalf              MusicFontSymbol = 0xe0a3
	NoteheadBlack             MusicFontSymbol = 0xe0a4
	NoteheadNull              MusicFontSymbol = 0xe0a5
	NoteheadXOrnate           MusicFontSymbol = 0xe0aa
	NoteheadTriangleUpWhole   MusicFontSymbol = 0xe0bb
	NoteheadTriangleUpHalf    MusicFontSymbol = 0xe0bc
	NoteheadTriangleUpBlack   MusicFontSymbol = 0xe0be
	NoteheadDiamondBlackWide  MusicFontSymbol = 0xe0dc
	NoteheadDiamondWhite      MusicFontSymbol = 0xe0dd
	NoteheadDiamondWhiteWide  MusicFontSymbol = 0xe0de
	NoteheadCircleX           MusicFontSymbol = 0xe0b3
	NoteheadXWhole            MusicFontSymbol = 0xe0a7
	NoteheadXHalf             MusicFontSymbol = 0xe0a8
	NoteheadXBlack            MusicFontSymbol = 0xe0a9
	NoteheadParenthesis       MusicFontSymbol = 0xe0ce
	NoteheadSlashedBlack2     MusicFontSymbol = 0xe0d0
	NoteheadCircleSlash       MusicFontSymbol = 0xe0f7
	NoteheadHeavyX            MusicFontSymbol = 0xe0f8
	NoteheadHeavyXHat         MusicFontSymbol = 0xe0f9

	NoteheadSlashVerticalEnds MusicFontSymbol = 0xe100
	NoteheadSlashWhiteWhole   MusicFontSymbol = 0xe102
	NoteheadSlashWhiteHalf    MusicFontSymbol = 0xe103

	NoteQuarterUp MusicFontSymbol = 0xe1d5
	NoteEighthUp  MusicFontSymbol = 0xe1d7

	Tremolo3 MusicFontSymbol = 0xe222
	Tremolo2 MusicFontSymbol = 0xe221
	Tremolo1 MusicFontSymbol = 0xe220

	FlagEighthUp                   MusicFontSymbol = 0xe240
	FlagEighthDown                 MusicFontSymbol = 0xe241
	FlagSixteenthUp                MusicFontSymbol = 0xe242
	FlagSixteenthDown              MusicFontSymbol = 0xe243
	FlagThirtySecondUp             MusicFontSymbol = 0xe244
	FlagThirtySecondDown           MusicFontSymbol = 0xe245
	FlagSixtyFourthUp              MusicFontSymbol = 0xe246
	FlagSixtyFourthDown            MusicFontSymbol = 0xe247
	FlagOneHundredTwentyEighthUp   MusicFontSymbol = 0xe248
	FlagOneHundredTwentyEighthDown MusicFontSymbol = 0xe249
	FlagTwoHundredFiftySixthUp     MusicFontSymbol = 0xe24a
	FlagTwoHundredFiftySixthDown   MusicFontSymbol = 0xe24b

	AccidentalFlat                      MusicFontSymbol = 0xe260
	AccidentalNatural                   MusicFontSymbol = 0xe261
	AccidentalSharp                     MusicFontSymbol = 0xe262
	AccidentalDoubleSharp               MusicFontSymbol = 0xe263
	AccidentalDoubleFlat                MusicFontSymbol = 0xe264
	AccidentalQuarterToneFlatArrowUp    MusicFontSymbol = 0xe270
	AccidentalQuarterToneSharpArrowUp   MusicFontSymbol = 0xe274
	AccidentalQuarterToneNaturalArrowUp MusicFontSymbol = 0xe272

	ArticAccentAbove   MusicFontSymbol = 0xe4a0
	ArticAccentBelow   MusicFontSymbol = 0xe4a1
	ArticStaccatoAbove MusicFontSymbol = 0xe4a2
	ArticStaccatoBelow MusicFontSymbol = 0xe4a3
	ArticTenutoAbove   MusicFontSymbol = 0xe4a4
	ArticTenutoBelow   MusicFontSymbol = 0xe4a5
	ArticMarcatoAbove  MusicFontSymbol = 0xe4ac
	ArticMarcatoBelow  MusicFontSymbol = 0xe4aD

	FermataAbove      MusicFontSymbol = 0xe4c0
	FermataShortAbove MusicFontSymbol = 0xe4c4
	FermataLongAbove  MusicFontSymbol = 0xe4c6

	RestLonga                  MusicFontSymbol = 0xe4e1
	RestDoubleWhole            MusicFontSymbol = 0xe4e2
	RestWhole                  MusicFontSymbol = 0xe4e3
	RestHalf                   MusicFontSymbol = 0xe4e4
	RestQuarter                MusicFontSymbol = 0xe4e5
	RestEighth                 MusicFontSymbol = 0xe4e6
	RestSixteenth              MusicFontSymbol = 0xe4e7
	RestThirtySecond           MusicFontSymbol = 0xe4e8
	RestSixtyFourth            MusicFontSymbol = 0xe4e9
	RestOneHundredTwentyEighth MusicFontSymbol = 0xe4ea
	RestTwoHundredFiftySixth   MusicFontSymbol = 0xe4eb

	Repeat1Bar  MusicFontSymbol = 0xe500
	Repeat2Bars MusicFontSymbol = 0xe501

	Ottava           MusicFontSymbol = 0xe510
	OttavaAlta       MusicFontSymbol = 0xe511
	OttavaBassaVb    MusicFontSymbol = 0xe51c
	Quindicesima     MusicFontSymbol = 0xe514
	QuindicesimaAlta MusicFontSymbol = 0xe515

	DynamicPPP   MusicFontSymbol = 0xe52a
	DynamicPP    MusicFontSymbol = 0xe52b
	DynamicPiano MusicFontSymbol = 0xe520
	DynamicMP    MusicFontSymbol = 0xe52c
	DynamicMF    MusicFontSymbol = 0xe52d
	DynamicForte MusicFontSymbol = 0xe522
	DynamicFF    MusicFontSymbol = 0xe52f
	DynamicFFF   MusicFontSymbol = 0xe530

	OrnamentTrill MusicFontSymbol = 0xe566

	StringsDownBow MusicFontSymbol = 0xe610
	StringsUpBow   MusicFontSymbol = 0xe612

	PictEdgeOfCymbal MusicFontSymbol = 0xe729

	GuitarString0 MusicFontSymbol = 0xe833
	GuitarString1 MusicFontSymbol = 0xe834
	GuitarString2 MusicFontSymbol = 0xe835
	GuitarString3 MusicFontSymbol = 0xe836
	GuitarString4 MusicFontSymbol = 0xe837
	GuitarString5 MusicFontSymbol = 0xe838
	GuitarString6 MusicFontSymbol = 0xe839
	GuitarString7 MusicFontSymbol = 0xe83A
	GuitarString8 MusicFontSymbol = 0xe83B
	GuitarString9 MusicFontSymbol = 0xe83C

	GuitarGolpe MusicFontSymbol = 0xe842

	FretboardX MusicFontSymbol = 0xe859
	FretboardO MusicFontSymbol = 0xe85a

	WiggleTrill             MusicFontSymbol = 0xeaa4
	WiggleVibratoMediumFast MusicFontSymbol = 0xeade

	OctaveBaselineM MusicFontSymbol = 0xec95
	OctaveBaselineB MusicFontSymbol = 0xec93
)
