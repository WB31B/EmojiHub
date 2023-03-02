package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Store []EmojiData

type EmojiData struct {
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Group    string   `json:"group"`
	HtmlCode []string `json:"htmlCode"`
	Unicode  []string `json:"unicode"`
}

func (ed EmojiData) Info() string {
	return fmt.Sprintf("[Name]: %s - [Category]: %s - [Group]: %s - [HtmlCode]: %s - [Unicode]: %s",
		ed.Name, ed.Category, ed.Group, ed.HtmlCode, ed.Unicode)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECT")
			return nil
		},
		Timeout: time.Second * 15,
	}

	resp, err := client.Get("https://emojihub.yurace.pro/api/all")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Responce status:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var store Store
	if err = json.Unmarshal(body, &store); err != nil {
		log.Fatal(err)
	}

	for _, item := range store {
		fmt.Println(item.Info())
	}
}
