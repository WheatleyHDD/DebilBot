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
}

var (
	commandList map[string]Command
)

func AllCommands(messageData object.MessagesMessage, args []string) {
	resultText := "Все команды:\n\n"

	for k, v := range commandList {
		resultText = resultText + v.Icon + " " + v.Name + " - " + v.Description + "\n"
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
	}
	commandList["команды"] = Command{
		Name:        "команды",
		Description: "Все команды",
		Icon:        "📱",
		Function:    AllCommands,
	}
	commandList["тест"] = Command{
		Name:        "тест",
		Description: "Тестовая команда",
		Icon:        "💡",
		Function:    commands.TestCommand,
	}
	commandList["видео"] = Command{
		Name:        "видео [запрос]",
		Description: "Ищу видео в ВК по запросу",
		Icon:        "🎬",
		Function:    commands.FindVideo,
	}
	commandList["фото"] = Command{
		Name:        "фото [запрос]",
		Description: "Ищу фото в ВК по запросу",
		Icon:        "📷",
		Function:    commands.FindPhoto,
	}
	commandList["гиф"] = Command{
		Name:        "гиф [запрос]",
		Description: "Ищу гиф-анимации в ВК по запросу",
		Icon:        "🎞",
		Function:    commands.FindGIF,
	}
}
