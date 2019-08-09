// +build ignore

package main

import (
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(
		http.Dir("../../../web/dist"),
		vfsgen.Options{
			Filename:     "static_data.go",
			PackageName:  "static",
			VariableName: "assets",
			BuildTags:    "!dev",
		})

	if err != nil {
		panic(err)
	}
}
