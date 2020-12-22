package aria2

// Status is the status of Aria2.
type Status struct {
	GID             string `json:"gid"`
	Status          State  `json:"status"`
	InfoHash        string `json:"infoHash"`
	CompletedLength int64  `json:"completedLength,string"`
	TotalLength     int64  `json:"totalLength"`
	DownloadSpeed   int64  `json:"downloadSpeed,string"`
	ErrorCode       int    `json:"errorCode,string"`
	ErrorMessage    string `json:"errorMessage"`
}

// GlobalStatistics is the overall download and upload speeds.
type GlobalStatistics struct {
	DownloadSpeed int64 `json:"downloadSpeed,string"`
	UploadSpeed   int64 `json:"uploadSpeed,string"`
	NumActive     int   `json:"numActive,string"`
	NumWaiting    int   `json:"numWaiting,string"`
	NumStopped    int   `json:"numStopped,string"`
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
