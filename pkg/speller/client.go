package speller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	models "github.com/gitkoDev/KODE-test-task/models/speller"
	"github.com/sirupsen/logrus"
)

const spellerUrl = "http://speller.yandex.net/services/spellservice.json/checkText"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) CheckText(postData models.SpellerRequest) (models.SpellerResponse, error) {
	response, err := http.PostForm(spellerUrl, url.Values{
		"text":   {postData.Text},
		"lang":   {postData.Lang},
		"format": {postData.Format},
	})
	if err != nil {
		return models.SpellerResponse{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.SpellerResponse{}, err
	}

	if err = response.Body.Close(); err != nil {
		return models.SpellerResponse{}, err
	}

	var spellerResponse models.SpellerResponse
	if err = json.Unmarshal(body, &spellerResponse); err != nil {
		return models.SpellerResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		logrus.Println(string(body))
		return spellerResponse, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return spellerResponse, nil
}
