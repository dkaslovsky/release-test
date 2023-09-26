package main

import "fmt"

var (
	name    string = "release-test"
	version string
)

func main() {
	v := "unset"
	if version != "" {
		v = "v" + version
	}
	fmt.Printf("%s(%s): hello world\n", name, v)
}
