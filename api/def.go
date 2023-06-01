package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Definition struct {
	En        string
	Es        string
	Categorie string
	En_def    string
}

func DefinitionLogic(word string) Definition {
	url := fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english-spanish/%s", word)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var def Definition
	doc.Find("#page-content > div.page").Each(func(i int, s *goquery.Selection) {
		title := s.Find("#dataset-caldes > div.pr.entry-body > div.di.english-spanish > div > span > div > span > div > div.di-title > h2 > span").Text()
		categorie := s.Find("#dataset-caldes > div.pr.entry-body > div.di.english-spanish > div > span > div > div.pos-body > div > h3 > span").Text()
		description := s.Find("#dataset-caldes > div.pr.entry-body > div.di.english-spanish > div > span > div > div.pos-body > div > div.sense-body.dsense_b > div.def-block.ddef_block > div.def-head.ddef_h > div").Text()
		es := s.Find("#dataset-caldes > div.pr.entry-body > div.di.english-spanish > div > span > div > div.pos-body > div > div.sense-body.dsense_b > div.def-block.ddef_block > div.def-body.ddef_b > span").Text()
		def = Definition{
			En:        title,
			Es:        es,
			Categorie: categorie,
			En_def:    description,
		}
	})
	return def
}
