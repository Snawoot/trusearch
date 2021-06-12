package xml

import (
	"encoding/xml"
	"io"

	"github.com/Snawoot/trusearch/def"
)

type XMLScanner struct {
	inp io.ReadCloser
	dec xml.Decoder
	end bool
}

func NewXMLScanner(input io.ReadCloser) *XMLScanner {
	s := XMLScanner{
		inp: input,
		dec: xml.NewDecoder(input),
	}
	return s
}

func (s *XMLScanner) Close() error {
	if !end {
		end = true
		return s.inp.Close()
	} else {
		return nil
	}
}

func (s *XMLScanner) Scan() (*def.TorrentRecord, error) {
	if s.end {
		return nil, io.EOF
	}

	for {
		t, err := s.dec.Token()
		if err == io.EOF {
			s.Close()
		}
		if err != nil {
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name == "torrent" {
				var tr def.TorrentRecord
				err = s.dec.DecodeElement(&tr, se)
				if err == io.EOF {
					s.Close()
				}
				if err != nil {
					return nil, err
				}
				return &tr, err
			}
		default:
		}
	}
}
