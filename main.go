package main

import (
	"log"
	"strings"
	"time"

	"os"

	"github.com/ztrue/shutdown"

	"DebilBot/chatbot"
	"DebilBot/globals"

	"github.com/SevereCloud/vksdk/v2/api"
	longpoll "github.com/SevereCloud/vksdk/v2/longpoll-user"
	wrapper "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"
)

var (
	lp *longpoll.LongPoll
)

func main() {

	log.Println("Загрузка конфигов...")
	LoadConfig()

	log.Println("Загрузка команд...")
	LoadCommands()

	log.Println("Загрузка базы ответов...")
	LoadAnswers()

	globals.VK = api.NewVK(globals.AccessToken)

	go GoToOnline()

	go StartLongPoll()

	shutdown.Listen()
}

func StartLongPoll() {
	mode := longpoll.ReceiveAttachments + longpoll.ExtendedEvents
	lp, err := longpoll.NewLongPoll(globals.VK, mode)
	if err != nil {
		panic(err)
	}

	shutdown.Add(func() {
		// Безопасное завершение
		// Ждет пока соединение закроется и события обработаются
		lp.Shutdown()

		// Закрыть соединение
		// Требует lp.Client.Transport = &http.Transport{DisableKeepAlives: true}
		lp.Client.CloseIdleConnections()

		// log.Println("Отключение модулей...")
		// CloseModules()
		// log.Println("Модули отключены")
		log.Println("Пока :(")
		os.Exit(1)
	})

	w := wrapper.NewWrapper(lp)

	// event with code 4
	w.OnNewMessage(OnMessage)

	log.Println("Лонгпул запущен")

	if err := lp.Run(); err != nil {
		StartLongPoll()
	}
}

func GoToOnline() {
	for {
		globals.VK.AccountSetOnline(api.Params{
			"voip": 0,
		})
		time.Sleep(time.Minute * 5)
	}
}

func OnMessage(m wrapper.NewMessage) {
	mText := strings.ToLower(m.Text)
	for _, a := range Appeals {
		if strings.HasPrefix(mText, strings.ToLower(a.(string))) {
			go OnMessageToBot(m, strings.ToLower(a.(string)))
			break
		}
	}
}

func OnMessageToBot(m wrapper.NewMessage, appeal string) {
	rawText := strings.Replace(strings.ToLower(m.Text), appeal, "", 1)
	args := strings.Split(strings.ToLower(rawText), " ")

	messageInfo, err := globals.VK.MessagesGetByID(api.Params{
		"message_ids": m.MessageID,
	})
	if err != nil {
		return
	}
	if len(messageInfo.Items) == 0 {
		return
	}
	if val, ok := commandList[args[0]]; ok {
		val.Function(messageInfo.Items[0], args)
	} else {
		if globals.HasAnswers {
			chatbot.FindAndSendAnswer(messageInfo.Items[0], rawText)
		}
	}
}
