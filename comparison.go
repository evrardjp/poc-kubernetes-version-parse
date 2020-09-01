package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/version"
)

const (
	StringOne = "4.2.3"
	StringTwo = "4.2.3-rev1"
)

func main() {
	var baseVersion *version.Version
	baseVersion, _ = version.ParseSemantic(StringOne)
	cmpResult, _ := baseVersion.Compare(StringTwo)
	switch {
	case cmpResult == 1:
		fmt.Printf("%v is newer than %v", StringOne, StringTwo)
	case cmpResult == -1:
		fmt.Printf("%v is older than %v", StringOne, StringTwo)
	}
}
