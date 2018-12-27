package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var supportedLanguages = map[string]struct{}{"en": struct{}{}, "nl": struct{}{}}

type document struct {
	Language string `json:"language"`
	ID       string `json:"id"`
	Text     string `json:"text"`
}

type documents struct {
	Documents []document `json:"documents"`
}

type sentiment struct {
	ID     string   `json:"id"`
	Score  float64  `json:"score"`
	Errors []string `json:"errors"`
}

type sentiments struct {
	Documents []sentiment `json:"documents"`
	Errors    []string    `json:"errors"`
}

func main() {
	language := flag.String("language", "en", "A language code such as en or nl")
	text := flag.String("text", "", "Text for which we return the sentiment")
	url := flag.String("url", "http://localhost:5000", "URL prefix to Sentiment Analysis container")

	flag.Parse()

	// check supported language
	if _, ok := supportedLanguages[*language]; ok != true {
		log.Fatal("Unsupported language ", *language)
	}

	// check for text
	if *text == "" {
		log.Fatal("Please enter some text")
	}

	color.Red("\nGetting sentiment with Cognitive Services\n")
	color.Red("=========================================\n\n")
	color.Cyan("Language set to: %s\n", *language)
	color.Cyan("Text set to: %s\n\n", *text)

	// create the documents for sentiment analysis
	// only one document for now
	var docs = documents{[]document{{*language, "1", *text}}}

	// http POST requires JSON text to be posted
	docsJSON, _ := json.Marshal(docs)

	resp, err := http.Post(*url+"/text/analytics/v2.0/sentiment", "application/json", bytes.NewBuffer(docsJSON))
	if err != nil {
		log.Fatal("Error posting document ", err)
	}

	// process the response
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var sentimentResponse sentiments

	json.Unmarshal(body, &sentimentResponse)

	color.Yellow("Sentiment score is %f\n\n", sentimentResponse.Documents[0].Score)

}
