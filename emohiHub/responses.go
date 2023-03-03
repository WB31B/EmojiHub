package emohihub

import "fmt"

// type emojiesStore []EmojiData
type randomEmoji EmojiData
type emojiesStore []EmojiData

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
