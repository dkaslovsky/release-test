package main

import "fmt"

var (
	name    = "release-test"
	version string
)

func main() {
	v := "unset"
	if version != "" {
		v = "v" + version
	}
	fmt.Printf("%s(%s): hello world\n", name, v)
}
