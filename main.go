package main

import (
	"archive/zip"
	"os"

	"github.com/charmbracelet/log"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           log.DebugLevel,
	TimeFormat:      "15:04:05",
	ReportTimestamp: true,
})

func main() {
	archive, err := zip.OpenReader("BreakThoseBones.gp")
	if err != nil {
		log.Error(err.Error())
	}
	defer archive.Close()

	for _, f := range archive.File {
		if f.Name == "Content/score.gpif" {
			file, err := f.Open()
			if err != nil {
				log.Error(err.Error())
			}
			buf := make([]byte, f.UncompressedSize64)
			defer file.Close()
			file.Read(buf)
			log.Info(string(buf))
		}
		log.Info(f.Name)
	}
}
