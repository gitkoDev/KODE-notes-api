package models

import (
	"strings"

	models "github.com/gitkoDev/KODE-test-task/models/speller"
	"github.com/gitkoDev/KODE-test-task/pkg/speller"
	"github.com/sirupsen/logrus"
)

type Note struct {
	Content string `json:"content"`
}

var client = speller.NewClient()

func (n *Note) Validate() {
	request := models.SpellerRequest{Lang: "ru", Text: n.Content}
	data, err := client.CheckText(request)
	if err != nil {
		logrus.Fatalln("error validating:", err)
	}

	for _, misspelling := range data {
		n.Content = strings.Replace(n.Content, misspelling.Word, misspelling.Suggestions[0], -1)
	}

}
