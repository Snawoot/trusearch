package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Snawoot/trusearch/def"
	"github.com/Snawoot/trusearch/task/forums"
	"github.com/Snawoot/trusearch/util"
	xmlscanner "github.com/Snawoot/trusearch/scanner/xml"
	"github.com/Snawoot/trusearch/scanner/multiscanner"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Scan struct {
		ScriptFile *os.File   `arg help:"JS script executed for each torrent object"`
		Xmls       []*os.File `arg help:"XML files to process"`
	} `cmd help:"Scan XML and apply JS function defined by script file"`
	Forums struct {
		Xmls []*os.File `arg help:"XML files to process"`
	} `cmd help:"Scan XML and print CSV with forum IDs and names"`
	Help struct{} `cmd default:"1" help:"Prints CLI synopsis"`
}

func run() int {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "scan <script-file> <xmls>":
		fmt.Printf("%#v", CLI.Scan)
	case "forums <xmls>":
		return forums.Forums(wrapInputs(CLI.Forums.Xmls), os.Stdout)
	case "help":
		parser, err := kong.New(&CLI)
		if err != nil {
			log.Fatal(err)
		}
		_, err = parser.Parse([]string{"--help"})
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown command:", ctx.Command())
	}
	return 0
}

func main() {
	os.Exit(run())
}

func wrapInputs(inputs []*os.File) def.RecordScanner {
	wrappedInputs := make([]def.RecordScanner, len(inputs))
	for i, inp := range inputs {
		wrappedInputs[i] = xmlscanner.NewXMLScanner(util.NewFileWrapper(inp))
	}
	return multiscanner.NewMultiScanner(wrappedInputs)
}
