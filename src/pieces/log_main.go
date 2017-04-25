package main

import (
	"log"
)

// set log format
func init() {
	log.SetPrefix("[main] ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	log.Println("test log msg")
	log.Fatalln("fatal")
}
