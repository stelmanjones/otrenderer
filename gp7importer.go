package main

import (
	"archive/zip"
)

type GP7Importer struct {
	Data []*zip.File
}
