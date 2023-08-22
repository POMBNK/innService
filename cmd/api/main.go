package main

import (
	"fmt"
	"github.com/POMBNK/shtrafovNetTestTask/pkg/client/rusprofile"
	"log"
)

func main() {
	rprofile := rusprofile.NewClient()
	page, err := rprofile.ParsePage("inn_number")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := rprofile.ParseInfo(page)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}
