package main

import (
	"DebilBot/commands"
	"DebilBot/globals"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

type commandFunction = func(messageData object.MessagesMessage, args []string)

type Command struct {
	Name        string
	Description string
	Icon        string
	Function    commandFunction
	isHidden    bool
}

var (
	commandList map[string]Command
)

func AllCommands(messageData object.MessagesMessage, args []string) {
	resultText := "Все команды:\n\n"

	for _, v := range commandList {
		if !v.isHidden {
			resultText = resultText + v.Icon + " " + v.Name + " - " + v.Description + "\n"
		}
	}

	globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   resultText,
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})
}

func HelpCommand(messageData object.MessagesMessage, args []string) {
	globals.VK.MessagesSend(api.Params{
		"peer_id":   messageData.PeerID,
		"message":   globals.BotSettings.Get("info_text").(string),
		"random_id": 0,
		"reply_to":  globals.CanReply(messageData.ID),
	})
}

func LoadCommands() {
	commandList = make(map[string]Command)

	commandList["помощь"] = Command{
		Name:        "помощь",
		Description: "Информация о боте",
		Icon:        "🚑",
		Function:    HelpCommand,
		isHidden:    false,
	}
	commandList["команды"] = Command{
		Name:        "команды",
		Description: "Все команды",
		Icon:        "📱",
		Function:    AllCommands,
		isHidden:    false,
	}
	commandList["тест"] = Command{
		Name:        "пинг",
		Description: "Проверяет бота",
		Icon:        "💡",
		Function:    commands.TestCommand,
		isHidden:    false,
	}
	commandList["видео"] = Command{
		Name:        "видео [запрос]",
		Description: "Ищу видео в ВК по запросу",
		Icon:        "🎬",
		Function:    commands.FindVideo,
		isHidden:    false,
	}
	commandList["фото"] = Command{
		Name:        "фото [запрос]",
		Description: "Ищу фото в ВК по запросу",
		Icon:        "📷",
		Function:    commands.FindPhoto,
		isHidden:    false,
	}
	commandList["гиф"] = Command{
		Name:        "гиф [запрос]",
		Description: "Ищу гиф-анимации в ВК по запросу",
		Icon:        "🎞",
		Function:    commands.FindGIF,
		isHidden:    false,
	}
	commandList["время"] = Command{
		Name:        "время",
		Description: "Показать время у бота",
		Icon:        "⏰",
		Function:    commands.Time,
		isHidden:    false,
	}
	commandList["оцени"] = Command{
		Name:        "оцени",
		Description: "Объективная оценка от бота",
		Icon:        "💯",
		Function:    commands.Rate,
		isHidden:    false,
	}
}
