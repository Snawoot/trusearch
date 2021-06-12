package split

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/Snawoot/trusearch/def"
	"github.com/Snawoot/trusearch/util"
)

func initWriter(rec *def.TorrentRecord, dirPath string) (*bufio.Writer, error) {
	filename := filepath.Join(dirPath, fmt.Sprintf("forum_%s.xml", rec.Forum.ID))
	wr := util.NewBufferedAppendWriter(filename)
	_, err := wr.Write([]byte("<torrents>\n"))
	if err != nil {
		return nil, err
	}
	return wr, nil
}

func closeWriter(wr *bufio.Writer) error {
	_, err := wr.Write([]byte("</torrents>\n"))
	if err != nil {
		return err
	}
	return wr.Flush()
}

func writeElement(rec *def.TorrentRecord, wr io.Writer) error {
	_, err := wr.Write([]byte("<torrent "))
	if err != nil {
		return err
	}

	pairs := make([]string, len(rec.RawAttrs))
	for i, attr := range rec.RawAttrs {
		pairs[i] = fmt.Sprintf("%s=\"%s\"", attr.Name.Local, attr.Value)
	}
	_, err = wr.Write([]byte(strings.Join(pairs, " ")))
	if err != nil {
		return err
	}

	_, err = wr.Write([]byte(">"))
	if err != nil {
		return err
	}

	_, err = wr.Write(rec.RawContent)
	if err != nil {
		return err
	}

	_, err = wr.Write([]byte("</torrent>\n"))
	if err != nil {
		return err
	}

	return nil
}

func Split(scanner def.RecordScanner, dirPath string, whitelist []string) int {
	wl := util.StringSetFromSlice(whitelist)
	m := make(map[string]*bufio.Writer)
	for {
		rec, err := scanner.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Got error on input scan: %v", err)
			return 3
		}

		if wl.Count() > 0 && !wl.Has(rec.Forum.ID) {
			continue
		}

		wr, ok := m[rec.Forum.ID]
		if !ok {
			wr, err = initWriter(rec, dirPath)
			if err != nil {
				log.Printf("Got error on initializing output writer: %v", err)
				return 4
			}
			m[rec.Forum.ID] = wr
		}
		writeElement(rec, wr)
	}

	for _, wr := range m {
		closeWriter(wr)
	}
	return 0
}
