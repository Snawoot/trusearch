package split

import (
	//"bufio"
	"fmt"
	"io"
	"log"

	"github.com/Snawoot/trusearch/def"
	//"github.com/Snawoot/trusearch/util"
)

func Split(scanner def.RecordScanner, dirPath string, whitelist []string) int {
	//m := make(map[string]*bufio.Writer)
	for {
		rec, err := scanner.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Got error on input scan: %v", err)
			return 3
		}

		//_, ok := m[rec.Forum.ID]
		//if !ok {
		//	m[rec.Forum.ID] = struct{}{}
		//	csvWr.Write([]string{rec.Forum.ID, rec.Forum.Name})
		//}
		fmt.Println(string(rec.Raw))
	}

	// Flush outputs
	return 0
}
