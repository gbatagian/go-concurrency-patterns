package main

import (
	"conc/conc"
	"flag"
	"fmt"
)

func main() {

	mode := flag.String("mode", "workerpool", "Example concorrency mode to select. Supported entries: workerpool, channels, waitgroups")
	flag.Parse()

	if *mode == "workerpool" {

		conc.WorkerPoolMainExample()

	} else if *mode == "channels" {

		conc.ChannelsMainExample()

	} else if *mode == "waitgroups" {

		conc.WaitgroupsMainExample()

	} else {

		fmt.Printf("Unsupported mode: %s\n", *mode)

	}

}
