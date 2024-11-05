package stylesheet

import (
	"archive/zip"

	crunch "github.com/superwhiskers/crunch/v3"
)

type BinaryStylesheetType int

const (
	Bool BinaryStylesheetType = iota
	Int32
	Float
	String
	Point
	Size
	Rectangle
	Color
	Invalid
)

/*
 * A BinaryStylesheet from Guitar Pro 6 and 7 files.
 * The BinaryStylesheet is a simple binary key-value store for additional settings
 * related to the display of the music sheet.
 *
 * File:
 *     int32 (big endian) | Number of KeyValuePairs
 *     KeyValuePair[]     | The raw records
 *
 * KeyValuePair:
 *     1 Byte  | length of the key
 *     n Bytes | key as utf8 encoded string
 *     1 Byte  | Data Type
 *     n Bytes | Value
 *
 * Values based on Data Type:
 *     0 = bool
 *         0===false
 *     1 = int32 (big endian)
 *     2 = float (big endian, IEEE)
 *     3 = string
 *       int16 (big endian) | length of string
 *       n bytes            | utf-8 encoded string
 *     4 = point
 *       int32 (big endian) | X-coordinate
 *       int32 (big endian) | Y-coordinate
 *     5 = size
 *       int32 (big endian) | Width
 *       int32 (big endian) | Height
 *     6 = rectangle
 *       int32 (big endian) | X-coordinate
 *       int32 (big endian) | Y-coordinate
 *       int32 (big endian) | Width
 *       int32 (big endian) | Height
 *     7 = color
 *       1 byte | Red
 *       1 byte | Green
 *       1 byte | Blue
 *       1 byte | Alpha
 */

type KeyValuePair[V comparable] struct {
	Value V
	Key   string
}

type BinaryStylesheet struct {
	Data    []KeyValuePair[any]
	buf     crunch.MiniBuffer
	entries int32
}

func New(sheet *zip.File) *BinaryStylesheet {
	raw, _ := sheet.Open()
	defer raw.Close()
	by := make([]byte, sheet.UncompressedSize64)
	b := crunch.MiniBuffer{}
	b.WriteBytes(0, by)
	return &BinaryStylesheet{
		buf:  b,
		Data: make([]KeyValuePair[any], 0),
	}
}

func (b *BinaryStylesheet) determineDataType(i byte) BinaryStylesheetType {
	switch i {
	case 0:
		return Bool
	case 1:
		return Int32
	case 2:
		return Float
	case 3:
		return String
	case 4:
		return Point
	case 5:
		return Size
	case 6:
		return Rectangle
	case 7:
		return Color
	default:
		return Invalid
	}
}

func (b *BinaryStylesheet) readEntryCount() {
	var out []int32
	b.buf.ReadI32BENext(&out, 1)
	b.entries = out[0]
}

func (b *BinaryStylesheet) readDataType() BinaryStylesheetType {
	var out []byte
	b.buf.ReadBytesNext(&out, 1)
	return b.determineDataType(out[0])
}

func (b *BinaryStylesheet) readEntry() {
	var keyLen []byte
	var key []byte
	var rawtype []byte

	b.buf.ReadBytesNext(&keyLen, 1)
	b.buf.ReadBytesNext(&key, int64(keyLen[0]))
	b.buf.ReadBytesNext(&rawtype, 1)
	t := b.determineDataType(rawtype[0])
}
