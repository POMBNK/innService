package rusprofile

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Info struct {
	Inn         string `json:"inn"`
	Kpp         string `json:"kpp"`
	CompanyName string `json:"company_name"`
	Fio         string `json:"fio"`
}

func (c *Client) ParseInfo(htmlPage []byte) (Info, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlPage)))

	if err != nil {
		log.Fatal(err)
	}

	//companyName
	companyName := doc.Find("div.company-name").Each(func(i int, s *goquery.Selection) {
		s.Find("span")
	}).Text()

	//fio
	fio := doc.Find("a.link-arrow.gtm_main_fl").Each(func(i int, s *goquery.Selection) {
		s.Find("span")
	}).Text()

	//inn
	inn := doc.Find("span#clip_inn.copy_target").Each(func(i int, s *goquery.Selection) {
		s.Find("span")
	}).Text()

	//kpp
	kpp := doc.Find("span#clip_kpp.copy_target").Each(func(i int, s *goquery.Selection) {
		s.Find("span")
	}).Text()

	return Info{
		Inn:         inn,
		Kpp:         kpp,
		CompanyName: companyName,
		Fio:         fio,
	}, nil
}

func (c *Client) ToJSON(info Info) error {
	inf, err := json.Marshal(info)
	if err != nil {
		return err
	}
	fmt.Printf("%s", inf)
	return nil
}
