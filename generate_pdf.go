package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"os"
)

func main() {
	doc := pdf.New()
	canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
	canvas.Translate(500, 500)
	// canvas.SetColor(230, 100, 30)

	canvas.SetStrokeColor(20, 40, 60)
	path := new(pdf.Path)
	path.Move(pdf.Point{0, 0})
	path.Line(pdf.Point{0, 50})
	canvas.Stroke(path)

	text := new(pdf.Text)
	text.SetFont(pdf.Helvetica, 14)
	text.Text("Hello, World!")
	canvas.DrawText(text)

	canvas.Close()

	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
