package version

var (
	version   string
	goVersion string
	commit    string
)

// VersionInfo holds information about the current version
type VersionInfo struct {
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
	Commit    string `json:"commit"`
}

// Info returns the version information
func Info() VersionInfo {
	return VersionInfo{
		Version:   version,
		GoVersion: goVersion,
		Commit:    commit,
	}
}
