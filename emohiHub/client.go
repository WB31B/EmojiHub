package emohihub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("Timeout can not be zero!")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetEmojies() ([]EmojiData, error) {
	resp, err := c.client.Get("https://emojihub.yurace.pro/api/all")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var emStore emojiesStore
	if err = json.Unmarshal(body, &emStore); err != nil {
		return nil, err
	}

	return emStore, nil
}

func (c Client) GetCategoryEmojies(name string) ([]EmojiData, error) {
	url := fmt.Sprintf("https://emojihub.yurace.pro/api/all/category/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var emStore emojiesStore
	if err = json.Unmarshal(body, &emStore); err != nil {
		return nil, err
	}

	return emStore, nil
}

func (c Client) GetGroupEmojies(name string) ([]EmojiData, error) {
	url := fmt.Sprintf("https://emojihub.yurace.pro/api/all/group/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var emStore emojiesStore
	if err = json.Unmarshal(body, &emStore); err != nil {
		return nil, err
	}

	return emStore, nil
}

func (c Client) GetRandomEmoji() (EmojiData, error) {
	resp, err := c.client.Get("https://emojihub.yurace.pro/api/random")
	if err != nil {
		return EmojiData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmojiData{}, err
	}

	var emStore randomEmoji
	if err = json.Unmarshal(body, &emStore); err != nil {
		return EmojiData{}, err
	}

	return EmojiData(emStore), nil
}

func (c Client) GetGroupRandomEmoji(name string) (EmojiData, error) {
	url := fmt.Sprintf("https://emojihub.yurace.pro/api/random/group/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return EmojiData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmojiData{}, err
	}

	var emStore randomEmoji
	if err = json.Unmarshal(body, &emStore); err != nil {
		return EmojiData{}, err
	}

	return EmojiData(emStore), nil
}

func (c Client) GetCategoryRandomEmoji(name string) (EmojiData, error) {
	url := fmt.Sprintf("https://emojihub.yurace.pro/api/random/category/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return EmojiData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmojiData{}, err
	}

	var emStore randomEmoji
	if err = json.Unmarshal(body, &emStore); err != nil {
		return EmojiData{}, err
	}

	return EmojiData(emStore), nil
}