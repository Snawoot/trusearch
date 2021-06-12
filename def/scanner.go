package def

type RecordScanner interface {
	Scan() (*TorrentRecord, error)
}
