package def

type Torrent struct {
	Hash      string `xml:"hash,attr"`
	TrackerID string `xml:"tracker_id,attr"`
}

type TorrentRecord struct {
	ID           string  `xml:"id,attr"`
	RegisteredAt string  `xml:"registered_at,attr"`
	Size         string  `xml:"size,attr"`
	Torrent      Torrent `xml:"torrent"`
	Content      string  `xml:"content"`
}
