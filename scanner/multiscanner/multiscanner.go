package multiscanner

import (
	"io"

	"github.com/Snawoot/trusearch/def"
)

type MultiScanner struct {
	scanners []def.RecordScanner
}

func NewMultiScanner(scanners []def.RecordScanner) *MultiScanner {
	return &MultiScanner{scanners}
}

func (s *MultiScanner) Scan() (*def.TorrentRecord, error) {
	for {
		if len(s.scanners) == 0 {
			return nil, io.EOF
		}

		scanres, err := s.scanners[0].Scan()
		if err != nil {
			if err == io.EOF {
				s.scanners = s.scanners[1:]
				continue
			} else {
				return nil, err
			}
		}
		return scanres, err
	}
}
