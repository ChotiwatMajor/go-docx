package docx

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"w:h"`
	Data    []interface{}

	file *File
}

func (h *Header) AddText(text string) *Run {
	t := &Text{
		Text: text,
	}

	run := &Run{
		Text:          t,
		RunProperties: &RunProperties{},
	}

	h.Data = append(h.Data, run)

	return run
}
