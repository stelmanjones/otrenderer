package main

import (
	"bytes"
)

const (
	HeaderBcFS string = "BCFS"
	HeaderBcFz string = "BCFZ"
)

type GPXFile struct {
	name string
	size uint64
	data bytes.Buffer
}

type GPXFileSystem struct {
	files []GPXFile
}

func (fs *GPXFileSystem) Load(path string) error {
	return nil
}

func (fs *GPXFileSystem) ReadHeader(buf *bytes.Buffer) string {
	return string(buf.Next(4))
}
