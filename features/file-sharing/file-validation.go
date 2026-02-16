package filesharing

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ValidateFile(fileid string, update *tgbotapi.Update) bool {

	switch {
	case update.Message.Document != nil:
		mime := update.Message.Document.MimeType
		fmt.Println(mime)
		if mime == "image/jpeg" {
			return false
		}
		return true

	}

	return false

}
