package def

import (
	"encoding/xml"
)

type Torrent struct {
	Hash      string `xml:"hash,attr"`
	TrackerID string `xml:"tracker_id,attr"`
}

type Forum struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type TorrentRecord struct {
	ID           string     `xml:"id,attr"`
	RegisteredAt string     `xml:"registred_at,attr"`
	Size         string     `xml:"size,attr"`
	Torrent      Torrent    `xml:"torrent"`
	Forum        Forum      `xml:"forum"`
	Content      string     `xml:"content"`
	RawAttrs     []xml.Attr `xml:-`
	RawContent   []byte     `xml:",innerxml"`
}
