package xml

import (
	"encoding/xml"
	"io"

	"github.com/Snawoot/trusearch/def"
)

type XMLScanner struct {
	inp io.ReadCloser
	dec *xml.Decoder
	end bool
}

func NewXMLScanner(input io.ReadCloser) *XMLScanner {
	return &XMLScanner{
		inp: input,
		dec: xml.NewDecoder(input),
	}
}

func (s *XMLScanner) Close() error {
	if !s.end {
		s.end = true
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
			if se.Name.Local == "torrent" {
				var tr def.TorrentRecord
				err = s.dec.DecodeElement(&tr, &se)
				if err == io.EOF {
					s.Close()
				}
				if err != nil {
					return nil, err
				}
				tr.RawAttrs = se.Attr
				return &tr, err
			}
		default:
		}
	}
}
