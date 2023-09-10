package main

import (
	"github.com/POMBNK/shtrafovNetTestTask/internal/server/gateAwayServer"
	"log"
)

func main() {
	//rprofile := rusprofile.NewClient()
	//page, err := rprofile.ParsePage("inn_number")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//resp, err := rprofile.ParseInfo(page)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(resp)
	server := gateAwayServer.NewServer()
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
