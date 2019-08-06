// Copyright 2019 Anton Johansson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
