package forums

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/Snawoot/trusearch/def"
)

func Forums(scanner def.RecordScanner, wr io.Writer) int {
	m := make(map[string]struct{})
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

		_, ok := m[rec.Forum.ID]
		if !ok {
			m[rec.Forum.ID] = struct{}{}
			csvWr.Write([]string{rec.Forum.ID, rec.Forum.Name})
		}
	}
	return 0
}
