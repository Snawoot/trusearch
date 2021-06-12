package forums

import (
	"io"
	"log"

	"github.com/Snawoot/trusearch/def"
)

func Forums(scanner def.RecordScanner) int {
	for {
		rec, err := scanner.Scan()
		if err == io.EOF {
			return 0
		}
		if err != nil {
			log.Printf("Got error on input scan: %v", err)
			return 3
		}

		log.Printf("%#v", rec)
	}
	return 0
}
