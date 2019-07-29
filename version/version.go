package version

import (
	"fmt"
	"time"
)

var (
	version   string
	goVersion string
	commit    string
	buildDate string
)

// VersionInfo holds information about the current version
type VersionInfo struct {
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
}

// Info returns the version information
func Info() VersionInfo {
	return VersionInfo{
		Version:   version,
		GoVersion: goVersion,
		Commit:    commit,
		BuildDate: buildDate,
	}
}

func (info *VersionInfo) BuildDateHumanReadable() string {
	timestamp, err := time.Parse(time.RFC3339, info.BuildDate)
	if err != nil {
		fmt.Println("Could not parse build date:", err)
		return info.BuildDate
	}
	return timestamp.Format(time.RFC1123Z)
}
