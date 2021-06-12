package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Scan struct {
		ScriptFile *os.File   `arg help:"JS script executed for each torrent object"`
		Xmls       []*os.File `arg help:"XML files to process"`
	} `cmd help:"Scan XML and apply JS function defined by script file"`
	Help struct {} `cmd default:"1" help:"Prints CLI synopsis"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "scan <script-file> <xmls>":
		fmt.Printf("%#v", CLI.Scan)
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
}
