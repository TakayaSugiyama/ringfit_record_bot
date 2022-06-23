package ringfit_record_bot

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"

	"github.com/joho/godotenv"
)

type TwitterApiResult struct {
	Data struct {
		ID          string `json:"id"`
		Text        string `json:"text"`
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments"`
	} `json:"data"`
	Includes struct {
		Media []struct {
			MediaKey string `json:"media_key"`
			Type     string `json:"type"`
			URL      string `json:"url"`
		} `json:"media"`
	} `json:"includes"`
}

var rule = regexp.MustCompile("[0-9]+")

func fetchPhotoURL() string {
	godotenv.Load(".env")
	var twUrl string = "https://twitter.com/tky_7201/status/1539739471054135296"
	var twitttID = rule.FindAllStringSubmatch(twUrl, -1)[1][0]
	var endpoint = "https://api.twitter.com/2/tweets?tweet.fields=attachments&expansions=attachments.media_keys&media.fields=url"
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	copiedURL := *u
	copiedURL.Path = path.Join(copiedURL.Path, twitttID)
	endpoint = copiedURL.String()

	client := &http.Client{}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)

	resq, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var r io.Reader = resq.Body
	//r = io.TeeReader(r, os.Stderr)

	var twResult TwitterApiResult
	err = json.NewDecoder(r).Decode(&twResult)

	if err != nil {
		log.Fatal(err)
	}

	var photoURL = twResult.Includes.Media[0].URL
	fmt.Println(photoURL)
	return photoURL
}
