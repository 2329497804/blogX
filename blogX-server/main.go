package main

import (
	"fmt"

	"github.com/2329497804/blogX-server/core"
	"github.com/2329497804/blogX-server/flags"
)

func main() {
	flags.Parse()
	fmt.Println(flags.FlagOptions)
	core.ReadConfig()
}
