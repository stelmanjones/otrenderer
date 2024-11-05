package renderer

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func RgbToHex(rgb *Color) string {
	hexNodes := make([]string, 0, 3)

	for i, v := range rgb {
		if i > 2 {
			break
		}
		hexNodes = append(hexNodes, strconv.FormatInt(int64(v), 16))
	}
	return strings.Join(hexNodes, "")
}

type Color [3]uint8

func (c *Color) HEX() string {
	return "#" + RgbToHex(c)
}

type Renderer struct {
	svg    *svg.SVG
	docBuf bytes.Buffer
	buf    bytes.Buffer
	color  Color
}

func New(w, h int) *Renderer {
	buf := bytes.Buffer{}
	return &Renderer{
		docBuf: buf,
		svg:    svg.New(&buf),
		color:  Color{0, 0, 0},
	}
}

func (r *Renderer) fillmusicSymbol(x, y, scale float32, symbols string, center bool) {
	r.svg.Gtransform(fmt.Sprintf("translate(%d,%d)", x, y))

	style := `style="stroke:none"`
	centerStr := ""
	fill := ""
	if scale != 1 {
		style = fmt.Sprintf(`style="font-size: d%; stroke:none"`, scale*100)
	}

	// TODO: Implement SVGTextAlignment method
	if center {
		centerStr = ``
	}
	if r.color.HEX() != "#000000" {
		fill = fmt.Sprintf(` fill="%s"`, r.color.HEX())
	}

	r.svg.Text(1, 1, symbols, style, fill, centerStr)
	r.svg.TextEnd()
	r.svg.Gend()
}
