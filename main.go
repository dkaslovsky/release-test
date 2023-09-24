package main

import "fmt"

var (
	name    = "release-test"
	version = "0.0.3"
)

func main() {
	v := "unset"
	if version != "" {
		v = "v" + version
	}
	fmt.Printf("%s(%s): hello world\n", name, v)
}
