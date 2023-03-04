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

const URL = "https://emojihub.yurace.pro/api"

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
	url := fmt.Sprintf("%s/all", URL)
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

func (c Client) GetCategoryEmojies(caterogy string) ([]EmojiData, error) {
	url := fmt.Sprintf("%s/all/category/%s", URL, caterogy)
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

func (c Client) GetGroupEmojies(group string) ([]EmojiData, error) {
	url := fmt.Sprintf("%s/all/group/%s", URL, group)
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
	url := fmt.Sprintf("%s/random", URL)
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

func (c Client) GetRandomGroup(group string) (EmojiData, error) {
	url := fmt.Sprintf("%s/random/group/%s", URL, group)
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

func (c Client) GetRandomCategory(category string) (EmojiData, error) {
	url := fmt.Sprintf("%s/random/category/%s", URL, category)
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