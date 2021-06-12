package def

type Torrent struct {
	Hash      string `xml:"hash,attr"`
	TrackerID string `xml:"tracker_id,attr"`
}

type Forum struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type TorrentRecord struct {
	ID           string  `xml:"id,attr"`
	RegisteredAt string  `xml:"registered_at,attr"`
	Size         string  `xml:"size,attr"`
	Torrent      Torrent `xml:"torrent"`
	Forum        Forum   `xml:"forum"`
	Content      string  `xml:"content"`
	Raw          []byte  `xml:",innerxml"`
}
