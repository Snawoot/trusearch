package forums

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/Snawoot/trusearch/def"
	"github.com/Snawoot/trusearch/util"
)

func Forums(scanner def.RecordScanner, wr io.Writer) int {
	m := util.NewStringSet()
	csvWr := csv.NewWriter(wr)
	defer csvWr.Flush()
	for {
		rec, err := scanner.Scan()
		if err == io.EOF {
			return 0
		}
		if err != nil {
			log.Printf("Got error on input scan: %v", err)
			return 3
		}

		if !m.Has(rec.Forum.ID) {
			m.Add(rec.Forum.ID)
			csvWr.Write([]string{rec.Forum.ID, rec.Forum.Name})
		}
	}
	return 0
}
