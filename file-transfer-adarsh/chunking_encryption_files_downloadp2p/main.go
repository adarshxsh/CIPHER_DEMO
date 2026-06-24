package main

import (
	"flag"
	"fmt"
	"os"
)

// starting point that reads flags to run as tracker, seeder, or leecher

func main() {
	mode := flag.String("mode", "", "run mode: tracker, seeder, or leecher")
	filePtr := flag.String("file", "../demo.txt", "file to seed or leech")
	trackerPtr := flag.String("tracker", "127.0.0.1:8080", "tracker address")
	portPtr := flag.String("port", "8080", "port to listen on (for tracker or seeder)")
	flag.Parse()

	if *mode == "tracker" {
		RunTrackerServer(*portPtr)
	} else if *mode == "seeder" {
		RunSeeder(*trackerPtr, *filePtr, *portPtr)
	} else if *mode == "leecher" {
		RunLeecher(*trackerPtr, *filePtr)
	} else {
		fmt.Println("Please specify a mode: -mode=tracker, -mode=seeder, or -mode=leecher")
		os.Exit(1)
	}
}
