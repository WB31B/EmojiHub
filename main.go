package main

import (
	"emojihub/httpclient/emohihub"
	"fmt"
	"log"
	"time"
)

func printError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	emojiClient, err := emohihub.NewClient(time.Second * 10)
	printError(err)

	emojies, err := emojiClient.GetEmojies()
	printError(err)

	emojiesCategory, err := emojiClient.GetCategoryEmojies("travel-and-places")
	printError(err)

	emojiesGroup, err := emojiClient.GetGroupEmojies("animal-bird")
	printError(err)

	randomEmoji, err := emojiClient.GetRandomEmoji()
	printError(err)

	groupRandomEmoji, err := emojiClient.GetRandomGroup("face-positive")
	printError(err)

	categoryRandomEmoji, err := emojiClient.GetRandomCategory("food-and-drink")
	printError(err)

	printEmojies("Emojies:", emojies)
	printEmojies("Emojies category:", emojiesCategory)
	printEmojies("Emojies group:", emojiesGroup)

	fmt.Println("Random emoji:", randomEmoji.Info())
	fmt.Println("Random group:", groupRandomEmoji.Info())
	fmt.Println("Random category:", categoryRandomEmoji.Info())
}

func printEmojies(title string, typeEmoji []emohihub.EmojiData) {
	for _, item := range typeEmoji {
		fmt.Println(title, item.Info())
	}
}
