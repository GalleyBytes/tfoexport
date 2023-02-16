package main

import (
	"log"

	"github.com/galleybytes/tfohttpclient"
)

func main() {
	b, err := tfohttpclient.ResourceSpec()
	if err != nil {
		panic(err)
	}
	log.Print(string(b))
	// TODO scrape TF_VARs
	// TODO download tfvar files
	// TODO consolidate
	// TODO push to github
}
