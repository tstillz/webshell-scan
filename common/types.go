package common

import "time"

func FormatTimestamp(dts time.Time)(cts string){
	cts = dts.Format("2006-01-02 15:04:05")
	return cts
}

type OSInfo struct {
	Hostname    string   `json:"hostname"`
	EnvVars     []string `json:"envVars"`
	Username    string   `json:"username"`
	UserID      string   `json:"userID"`
	RealName    string   `json:"realName"`
	UserHomeDir string   `json:"userHomeDir"`
}

type FileObj struct {
	FilePath    string         `json:"filePath"`
	Size        int64          `json:"size"`
	MD5         string         `json:"md5"`
	Timestamps  FileTimes   `json:"timestamps"`
	Matches     map[string]int `json:"matches"`
	RawContents string         `json:"rawContents,omitempty"`
}
type Metrics struct {
	Scanned    int     `json:"scanned"`
	Matched    int     `json:"matches"`
	Clear      int     `json:"noMatches"`
	ScannedDir string  `json:"directory"`
	ScanTime   float64 `json:"scanDuration"`
	SystemInfo OSInfo  `json:"systemInfo"`
}

type FileTimes struct {
	Birth    string `json:"birth,omitempty"`
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Accessed string `json:"accessed,omitempty"`
}