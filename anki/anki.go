package anki

import (
	"context"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type Card struct {
	Kanji    string
	Furigana string
	English  string
}

func SaveCard(filepath string, word string) error {
	ctx := context.Background()

	lang, err := language.Parse("en-US")
	if err != nil {
		return err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{word}, lang, &translate.Options{
		Source: language.Japanese,
		Format: translate.Text,
	})
	if err != nil {
		return err
	}
	if len(resp) == 0 {
		return err
	}
	cardData := []string{word, resp[0].Text}
	delimitedData := strings.Join(cardData, ",")
	byteData := []byte(delimitedData)
	err = os.WriteFile(filepath, byteData, 0644)
	if err != nil {
		return err
	}
	return nil
}
