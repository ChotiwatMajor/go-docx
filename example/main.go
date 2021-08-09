package main

import (
	"github.com/srdolor/docx"
)

func main() {
	f := docx.NewFile()
	// add new table
	t := f.AddTable()

	for i := 0; i < 3; i++ {
		row := t.AddRow()
		for i := 0; i < 3; i++ {
			c := row.AddCell()
			c.AddText("Hello").Size(20).Color("4900db")
		}
	}
	row := t.AddRow()
	c := row.AddCell()
	c.AddText("One column")
	row = t.AddRow()
	c = row.AddCell()
	c.AddText("Two columns")
	c = row.AddCell()
	c.AddText("Two columns")
	// Adding new paragraph
	para := f.AddParagraph()
	// Adding new page
	para.AddNewPage()
	// Setting Justification/Alignment of paragraph
	para.Properties.Justification(docx.JUSTIFY_END)
	para.AddText("Paragraph starting from right")
	para = f.AddParagraph()
	para.Properties.Justification(docx.JUSTIFY_START)
	para.AddText("Paragraph starting from left")
	para.AddNewLine()
	para = f.AddParagraph()
	para.Properties.Justification(docx.JUSTIFY_BOTH)
	para.AddText("Paragraph distributed over whole line with very long lines").Size(18)
	para.AddNewLine()
	para = f.AddParagraph()
	para.Properties.Justification(docx.JUSTIFY_CENTER)
	para.AddText("Paragraph with centered text").Size(18)
	para.AddNewLine()
	para = f.AddParagraph()
	para.AddText("test font size").Size(22)
	para.AddNewLine()
	para.AddText("test color").Color("808080")
	para.AddNewLine()
	para.AddText("test font size and color").Size(22).Color("ff0f0f")
	nextPara := f.AddParagraph()
	nextPara.AddNewPage()
	nextPara.AddLink("google", `http://google.com`)

	f.Save("./test.docx")
}
