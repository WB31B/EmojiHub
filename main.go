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

	emojiesCategory, err := emojiClient.GetCategoryEmojies("travel-and-places")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range emojiesCategory {
		fmt.Println("[Emojo-Category]:", item.Info())
	}

	emojiesGroup, err := emojiClient.GetGroupEmojies("animal-bird")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range emojiesGroup {
		fmt.Println("[Emoji-Group]:", item.Info())
	}

	randomEmoji, err := emojiClient.GetRandomEmoji()
	if err != nil {
		log.Fatal(err)
	}

	groupRandomEmoji, err := emojiClient.GetGroupRandomEmoji("face-positive")
	if err != nil {
		log.Fatal(err)
	}

	categoryRandomEmoji, err := emojiClient.GetCategoryRandomEmoji("food-and-drink")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Random-Emoji]:", randomEmoji.Info())
	fmt.Println("[Group-Random]:", groupRandomEmoji.Info())
	fmt.Println("[Category-Random]: ", categoryRandomEmoji.Info())
}
