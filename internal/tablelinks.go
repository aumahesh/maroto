package internal

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// TableList is the abstraction to create a table with header and contents
type TableLink interface {
	Create(header []string, contents [][]string, links [][]int, prop ...props.TableLink)
	BindGrid(part MarotoGridPart)
}

type tableLink struct {
	pdf  MarotoGridPart
	text Text
	font Font
}

// NewTableList create a TableList
func NewTableLink(text Text, font Font) *tableLink {
	return &tableLink{
		text: text,
		font: font,
	}
}

// BindGrid bind the grid system to TableList
func (s *tableLink) BindGrid(pdf MarotoGridPart) {
	s.pdf = pdf
}

// Create create a header section with a list of strings and
// create many rows with contents
func (s *tableLink) Create(header []string, contents [][]string, links [][]int, prop ...props.TableLink) {
	if len(header) == 0 {
		return
	}

	if len(contents) == 0 {
		return
	}

	if len(links) == 0 {
		return
	}

	tableProp := props.TableLink{}

	if len(prop) > 0 {
		tableProp = prop[0]
	}

	tableProp.MakeValid(header, contents)
	headerHeight := s.calcLinesHeight(header, tableProp.HeaderProp, tableProp.Align)

	// Draw header
	s.pdf.Row(headerHeight+1, func() {
		for i, h := range header {
			hs := h

			s.pdf.Col(tableProp.HeaderProp.GridSizes[i], func() {
				reason := hs
				s.pdf.Text(reason, tableProp.HeaderProp.ToTextProp(tableProp.Align, 0, false, 0.0))
			})
		}
	})

	// Define space between header and contents
	s.pdf.Row(tableProp.HeaderContentSpace, func() {
		s.pdf.ColSpace(0)
	})

	var col color.Color

	// Draw contents
	for index, content := range contents {
		link := links[index]
		contentHeight := s.calcLinesHeight(content, tableProp.ContentProp, tableProp.Align)

		if tableProp.AlternatedBackground != nil && index%2 == 0 {
			col = *tableProp.AlternatedBackground
			s.pdf.SetBackgroundColor(col)
		}

		s.pdf.Row(contentHeight+1, func() {
			for i, c := range content {
				cs := c
				s.pdf.Col(tableProp.ContentProp.GridSizes[i], func() {
					p := tableProp.ContentProp.ToTextProp(tableProp.Align, 0, false, 0.0)
					if i == tableProp.HighlightColumn {
						s.pdf.SetTextColor(tableProp.HighlightColors[index])
						p.Style = consts.Bold
					} else {}
					if  link[i] == -1 {
						s.pdf.Text(cs, p)
					} else {
						s.pdf.TextWithLink(cs, link[i], col, p)
					}
					if i == tableProp.HighlightColumn {
						s.pdf.SetTextColor(color.Color{0,0,0})
					}
				})
			}
		})

		if tableProp.AlternatedBackground != nil && index%2 == 0 {
			col = color.NewWhite()
			s.pdf.SetBackgroundColor(col)
		}

		if tableProp.Line {
			s.pdf.Line(1.0)
		}
	}
}

func (s *tableLink) calcLinesHeight(textList []string, contentProp props.TableListContent, align consts.Align) float64 {
	maxLines := 1.0

	left, _, right, _ := s.pdf.GetPageMargins()
	width, _ := s.pdf.GetPageSize()
	usefulWidth := float64(width - left - right)

	textProp := contentProp.ToTextProp(align, 0, false, 0.0)

	for i, text := range textList {
		gridSize := float64(contentProp.GridSizes[i])
		percentSize := gridSize / consts.MaxGridSum
		colWidth := usefulWidth * percentSize
		qtdLines := float64(s.text.GetLinesQuantity(text, textProp, colWidth))
		if qtdLines > maxLines {
			maxLines = qtdLines
		}
	}

	_, _, fontSize := s.font.GetFont()

	// Font size corrected by the scale factor from "mm" inside gofpdf f.k
	fontHeight := fontSize / s.font.GetScaleFactor()

	return fontHeight * maxLines
}
