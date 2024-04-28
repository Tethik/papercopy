package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Tethik/go-template/internal/niceware"
	"github.com/raceresult/gopdf"
	"github.com/raceresult/gopdf/types"
)

var (
	version string
	commit  string
	build   string
)

var versionFlag = flag.Bool("v", false, "Print version information and quit")
var versionString = fmt.Sprintf("niceware %s, commit %s, build %s", version, commit, build)
var passphraseFlag = flag.String("p", "", "Passphrase to encrypt mnemonic")
var filenameFlag = flag.String("f", "paper-copy.pdf", "Filename to save PDF to")
var helpString = `Usage: papercopy [options]

Options:
  -v	Print version information and quit
  -f 	Filename to save QR code to (default "repo-qrcode.jpeg")
  -p 	Passphrase to encrypt
`
var instructions = `This document contains a passphrase that can be used to recover a cryptographic key. The passphrase is in the "niceware" format. The QR code contains the same passphrase. 
Keep this document in a safe place and do not share it with anyone.`

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		return
	}

	if len(*filenameFlag) == 0 {
		fmt.Println(helpString)
		return
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("could not read input: %v", err)
		return
	}

	words, err := niceware.BytesToWords(input)
	if err != nil {
		fmt.Printf("could not convert bytes to words: %v", err)
		return
	}

	passphrase := strings.Join(words, " ")

	// create new PDF Builder
	pb := gopdf.New()

	// use a built-in standard fontm
	f, err := pb.NewStandardFont(types.StandardFont_Helvetica, types.EncodingWinAnsi)
	if err != nil {
		fmt.Printf("could not create font: %v", err)
		return
	}

	// add first page
	bounds := gopdf.GetStandardPageSize(gopdf.PageSizeA4, false)
	p := pb.NewPage(bounds)

	// add border
	p.AddElement(&gopdf.RectElement{
		Width:     gopdf.MM(210 - 10),
		Height:    gopdf.MM(297 - 10),
		Left:      gopdf.MM(5),
		Top:       gopdf.MM(5),
		LineWidth: gopdf.MM(1),
		FillColor: gopdf.ColorRGB{R: 255, G: 255, B: 255},
		LineColor: gopdf.ColorRGB{R: 0, G: 0, B: 0},
	})

	// add inner border
	p.AddElement(&gopdf.RectElement{
		Width:     gopdf.MM(210 - 20),
		Height:    gopdf.MM(100),
		Left:      gopdf.MM(10),
		Top:       gopdf.MM(89),
		LineWidth: gopdf.MM(1),
		FillColor: gopdf.ColorRGB{R: 255, G: 255, B: 255},
		LineColor: gopdf.ColorRGB{R: 0, G: 0, B: 0},
	})

	p.AddElement(&gopdf.RectElement{
		Width:     gopdf.MM(35),
		Height:    gopdf.MM(20),
		Left:      gopdf.MM(13),
		Top:       gopdf.MM(88),
		LineWidth: gopdf.MM(1),
		FillColor: gopdf.ColorRGB{R: 255, G: 255, B: 255},
		LineColor: gopdf.ColorRGB{R: 255, G: 255, B: 255},
	})

	// Add instructions to the PDF
	p.AddElement(&gopdf.TextBoxElement{
		TextElement: gopdf.TextElement{
			TextChunk: gopdf.TextChunk{
				Text:     instructions,
				Font:     f,
				FontSize: 12,
			},
			Left: gopdf.MM(10),
			Top:  gopdf.MM(40),
		},
		Width:  gopdf.MM(190),
		Height: gopdf.MM(100),
	})

	p.AddElement(&gopdf.TextElement{
		TextChunk: gopdf.TextChunk{
			Text:     "Passphrase",
			Font:     f,
			FontSize: 14,
			Bold:     true,
		},
		Left: gopdf.MM(15),
		Top:  gopdf.MM(91),
	})

	p.AddElement(&gopdf.TextBoxElement{
		TextElement: gopdf.TextElement{
			TextChunk: gopdf.TextChunk{
				Text:     passphrase,
				Font:     f,
				FontSize: 16,
			},
			Left: gopdf.MM(20),
			Top:  gopdf.MM(160),
		},
		Width:         gopdf.MM(170),
		Height:        gopdf.MM(50),
		VerticalAlign: gopdf.VerticalAlignTop,
	})

	p.AddElement(&gopdf.QRCodeElement{
		Text: passphrase,
		Left: gopdf.MM(80),
		Top:  gopdf.MM(105),
		Size: gopdf.MM(50),
	})

	// Add footer
	p.AddElement(&gopdf.TextBoxElement{
		TextElement: gopdf.TextElement{
			TextChunk: gopdf.TextChunk{
				Text: fmt.Sprintf(`This document was generated using the 'papercopy' tool - version %v 
				For more information, see https://github.com/Tethik/papercopy`, version),
				Font:     f,
				FontSize: 12,
			},
			Left: gopdf.MM(10),
			Top:  gopdf.MM(277),
		},
		Width:  gopdf.MM(190),
		Height: gopdf.MM(50),
	})

	// output
	bts, err := pb.Build()
	if err != nil {
		fmt.Printf("could not build PDF: %v", err)
		return
	}

	os.WriteFile(*filenameFlag, bts, 0644)
	fmt.Printf("PDF saved to %s\n", *filenameFlag)
}
