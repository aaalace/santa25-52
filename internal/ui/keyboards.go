package ui

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreatePeopleListKeyboard(names []string) api.InlineKeyboardMarkup {
	var keyboardRows [][]api.InlineKeyboardButton

	for i := 0; i < len(names); i++ {
		button := api.NewInlineKeyboardButtonData(names[i], names[i])
		keyboardRows = append(keyboardRows, []api.InlineKeyboardButton{button})
	}

	return api.NewInlineKeyboardMarkup(keyboardRows...)
}
