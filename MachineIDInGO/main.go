package main

import (
	"flag"
	"fmt"
	"log"
)

const usageStr = `
Usage: machineid [options]

Options:
  --appid    <AppID>    Protect machine id by hashing it together with an app id.

Try:
  machineid
  machineid --appid MyAppID
`

func usage() {
	log.Fatalln(usageStr)
}

func main() {
	var appID string
	flag.StringVar(&appID, "appid", "", "Protect machine id by hashing it together with an app id.")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	var id string
	var err error
	if appID != "" {
		id, err = id.ProtectedID(appID)
	} else {
		id, err = id.ID()
	}
	if err != nil {
		log.Fatalf("Failed to read machine id with error: %s\n", err)
	}
	fmt.Println(id)
}
