package main

import (
	"github.com/gmbh-micro/gmbh"
)

func main() {

	client, err := gmbh.NewService().Config("../gmbh.yaml")
	if err != nil {
		panic(err)
	}

	client.Route("about", handleAbout)
	client.Verbose().Start()
}

func handleAbout(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "About: This is a remote tester service."
	return
}