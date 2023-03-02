package main

import (
	"emojihub/httpclient/emohihub"
	"fmt"
	"log"
	"time"
)

func main() {
	emojiClient, err := emohihub.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	emojies, err := emojiClient.GetEmojies()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range emojies {
		fmt.Println(item.Info())
	}
}
