package internal

import (
	"github.com/jung-kurt/gofpdf"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

// TableList is the abstraction to create a table with header and contents
type TableLink interface {
	Create(header []string, contents [][]string, links [][]int, prop ...props.TableList)
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
func (s *tableLink) Create(header []string, contents [][]string, links [][]int, prop ...props.TableList) {
	if len(header) == 0 {
		return
	}

	if len(contents) == 0 {
		return
	}

	if len(links) == 0 {
		return
	}

	tableProp := props.TableList{}

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

	fPdf, ok := s.pdf.Fpdf().(*gofpdf.Fpdf)
	if !ok {
		return
	}
	// Draw contents
	for index, content := range contents {
		link := links[index]
		contentHeight := s.calcLinesHeight(content, tableProp.ContentProp, tableProp.Align)

		if tableProp.AlternatedBackground != nil && index%2 == 0 {
			s.pdf.SetBackgroundColor(*tableProp.AlternatedBackground)
		}

		s.pdf.Row(contentHeight+1, func() {
			for i, c := range content {
				cs := c
				l := link[i]
				if l == -1 {
					s.pdf.Col(tableProp.ContentProp.GridSizes[i], func() {
						s.pdf.Text(cs, tableProp.ContentProp.ToTextProp(tableProp.Align, 0, false, 0.0))
					})
				} else {
					fPdf.SetFont(string(tableProp.ContentProp.Family), "U", tableProp.ContentProp.Size)
					fPdf.WriteLinkID(tableProp.ContentProp.Size, cs, l)
					fPdf.SetFont("", "", 0)
					s.pdf.Col(tableProp.ContentProp.GridSizes[i], func() {})
				}
			}
		})

		if tableProp.AlternatedBackground != nil && index%2 == 0 {
			s.pdf.SetBackgroundColor(color.NewWhite())
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