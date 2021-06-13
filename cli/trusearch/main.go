package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Snawoot/trusearch/def"
	"github.com/Snawoot/trusearch/scanner/multiscanner"
	xmlscanner "github.com/Snawoot/trusearch/scanner/xml"
	"github.com/Snawoot/trusearch/task/forums"
	"github.com/Snawoot/trusearch/task/scan"
	"github.com/Snawoot/trusearch/task/split"
	"github.com/Snawoot/trusearch/util"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Scan struct {
		Inline     bool       `help:"specify script directly as argument instead of it's filename"`
		ScriptFile string     `arg help:"JS script executed for each torrent object"`
		Xmls       []*os.File `arg help:"XML files to process"`
	} `cmd help:"Scan XML and apply JS function defined by script file"`
	Forums struct {
		Xmls []*os.File `arg help:"XML files to process"`
	} `cmd help:"Scan XML and print CSV with forum IDs and names"`
	Split struct {
		Whitelist []string   `help:"Comma-separated list of Forum IDs to process. Default is to process everything"`
		Dir       string     `help:"Output directory" default:"." type:"existingdir"`
		Xmls      []*os.File `arg help:"XML files to process"`
	} `cmd help:"Divide XML file into smaller ones by Forum ID"`
	Version struct{} `cmd help:"Print program version and exit"`
	Help    struct{} `cmd default:"1" help:"Prints CLI synopsis"`
}

var version = "undefined"

func run() int {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "scan <script-file> <xmls>":
		script, err := loadScriptText(CLI.Scan.ScriptFile, CLI.Scan.Inline)
		if err != nil {
			log.Fatal(err)
		}
		return scan.Scan(wrapInputs(CLI.Scan.Xmls), script)
	case "forums <xmls>":
		return forums.Forums(wrapInputs(CLI.Forums.Xmls), os.Stdout)
	case "split <xmls>":
		return split.Split(wrapInputs(CLI.Split.Xmls), CLI.Split.Dir, CLI.Split.Whitelist)
	case "help":
		parser, err := kong.New(&CLI)
		if err != nil {
			log.Fatal(err)
		}
		_, err = parser.Parse([]string{"--help"})
		if err != nil {
			log.Fatal(err)
		}
	case "version":
		fmt.Println(version)
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

func loadScriptText(script string, inline bool) (string, error) {
	if inline {
		return script, nil
	}
	scriptBytes, err := ioutil.ReadFile(script)
	if err != nil {
		return "", err
	}
	return string(scriptBytes), nil
}
