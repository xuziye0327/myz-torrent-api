package aria2

// Status is the status of Aria2.
type Status struct {
	GID             string     `json:"gid"`
	Status          State      `json:"status"`
	InfoHash        string     `json:"infoHash"`
	CompletedLength int64      `json:"completedLength,string"`
	TotalLength     int64      `json:"totalLength,string"`
	DownloadSpeed   int64      `json:"downloadSpeed,string"`
	ErrorCode       int        `json:"errorCode"`
	ErrorMessage    string     `json:"errorMessage"`
	Bittorrent      Bittorrent `json:"bittorrent"`
}

// GlobalStatistics is the overall download and upload speeds.
type GlobalStatistics struct {
	DownloadSpeed int64 `json:"downloadSpeed"`
	UploadSpeed   int64 `json:"uploadSpeed"`
	NumActive     int   `json:"numActive"`
	NumWaiting    int   `json:"numWaiting"`
	NumStopped    int   `json:"numStopped"`
}

// Bittorrent struct
type Bittorrent struct {
	Comment string         `json:"comment"`
	Info    BittorrentInfo `json:"info"`
}

// BittorrentInfo struct
type BittorrentInfo struct {
	Name string `json:"name"`
}

// State is running status of aria
type State string

const (
	active   State = "active"
	waiting  State = "waiting"
	paused   State = "paused"
	err      State = "error"
	complete State = "complete"
	removed  State = "removed"
)

// Method struct
type Method struct {
	MethodName string        `json:"methodName"`
	Params     []interface{} `json:"params,omitempty"`
}
