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

package main

import (
	"fmt"

	"kertificate.io/kertificate/pkg/version"

	"github.com/spf13/cobra"
)

var short bool

func init() {
	var command = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of Kertificate",
		Run: func(command *cobra.Command, args []string) {
			info := version.Info()
			if short {
				fmt.Println(info.Version)
			} else {
				fmt.Println(info.Version + " (go version: " + info.GoVersion + ", commit: " + info.Commit + "), built on " + info.BuildDateHumanReadable())
			}
		},
	}

	command.Flags().BoolVarP(&short, "short", "s", false, "Prints application version only")
	rootCommand.AddCommand(command)
}
